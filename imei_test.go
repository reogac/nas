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
	"testing"
)

func Test_Imei(t *testing.T) {
	imeiStr := "123203209090402"
	imei := new(Imei)
	var err error
	if err = imei.Parse(imeiStr); err != nil {
		t.Errorf("Parse IMEI fails: %+v", err)
	}
	var wire []byte
	if wire, err = imei.encode(); err != nil {
		t.Errorf("encode IMEI fails: %+v", err)
	}
	var newImei Imei
	if err = newImei.decode(wire); err != nil {
		t.Errorf("decode IMEI fails: %+v", err)
	}
	if imei.String() != newImei.String() {
		t.Errorf("not equal: %s != %s", imei.String(), newImei.String())
	}
}

// imeisv is the IMEISV 3569380341298401 as a 5GS mobile identity: digit 1
// beside the even flag and the type, then the other fifteen in pairs, the last
// beside the filler 1111.
var imeisv = []byte{0x30 | MobileIdentity5GSTypeImeisv, 0x65, 0x39, 0x08, 0x43, 0x21, 0x89, 0x04, 0xf1}

// The number of digits follows the odd/even flag, not the first digit. The
// flag was read from the lowest bit of the first digit, so an IMEISV starting
// with an odd digit took its filler for a 17th digit and failed the Security
// Mode Complete that carried it, and an IMEI starting with an even digit lost
// its last digit.
func Test_ImeiRoundTripsWhateverItsFirstDigit(t *testing.T) {
	for first := byte('0'); first <= '9'; first++ {
		for _, c := range []struct {
			sv     bool
			digits string
		}{{false, "56938034129842"}, {true, "569380341298401"}} {
			in := Imei{IsSv: c.sv}
			s := string(first) + c.digits
			if err := in.Parse(s); err != nil {
				t.Fatal(err)
			}
			wire, err := in.encode()
			if err != nil {
				t.Fatal(err)
			}
			out := Imei{IsSv: c.sv}
			if err := out.decode(wire); err != nil {
				t.Errorf("%s: decode: %v", s, err)
			} else if out.String() != s {
				t.Errorf("%s decoded as %s", s, out.String())
			}
		}
	}
}

// An even number of digits is encoded with the filler 1111 in the last
// half-octet (TS 24.501 9.11.3.4); it was encoded 0000.
func Test_ImeisvIsEncodedWithTheFiller(t *testing.T) {
	sv := Imei{IsSv: true}
	if err := sv.Parse("3569380341298401"); err != nil {
		t.Fatal(err)
	}
	wire, err := sv.encode()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wire, imeisv) {
		t.Errorf("IMEISV encoded as %x, want %x", wire, imeisv)
	}
}

// A Security Mode Complete carrying an IMEISV decodes, with the IMEISV in it.
func Test_ASecurityModeCompleteWithAnImeisvDecodes(t *testing.T) {
	//EPD, plain, message type, then the IMEISV IE: IEI 77, two-octet length
	wire := append([]byte{EPD_5GMM, NasSecNone, SecurityModeCompleteMsgType, 0x77, 0x00, byte(len(imeisv))}, imeisv...)
	msg, err := Decode(nil, wire, true)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	cmpl := msg.Gmm.SecurityModeComplete
	if cmpl == nil || cmpl.Imeisv == nil {
		t.Fatal("the Complete has no IMEISV")
	}
	if id, ok := cmpl.Imeisv.Id.(*Imei); !ok || !id.IsSv || id.String() != "3569380341298401" {
		t.Errorf("the Complete carries %+v, want IMEISV 3569380341298401", cmpl.Imeisv.Id)
	}
}
