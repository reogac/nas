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
	"testing"
)

func Test_UeSecurityCapability(t *testing.T) {
	secCap := &UeSecurityCapability{}

	secCap.SetEA(0, true)
	secCap.SetIA(5, true)
	secCap.SetExtra([]byte{1, 2, 3})

	if secCap.GetEA(1) || secCap.GetIA(7) {
		t.Errorf("Get bit return invalid values")
	}
	if !secCap.GetEA(0) || !secCap.GetIA(5) {
		t.Errorf("Set bit fails")
	}
	extra := secCap.GetExtra()
	if len(extra) != 3 || extra[2] != 3 {
		t.Errorf("set/get extra fails")
	}

	if buf, err := secCap.encode(); err != nil {
		t.Errorf("encode fails")
	} else {
		var newSecCap UeSecurityCapability
		if err = newSecCap.decode(buf); err != nil {
			t.Errorf("decode fails")
		} else {
			if !newSecCap.GetEA(0) || !newSecCap.GetIA(5) {
				t.Errorf("decode wrong")
			}
			extra = newSecCap.GetExtra()
			if len(extra) != 3 || extra[2] != 3 {
				t.Errorf("decode wrong (extra)")
			}
		}
	}

}
