package vmctl

import (
	"errors"
	"strings"
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
	IfIde    DriveInterface = "ide"
	IfScsi                  = "scsi"
	IfSd                    = "sd"
	IfMtd                   = "mtd"
	IfFloppy                = "floppy"
	IfPflash                = "pflash"
	IfVirtio                = "virtio"
	IfNone                  = "none"
)

// DriveMedia is the media type of the disk
type DriveMedia string

const (
	// Disk set drive as a disk
	Disk DriveMedia = "disk"
	// CdRom set drive as a cdrom
	CdRom = "cdrom"
)

// DriveCache is the cache used by the disk
type DriveCache string

const (
	// CacheNone disables cache
	CacheNone DriveCache = "none"
	// CacheWriteback enables writeback
	CacheWriteback = "writeback"
	// CacheUnsafe enables writeback and no flush
	CacheUnsafe = "unsafe"
	// CacheDirectSync enables direct sync
	CacheDirectSync = "directsync"
	// CacheWriteThrough disables direct sync
	CacheWriteThrough = "writethrough"
)

// TODO: implement unmarshal to check for invalid values

// Drive information to attach and/or create images
type Drive struct {
	Name      string         `yaml:"name"`
	Type      DriveType      `yaml:"type"`
	Path      string         `yaml:"path"`
	ReadOnly  bool           `yaml:"readonly"`
	Interface DriveInterface `yaml:"interface"`
	Media     DriveMedia     `yaml:"media"`
	Cache     DriveCache     `yaml:"cache"`
	Size      string         `yaml:"size"`
}

// Create return command line to create drive
func (d Drive) Create() ([]string, error) {
	if len(d.Size) == 0 {
		return nil, errors.New("missing disk size")
	}
	if len(d.Type) == 0 {
		return nil, errors.New("missing disk type")
	}
	if len(d.Path) == 0 {
		return nil, errors.New("missing disk path")
	}

	var res []string

	res = append(res, "qemu-img")
	res = append(res, "create")
	res = append(res, "-f")
	res = append(res, string(d.Type))
	res = append(res, d.Path)
	res = append(res, d.Size)

	return res, nil
}

// ToQemu convert struct to command line arguments for qemu
func (d Drive) ToQemu() ([]string, error) {
	if len(d.Path) == 0 {
		return nil, errors.New("missing path for disk")
	}
	if len(d.Interface) == 0 {
		d.Interface = IfVirtio
	}
	if len(d.Media) == 0 {
		d.Media = Disk
	}
	if len(d.Cache) == 0 {
		d.Cache = CacheWriteback
	}

	var res []string
	var buf []string

	res = append(res, "-drive")
	buf = append(buf, "file="+d.Path)

	if len(d.Type) > 0 {
		buf = append(buf, "format="+string(d.Type))
	}
	if d.ReadOnly {
		buf = append(buf, "readonly")
	}
	buf = append(buf, "if="+string(d.Interface))
	buf = append(buf, "media="+string(d.Media))
	buf = append(buf, "cache="+string(d.Cache))

	res = append(res, strings.Join(buf, ","))
	return res, nil
}
