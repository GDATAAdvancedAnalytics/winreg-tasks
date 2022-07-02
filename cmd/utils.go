// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"strings"
)

func toGolangHexBytes(data []byte) string {
	str := ""

	for _, c := range data {
		str += fmt.Sprintf(`0x%02x, `, c)
	}

	return strings.TrimSuffix(str, ` `)
}
