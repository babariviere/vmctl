# vmctl: A tool to manage your VMs with Qemu

## VM Configuration

The configuration is written in YAML for easy reading.

Here is an example

```yaml
name: Linux # name of your VM
system: x86_64 # system architecture
# All of your drivers
drives:
  - name: OS # optional: name of the drive (useful for creation)
    type: qcow2 # default to auto
    path: path/to/image
    readonly: true # if the disk should be readonly
    interface: ide # interface of the disk
    media: disk # or cdrom
    size: 10G # useful for disk creation
# Cpu configuration
cpu:
  count: 2
  arch: Haswell
# Ram configuration
memory: 10G
# VGA configuration
vga: std
# Enable KVM
kvm: true
```

All options can be found in the manual of qemu.

## Usage

Spawn a VM: `vmctl run <config.yml>`

Create a disk: `vmctl create <config.yml> [disk_name]`

## TODO

- [ ] Spawn VM by name and not by configuration file
- [ ] Spawn editor to edit VM configuration
- [ ] Keep VM files in a folder
- [ ] Commands: add, list, edit, delete and info
