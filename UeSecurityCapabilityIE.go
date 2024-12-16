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

// 9.11.3.54 UE security capability [O TLV 4-10]
type UeSecurityCapability struct {
	bytes []byte
}

func (ie *UeSecurityCapability) Bytes() []byte {
	if wire, err := ie.encode(); err == nil {
		return wire
	}
	return []byte{}
}
func (ie *UeSecurityCapability) encode() (wire []byte, err error) {
	l := len(ie.bytes)
	if l < 2 || l == 3 {
		err = ErrInvalidSize
		return
	}

	if l == 4 && ie.bytes[2] == 0 && ie.bytes[3] == 0 {
		//discard last two bytes
		wire = ie.bytes[0:2]
	} else {
		wire = ie.bytes
	}
	return
}

func (ie *UeSecurityCapability) decode(wire []byte) (err error) {
	l := len(wire)
	//at least 2 bytes, can't be 3 bytes
	if l < 2 || l == 3 {
		err = ErrIncomplete
		return
	}
	ie.bytes = make([]byte, l)
	copy(ie.bytes, wire)
	return
}

func (ie *UeSecurityCapability) getBit(row uint8, pos uint8) bool {
	if pos >= 8 {
		return false
	}

	if row > 4 {
		return false
	}

	if int(row) > len(ie.bytes) {
		return false
	}

	return getBit(ie.bytes[row], (7-pos)) == 1
}

func (ie *UeSecurityCapability) setBit(row uint8, pos uint8, v bool) {
	if row >= 4 {
		return //ignore
	}

	if len(ie.bytes) == 0 {
		ie.bytes = make([]byte, 4) //always make 4 bytes, discard the last two if needed in encoding
	}

	if v {
		ie.bytes[row] |= 1 << (7 - pos) //set
	} else {
		ie.bytes[row] &= ^(1 << (7 - pos)) //unset
	}
}

func (ie *UeSecurityCapability) GetExtra() []byte {
	if len(ie.bytes) > 4 {
		return ie.bytes[4:]
	}
	return []byte{}
}

func (ie *UeSecurityCapability) SetExtra(extra []byte) {
	buf := make([]byte, 4+len(extra))
	copy(buf[4:], extra)
	if len(ie.bytes) == 0 {
		ie.bytes = buf
	} else {
		copy(buf[0:4], ie.bytes)
		ie.bytes = buf
	}
}

// index must be 0-7
func (ie *UeSecurityCapability) GetEA(index uint8) bool {
	return ie.getBit(0, index)
}

// index must be 0-7
func (ie *UeSecurityCapability) SetEA(index uint8, v bool) {
	ie.setBit(0, index, v)
}

// index must be 0-7
func (ie *UeSecurityCapability) GetIA(index uint8) bool {
	return ie.getBit(1, index)
}

// index must be 0-7
func (ie *UeSecurityCapability) SetIA(index uint8, v bool) {
	ie.setBit(1, index, v)

}

// index must be 0-7
func (ie *UeSecurityCapability) GetEEA(index uint8) bool {
	return ie.getBit(2, index)
}

// index must be 0-7
func (ie *UeSecurityCapability) SetEEA(index uint8, v bool) {
	ie.setBit(2, index, v)
}

// index must be 0-7
func (ie *UeSecurityCapability) GetEIA(index uint8) bool {
	return ie.getBit(3, index)
}

// index must be 0-7
func (ie *UeSecurityCapability) SetEIA(index uint8, v bool) {
	ie.setBit(3, index, v)
}
