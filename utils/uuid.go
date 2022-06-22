package utils

import (
	"github.com/google/uuid"
)

func UuidFromMemory(raw []byte) (uuid.UUID, error) {
	fixed := make([]byte, 16)

	// invert DWORD Data1
	fixed[0] = raw[3]
	fixed[1] = raw[2]
	fixed[2] = raw[1]
	fixed[3] = raw[0]

	// invert WORD Data2
	fixed[4] = raw[5]
	fixed[5] = raw[4]

	// invert WORD Data3
	fixed[6] = raw[7]
	fixed[7] = raw[6]

	// Data4 has correct order because it's a char[8]
	copy(fixed[8:16], raw[8:16])

	return uuid.FromBytes(fixed)
}
