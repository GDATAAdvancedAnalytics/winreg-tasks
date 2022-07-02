// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
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

	dynamicInfo := generated.NewDynamicInfo()
	if err = dynamicInfo.Read(kaitai.NewStream(bytes.NewReader(dynamicInfoRaw)), dynamicInfo, dynamicInfo); err != nil {
		log.Fatalf("cannot parse dynamic info: %v", err)
	}

	creationTime := utils.TimeFromFILETIME(int64(dynamicInfo.CreationTime))
	log.Printf("Creation Time: %s", creationTime)

	lastRunTime := utils.TimeFromFILETIME(int64(dynamicInfo.LastRunTime))
	log.Printf("Last Run Time: %s", lastRunTime)

	taskState := dynamicInfo.TaskState
	log.Printf("Task State: 0x%08x", taskState)

	lastErrorCode := dynamicInfo.LastErrorCode
	log.Printf("Last Error Code: 0x%08x", lastErrorCode)

	lastSuccessfulRunTime := utils.TimeFromFILETIME(int64(dynamicInfo.LastSuccessfulRunTime))
	log.Printf("Last Successful Run Time: %s", lastSuccessfulRunTime)
}

func init() {
	registerCommand(Command{
		Name:             "dynamicinfo",
		Args:             []string{"<task id>", "[-d|--dump]"},
		RequiredArgCount: 1,
		Func:             dynamicInfo,
	})
}
