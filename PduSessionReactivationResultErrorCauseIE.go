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

import (
	"fmt"
)

// 9.11.3.43 PDU session reactivation result error cause [O TLV-E 5-515]
type PduSessionReactivationResultErrorCause struct {
	bytes []byte
}

func (ie *PduSessionReactivationResultErrorCause) encode() (wire []byte, err error) {
	wire = ie.bytes
	return
}

func (ie *PduSessionReactivationResultErrorCause) decode(wire []byte) (err error) {
	if len(wire)%2 == 1 {
		return fmt.Errorf("Expect even number of bytes")
	}
	ie.bytes = make([]byte, len(wire))
	copy(ie.bytes, wire)
	return
}

func (ie *PduSessionReactivationResultErrorCause) Set(idList, causeList []uint8) (err error) {
	if len(idList) != len(causeList) {
		return fmt.Errorf("Mismatched id list and cause list")
	}
	ie.bytes = make([]byte, 2*len(idList))
	for i, id := range idList {
		ie.bytes[2*i] = id
		ie.bytes[2*i+1] = causeList[i]
	}
	return
}
