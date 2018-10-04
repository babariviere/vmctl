package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/babariviere/vmctl"
	"gopkg.in/yaml.v2"
)

const usage = `usage: vmctl <command> [arguments]

Commands:
	create	create disk images with qemu-img
	run		run a VM in Qemu (alias: spawn)
`

var commands = map[string]Command{
	"run":    &RunCommand{},
	"create": &CreateCommand{},
}

var aliases = map[string]string{
	"spawn": "run",
}

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

// TODO: commands add, list, remove, info and run/spawn
func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	cmd := os.Args[1]
	if val, ok := aliases[cmd]; ok {
		cmd = val
	}

	if val, ok := commands[cmd]; ok {
		if err := val.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			fmt.Println(val.Usage())
			os.Exit(1)
		}

		if err := val.Spawn(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if cmd == "help" {
		if len(os.Args) > 2 {
			cmd := os.Args[2]
			if val, ok := aliases[cmd]; ok {
				cmd = val
			}
			if val, ok := commands[cmd]; ok {
				fmt.Println(val.Usage())
			} else {
				fmt.Println(usage)
			}
		} else {
			fmt.Println(usage)
		}
	} else {
		fmt.Println(usage)
		os.Exit(1)
	}
}
