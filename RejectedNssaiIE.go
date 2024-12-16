/*
* Copyright [2024] [Quang Tung Thai <tqtung@etri.re.kr>]
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
package nas

import (
	"encoding/hex"
	"fmt"
)

type RejectedSnssai struct {
	Cause uint8
	Sst   uint8
	Sd    []byte //zero or three bytes
}

func (slice *RejectedSnssai) SetSd(sd string) error {
	if buf, err := hex.DecodeString(sd); err != nil {
		return fmt.Errorf("decode Sd from hex string fails: %+v", err)
	} else if len(buf) != 3 {
		return fmt.Errorf("Sd must be 3 bytes")
	} else {
		slice.Sd = buf
	}
	return nil
}

// 9.11.3.46 Rejected NSSAI [O TLV 4-42]
type RejectedNssai struct {
	List []RejectedSnssai
}

func (ie *RejectedNssai) encode() (wire []byte, err error) {
	for i, slice := range ie.List {
		if i == 7 { //no more than 8 rejected slice
			break
		}
		header := slice.Cause & 0x0f
		bufSize := uint8(2)
		if len(slice.Sd) == 3 {
			bufSize = uint8(5)
		}
		header |= bufSize << 4
		wire = append(wire, []byte{header, slice.Sst}...)
		if len(slice.Sd) == 3 {
			wire = append(wire, slice.Sd...)
		}
	}

	return
}

func (ie *RejectedNssai) decode(wire []byte) (err error) {
	wireLen := len(wire)
	offset := 0
	for offset < wireLen {
		sliceSize := wire[offset] >> 4
		if sliceSize != 2 && sliceSize != 5 {
			return fmt.Errorf("Unexpected rejected Snssai size (%d), expected [2 or 5]", sliceSize)
		}
		if offset+int(sliceSize) > wireLen {
			return ErrIncomplete
		}
		slice := new(RejectedSnssai)
		slice.Cause = wire[offset] & 0x0f
		slice.Sst = wire[offset+1]
		if sliceSize == 5 {
			slice.Sd = make([]byte, 3)
			copy(slice.Sd, wire[offset+2:offset+5])
		}
		ie.List = append(ie.List, *slice)
		offset += int(sliceSize)

		if len(ie.List) >= 8 { //no more than 8 rejected slices
			break
		}
	}

	return
}
