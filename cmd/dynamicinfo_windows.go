// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/dynamicinfo"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

func dynamicInfo(args ...string) {
	key := openTaskKey(args[0])
	if key == 0 {
		log.Fatalln("cannot open task key")
	}
	defer key.Close()

	dump := false
	if len(args) > 1 && (args[1] == "-d" || args[1] == "--dump") {
		dump = true
	}

	dynamicInfoRaw, _, err := key.GetBinaryValue("DynamicInfo")
	if err != nil {
		log.Fatalf("cannot get dynamic info for task: %v", err)
	}

	if dump {
		hex := utils.Hexdump(dynamicInfoRaw, 16)
		fmt.Println(hex)
	}

	dynamicInfo, err := dynamicinfo.FromBytes(dynamicInfoRaw)
	if err != nil {
		log.Fatalf("cannot parse DynamicInfo: %v", err)
	}

	log.Printf("Creation Time: %s", dynamicInfo.CreationTime.String())
	log.Printf("Last Run Time: %s", dynamicInfo.LastRunTime.String())
	log.Printf("Task State: 0x%08x", dynamicInfo.TaskState)
	log.Printf("Last Error Code: 0x%08x", dynamicInfo.LastErrorCode)
	log.Printf("Last Successful Run Time: %s", dynamicInfo.LastSuccessfulRunTime.String())
}

func init() {
	registerCommand(Command{
		Name:             "dynamicinfo",
		Args:             []string{"<task id>", "[-d|--dump]"},
		RequiredArgCount: 1,
		Func:             dynamicInfo,
	})
}
