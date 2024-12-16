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
/** this file was generated at 2024-12-16 17:55:27.328393 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * 5GMM STATUS
 ******************************************************/
type GmmStatus struct {
	MmHeader
	GmmCause uint8 //M: V [1]
}

func (msg *GmmStatus) encode() (wire []byte, err error) {
	// M: V[1]
	wire = append(wire, uint8(msg.GmmCause))

	msg.msgType = GmmStatusMsgType //set message type to 5GMM STATUS
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *GmmStatus) decodeBody(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = wire[offset]
	offset++

	return
}
