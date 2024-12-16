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
/** this file was generated at 2024-12-16 17:55:27.322699 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * DEREGISTRATION ACCEPT TO UE
 ******************************************************/
type DeregistrationAcceptToUe struct {
	MmHeader
}

func (msg *DeregistrationAcceptToUe) encode() (wire []byte, err error) {
	msg.msgType = DeregistrationAcceptToUeMsgType //set message type to DEREGISTRATION ACCEPT TO UE
	wire = msg.headerBytes()
	return
}
func (msg *DeregistrationAcceptToUe) decodeBody(wire []byte) (err error) {
	if len(wire) > 0 {
		err = ErrTail
	}
	return
}
