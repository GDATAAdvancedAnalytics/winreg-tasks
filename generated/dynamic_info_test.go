// SPDX-License-Identifier: MIT

package generated_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func TestDynamicInfo_Read_WithErrorCode(t *testing.T) {
	blob := []byte{
		0x03, 0x00, 0x00, 0x00, 0xe9, 0x79, 0x2d, 0xf1,
		0x31, 0x1c, 0xd8, 0x01, 0xb4, 0x4a, 0x8b, 0x8d,
		0x3e, 0x1c, 0xd8, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x02, 0x00, 0x07, 0x80, 0xe4, 0x70, 0xd5, 0x67,
		0x34, 0x1c, 0xd8, 0x01,
	}

	tt := generated.NewDynamicInfo()
	if err := tt.Read(kaitai.NewStream(bytes.NewReader(blob)), tt, tt); err != nil {
		t.Error(fmt.Sprintf("Read (%v)", err))
	}

	if tt.CreationTime != 0x01d81c31f12d79e9 {
		t.Error("CreationTime")
	}

	if tt.LastRunTime != 0x01d81c3e8d8b4ab4 {
		t.Error("LastRunTime")
	}

	if tt.TaskState != 0x00000000 {
		t.Error("TaskState")
	}

	if tt.LastErrorCode != 0x80070002 {
		t.Error("LastErrorCode")
	}

	if tt.LastSuccessfulRunTime != 0x01d81c3467d570e4 {
		t.Error("LastSuccessfulRunTime")
	}
}

func TestDynamicInfo_Read_WithSuccess(t *testing.T) {
	blob := []byte{
		0x03, 0x00, 0x00, 0x00, 0xe9, 0x79, 0x2d, 0xf1,
		0x31, 0x1c, 0xd8, 0x01, 0x96, 0xc1, 0x54, 0xa5,
		0x3e, 0x1c, 0xd8, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xf6, 0x79, 0xa3, 0xa5,
		0x3e, 0x1c, 0xd8, 0x01,
	}

	tt := generated.NewDynamicInfo()
	if err := tt.Read(kaitai.NewStream(bytes.NewReader(blob)), tt, tt); err != nil {
		t.Error(fmt.Sprintf("Read (%v)", err))
	}

	if tt.CreationTime != 0x01d81c31f12d79e9 {
		t.Error("CreationTime")
	}

	if tt.LastRunTime != 0x01d81c3ea554c196 {
		t.Error("LastRunTime")
	}

	if tt.TaskState != 0x00000000 {
		t.Error("TaskState")
	}

	if tt.LastErrorCode != 0x00000000 {
		t.Error("LastErrorCode")
	}

	if tt.LastSuccessfulRunTime != 0x01d81c3ea5a379f6 {
		t.Error("LastSuccessfulRunTime")
	}
}
