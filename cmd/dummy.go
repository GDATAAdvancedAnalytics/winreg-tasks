//go:build debug
// +build debug

// SPDX-License-Identifier: MIT

package main

import "fmt"

func dummy(args ...string) {
	switch len(args) {
	case 0:
		fmt.Println("no args")

	case 1:
		fmt.Printf("one arg: %s\n", args[0])

	default:
		fmt.Printf("more than one arg: %v\n", args)
	}
}

func init() {
	registerCommand(Command{
		Name: "dummy",
		Args: []string{"[foo]", "[bar]"},
		// RequiredArgCount: 1,
		Func: dummy,
	})
}
