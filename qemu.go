package vmctl

// Qemu interface to create a full qemu command
type Qemu interface {
	ToQemu() string
}
