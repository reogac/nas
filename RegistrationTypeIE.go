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

// 9.11.3.7 5GS registration type [M V 1/2]
type RegistrationType struct {
	value uint8
}

func NewRegistrationType(forValue bool, typeValue uint8) (v RegistrationType) {
	if forValue {
		v.value = 1<<3 + typeValue&0x07
	} else {
		v.value = typeValue & 0x07
	}
	return
}

func (ie *RegistrationType) encode() (wire []byte, err error) {
	wire = []byte{ie.value & 0x0f}
	return
}

func (ie *RegistrationType) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.value = wire[0] & 0x0f
	return
}

func (ie *RegistrationType) GetFor() bool {
	return (ie.value&0x0f)>>3 == 1
}

func (ie *RegistrationType) GetType() uint8 {
	return ie.value & 0x07
}
