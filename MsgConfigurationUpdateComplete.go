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
/** this file was generated at 2024-12-16 17:55:27.322720 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * CONFIGURATION UPDATE COMPLETE
 ******************************************************/
type ConfigurationUpdateComplete struct {
	MmHeader
}

func (msg *ConfigurationUpdateComplete) encode() (wire []byte, err error) {
	msg.msgType = ConfigurationUpdateCompleteMsgType //set message type to CONFIGURATION UPDATE COMPLETE
	wire = msg.headerBytes()
	return
}
func (msg *ConfigurationUpdateComplete) decodeBody(wire []byte) (err error) {
	if len(wire) > 0 {
		err = ErrTail
	}
	return
}
