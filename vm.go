package vmctl

type Vm struct {
	Name  string `yaml:"name"`
	Disks []Disk `yaml:"disks"`
	Cpu   Cpu    `yaml:"cpu"`
}
