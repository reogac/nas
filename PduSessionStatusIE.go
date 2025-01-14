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

// 9.11.3.44 PDU session status [O TLV 4-34]
type PduSessionStatus struct {
	bytes  [2]byte
	Spares []byte
}

func (ie *PduSessionStatus) encode() (wire []byte, err error) {
	wire = ie.bytes[:]
	if len(ie.Spares) > 0 {
		wire = append(wire, ie.Spares...)
	}
	return
}

func (ie *PduSessionStatus) decode(wire []byte) (err error) {
	if len(wire) < 2 {
		return ErrIncomplete
	}
	copy(ie.bytes[:], wire[0:2])
	if len(wire) > 2 {
		ie.Spares = make([]byte, len(wire)-2)
		copy(ie.Spares[:], wire[2:])
	}
	return
}

func (ie *PduSessionStatus) Set(flags [16]bool) {
	var pos uint8
	var row int
	for i, f := range flags {
		row = i / 8
		pos = uint8(i % 8)
		if f {
			ie.bytes[row] = setBit(ie.bytes[row], pos)
		} else {
			ie.bytes[row] = clearBit(ie.bytes[row], pos)
		}
	}
}

func (ie *PduSessionStatus) Get() (flags [16]bool) {
	var pos uint8
	var row int
	for i := 0; i < 16; i++ {
		row = i / 8
		pos = uint8(i % 8)
		flags[i] = getBit(ie.bytes[row], pos) == 1
	}
	return
}
