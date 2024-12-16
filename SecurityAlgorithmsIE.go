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

// 9.11.3.34 security algorithms [M V 1]
type SecurityAlgorithms struct {
	value uint8
}

func NewSecurityAlgorithms(intAlg, encAlg uint8) SecurityAlgorithms {
	v := intAlg & 0x0f
	v += (encAlg & 0x0f) << 4
	return SecurityAlgorithms{
		value: v,
	}
}

func (ie *SecurityAlgorithms) encode() (wire []byte, err error) {
	wire = []byte{ie.value}
	return
}

func (ie *SecurityAlgorithms) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.value = wire[0]
	return
}

func (ie *SecurityAlgorithms) IntAlg() uint8 {
	return ie.value & 0x0f
}

func (ie *SecurityAlgorithms) EncAlg() uint8 {
	return ie.value & 0xf0 >> 4
}
