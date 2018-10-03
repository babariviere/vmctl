package vmctl

// TODO: list archs?

type Cpu struct {
	Count uint   `yaml:"count"`
	Arch  string `yaml:"arch"`
}
