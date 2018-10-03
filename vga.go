package vmctl

// VGA graphics of the vm
type VGA string

// ToQemu convert vga to qemu command line parameter
func (v VGA) ToQemu() (string, error) {
	if len(v) == 0 {
		return "-vga std", nil
	}
	return "-vga " + string(v), nil
}
