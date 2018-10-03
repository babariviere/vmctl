package vmctl

import (
	"bytes"
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
	CPU CPU `yaml:"cpu"`
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

	var buf bytes.Buffer

	buf.WriteString("qemu-system-" + v.System)

	for i := 0; i < len(v.Drives); i++ {
		dbuf, err := v.Drives[i].ToQemu()
		if err != nil {
			return "", err
		}
		buf.WriteString(" " + dbuf)
	}

	cbuf, err := v.CPU.ToQemu()
	if err != nil {
		return "", err
	}
	buf.WriteString(" " + cbuf)

	return buf.String(), nil
}
