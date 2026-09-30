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

// A UE may send the capability with its two 5GS octets alone. Reading or
// setting an E-UTRA algorithm on one indexed past its end.
func Test_UeSecurityCapabilityWithoutEutraOctets(t *testing.T) {
	secCap := &UeSecurityCapability{}
	if err := secCap.decode([]byte{0xf0, 0x70}); err != nil {
		t.Fatalf("decode: %v", err)
	}
	for i := uint8(0); i < 8; i++ {
		if secCap.GetEEA(i) || secCap.GetEIA(i) {
			t.Errorf("a two-octet capability reports E-UTRA algorithm %d", i)
		}
	}
	if !secCap.GetEA(1) || !secCap.GetIA(1) {
		t.Errorf("the 5GS octets were not read")
	}
	if wire := secCap.Bytes(); len(wire) != 2 {
		t.Errorf("an untouched two-octet capability encodes as % x", wire)
	}

	secCap.SetEEA(2, true)
	if !secCap.GetEEA(2) {
		t.Errorf("EEA2 was not set")
	}
	if wire := secCap.Bytes(); len(wire) != 4 || wire[0] != 0xf0 || wire[1] != 0x70 || wire[2] != 0x20 {
		t.Errorf("with EEA2 set it encodes as % x, want f0 70 20 00", wire)
	}
}
