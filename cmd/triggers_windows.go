// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func triggers(args ...string) {
	key := openTaskKey(args[0])
	if key == 0 {
		log.Fatalln("cannot open task key")
	}
	defer key.Close()

	dump := false
	if len(args) > 1 && (args[1] == "-d" || args[1] == "--dump") {
		dump = true
	}

	triggersRaw, _, err := key.GetBinaryValue("Triggers")
	if err != nil {
		log.Fatalf("cannot get triggers for task: %v", err)
	}

	if dump {
		hex := hexdump(triggersRaw)
		fmt.Println(hex)
	}

	triggers := generated.NewTriggers()
	if err = triggers.Read(kaitai.NewStream(bytes.NewReader(triggersRaw)), triggers, triggers); err != nil {
		log.Fatalf("cannot parse triggers: %v", err)
	}

	if len(triggers.Triggers) == 0 {
		return
	}

	log.Println("Triggers:")
	for _, trigger := range triggers.Triggers {
		switch props := trigger.Properties.(type) {
		case *generated.Triggers_BootTrigger:
			log.Printf("\tBootTrigger -> Id: \"%s\"", props.GenericData.TriggerId.Str)
		case *generated.Triggers_EventTrigger:
			log.Printf("\tEventTrigger -> Id \"%s\", Subscription: \"%s\"", props.GenericData.TriggerId.Str, props.Subscription.Content)
		case *generated.Triggers_IdleTrigger:
			log.Printf("\tIdleTrigger -> Id: \"%s\"", props.GenericData.TriggerId.Str)
		case *generated.Triggers_LogonTrigger:
			username := ""
			if props.User.SkipUser.Value != 0 {
				username = "<skipped>"
			} else {
				username = props.User.Username.String
				if username == "" && props.User.SkipSid.Value == 0 {
					username = binarySidToString(props.User.Sid.Data)
				}
			}
			log.Printf("\tLogonTrigger -> Id: \"%s\", User: \"%s\"", props.GenericData.TriggerId.Str, username)
		case *generated.Triggers_RegistrationTrigger:
			log.Printf("\tRegistrationTrigger -> Id: \"%s\"", props.GenericData.TriggerId.Str)
		case *generated.Triggers_SessionChangeTrigger:
			username := ""
			if props.User.SkipUser.Value == 0 {
				username = props.User.Username.String
			}
			log.Printf("\tSessionChangeTrigger -> Id: \"%s\", User: \"%s\", State: %d", props.GenericData.TriggerId.Str, username, props.StateChange)
		case *generated.Triggers_TimeTrigger:
			log.Printf("\tTimeTrigger -> Id: \"%s\"", props.TriggerId.Str)
		case *generated.Triggers_WnfStateChangeTrigger:
			log.Printf("\tWnfStateChangeTrigger -> Id: \"%s\", StateName: \"%s\", Data: \"%s\"", props.GenericData.TriggerId.Str, hexdump(props.StateName), hexdump(props.Data))
		default:
			log.Printf("unhandled trigger type %v", reflect.TypeOf(trigger.Properties))
		}
	}
}

func init() {
	registerCommand(Command{
		Name:             "triggers",
		RequiredArgCount: 1,
		Args:             []string{"<task id>", "[-d|--dump]"},
		Func:             triggers,
	})
}
