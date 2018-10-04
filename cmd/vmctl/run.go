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

// RunCommand is a struct holding the run subcommand
type RunCommand struct {
	file string
}

// Parse arguments for run command
func (r *RunCommand) Parse(args []string) error {
	if len(args) != 1 {
		return errors.New("expected 1 argument")
	}
	r.file = args[0]
	return nil
}

// Usage returns usage of subcommand
func (r RunCommand) Usage() string {
	return `usage: vmctl run <file>

Arguments:	
	file		path to vm's config
`
}

// Spawn subcommand
func (r RunCommand) Spawn() error {
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
