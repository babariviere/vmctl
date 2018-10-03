package vmctl

import (
	"bytes"
	"errors"
)

// DriveType is a list of all disk type supported by qemu
type DriveType string

const (
	// Raw format is a plain binary image of the disc image, and is very portable. On filesystems that support sparse files, images in this format only use the space actually used by the data recorded in them.
	Raw DriveType = "raw"
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

// DriveInterface is the interface used by the disk
type DriveInterface string

const (
	Ide    DriveInterface = "ide"
	Scsi                  = "scsi"
	Sd                    = "sd"
	Mtd                   = "mtd"
	Floppy                = "floppy"
	Pflash                = "pflash"
	Virtio                = "virtio"
	None                  = "none"
)

// DriveMedia is the media type of the disk
type DriveMedia string

const (
	// Disk set drive as a disk
	Disk DriveMedia = "disk"
	// CdRom set drive as a cdrom
	CdRom = "cdrom"
)

// Drive information to attach and/or create images
type Drive struct {
	Type      DriveType      `yaml:"type"`
	Path      string         `yaml:"path"`
	ReadOnly  bool           `yaml:"readonly"`
	Interface DriveInterface `yaml:"interface"`
	Media     DriveMedia     `yaml:"media"`
}

// ToQemu convert struct to command line arguments for qemu
func (d Drive) ToQemu() (string, error) {
	if len(d.Path) == 0 {
		return "", errors.New("missing path for disk")
	}
	if len(d.Interface) == 0 {
		d.Interface = Virtio
	}
	if len(d.Media) == 0 {
		d.Media = Disk
	}

	var buf bytes.Buffer

	buf.WriteString("-drive file=" + d.Path)

	if len(d.Type) > 0 {
		buf.WriteString(",format=" + string(d.Type))
	}
	if d.ReadOnly {
		buf.WriteString(",readonly")
	}
	buf.WriteString(",if=" + string(d.Interface))
	buf.WriteString(",media=" + string(d.Media))

	return buf.String(), nil
}
