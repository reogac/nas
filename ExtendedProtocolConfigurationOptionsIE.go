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
	"encoding/binary"
	"math"
)

// 9.11.4.6 Extended protocol configuration options [O TLV-E 4-65538]
type ExtendedProtocolConfigurationOptions struct {
	units []PcoUnit
}

func (ie *ExtendedProtocolConfigurationOptions) Units() []PcoUnit {
	return ie.units
}

func (ie *ExtendedProtocolConfigurationOptions) AddUnit(unit PcoUnit) {
	ie.units = append(ie.units, unit)
}

func (ie *ExtendedProtocolConfigurationOptions) encode() (wire []byte, err error) {
	var extension uint8 = 1
	var spare uint8 = 0
	var configurationProtocol uint8 = 0
	var buf []byte

	wire = []byte{(extension << 7) | (spare << 6) | (configurationProtocol)}
	for _, u := range ie.units {
		if buf, err = u.toBytes(); err != nil {
			return
		}
		wire = append(wire, buf...)
	}

	return
}

// inserted by shchoi
const (
	ReadingID = iota
	ReadingLength
)

func (ie *ExtendedProtocolConfigurationOptions) decode(wire []byte) error {
	if len(wire) < 1 {
		return ErrIncomplete
	}

	offset := 1 // 첫 번째 바이트는 스킵 (PCO 헤더 등)
	wireLen := len(wire)
	readingState := ReadingID
	var curUnit PcoUnit

	for offset < wireLen {
		switch readingState {
		case ReadingID:
			if offset+2 > wireLen {
				return ErrIncomplete
			}
			curUnit = PcoUnit{} // 새 유닛 초기화
			curUnit.Id = binary.BigEndian.Uint16(wire[offset : offset+2])
			offset += 2
			readingState = ReadingLength

		case ReadingLength:
			if offset+1 > wireLen {
				return ErrIncomplete
			}
			unitLength := int(wire[offset])
			offset += 1

			if unitLength == 0 {
				curUnit.Content = []byte{}
				ie.units = append(ie.units, curUnit)
				readingState = ReadingID
			} else {
				if offset+unitLength > wireLen {
					return ErrIncomplete
				}
				curUnit.Content = make([]byte, unitLength)
				copy(curUnit.Content, wire[offset:offset+unitLength])
				offset += unitLength
				ie.units = append(ie.units, curUnit)
				readingState = ReadingID
			}
		}
	}

	return nil
}

type PcoUnit struct {
	Id      uint16
	Content []byte
}

func (u *PcoUnit) toBytes() (buf []byte, err error) {
	if len(u.Content) > math.MaxUint8 {
		err = ErrInvalidSize
		return
	}
	buf = []byte{0, 0, uint8(len(u.Content))}
	binary.BigEndian.PutUint16(buf, u.Id)
	buf = append(buf, u.Content...)
	return
}
