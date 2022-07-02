// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/actions"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/triggers"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
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
	triggs, err := triggers.FromBytes(data, time.Local)
	if err != nil {
		return "", err
	}

	if len(triggs.Triggers) == 0 {
		return "<no triggers>", nil
	}

	var triggers []string

	for _, trigger := range triggs.Triggers {
		triggers = append(triggers, trigger.Name())
	}

	return strings.Join(triggers, ", "), nil
}

func parseActions(data []byte) (string, error) {
	actions, err := actions.FromBytes(data)
	if err != nil {
		return "", err
	}

	if len(actions.Properties) == 0 {
		return "<no actions>", nil
	}

	var propCollection []string

	for _, props := range actions.Properties {
		propCollection = append(propCollection, props.Name())
	}

	return strings.Join(propCollection, ", "), nil
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
