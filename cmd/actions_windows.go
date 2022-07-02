// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/actions"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

func actionsHandler(args ...string) {
	key := openTaskKey(args[0])
	if key == 0 {
		log.Fatalln("cannot open task key")
	}
	defer key.Close()

	dump := false
	if len(args) > 1 && (args[1] == "-d" || args[1] == "--dump") {
		dump = true
	}

	actionsRaw, _, err := key.GetBinaryValue("Actions")
	if err != nil {
		log.Fatalf("cannot get actions for task: %v", err)
	}

	if dump {
		hex := utils.Hexdump(actionsRaw, 16)
		fmt.Println(hex)
	}

	actions, err := actions.FromBytes(actionsRaw)
	if err != nil {
		log.Fatalf("cannot parse actions: %v", err)
	}

	log.Println("Context: " + actions.Context)
	log.Println(`Actions:`)

	if len(actions.Properties) == 0 {
		log.Println("\t<no actions>")
		return
	}

	for _, props := range actions.Properties {
		log.Println("\t" + props.String())
	}
}

func init() {
	registerCommand(Command{
		Name:             "actions",
		Args:             []string{"<task id>", "[-d|--dump]"},
		RequiredArgCount: 1,
		Func:             actionsHandler,
	})
}
