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
/** this file was generated at 2024-12-16 17:55:27.327641 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * AUTHENTICATION RESPONSE
 ******************************************************/
type AuthenticationResponse struct {
	MmHeader
	AuthenticationResponseParameter []byte //O: TLV [2D][18]
	EapMessage                      []byte //O: TLV-E [78][7-1503]
}

func (msg *AuthenticationResponse) encode() (wire []byte, err error) {
	var buf []byte
	// O: TLV[18]
	if len(msg.AuthenticationResponseParameter) > 0 {
		tmp := newBytesEncoder(msg.AuthenticationResponseParameter)
		if buf, err = encodeLV(false, uint16(16), uint16(16), tmp); err != nil {
			err = nasError("encoding AuthenticationResponseParameter [O TLV 18]", err)
			return
		}
		wire = append(append(wire, 0x2D), buf...)
	}

	// O: TLV-E[7-1503]
	if len(msg.EapMessage) > 0 {
		tmp := newBytesEncoder(msg.EapMessage)
		if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	msg.msgType = AuthenticationResponseMsgType //set message type to AUTHENTICATION RESPONSE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationResponse) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x2D: //O: TLV[18]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(16), uint16(16), v); err != nil {
				err = nasError("decoding AuthenticationResponseParameter [O TLV 18]", err)
				return
			}
			offset += consumed
			msg.AuthenticationResponseParameter = []byte(*v)
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
