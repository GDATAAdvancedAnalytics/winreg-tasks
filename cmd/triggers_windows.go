// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/triggers"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

func handleTriggers(args ...string) {
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
		hex := utils.Hexdump(triggersRaw, 16)
		fmt.Println(hex)
	}

	triggers, err := triggers.FromBytes(triggersRaw, time.Local)
	if err != nil {
		log.Fatalf("cannot parse triggers: %v", err)
	}

	log.Println("Header:")
	log.Printf("\tVersion: %d", triggers.Header.Version)
	log.Printf("\tStartBoundary: %s", triggers.Header.StartBoundary.String())
	log.Printf("\tEndBoundary: %s", triggers.Header.EndBoundary.String())

	log.Println("JobBucket:")
	log.Printf("\tFlags: %08x", triggers.JobBucket.Flags)
	log.Printf("\tCRC32: %08x", triggers.JobBucket.Crc32)
	log.Printf("\tPrincipal ID: %s", triggers.JobBucket.PrincipalId)
	log.Printf("\tDisplay Name: %s", triggers.JobBucket.DisplayName)
	log.Printf("\tUser: %s", triggers.JobBucket.UserInfo.UserToString())

	log.Println("Triggers:")
	if len(triggers.Triggers) == 0 {
		log.Println("\t<no triggers>")
		return
	}

	for _, trigger := range triggers.Triggers {
		log.Println("\t" + trigger.String())
	}
}

func init() {
	registerCommand(Command{
		Name:             "triggers",
		RequiredArgCount: 1,
		Args:             []string{"<task id>", "[-d|--dump]"},
		Func:             handleTriggers,
	})
}
