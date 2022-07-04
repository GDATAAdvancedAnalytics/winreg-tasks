// SPDX-License-Identifier: MIT

package actions

import (
	"fmt"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/generated"
	"github.com/GDATAAdvancedAnalytics/winreg-tasks/utils"
	"github.com/google/uuid"
)

const ComHandlerPropertiesMagic PropertiesMagic = 0x7777

type ComHandlerProperties struct {
	Id string

	Clsid uuid.UUID
	Data  string
}

func NewComHandlerProperties(id string, gen *generated.Actions_ComHandlerProperties) (*ComHandlerProperties, error) {
	clsid, err := utils.UuidFromMemory(gen.Clsid)
	if err != nil {
		return nil, err
	}

	return &ComHandlerProperties{
		Id:    id,
		Clsid: clsid,
		Data:  gen.Data.Str,
	}, nil
}

func IsComHandlerProperties(properties Properties) bool {
	return properties.Magic() == ComHandlerPropertiesMagic
}

func (c ComHandlerProperties) Magic() PropertiesMagic {
	return ComHandlerPropertiesMagic
}

func (e ComHandlerProperties) Name() string {
	return "ComHandler"
}

func (c ComHandlerProperties) String() string {
	return fmt.Sprintf(
		`<ComHandler id="%s" clsid="%s" data="%s">`,
		c.Id, c.Clsid.String(), c.Data,
	)
}
