package vmctl

// QemuBuilder interface to create a full qemu command
type QemuBuilder interface {
	ToQemu() ([]string, error)
}
