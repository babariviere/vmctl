package vmctl

import "fmt"

// TODO: list archs?

// CPU configuration, specify number of cpu allowed to use and it's architecture
type CPU struct {
	// Count is the number of cpu that the host will provide
	Count uint `yaml:"count"`
	// Arch is the target architecture
	Arch string `yaml:"arch"`
}

// ToQemu convert CPU to a qemu command line parameter
func (c CPU) ToQemu() ([]string, error) {
	if c.Count == 0 {
		c.Count = 1
	}
	if len(c.Arch) == 0 {
		c.Arch = "host"
	}

	return []string{"-smp", fmt.Sprint(c.Count), "-cpu", c.Arch}, nil
}
