package vmctl

import (
	"errors"
)

// VM represents the spawn parameters used by qemu
type VM struct {
	// Name of the VM
	Name string `yaml:"name"`
	// System of the VM, used to spawn with qemu
	System string `yaml:"system"`
	// Drives used by the vm
	Drives []Drive `yaml:"drives"`
	// CPU configuration
	CPU `yaml:"cpu"`
	// Memory used by the vm
	Memory `yaml:"memory"`
	// VGA graphics of the vm
	VGA `yaml:"vga"`
	// Net is net interfaces and redirections
	Net `yaml:"net"`
	// Kvm specify if we enable kvm or not
	Kvm bool `yaml:"kvm"`
	// Snapshot enables temporary snapshot
	Snapshot bool `yaml:"snapshot"`
}

type errBuilder struct {
	buf []string
	err error
}

func (e *errBuilder) add(builder QemuBuilder) {
	if e.err != nil {
		return
	}

	buf, err := builder.ToQemu()
	if err != nil {
		e.err = err
		return
	}
	e.buf = append(e.buf, buf...)
}

func (e *errBuilder) addString(s string) {
	e.buf = append(e.buf, s)
}

// ToQemu converts VM to a qemu command
func (v VM) ToQemu() ([]string, error) {
	if len(v.Name) == 0 {
		return nil, errors.New("missing vm name")
	}
	if len(v.System) == 0 {
		return nil, errors.New("missing vm system")
	}
	if len(v.Drives) == 0 {
		return nil, errors.New("no disk for vm")
	}

	var builder errBuilder

	builder.addString("qemu-system-" + v.System)

	for i := 0; i < len(v.Drives); i++ {
		builder.add(v.Drives[i])
	}

	builder.add(v.CPU)
	builder.add(v.Memory)
	builder.add(v.VGA)
	builder.add(v.Net)

	if v.Kvm {
		builder.addString("-enable-kvm")
	}
	if v.Snapshot {
		builder.addString("-snapshot")
	}

	return builder.buf, builder.err
}
