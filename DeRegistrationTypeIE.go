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

// 9.11.3.20 De-registration type [M V 1/2]
type DeRegistrationType struct {
	v uint8
}

func (ie *DeRegistrationType) encode() (wire []byte, err error) {
	wire = []byte{ie.v}
	return
}

func (ie *DeRegistrationType) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.v = wire[0]
	return
}

func (ie *DeRegistrationType) GetSwitchOff() bool {
	return getBit(ie.v, 3) == 1
}

func (ie *DeRegistrationType) SetSwitchOff(v bool) {
	if v {
		ie.v = setBit(ie.v, 3)
	} else {
		ie.v = clearBit(ie.v, 3)
	}
}

func (ie *DeRegistrationType) GetReregistration() bool {
	return getBit(ie.v, 2) == 1
}

func (ie *DeRegistrationType) SetReregistration(v bool) {
	if v {
		ie.v = setBit(ie.v, 2)
	} else {
		ie.v = clearBit(ie.v, 2)
	}
}

func (ie *DeRegistrationType) GetAccessType() uint8 {
	return ie.v & 0x03
}

func (ie *DeRegistrationType) SetAccessType(v uint8) {
	mask := uint8(0x03)
	ie.v = (ie.v & (^mask)) | (v & mask)
}
