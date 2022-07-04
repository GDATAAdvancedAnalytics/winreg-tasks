// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/actions"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/dynamicinfo"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/triggers"
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
	dynamicInfo, err := dynamicinfo.FromBytes(data)
	if err != nil {
		return "", err
	}

	ret := fmt.Sprintf(
		"Creation Time: %s, Last Run Time: %s, Last Error Code: 0x%08x",
		dynamicInfo.CreationTime.String(), dynamicInfo.LastRunTime.String(),
		dynamicInfo.LastErrorCode,
	)

	return ret, nil
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
