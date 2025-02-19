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
/** this file was generated at 2024-12-16 17:55:27.324844 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * REGISTRATION COMPLETE
 ******************************************************/
type RegistrationComplete struct {
	MmHeader
	SorTransparentContainer []byte //O: TLV-E [73][20]
}

func (msg *RegistrationComplete) encode() (wire []byte, err error) {
	var buf []byte
	// O: TLV-E[20]
	if len(msg.SorTransparentContainer) > 0 {
		tmp := newBytesEncoder(msg.SorTransparentContainer)
		if buf, err = encodeLV(true, uint16(17), uint16(17), tmp); err != nil {
			err = nasError("encoding SorTransparentContainer [O TLV-E 20]", err)
			return
		}
		wire = append(append(wire, 0x73), buf...)
	}

	msg.msgType = RegistrationCompleteMsgType //set message type to REGISTRATION COMPLETE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *RegistrationComplete) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x73: //O: TLV-E[20]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(17), uint16(17), v); err != nil {
				err = nasError("decoding SorTransparentContainer [O TLV-E 20]", err)
				return
			}
			offset += consumed
			msg.SorTransparentContainer = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
