package vmctl

type DiskType string

const (
	Raw   DiskType = "raw"
	CLoop          = "cloop"
	Cow            = "cow"
	QCow           = "qcow"
	QCow2          = "qcow2"
	Vmdk           = "vmdk"
	Vdi            = "vdi"
	Vhdx           = "vhdx"
	Vpc            = "vpc"
)

type Disk struct {
	Type DiskType `yaml:"type"`
	Path string   `yaml:"path"`
}
