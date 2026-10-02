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
	"regexp"
	"testing"
)

func Test_SuciNai(t *testing.T) {
	testCases := make(map[string]string)
	testCases["gci-"] = "gci-"
	testCases["nai-324024243243"] = "nai-324024243243"
	testCases["suci-0-129-29-0001-1-1-6783198712"] = "suci-0-129-29-0001-1-1-6783198712"
	testCases["suci-0-129-29-0001-0-0-6783198712"] = "suci-0-129-29-0001-0-0-6783198712"
	testCases["suci-0-001-01-0-1-1-aabbcc"] = "suci-0-001-01-0-1-1-aabbcc"
	testCases["suci-0-001-01-12-1-1-aabbcc"] = "suci-0-001-01-12-1-1-aabbcc"
	//the forms this package rendered before, read so that a peer on an older
	//release is understood, and rendered in the form TS 29.571 defines
	testCases["suci-12929-6783198712"] = "suci-0-129-29-0000-0-0-6783198712"
	testCases["imsi-129290-0000000001"] = "suci-0-129-290-0000-0-0-0000000001"
	testCases["suci-129-29-0001-1-1-6783198712"] = "suci-0-129-29-0001-1-1-6783198712"
	testCases["suci-129-29-0001-0-0-6783198712"] = "suci-0-129-29-0001-0-0-6783198712"
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

// tsSuci is the SUCI pattern of TS 29.571 (SupiOrSuci) and TS 29.509 (Suci),
// without the alternative that accepts any string
var tsSuci = regexp.MustCompile(`^suci-(0-[0-9]{3}-[0-9]{2,3}|[1-7]-.+)-[0-9]{1,4}-(0-0-.+|[a-fA-F1-9]-([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])-[a-fA-F0-9]+)$`)

// A SUCI renders in the form TS 29.571 defines, whatever its scheme. It lacked
// the SUPI type, and a null-scheme one rendered as imsi-plmnId-msin, so a peer
// outside this package could read neither.
func Test_ASuciRendersAsTs29571Defines(t *testing.T) {
	for _, in := range []string{
		"suci-0-208-93-0000-0-0-0000000001",
		"suci-0-208-93-12-0-0-0000000001",
		"suci-0-001-001-1234-1-1-aabbcc",
		"suci-0-208-93-0-2-255-aabbcc",
	} {
		suci := new(Suci)
		if err := suci.Parse(in); err != nil {
			t.Errorf("%s: %+v", in, err)
			continue
		}
		got := suci.String()
		if got != in {
			t.Errorf("%s rendered as %s", in, got)
		}
		if !tsSuci.MatchString(got) {
			t.Errorf("%s is not a SUCI TS 29.571 defines", got)
		}
	}
	//a SUPI type that is not an IMSI is refused rather than read as one
	if err := new(Suci).Parse("suci-1-208-93-0000-0-0-0000000001"); err == nil {
		t.Error("a SUCI of SUPI type 1 was read as an IMSI-based one")
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
