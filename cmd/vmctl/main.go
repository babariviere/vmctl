package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"vmctl"
)

// TODO: commands add, list, remove, info and run/spawn
func main() {
	fmt.Println("vmctl tool")
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

	fmt.Printf("%+v\n", vm)
}
