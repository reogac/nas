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

type uint8Encoder uint8

func newUint8Encoder(v uint8) *uint8Encoder {
	ret := uint8Encoder(v)
	return &ret
}
func (ie *uint8Encoder) encode() ([]byte, error) {
	return []byte{uint8(*ie)}, nil
}

type uint8Decoder uint8

func (ie *uint8Decoder) decode(wire []byte) error {
	if len(wire) != 1 {
		return ErrInvalidSize
	}
	*ie = uint8Decoder(wire[0])
	return nil
}

type uint16Encoder uint16

func newUint16Encoder(v uint16) *uint16Encoder {
	ret := uint16Encoder(v)
	return &ret
}
func (ie *uint16Encoder) encode() ([]byte, error) {
	var wire [2]byte
	binary.BigEndian.PutUint16(wire[:], uint16(*ie))
	return wire[:], nil
}

type uint16Decoder uint16

func (ie *uint16Decoder) decode(wire []byte) error {
	if len(wire) != 2 {
		return ErrInvalidSize
	}
	*ie = uint16Decoder(binary.BigEndian.Uint16(wire))
	return nil
}

type bytesEncoder []byte

func newBytesEncoder(dat []byte) *bytesEncoder {
	ret := bytesEncoder(dat)
	return &ret
}

func (ie *bytesEncoder) encode() ([]byte, error) {
	return []byte(*ie), nil
}

type bytesDecoder []byte

func (ie *bytesDecoder) decode(wire []byte) error {
	*ie = make([]byte, len(wire))
	copy(*ie, wire)
	return nil
}
