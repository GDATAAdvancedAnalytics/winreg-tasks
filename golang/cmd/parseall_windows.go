// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/golang/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/golang/utils"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"golang.org/x/sys/windows/registry"
)

func readAndParse(taskId string, key registry.Key, value string, parserCallback func(data []byte) (string, error), quiet bool) error {
	rawData, _, err := key.GetBinaryValue(value)
	if err != nil {
		return err
	}

	result, err := parserCallback(rawData)
	if err != nil {
		return err
	}

	if !quiet {
		log.Printf("Task %s - %s: %s", taskId, value, result)
	}

	return nil
}

func parseTriggers(data []byte) (string, error) {
	tt := generated.NewTriggers()
	s := kaitai.NewStream(bytes.NewReader(data))
	if err := tt.Read(s, tt, tt); err != nil {
		return "", fmt.Errorf("cannot parse trigger data: %v", err)
	}

	if len(tt.Triggers) == 0 {
		return "<no triggers>", nil
	}

	var triggers []string

	for _, trigger := range tt.Triggers {
		switch trigger.Properties.(type) {
		case *generated.Triggers_BootTrigger:
			triggers = append(triggers, "BootTrigger")
		case *generated.Triggers_EventTrigger:
			triggers = append(triggers, "EventTrigger")
		case *generated.Triggers_IdleTrigger:
			triggers = append(triggers, "IdleTrigger")
		case *generated.Triggers_LogonTrigger:
			triggers = append(triggers, "LogonTrigger")
		case *generated.Triggers_RegistrationTrigger:
			triggers = append(triggers, "RegistrationTrigger")
		case *generated.Triggers_SessionChangeTrigger:
			triggers = append(triggers, "SessionChangeTrigger")
		case *generated.Triggers_TimeTrigger:
			triggers = append(triggers, "TimeTrigger")
		case *generated.Triggers_WnfStateChangeTrigger:
			triggers = append(triggers, "WnfStateChangeTrigger")
		default:
			return "", fmt.Errorf("unhandled trigger type %v", reflect.TypeOf(trigger.Properties))
		}
	}

	if eof, err := s.EOF(); !eof {
		return "", fmt.Errorf("did not parse all data")
	} else if err != nil {
		return "", fmt.Errorf("error trying to eof-check: %v", err)
	}

	return strings.Join(triggers, ", "), nil
}

func parseActions(data []byte) (string, error) {
	actions := generated.NewActions()
	s := kaitai.NewStream(bytes.NewReader(data))
	if err := actions.Read(s, actions, actions); err != nil {
		return "", err
	}

	if len(actions.Actions) == 0 {
		return "<no actions>", nil
	}

	var actionCollection []string

	for _, action := range actions.Actions {
		switch action.Properties.(type) {
		case *generated.Actions_ComHandlerProperties:
			actionCollection = append(actionCollection, "ComHandler")
		case *generated.Actions_EmailTaskProperties:
			actionCollection = append(actionCollection, "EmailTask")
		case *generated.Actions_ExeTaskProperties:
			actionCollection = append(actionCollection, "ExeTask")
		case *generated.Actions_MessageboxTaskProperties:
			actionCollection = append(actionCollection, "MessageBoxTask")
		default:
			return "", fmt.Errorf("unhandled action type %v", reflect.TypeOf(action.Properties))
		}
	}

	if eof, err := s.EOF(); !eof {
		return "", fmt.Errorf("did not parse all data")
	} else if err != nil {
		return "", fmt.Errorf("error trying to eof-check: %v", err)
	}

	return strings.Join(actionCollection, ", "), nil
}

func parseDynamicInfo(data []byte) (string, error) {
	dynamicInfo := generated.NewDynamicInfo()
	s := kaitai.NewStream(bytes.NewReader(data))
	if err := dynamicInfo.Read(s, dynamicInfo, dynamicInfo); err != nil {
		return "", err
	}

	var info []string

	creationTime := utils.TimeFromFILETIME(int64(dynamicInfo.CreationTime))
	info = append(info, fmt.Sprintf("Creation: %s", creationTime))

	lastRunTime := utils.TimeFromFILETIME(int64(dynamicInfo.LastRunTime))
	info = append(info, fmt.Sprintf("Last Run: %s", lastRunTime))

	taskState := dynamicInfo.TaskState
	info = append(info, fmt.Sprintf("Task State: 0x%08x", taskState))

	lastErrorCode := dynamicInfo.LastErrorCode
	info = append(info, fmt.Sprintf("Last Error: 0x%08x", lastErrorCode))

	lastSuccessfulRunTime := utils.TimeFromFILETIME(int64(dynamicInfo.LastSuccessfulRunTime))
	info = append(info, fmt.Sprintf("Last Successful Run: %s", lastSuccessfulRunTime))

	if eof, err := s.EOF(); !eof {
		return "", fmt.Errorf("did not parse all data")
	} else if err != nil {
		return "", fmt.Errorf("error trying to eof-check: %v", err)
	}

	return strings.Join(info, ", "), nil
}

func parseAll(args ...string) {
	quiet := false
	if len(args) > 0 && (args[0] == "-q" || args[0] == "--quiet") {
		quiet = true
	}

	taskDir, err := openKey(`Tasks`)
	if err != nil {
		log.Println(err)
		return
	}
	defer taskDir.Close()

	tasks, err := taskDir.ReadSubKeyNames(-1)
	if err != nil {
		log.Printf("cannot get task list from registry: %v\n", err)
		return
	}

	for _, taskId := range tasks {
		key := openTaskKey(taskId)

		if err := readAndParse(taskId, key, "Actions", parseActions, quiet); err != nil {
			log.Printf("error reading Actions of task %s: %v", taskId, err)
		}

		if err := readAndParse(taskId, key, "Triggers", parseTriggers, quiet); err != nil {
			log.Printf("error reading Triggers of task %s: %v", taskId, err)
		}

		if err := readAndParse(taskId, key, "DynamicInfo", parseDynamicInfo, quiet); err != nil {
			log.Printf("error reading DynamicInfo of task %s: %v", taskId, err)
		}
	}

}

func init() {
	registerCommand(Command{
		Name: "parseall",
		Args: []string{"[-q|--quiet]"},
		Func: parseAll,
	})
}
