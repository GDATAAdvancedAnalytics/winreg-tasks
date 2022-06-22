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

func hexdump(data []byte) string {
	str := ""

	for i, c := range data {
		if i > 0 && i%16 == 0 {
			str = str[:len(str)-1] + "\n"
		}

		str += fmt.Sprintf("%02x ", c)
	}

	return strings.TrimSuffix(strings.TrimSuffix(str, "\n"), " ")
}
