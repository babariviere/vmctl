package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
	"vmctl"
)

// TODO: commands add, list, remove, info and run/spawn
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: vmctl vm.yaml")
		return
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	vm := vmctl.VM{}
	err = yaml.Unmarshal(buf, &vm)
	if err != nil {
		log.Fatalln(err)
	}

	spew.Dump(vm)

	cmd, err := vm.ToQemu()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Result command:", cmd)
}
