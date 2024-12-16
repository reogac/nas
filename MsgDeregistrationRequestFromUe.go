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
/** this file was generated at 2024-12-16 17:55:27.325677 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * DEREGISTRATION REQUEST FROM UE
 ******************************************************/
type DeregistrationRequestFromUe struct {
	MmHeader
	DeRegistrationType DeRegistrationType //M: V [1/2]
	Ngksi              KeySetIdentifier   //M: V [1/2]
	MobileIdentity     MobileIdentity     //M: LV-E [6-n]
}

func (msg *DeregistrationRequestFromUe) encode() (wire []byte, err error) {
	var buf []byte
	// M: V[1/2]
	if buf, err = msg.DeRegistrationType.encode(); err != nil {
		err = nasError("encoding DeRegistrationType [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding DeRegistrationType [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// M: V[1/2]
	if buf, err = msg.Ngksi.encode(); err != nil {
		err = nasError("encoding Ngksi [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding Ngksi [M V 1/2]", ErrInvalidSize)
		return
	}
	v |= (buf[0] & 0x0f) << 4 //fill lefthalf
	wire = append(wire, v)

	// M: LV-E[6-n]
	if buf, err = encodeLV(true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("encoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	wire = append(wire, buf...)

	msg.msgType = DeregistrationRequestFromUeMsgType //set message type to DEREGISTRATION REQUEST FROM UE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *DeregistrationRequestFromUe) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding DeRegistrationType [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.DeRegistrationType.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding DeRegistrationType [M V 1/2]", err)
		return
	}
	// M V[1/2]
	if err = msg.Ngksi.decode([]byte{(0xf0 & wire[offset]) >> 4 /*lefthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	offset++

	// M LV-E[6-n]
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("decoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	offset += consumed
	return
}
