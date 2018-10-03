package vmctl

// DiskType is a list of all disk type supported by qemu
type DiskType string

const (
	// Raw format is a plain binary image of the disc image, and is very portable. On filesystems that support sparse files, images in this format only use the space actually used by the data recorded in them.
	Raw DiskType = "raw"
	// CLoop format, mainly used for reading Knoppix and similar live CD image formats
	CLoop = "cloop"
	// Cow format, supported for historical reasons only and not available to QEMU on Windows
	Cow = "cow"
	// QCow format, supported for historical reasons and superseded by qcow2
	QCow = "qcow"
	// QCow2 format with a range of special features, including the ability to take multiple snapshots, smaller images on filesystems that don't support sparse files, optional AES encryption, and optional zlib compression
	QCow2 = "qcow2"
	// Vmdk VMware 3 & 4, or 6 image format, for exchanging images with that product
	Vmdk = "vmdk"
	// Vdi VirtualBox 1.1 compatible image format, for exchanging images with VirtualBox.
	Vdi = "vdi"
	// Vhdx Hyper-V compatible image format, for exchanging images with Hyper-V 2012 or later.
	Vhdx = "vhdx"
	// Vpc Hyper-V legacy image format, for exchanging images with Virtual PC / Virtual Server / Hyper-V 2008.
	Vpc = "vpc"
	// Auto format, choose raw by default
	Auto = "auto"
)

// Disk information to attach and/or create images
type Disk struct {
	Type     DiskType `yaml:"type"`
	Path     string   `yaml:"path"`
	ReadOnly bool     `yaml:"readonly"`
}
