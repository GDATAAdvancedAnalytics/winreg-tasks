// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Command struct {
	Name             string
	Args             []string
	RequiredArgCount int
	Func             func(args ...string)
}

func help() {
	fmt.Printf("Usage: %s <command> [args...]\n", os.Args[0])

	fmt.Println("Available Commands:")
	for _, command := range commands {
		fmt.Printf("\t%s %s\n", command.Name, strings.Join(command.Args, " "))
	}
}

// noreturn if argument count mismatch
func ensureArgs(n int) {
	if len(os.Args) < 2+n {
		fmt.Println("argument count mismatch!")
		help()
		os.Exit(1)
	}
}

var (
	commands map[string]Command
)

func registerCommand(command Command) {
	if commands == nil {
		commands = make(map[string]Command)
	}
	commands[command.Name] = command
}

func main() {
	switch len(os.Args) {
	case 1:
		help()
		return
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			help()
			return
		}
	}

	if command, ok := commands[os.Args[1]]; ok {
		ensureArgs(command.RequiredArgCount)

		args := []string{}
		if len(os.Args) > 2 {
			args = os.Args[2:]
		}

		command.Func(args...)
	} else {
		log.Println("Invalid command!")
		help()
	}
}
