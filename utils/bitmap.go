// SPDX-License-Identifier: MIT

package utils

import "fmt"

func BitmapToString(bitmap uint64) string {
	ret := ""

	for i := 1; bitmap != 0; i++ {
		if (bitmap & 1) == 1 {
			ret += fmt.Sprintf("%d,", i)
		}
		bitmap >>= 1
	}

	return ret
}
