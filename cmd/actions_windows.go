// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func actions(args ...string) {
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
		hex := hexdump(actionsRaw)
		fmt.Println(hex)
	}

	actions := generated.NewActions()
	if err = actions.Read(kaitai.NewStream(bytes.NewReader(actionsRaw)), actions, actions); err != nil {
		log.Fatalf("cannot parse actions: %v", err)
	}

	log.Println("Context: " + actions.Context.Str)
	log.Println(`Actions:`)

	if len(actions.Actions) == 0 {
		log.Println("\t<no actions>")
		return
	}

	for _, action := range actions.Actions {
		switch props := action.Properties.(type) {
		case *generated.Actions_ComHandlerProperties:
			if clsid, err := utils.UuidFromMemory(props.Clsid); err == nil {
				log.Printf("\t"+`ComHandler -> Id "%s", CLSID {%s}, Data "%s"`, action.Id.Str, clsid.String(), props.Data.Str)
			} else {
				log.Printf("\tComHandler error - cannot convert clsid %v: %v\n", props.Clsid, err)
			}

		case *generated.Actions_EmailTaskProperties:
			log.Printf("\t"+`Email Task -> Id "%s", Sender "%s", Receiver "%s", Subject "%s"`, action.Id.Str, props.Server.Str, props.To.Str, props.Subject.Str)

		case *generated.Actions_ExeTaskProperties:
			log.Printf("\t"+`Execution Task -> Id "%s", Executable "%s", Args "%s"`, action.Id.Str, props.Command.Str, props.Arguments.Str)

		case *generated.Actions_MessageboxTaskProperties:
			log.Printf("\t"+`MessageBox Task -> Id "%s", Caption "%s", Content "%s"`, action.Id.Str, props.Caption.Str, props.Content.Str)
		default:
			log.Printf("\tunhandled action type %v", reflect.TypeOf(action.Properties))
		}
	}
}

func init() {
	registerCommand(Command{
		Name:             "actions",
		Args:             []string{"<task id>", "[-d|--dump]"},
		RequiredArgCount: 1,
		Func:             actions,
	})
}
