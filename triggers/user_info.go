// SPDX-License-Identifier: MIT

package triggers

import (
	"fmt"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
)

type SidType int

const (
	SidTypeUser           SidType = 1
	SidTypeGroup          SidType = 2
	SidTypeDomain         SidType = 3
	SidTypeAlias          SidType = 4
	SidTypeWellKnownGroup SidType = 5
	SidTypeDeletedAccount SidType = 6
	SidTypeInvalid        SidType = 7
	SidTypeUnknown        SidType = 8
	SidTypeComputer       SidType = 9
	SidTypeLabel          SidType = 10
	SidTypeLogonSession   SidType = 11
)

type UserInfo struct {
	HasUser  bool
	HasSid   bool
	SidType  SidType
	Sid      []byte
	Username string
}

func NewUserInfo(gen *generated.UserInfo) (*UserInfo, error) {
	if gen.SkipUser.Value != 0 {
		return &UserInfo{HasUser: false}, nil
	}

	userInfo := &UserInfo{
		HasUser:  true,
		Username: gen.Username.String,
	}

	if gen.SkipSid.Value == 0 {
		userInfo.HasSid = true
		userInfo.Sid = gen.Sid.Data[:]
		userInfo.SidType = SidType(gen.SidType.Value)
	}

	return userInfo, nil
}

func (u UserInfo) UserToString() string {
	user := "<unset>"
	if u.HasUser {
		if u.Username != "" {
			user = u.Username
		} else if u.HasSid {
			if sid, err := utils.SidFromBytes(u.Sid); err == nil {
				user = sid.String()
			} else {
				user = fmt.Sprintf("<error converting sid: %s>", err)
			}
		}
	}
	return user
}
