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

// 9.11.4.10 PDU address [O TLV 11]
type PduAddress struct {
	bytes []byte
}

func NewPduAddress(t uint8, content []byte) *PduAddress {
	v := &PduAddress{
		bytes: make([]byte, len(content)+1),
	}
	v.bytes[0] = t
	copy(v.bytes[1:], content)
	return v
}

func (ie *PduAddress) AddressType() uint8 {
	return ie.bytes[0] & 0x0f
}

func (ie *PduAddress) Content() []byte {
	return ie.bytes[1:]
}

func (ie *PduAddress) encode() (wire []byte, err error) {
	wire = make([]byte, len(ie.bytes))
	copy(wire, ie.bytes)
	return
}

func (ie *PduAddress) decode(wire []byte) (err error) {
	if len(wire) < 1 {
		err = ErrInvalidSize
		return
	}
	ie.bytes = make([]byte, len(wire))
	copy(ie.bytes, wire)
	return
}
