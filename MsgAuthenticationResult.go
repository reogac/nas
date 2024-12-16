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
/** this file was generated at 2024-12-16 17:55:27.327841 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * AUTHENTICATION RESULT
 ******************************************************/
type AuthenticationResult struct {
	MmHeader
	Ngksi      KeySetIdentifier //M: V [1/2]
	EapMessage []byte           //M: LV-E [6-1502]
	Abba       []byte           //O: TLV [38][4-n]
}

func (msg *AuthenticationResult) encode() (wire []byte, err error) {
	var buf []byte
	// M: V[1/2]
	if buf, err = msg.Ngksi.encode(); err != nil {
		err = nasError("encoding Ngksi [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding Ngksi [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// M: LV-E[6-1502]
	wire = append(wire, v)

	tmp := newBytesEncoder(msg.EapMessage)
	if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
		err = nasError("encoding EapMessage [M LV-E 6-1502]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TLV[4-n]
	if len(msg.Abba) > 0 {
		tmp := newBytesEncoder(msg.Abba)
		if buf, err = encodeLV(false, uint16(2), uint16(0), tmp); err != nil {
			err = nasError("encoding Abba [O TLV 4-n]", err)
			return
		}
		wire = append(append(wire, 0x38), buf...)
	}

	msg.msgType = AuthenticationResultMsgType //set message type to AUTHENTICATION RESULT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationResult) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding Ngksi [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.Ngksi.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	// M LV-E[6-1502]
	offset++

	v := new(bytesDecoder)
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
		err = nasError("decoding EapMessage [M LV-E 6-1502]", err)
		return
	}
	offset += consumed
	msg.EapMessage = []byte(*v)
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x38: //O: TLV[4-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(0), v); err != nil {
				err = nasError("decoding Abba [O TLV 4-n]", err)
				return
			}
			offset += consumed
			msg.Abba = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
