package vmctl

type VM struct {
	Name  string `yaml:"name"`
	Disks []Disk `yaml:"disks"`
}
