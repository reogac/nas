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
	"fmt"
)

func getIei(v uint8) uint8 {
	if v >= 0x80 {
		return v & 0xf0 >> 4
	}
	return v
}

func checkLengthBounds(v, l, u uint16) bool {
	if v < l { //check lower bound
		log.Errorf("fail check %d<%d\n", v, l)
		return false
	}
	//NOTE: must check a corner case where l=u=0
	//u = 0 means no upper bound
	if u != 0 && u >= l { //check if upper bound is valid
		if v > u {
			log.Errorf("fail check %d>%d\n", v, u)
			return false
		}
	}
	return true
}

// decode LV, LV-E
func decodeLV(buf []byte, extend bool, l, u uint16, v Decoder) (consumed int, err error) {
	bufSize := len(buf)
	lenSize := 1 //TLV
	if extend {
		lenSize = 2 //TLV-E
	}
	if lenSize > bufSize {
		err = ErrIncomplete
		return
	}
	lenValue := uint16(buf[0])
	if extend {
		lenValue = binary.BigEndian.Uint16(buf[0:2])
	}

	log.Tracef("DecodeLV %d byte\n", lenValue)

	if !checkLengthBounds(lenValue, l, u) {
		err = ErrInvalidSize
		return
	}

	consumed = int(uint16(lenSize) + lenValue)
	if consumed > bufSize {
		consumed = bufSize
		err = ErrIncomplete
		return
	}
	err = v.decode(buf[lenSize:consumed])
	return
}

// encode LV, LV-E
func encodeLV(extend bool, l, u uint16, v Encoder) (wire []byte, err error) {

	var buf []byte
	if buf, err = v.encode(); err != nil {
		return
	}
	vLen := uint16(len(buf))

	if !checkLengthBounds(vLen, l, u) {
		err = ErrInvalidSize
		return
	}

	log.Tracef("EncodeLV %d byte\n", vLen)
	if extend {
		prefix := [2]byte{0, 0}
		binary.BigEndian.PutUint16(prefix[:], vLen)
		wire = append(prefix[:], buf...)
	} else {
		wire = append([]byte{uint8(vLen)}, buf...)
	}
	return
}

// lastBit 0[rightmost][leftmost]:7; numBits:0:8
func bitMask(lastBit, numBits uint8) uint8 {
	return (uint8(1)<<numBits - uint8(1)) << lastBit
}

// pos must be 0:7
func getBit(v uint8, pos uint8) uint8 {
	return (v & (uint8(1) << pos)) >> pos
}

// pos must be 0:7
func setBit(v uint8, pos uint8) uint8 {
	return (v | (uint8(1) << pos))
}

// pos must be 0:7
func clearBit(v uint8, pos uint8) uint8 {
	mask := ^(uint8(1) << pos)
	return v & mask
}

// decimal string to byte array
func decimalBytes(s string) (buf []byte, err error) {
	tmp := make([]byte, len(s))
	for i, c := range s {
		d := uint8(c - '0')
		if d >= 0 && d <= 9 {
			tmp[i] = d
		} else {
			err = fmt.Errorf("\"%c\" is not decimal number", c)
			return
		}
	}
	buf = tmp
	return
}

// skipLV returns the offset just past the length and value of an IE whose
// length octets start at offset: one for a TLV, two for a TLV-E. It reads only
// the length, so it steps over a value its own decoder refused.
func skipLV(wire []byte, offset int, extend bool) (int, error) {
	lenSize := 1
	if extend {
		lenSize = 2
	}
	if offset+lenSize > len(wire) {
		return offset, ErrIncomplete
	}
	l := int(wire[offset])
	if extend {
		l = int(binary.BigEndian.Uint16(wire[offset : offset+2]))
	}
	end := offset + lenSize + l
	if end > len(wire) {
		return offset, ErrIncomplete
	}
	return end, nil
}

// skipUnknownIe returns the offset just past an IE whose IEI the message does
// not define, read by the format its IEI implies (TS 24.007 11.2.4): bit 8 set
// is a type 1 or 2 IE of one octet, 0111 in bits 8 to 5 a TLV-E, and any other
// a TLV. An IEI with 0000 in bits 8 to 5 is encoded as comprehension required,
// and refuses the message.
func skipUnknownIe(wire []byte, offset int) (int, error) {
	iei := wire[offset]
	switch {
	case iei&0x80 != 0:
		return offset + 1, nil
	case iei&0xf0 == 0:
		return offset, ErrUnknownIei
	default:
		return skipLV(wire, offset+1, iei&0xf0 == 0x70)
	}
}
