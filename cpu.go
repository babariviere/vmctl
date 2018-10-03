package vmctl

// TODO: list archs?

// CPU configuration, specify number of cpu allowed to use and it's architecture
type CPU struct {
	// Count is the number of cpu that the host will provide
	Count uint `yaml:"count"`
	// Arch is the target architecture
	Arch string `yaml:"arch"`
}
