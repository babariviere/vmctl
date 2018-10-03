package vmctl

import (
	"errors"
	"strings"
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
	CPU CPU `yaml:"cpu"`
	// Memory used by the vm
	Memory Memory `yaml:"memory"`
	// VGA graphics of the vm
	VGA VGA `yaml:"vga"`
	// Kvm specify if we enable kvm or not
	Kvm bool `yaml:"kvm"`
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
	e.buf = append(e.buf, buf)
}

func (e *errBuilder) addString(s string) {
	e.buf = append(e.buf, s)
}

// ToQemu converts VM to a qemu command
func (v VM) ToQemu() (string, error) {
	if len(v.Name) == 0 {
		return "", errors.New("missing vm name")
	}
	if len(v.System) == 0 {
		return "", errors.New("missing vm system")
	}
	if len(v.Drives) == 0 {
		return "", errors.New("no disk for vm")
	}

	var builder errBuilder

	builder.addString("qemu-system-" + v.System)

	for i := 0; i < len(v.Drives); i++ {
		builder.add(v.Drives[i])
	}

	builder.add(v.CPU)
	builder.add(v.Memory)
	builder.add(v.VGA)

	if v.Kvm {
		builder.addString("--enable-kvm")
	}

	return strings.Join(builder.buf, " "), builder.err
}
