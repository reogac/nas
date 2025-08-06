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

// 9.11.3.12 Additional 5G security information [O TLV 3]
type AdditionalSecurityInformation struct {
	value uint8
}

func NewAdditionalSecurityInformation(retransmission, hoDerivation bool) *AdditionalSecurityInformation {
	ie := &AdditionalSecurityInformation{}
	ie.SetRetransmission(retransmission)
	ie.SetHoDerivation(hoDerivation)
	return ie
}

func (ie *AdditionalSecurityInformation) encode() (wire []byte, err error) {
	wire = []byte{ie.value}
	return
}

func (ie *AdditionalSecurityInformation) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.value = wire[0]
	return
}

func (ie *AdditionalSecurityInformation) GetRetransmission() bool {
	return getBit(ie.value, 1) == 1
}

func (ie *AdditionalSecurityInformation) SetRetransmission(flag bool) {
	if flag {
		ie.value = setBit(ie.value, 1)
	} else {
		ie.value = clearBit(ie.value, 1)
	}
}
func (ie *AdditionalSecurityInformation) GetHoDerivation() bool {
	return getBit(ie.value, 0) == 1
}

func (ie *AdditionalSecurityInformation) SetHoDerivation(flag bool) {
	if flag {
		ie.value = setBit(ie.value, 0)
	} else {
		ie.value = clearBit(ie.value, 0)
	}

}
