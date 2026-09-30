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
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

var msgStr = "7e0042010177000b0102030405060708090a0b4a2d010100000000000000000000000000000000000000000000000000000000000000000000000000000000000000540a010101010101010101011502010131020101210201015002010126020101720002010179000a01010101010101010101b090270a010101010101010101015e01015d01011601013404010101017a000401010101730014010101010101010101010101010101010101010178000401010101a07600020101510101"

func TestRegistrationAccept(t *testing.T) {
	var nasPdu, newNasPdu []byte
	var err error
	var nasMsg NasMessage
	if nasPdu, err = hex.DecodeString(msgStr); err != nil {
		t.Errorf("Fail to  get message pdu: %+v", err)
		return
	}

	fmt.Printf("Pdu: %v\n", nasPdu)

	if nasMsg, err = Decode(nil, nasPdu, true); err != nil {
		t.Errorf("Fail to decode nas message: %+v", err)
	} else if nasMsg.Gmm == nil {
		t.Errorf("Empty Gmm")
	} else {
		if newNasPdu, err = EncodeMm(nil, nasMsg.Gmm.RegistrationAccept, true); err != nil {
			t.Errorf("Fail to encode nas message: %+v", err)
			return
		}
		if bytes.Compare(nasPdu, newNasPdu) != 0 {
			t.Errorf("Not equal")
			return
		}
	}
}
