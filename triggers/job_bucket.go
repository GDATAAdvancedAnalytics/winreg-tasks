// SPDX-License-Identifier: MIT

package triggers

import "github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"

type JobBucket struct {
	Flags            uint32
	Crc32            uint32
	PrincipalId      string
	DisplayName      string
	UserInfo         *UserInfo
	OptionalSettings *OptionalSettings
}

func NewJobBucket(gen *generated.Triggers_JobBucket) (*JobBucket, error) {
	userInfo, err := NewUserInfo(gen.UserInfo)
	if err != nil {
		return nil, err
	}

	optionalSettings, err := NewOptionalSettings(gen.OptionalSettings)
	if err != nil {
		return nil, err
	}

	principalId := ""
	if gen.PrincipalId != nil {
		principalId = gen.PrincipalId.String
	}

	displayName := ""
	if gen.DisplayName != nil {
		displayName = gen.DisplayName.String
	}

	return &JobBucket{
		Flags:            gen.Flags.Value,
		Crc32:            gen.Crc32.Value,
		PrincipalId:      principalId,
		DisplayName:      displayName,
		UserInfo:         userInfo,
		OptionalSettings: optionalSettings,
	}, nil
}
