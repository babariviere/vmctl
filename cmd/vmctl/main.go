package main

import (
	"fmt"
	"os"
)

const usage = `usage: vmctl <command> [arguments]

Commands:
	run		run a VM in Qemu (alias: spawn)
`

var commands = map[string]Command{
	"run": &RunCommand{},
}

var aliases = map[string]string{
	"spawn": "run",
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
