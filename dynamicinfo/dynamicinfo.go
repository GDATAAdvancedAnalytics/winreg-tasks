// SPDX-License-Identifier: MIT

package dynamicinfo

import (
	"bytes"
	"fmt"
	"time"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

type DynamicInfo struct {
	Magic                 uint32
	CreationTime          time.Time
	LastRunTime           time.Time
	TaskState             uint32
	LastErrorCode         uint32
	LastSuccessfulRunTime time.Time
}

func FromBytes(raw []byte) (*DynamicInfo, error) {
	stream := kaitai.NewStream(bytes.NewReader(raw))
	generatedDynamicInfo := generated.NewDynamicInfo()

	if err := generatedDynamicInfo.Read(stream, nil, generatedDynamicInfo); err != nil {
		return nil, err
	}

	if eof, err := stream.EOF(); !eof {
		return nil, fmt.Errorf("did not parse all data")
	} else if err != nil {
		return nil, fmt.Errorf("error trying to eof-check (%v)", err)
	}

	magic := uint32(generatedDynamicInfo.Magic[3])<<24 |
		uint32(generatedDynamicInfo.Magic[2])<<16 |
		uint32(generatedDynamicInfo.Magic[1])<<8 |
		uint32(generatedDynamicInfo.Magic[0])

	return &DynamicInfo{
		Magic:                 magic,
		CreationTime:          utils.TimeFromFILETIME(generatedDynamicInfo.CreationTime),
		LastRunTime:           utils.TimeFromFILETIME(generatedDynamicInfo.LastRunTime),
		TaskState:             generatedDynamicInfo.TaskState,
		LastErrorCode:         generatedDynamicInfo.LastErrorCode,
		LastSuccessfulRunTime: utils.TimeFromFILETIME(generatedDynamicInfo.LastSuccessfulRunTime),
	}, nil
}

func (d DynamicInfo) String() string {
	return fmt.Sprintf(
		`<DynamicInfo creation_time="%s" last_run_time="%s" last_error_code=0x%08x>`,
		d.CreationTime.String(), d.LastRunTime.String(), d.LastErrorCode,
	)
}
