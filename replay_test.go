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

// sealed returns the UE's uplinks 0..n-1, protected in order, for the test to
// deliver in any order it likes.
func sealed(t *testing.T, ue *NasContext, n int) [][]byte {
	out := make([][]byte, n)
	out[0] = uplink(t, ue, NasSecBothNew)
	for i := 1; i < n; i++ {
		out[i] = uplink(t, ue, NasSecBoth)
	}
	return out
}

// A message overtaken by the one sent after it still verifies, once: the hops
// between the gNB and the AMF may deliver a little out of order, and a
// Registration Complete overtaken by the next uplink used to be refused as a
// replay, hanging the registration on T3550.
func TestAReorderedUplinkIsAcceptedOnce(t *testing.T) {
	ue, amf := peers(t)
	msgs := sealed(t, ue, 4)
	for i, k := range []int{0, 2, 1, 3} {
		if !verified(amf, msgs[k]) {
			t.Fatalf("delivery %d (message %d) did not verify", i, k)
		}
	}
	for k := range msgs {
		if verified(amf, msgs[k]) {
			t.Errorf("message %d verified a second time", k)
		}
	}
	if got, want := amf.UlCounter(), uint32(3); got != want {
		t.Errorf("uplink COUNT %d after a late message, want the highest, %d", got, want)
	}
}

// A message later than the window is refused: it is indistinguishable from a
// replay of one the window has forgotten.
func TestAnUplinkOlderThanTheWindowIsRefused(t *testing.T) {
	ue, amf := peers(t)
	msgs := sealed(t, ue, REPLAY_WINDOW+3)
	verified(amf, msgs[0])
	late := msgs[1]
	for _, m := range msgs[2:] {
		if !verified(amf, m) {
			t.Fatal("an in-order message did not verify")
		}
	}
	if verified(amf, late) {
		t.Errorf("a message %d behind the highest verified", REPLAY_WINDOW+1)
	}
}

// The last message inside the window is still accepted.
func TestTheWindowsOldestSlotIsAccepted(t *testing.T) {
	ue, amf := peers(t)
	msgs := sealed(t, ue, REPLAY_WINDOW+2)
	verified(amf, msgs[0])
	late := msgs[1]
	for _, m := range msgs[2:] {
		verified(amf, m)
	}
	//highest is REPLAY_WINDOW+1, and message 1 is REPLAY_WINDOW below it
	if !verified(amf, late) {
		t.Errorf("a message %d behind the highest was refused", REPLAY_WINDOW)
	}
}

// A forged message leaves the window as it was: a late message it pretends to
// be is still accepted afterwards, and so is the next one.
func TestAForgedUplinkLeavesTheWindowUnchanged(t *testing.T) {
	ue, amf := peers(t)
	msgs := sealed(t, ue, 4)
	for _, k := range []int{0, 1, 3} {
		verified(amf, msgs[k])
	}
	before := amf.pairs[0].remote
	forged := append([]byte(nil), msgs[2]...)
	forged[2] ^= 0xff
	if verified(amf, forged) {
		t.Fatal("forged message verified")
	}
	if amf.pairs[0].remote != before {
		t.Errorf("window moved from %+v to %+v on a message that failed its check", before, amf.pairs[0].remote)
	}
	if !verified(amf, msgs[2]) {
		t.Error("the real late message was refused after a forgery of it")
	}
}

// Past an sqn wrap, the window still spans the old overflow.
func TestTheWindowSpansAnOverflow(t *testing.T) {
	ue, amf := peers(t)
	msgs := sealed(t, ue, 260)
	for k := 0; k < 255; k++ {
		verified(amf, msgs[k])
	}
	//255 overtaken by 256 and 257, which carry the next overflow
	for _, k := range []int{256, 257, 255, 258} {
		if !verified(amf, msgs[k]) {
			t.Fatalf("message %d did not verify", k)
		}
	}
	if verified(amf, msgs[255]) {
		t.Error("the late message across the wrap verified twice")
	}
}

// A container a late message carries is ciphered with that message's COUNT,
// not with the highest: the decode keeps the COUNT it verified with.
func TestTheLastCountIsTheLateMessages(t *testing.T) {
	ue, amf := peers(t)
	msgs := sealed(t, ue, 3)
	for _, k := range []int{0, 2, 1} {
		verified(amf, msgs[k])
	}
	if amf.pairs[0].remoteLast != newCounter(0, 1) || amf.pairs[0].remote.highest != newCounter(0, 2) {
		t.Errorf("last %d highest %d, want 1 and 2", amf.pairs[0].remoteLast, amf.pairs[0].remote.highest)
	}
}

func TestCandidates(t *testing.T) {
	w := replayWindow{highest: newCounter(3, 200), seen: true, used: 1 << 1} //COUNT (3,198) used
	cases := []struct {
		sqn     uint8
		ahead   Counter
		late    Counter
		hasLate bool
	}{
		{201, newCounter(3, 201), 0, false},
		{200, newCounter(4, 200), 0, false},                 //the highest itself
		{199, newCounter(4, 199), newCounter(3, 199), true}, //just below
		{198, newCounter(4, 198), 0, false},                 //already used
		{168, newCounter(4, 168), newCounter(3, 168), true}, //the window's edge
		{167, newCounter(4, 167), 0, false},                 //past it
	}
	for _, c := range cases {
		ahead, late, hasLate := w.candidates(c.sqn)
		if ahead != c.ahead || hasLate != c.hasLate || (hasLate && late != c.late) {
			t.Errorf("candidates(%d) = %d, %d, %v; want %d, %d, %v",
				c.sqn, ahead, late, hasLate, c.ahead, c.late, c.hasLate)
		}
	}
	fresh := replayWindow{}
	if ahead, _, hasLate := fresh.candidates(5); ahead != newCounter(0, 5) || hasLate {
		t.Errorf("a fresh window read sqn 5 as %d (late %v)", ahead, hasLate)
	}
	//sqn above the highest at overflow 0 has nothing below COUNT 0 to be late for
	low := replayWindow{highest: newCounter(0, 3), seen: true}
	if _, _, hasLate := low.candidates(9); hasLate {
		t.Error("a COUNT below 0 was offered as late")
	}
}

// A retransmitted message under a new security context takes the next COUNT,
// as any retransmission does (TS 24.501 4.4.3.1), and the peer that took the
// first accepts it. Each one used to be sent with COUNT 0, so a UE that had
// accepted a SECURITY MODE COMMAND discarded its retransmissions as replays.
func TestARetransmittedNewContextMessageTakesTheNextCount(t *testing.T) {
	ue, amf := peers(t)
	caps := UeSecurityCapability{}
	caps.SetEA(2, true)
	caps.SetIA(2, true)
	cmd := &SecurityModeCommand{
		SelectedNasSecurityAlgorithms:  NewSecurityAlgorithms(AlgIntegrity128NIA2, AlgCiphering128NEA2),
		ReplayedUeSecurityCapabilities: caps,
	}
	cmd.SetSecurityHeader(NasSecIntegrityNew)
	var sent [2][]byte
	for i := range sent {
		wire, err := EncodeMm(amf, cmd, true)
		if err != nil {
			t.Fatal(err)
		}
		sent[i] = wire
	}
	if sent[0][6] == sent[1][6] {
		t.Fatalf("both commands carry sequence number %d", sent[0][6])
	}
	for i, wire := range sent {
		msg, err := Decode(ue, wire, true)
		if err != nil || msg.Gmm == nil || msg.Gmm.MacFailed {
			t.Errorf("command %d was not accepted: %v", i, err)
		}
	}
}

// A context used over both accesses keeps a pair of NAS COUNTs for each (TS
// 24.501 4.4.3.1): messages over 3GPP access leave the non-3GPP COUNTs where
// they were, so the first non-3GPP message, with COUNT 0, still verifies. With
// one pair it was taken for a replay of a 3GPP message's COUNT.
func TestTheAccessesKeepCountsOfTheirOwn(t *testing.T) {
	ue, amf := peers(t)
	for i := 0; i < 3; i++ {
		if !verified(amf, uplink(t, ue, NasSecBoth)) {
			t.Fatalf("3GPP message %d did not verify", i)
		}
	}
	if got := amf.UlCounterFor(false); got != 0 {
		t.Errorf("the non-3GPP uplink COUNT moved to %d with 3GPP messages", got)
	}

	msg := &RegistrationComplete{}
	msg.SetSecurityHeader(NasSecBoth)
	wire, err := EncodeMm(ue, msg, false)
	if err != nil {
		t.Fatal(err)
	}
	if wire[6] != 0 {
		t.Errorf("the first non-3GPP message carries sequence number %d, want 0", wire[6])
	}
	got, err := Decode(amf, wire, false)
	if err != nil || got.Gmm == nil || got.Gmm.MacFailed {
		t.Errorf("the first non-3GPP message did not verify: %v", err)
	}
	if amf.UlCounterFor(true) != 2 || amf.UlCounter() != 2 {
		t.Errorf("the 3GPP uplink COUNT is %d, want 2", amf.UlCounterFor(true))
	}
}
