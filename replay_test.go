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

// peers returns a UE and an AMF context keyed alike
func peers(t *testing.T) (ue, amf *NasContext) {
	kAmf := make([]byte, 32)
	for i := range kAmf {
		kAmf[i] = byte(i)
	}
	ue, amf = NewNasContext(false), NewNasContext(true)
	for _, ctx := range []*NasContext{ue, amf} {
		if err := ctx.DeriveKeys(AlgCiphering128NEA2, AlgIntegrity128NIA2, kAmf); err != nil {
			t.Fatalf("derive keys: %+v", err)
		}
	}
	return
}

func uplink(t *testing.T, ue *NasContext, secType uint8) []byte {
	msg := &RegistrationComplete{}
	msg.SetSecurityHeader(secType)
	wire, err := EncodeMm(ue, msg, true)
	if err != nil {
		t.Fatalf("encode: %+v", err)
	}
	return wire
}

// verified reports whether the AMF accepts wire as integrity protected
func verified(amf *NasContext, wire []byte) bool {
	msg, err := Decode(amf, wire, true)
	return err == nil && msg.Gmm != nil && !msg.Gmm.MacFailed
}

func TestReplayedUplinkIsRefused(t *testing.T) {
	ue, amf := peers(t)
	var sent [][]byte
	sent = append(sent, uplink(t, ue, NasSecBothNew))
	for i := 0; i < 5; i++ {
		sent = append(sent, uplink(t, ue, NasSecBoth))
	}
	for i, wire := range sent {
		if !verified(amf, wire) {
			t.Fatalf("message %d did not verify", i)
		}
	}
	//every one of them, the latest included, is within the old window of 20
	for i, wire := range sent {
		if verified(amf, wire) {
			t.Errorf("message %d verified a second time", i)
		}
	}
	//the replays moved nothing: the UE's next message still verifies
	if !verified(amf, uplink(t, ue, NasSecBoth)) {
		t.Errorf("next message from the UE refused after the replays")
	}
}

func TestIntegrityProtectedReplayIsReportedMacFailed(t *testing.T) {
	ue, amf := peers(t)
	wire := uplink(t, ue, NasSecIntegrity)
	if !verified(amf, wire) {
		t.Fatalf("message did not verify")
	}
	msg, err := Decode(amf, wire, true)
	if err != nil || msg.Gmm == nil || !msg.Gmm.MacFailed {
		t.Errorf("replayed integrity-protected message not flagged MacFailed: err=%v", err)
	}
}

func TestForgedUplinkLeavesCounterUnchanged(t *testing.T) {
	ue, amf := peers(t)
	for i := 0; i < 30; i++ {
		if !verified(amf, uplink(t, ue, NasSecBoth)) {
			t.Fatalf("message %d did not verify", i)
		}
	}
	before := amf.UlCounter()
	//an sqn more than 20 below the stored one used to bump the overflow
	//before the check, whatever the check then found
	forged := uplink(t, ue, NasSecIntegrity)
	forged[6] = 1
	forged[2] ^= 0xff
	if verified(amf, forged) {
		t.Fatalf("forged message verified")
	}
	if got := amf.UlCounter(); got != before {
		t.Errorf("uplink COUNT moved from %d to %d on a message that failed its check", before, got)
	}
	//the forgery consumed the UE's COUNT 30 on the UE side only; its next one
	//is still ahead of the stored COUNT
	if !verified(amf, uplink(t, ue, NasSecBoth)) {
		t.Errorf("real UE refused after a forged message")
	}
}

func TestUplinkCountCrossesSqnWrap(t *testing.T) {
	ue, amf := peers(t)
	for i := 0; i < 600; i++ {
		if !verified(amf, uplink(t, ue, NasSecBoth)) {
			t.Fatalf("message %d did not verify", i)
		}
	}
	if got, want := amf.UlCounter(), uint32(599); got != want {
		t.Errorf("uplink COUNT %d, want %d", got, want)
	}
}

func TestLostUplinksAreTolerated(t *testing.T) {
	ue, amf := peers(t)
	for i := 0; i < 3; i++ {
		verified(amf, uplink(t, ue, NasSecBoth))
	}
	for i := 0; i < 100; i++ {
		uplink(t, ue, NasSecBoth) //never delivered
	}
	if !verified(amf, uplink(t, ue, NasSecBoth)) {
		t.Errorf("message after lost ones refused")
	}
}

func TestNewContextHeaderOnUsedContextIsRefused(t *testing.T) {
	ue, amf := peers(t)
	first := uplink(t, ue, NasSecBothNew)
	if !verified(amf, first) {
		t.Fatalf("first message did not verify")
	}
	if verified(amf, first) {
		t.Errorf("new-context message verified again on a context that has taken it")
	}
}

func TestEstimate(t *testing.T) {
	stored := newCounter(3, 200)
	cases := []struct {
		sqn  uint8
		seen bool
		want Counter
	}{
		{201, true, newCounter(3, 201)},
		{255, true, newCounter(3, 255)},
		{200, true, newCounter(4, 200)},
		{199, true, newCounter(4, 199)},
		{0, true, newCounter(4, 0)},
		{5, false, newCounter(3, 5)},
	}
	for _, c := range cases {
		if got := stored.estimate(c.sqn, c.seen); got != c.want {
			t.Errorf("estimate(%d, %v) = %d, want %d", c.sqn, c.seen, got, c.want)
		}
	}
}
