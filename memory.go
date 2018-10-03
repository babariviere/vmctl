package vmctl

// Memory represents memory used by vm
type Memory string

// ToQemu convert memory to qemu command line argument
func (m Memory) ToQemu() ([]string, error) {
	return []string{"-m", string(m)}, nil
}
