package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/babariviere/vmctl"
	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v2"
)

var local = configdir.LocalConfig("vmctl")

// Config contains list of all vms
type Config struct {
	files []string
	vms   []vmctl.VM
}

// NewConfig creates a config struct and feed list of config files
func NewConfig() (Config, error) {
	configdir.MakePath(local)
	var config Config
	files, err := listVMFiles()
	if err != nil {
		return config, err
	}
	config.files = files
	return config, nil
}

// ListVMFiles returns all vm file name from config folder
func listVMFiles() ([]string, error) {
	files, err := ioutil.ReadDir(local)
	if err != nil {
		return nil, err
	}

	var buf []string

	for _, f := range files {
		buf = append(buf, f.Name())
	}
	return buf, nil
}

// ListVMs return a list of all vms in config folder
func (c *Config) ListVMs() ([]vmctl.VM, error) {
	if len(c.vms) == 0 && len(c.files) > 0 {
		for _, file := range c.files {
			cfg, err := OpenVMConfig(file)
			if err != nil {
				return nil, err
			}
			c.vms = append(c.vms, cfg)
		}
	}
	return c.vms, nil
}

// GetVM query vm from it's name
func (c *Config) GetVM(name string) (vmctl.VM, error) {
	if len(c.vms) == 0 {
		_, err := c.ListVMs()
		if err != nil {
			return vmctl.VM{}, err
		}
	}

	for _, vm := range c.vms {
		if vm.Name == name {
			return vm, nil
		}
	}
	return vmctl.VM{}, errors.New("cannot find VM with name " + name)
}

// GetVMOrRead query vm or read file if cannot found vm by it's name
func (c *Config) GetVMOrRead(name string) (vmctl.VM, error) {
	vm, err := c.GetVM(name)
	if err != nil {
		if _, errfile := os.Stat(name); !os.IsNotExist(errfile) {
			vm, errfile = OpenVMConfig(name)
			if errfile != nil {
				return vmctl.VM{}, errfile
			}
		} else {
			return vmctl.VM{}, err
		}
	}
	return vm, nil
}

// OpenVMConfig converts config to VM struct
func OpenVMConfig(path string) (vmctl.VM, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return vmctl.VM{}, err
	}

	vm := vmctl.VM{}
	err = yaml.Unmarshal(buf, &vm)
	if err != nil {
		return vmctl.VM{}, err
	}

	return vm, nil
}
