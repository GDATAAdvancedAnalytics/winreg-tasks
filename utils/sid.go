// Taken from https://github.com/bwmarrin/go-objectsid with small modifications
// to make the code more error resilient.

// SPDX-License-Identifier: BSD-2-Clause

package utils

import (
	"errors"
	"fmt"
)

var (
	ErrSidToShort          = errors.New("binary sid too short")
	ErrSidLengthUnexpected = errors.New("binary sid length unexpected")
)

type SID struct {
	RevisionLevel     uint8
	SubAuthorityCount int
	Authority         uint64
	SubAuthorities    []uint
}

func (sid SID) String() string {
	s := fmt.Sprintf("S-%d-%d", sid.RevisionLevel, sid.Authority)
	for _, v := range sid.SubAuthorities {
		s += fmt.Sprintf("-%d", v)
	}
	return s
}

func (sid SID) RID() uint {
	return sid.SubAuthorities[sid.SubAuthorityCount-1]
}

func SidFromBytes(b []byte) (*SID, error) {
	if len(b) < 2 {
		return nil, ErrSidToShort
	}

	sid := &SID{
		RevisionLevel:     b[0],
		SubAuthorityCount: int(b[1]),
	}

	if len(b) != int(1+1+6+sid.SubAuthorityCount*4) {
		return nil, ErrSidLengthUnexpected
	}

	for i := 2; i <= 7; i++ {
		sid.Authority = sid.Authority | uint64(b[i])<<(8*(5-(i-2)))
	}

	var offset = 8
	const size = 4

	for i := 0; i < sid.SubAuthorityCount; i++ {
		var subAuthority uint
		for k := 0; k < size; k++ {
			subAuthority = subAuthority | uint(b[offset+k])<<(8*k)
		}
		sid.SubAuthorities = append(sid.SubAuthorities, subAuthority)
		offset += size
	}

	return sid, nil
}
