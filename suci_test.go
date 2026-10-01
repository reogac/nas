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

func Test_SuciNai(t *testing.T) {
	testCases := make(map[string]string)
	testCases["gci-"] = "gci-"
	testCases["nai-324024243243"] = "nai-324024243243"
	testCases["suci-12929-6783198712"] = "imsi-12929-6783198712"
	testCases["imsi-129290-0000000001"] = "imsi-129290-0000000001"
	testCases["suci-129-29-0001-1-1-6783198712"] = "suci-129-29-0001-1-1-6783198712"
	testCases["suci-129-29-0001-0-1-6783198712"] = "imsi-12929-6783198712"
	testCases["suci-001-01-0-1-1-aabbcc"] = "suci-001-01-0-1-1-aabbcc"
	testCases["suci-001-01-12-1-1-aabbcc"] = "suci-001-01-12-1-1-aabbcc"
	var err error
	for suciStr, expectedSuciStr := range testCases {
		suci := new(Suci)
		if err = suci.Parse(suciStr); err != nil {
			t.Errorf("NAI parse fail: %+v", err)
			continue
		}
		var buf []byte
		if buf, err = suci.encode(); err != nil {
			t.Errorf("encode suci fail: %+v", err)
			return
		}
		var newSuci = new(Suci)

		if err = newSuci.decode(buf); err != nil {
			t.Errorf("decode suci fail: %+v", err)
			return
		}

		if expectedSuciStr != newSuci.String() {
			t.Errorf("expected suci=%s, parsed: %s", expectedSuciStr, newSuci.String())
			return
		}
	}
}

// A SUCI whose scheme output is empty is refused. It was taken, and rendering
// its MSIN indexed out of range.
func Test_ASuciWithNoSchemeOutputIsRefused(t *testing.T) {
	//an Identity Response carrying a SUCI of IMSI format, PLMN 001/01, routing
	//indicator 0, the null scheme, key id 0 - and nothing after
	suci := []byte{0x01, 0x00, 0xf1, 0x10, 0x00, 0x00, 0x00, 0x00}
	wire := append([]byte{EPD_5GMM, NasSecNone, IdentityResponseMsgType, 0x00, byte(len(suci))}, suci...)
	if _, err := Decode(nil, wire, true); err == nil {
		t.Fatal("a SUCI with no scheme output decoded")
	}
	//and with one octet of MSIN it decodes, and renders
	suci = append(suci, 0x21)
	wire = append([]byte{EPD_5GMM, NasSecNone, IdentityResponseMsgType, 0x00, byte(len(suci))}, suci...)
	msg, err := Decode(nil, wire, true)
	if err != nil {
		t.Fatalf("a SUCI with an MSIN failed to decode: %v", err)
	}
	_ = msg.Gmm.IdentityResponse.MobileIdentity.String()
}

// A routing indicator of fewer than 4 digits is read as the digits it has: a UE
// with none configured codes it 0 followed by the 1111 filler (TS 24.501
// 9.11.3.4). The filler rendered as "15", in a SUCI no parser took back, and
// the UDM could not de-conceal it.
func Test_AShortRoutingIndicatorIsItsDigits(t *testing.T) {
	for wire, want := range map[[2]byte]string{{0xf0, 0xff}: "0", {0x21, 0xff}: "12", {0x21, 0xf3}: "123", {0x21, 0x43}: "1234"} {
		ri := new(RoutingIndicator)
		if err := ri.decode(wire[:]); err != nil {
			t.Errorf("%x refused: %v", wire, err)
			continue
		}
		if got := ri.String(); got != want {
			t.Errorf("%x reads %q, want %q", wire, got, want)
		}
		back := new(RoutingIndicator)
		if err := back.Parse(want); err != nil || back.bytes != wire {
			t.Errorf("%q parses to %x (%v), want %x", want, back.bytes, err, wire)
		}
	}
}

// A routing indicator that is not 1 to 4 digits followed by filler is refused.
func Test_AMalformedRoutingIndicatorIsRefused(t *testing.T) {
	for _, wire := range [][2]byte{{0xff, 0xff}, {0x1f, 0xf2}, {0xa1, 0xff}} {
		if err := new(RoutingIndicator).decode(wire[:]); err == nil {
			t.Errorf("%x was taken", wire)
		}
	}
	for _, s := range []string{"", "12345"} {
		if err := new(RoutingIndicator).Parse(s); err == nil {
			t.Errorf("%q was taken", s)
		}
	}
}
