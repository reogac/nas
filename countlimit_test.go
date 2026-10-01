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
	"errors"
	"testing"
)

// The last COUNT protects one message, and none after it: the counter used to
// wrap to 0 and protect the next with a COUNT already used under the same keys
// (TS 24.501 4.4.3.5).
func TestASpentCountProtectsNothing(t *testing.T) {
	_, amf := peers(t)
	amf.pair(Bearer3GPP).local = NAS_COUNT_MAX
	msg := &DeregistrationAcceptFromUe{}
	msg.SetSecurityHeader(NasSecBoth)
	if _, err := EncodeMm(amf, msg, true); err != nil {
		t.Fatalf("the last COUNT was refused: %v", err)
	}
	if _, err := EncodeMm(amf, msg, true); !errors.Is(err, ErrCountSpent) {
		t.Errorf("a message past the last COUNT was protected (%v)", err)
	}
	//the other access has COUNTs of its own
	if _, err := EncodeMm(amf, msg, false); err != nil {
		t.Errorf("the other access's COUNT was refused: %v", err)
	}

}

// A message whose sqn would take the received COUNT past the last one is not
// read as a COUNT wrapped to 0.
func TestAReceivedCountDoesNotWrap(t *testing.T) {
	w := replayWindow{highest: newCounter(0xffff, 200), seen: true}
	if ahead, _, _ := w.candidates(100); !ahead.spent() {
		t.Errorf("sqn 100 after COUNT %x read as %x", uint32(w.highest), uint32(ahead))
	}
	if ahead, _, _ := w.candidates(201); ahead != newCounter(0xffff, 201) {
		t.Errorf("sqn 201 after COUNT %x read as %x", uint32(w.highest), uint32(ahead))
	}
}

// A COUNT within NAS_COUNT_MARGIN of the last, sent or received over either
// access, is close to the limit.
func TestACountNearTheLastIsReported(t *testing.T) {
	_, amf := peers(t)
	if amf.CountNearLimit() {
		t.Fatal("a fresh context is close to the limit")
	}
	amf.pair(BearerNon3GPP).remote.highest = NAS_COUNT_MAX - NAS_COUNT_MARGIN + 1
	if !amf.CountNearLimit() {
		t.Error("a received COUNT close to the limit is not reported")
	}
	_, amf = peers(t)
	amf.pair(Bearer3GPP).local = NAS_COUNT_MAX - NAS_COUNT_MARGIN + 1
	if !amf.CountNearLimit() {
		t.Error("a sent COUNT close to the limit is not reported")
	}
}
