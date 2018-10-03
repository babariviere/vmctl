package vmctl

// VM represents the spawn parameters used by qemu
type VM struct {
	// Name of the VM
	Name string `yaml:"name"`
	// System of the VM, used to spawn with qemu
	System string `yaml:"system"`
	// Disks used by the vm
	Disks []Disk `yaml:"disks"`
	// CPU configuration
	CPU CPU `yaml:"cpu"`
}
