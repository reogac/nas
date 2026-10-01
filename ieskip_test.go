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

// An Authentication Response from a UE of a later release carries IEs this
// codec does not know. They are skipped (TS 24.501 7.6.1) and the RES* after
// them is read: the message used to be refused whole, and the authentication
// ran to T3560's last expiry.
func TestUnknownIesAreSkipped(t *testing.T) {
	res := bytes.Repeat([]byte{0xab}, 16)
	wire := []byte{EPD_5GMM, NasSecNone, AuthenticationResponseMsgType,
		0x9a,             //an unknown type 1 IE
		0x5a, 0x02, 1, 2, //an unknown TLV
		0x7a, 0x00, 0x01, 3, //an unknown TLV-E
		0x2d, 16}
	wire = append(wire, res...)
	msg, err := Decode(nil, wire, true)
	if err != nil {
		t.Fatalf("a message with unknown IEs was refused: %v", err)
	}
	if got := msg.Gmm.AuthenticationResponse.AuthenticationResponseParameter; !bytes.Equal(got, res) {
		t.Errorf("RES* %x, want %x", got, res)
	}
}

// An IE encoded as comprehension required refuses the message, unknown as it is.
func TestAComprehensionRequiredIeIsNotSkipped(t *testing.T) {
	wire := []byte{EPD_5GMM, NasSecNone, AuthenticationResponseMsgType, 0x05, 0x01, 0x00}
	if _, err := Decode(nil, wire, true); err == nil {
		t.Error("an unknown IE encoded as comprehension required was skipped")
	}
}

// A known optional IE that is syntactically incorrect is taken as absent and
// the IEs after it are read (TS 24.501 7.7.1). A RES* of the wrong length, or a
// TV IE whose value does not decode, used to refuse the message.
func TestAMalformedOptionalIeIsTakenAsAbsent(t *testing.T) {
	eap := []byte{2, 1, 0, 4}
	wire := []byte{EPD_5GMM, NasSecNone, AuthenticationResponseMsgType,
		0x2d, 8, 1, 2, 3, 4, 5, 6, 7, 8, //a RES* of 8 octets
		0x78, 0x00, 4}
	wire = append(wire, eap...)
	msg, err := Decode(nil, wire, true)
	if err != nil {
		t.Fatalf("a message with a malformed optional IE was refused: %v", err)
	}
	rsp := msg.Gmm.AuthenticationResponse
	if rsp.AuthenticationResponseParameter != nil {
		t.Errorf("a malformed RES* was taken: %x", rsp.AuthenticationResponseParameter)
	}
	if !bytes.Equal(rsp.EapMessage, eap) {
		t.Errorf("EAP message %x, want %x", rsp.EapMessage, eap)
	}
}

// An optional IE that runs past the end of the message is not skipped: there is
// nothing after it to read, and the message is incomplete.
func TestATruncatedOptionalIeRefusesTheMessage(t *testing.T) {
	wire := []byte{EPD_5GMM, NasSecNone, AuthenticationResponseMsgType, 0x2d, 16, 1, 2}
	if _, err := Decode(nil, wire, true); err == nil {
		t.Error("a truncated IE was accepted")
	}
}
