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
/** this file was generated at 2024-12-16 17:55:27.327978 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * IDENTITY RESPONSE
 ******************************************************/
type IdentityResponse struct {
	MmHeader
	MobileIdentity MobileIdentity //M: LV-E [3-n]
}

func (msg *IdentityResponse) encode() (wire []byte, err error) {
	var buf []byte
	// M: LV-E[3-n]
	if buf, err = encodeLV(true, uint16(1), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("encoding MobileIdentity [M LV-E 3-n]", err)
		return
	}
	wire = append(wire, buf...)

	msg.msgType = IdentityResponseMsgType //set message type to IDENTITY RESPONSE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *IdentityResponse) decodeBody(wire []byte) (err error) {
	offset := 0
	consumed := 0
	// M LV-E[3-n]
	if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("decoding MobileIdentity [M LV-E 3-n]", err)
		return
	}
	offset += consumed
	return
}
