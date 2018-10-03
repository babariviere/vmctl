package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/babariviere/vmctl"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

const usage = `usage: vmctl <command> [arguments]

Commands:
	run		run a VM in Qemu (alias: spawn)
`

type runCommand struct {
	file string
}

func (r *runCommand) parse(args []string) error {
	if len(args) != 1 {
		return errors.New("expected 1 argument")
	}
	r.file = args[0]
	return nil
}

func (r runCommand) usage() string {
	return `usage: vmctl run <file>

Arguments:	
	file		path to vm's config
`
}

func (r runCommand) spawn() error {
	fmt.Println(r.file)
	buf, err := ioutil.ReadFile(r.file)
	if err != nil {
		return err
	}

	vm := vmctl.VM{}
	err = yaml.Unmarshal(buf, &vm)
	if err != nil {
		return err
	}

	spew.Dump(vm)

	cmdStr, err := vm.ToQemu()
	if err != nil {
		return err
	}
	fmt.Println("Result command:", cmdStr)
	if err != nil {
		return err
	}
	cmd := exec.Command(cmdStr[0], cmdStr[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("%+v\n", *cmd)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// TODO: commands add, list, remove, info and run/spawn
func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "spawn":
		fallthrough
	case "run":
		cmd := runCommand{}
		if err := cmd.parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := cmd.spawn(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "help":
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "spawn":
				fallthrough
			case "run":
				var cmd runCommand
				fmt.Println(cmd.usage())
			default:
				fmt.Println(usage)
			}
		} else {
			fmt.Println(usage)
		}
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
