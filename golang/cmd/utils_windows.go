// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

const (
	taskKeyBase       = `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Schedule\TaskCache\`
	secondsUntilEpoch = 11_644_473_600
)

func timeFromFILETIME(filetime int64) time.Time {
	epoch := filetime/10_000_000 - secondsUntilEpoch
	return time.Unix(epoch, 0)
}

func getUUIDFromTaskPath(path string) (string, error) {
	key, err := openKey(`Tree\` + path)
	if err != nil {
		return "", err
	}

	val, _, err := key.GetStringValue("Id")
	if err != nil {
		return "", err
	}

	return val, nil
}

func openKey(subKey string) (registry.Key, error) {
	return registry.OpenKey(registry.LOCAL_MACHINE, taskKeyBase+subKey, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
}

func openTaskKey(keyId string) registry.Key {
	var err error

	switch {
	case strings.HasPrefix(keyId, `\`):
		keyId, err = getUUIDFromTaskPath(keyId)
		if err != nil {
			log.Printf("cannot convert task path to uuid: %v\n", err)
			return 0
		}
		fallthrough

	case strings.HasPrefix(keyId, `{`):
		key, err := openKey(`Tasks\` + keyId)
		if err != nil {
			log.Printf("cannot open key %s: %v\n", keyId, err)
			return 0
		}
		return key

	default:
		log.Printf("task id unknown. must start with \\ or {")
		return 0
	}
}

func binarySidToString(raw []byte) string {
	sid := (*syscall.SID)(unsafe.Pointer(&raw[0]))

	strSid, err := sid.String()
	if err != nil {
		return fmt.Sprintf("<error: %s>", err)
	}
	return strSid
}
