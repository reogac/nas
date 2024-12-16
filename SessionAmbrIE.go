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
)

// 9.11.4.14 Session-AMBR [M LV 7]
type SessionAmbr struct {
	Bytes [6]byte
}

func NewSessionAmbr(ulU uint8, ulR uint16, dlU uint8, dlR uint16) *SessionAmbr {
	ambr := &SessionAmbr{}
	ambr.Bytes[0] = ulU
	binary.BigEndian.PutUint16(ambr.Bytes[1:], ulR)
	ambr.Bytes[3] = dlU
	binary.BigEndian.PutUint16(ambr.Bytes[4:], dlR)
	return ambr
}

func (ie *SessionAmbr) encode() (wire []byte, err error) {
	wire = ie.Bytes[:]
	return
}

func (ie *SessionAmbr) decode(wire []byte) (err error) {
	if len(wire) != 6 {
		err = ErrInvalidSize
		return
	}
	copy(ie.Bytes[:], wire)
	return
}
