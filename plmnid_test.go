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
	//	"fmt"
	"testing"
)

func Test_PlmnId(t *testing.T) {
	var plmnId PlmnId
	var err error
	mcc := "108"
	mnc := "90"
	if err = plmnId.Set(mcc, mnc); err != nil {
		t.Errorf("Fail to set PlmnId: %+v", err)
	}
	rMcc, rMnc := plmnId.Get()
	if mcc != rMcc || mnc != rMnc {
		t.Errorf("Wrong decoded  PlmnId")
	}
}
