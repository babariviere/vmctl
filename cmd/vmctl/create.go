package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// CreateCommand represents the create subcommand
type CreateCommand struct {
	name  string
	drive string
}

// Parse command line arguments for the create subcommand
func (c *CreateCommand) Parse(args []string) error {
	if len(args) == 0 {
		return errors.New("missing 1 argument")
	}
	c.name = args[0]
	if len(args) > 1 {
		c.drive = args[1]
	}
	return nil
}

// Usage for the create subcommand
func (c CreateCommand) Usage() string {
	return `usage: vmctl create <name/file> [drive_name]

Arguments:
	name			name of the vm
	file			path to vm's config
	drive_name		name of drive to create, if not specified, a prompt will be displayed
`
}

// Spawn launch the create subcommand
// TODO: allow for resize if file exists
func (c CreateCommand) Spawn(config *Config) error {
	vm, err := config.GetVMOrRead(c.name)
	if err != nil {
		return err
	}

	var cmdStr []string

	if len(c.drive) == 0 {
		fmt.Println("Select drive to create:")
		for i, drive := range vm.Drives {
			fmt.Println("[", i+1, "]", "type:", string(drive.Type), "path:", drive.Path)
		}

		var index int
		fmt.Scanln(&index)

		for index < 1 || index > len(vm.Drives) {
			fmt.Println("out of range")
			fmt.Scanln(&index)
		}

		cmdStr, err = vm.Drives[index-1].Create()
		if err != nil {
			return err
		}
	} else {
		for _, drive := range vm.Drives {
			if c.drive == drive.Name {
				cmdStr, err = drive.Create()
				if err != nil {
					return err
				}
			}
		}
	}
	if len(cmdStr) == 0 {
		return errors.New("drive not found")
	}
	fmt.Println(cmdStr)
	cmd := exec.Command(cmdStr[0], cmdStr[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("%+v\n", *cmd)
	return cmd.Run()
}
