package vmctl

// VGA graphics of the vm
type VGA string

// ToQemu convert vga to qemu command line parameter
func (v VGA) ToQemu() ([]string, error) {
	if len(v) == 0 {
		v = "std"
	}
	return []string{"-vga", string(v)}, nil
}
