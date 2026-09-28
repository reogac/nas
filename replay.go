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

// REPLAY_WINDOW is how far below the highest COUNT received a message may
// arrive and still be accepted, once.
const REPLAY_WINDOW = 32

// replayWindow decides which COUNT a received message was protected with, and
// remembers which COUNTs have been used, so that each is accepted at most once
// and only once its message has verified (TS 24.501 4.4.3.2).
//
// TS 24.501 4.4.3.1 assumes NAS arrives in order: an sqn at or below the stored
// one stands for the next overflow. Between the gNB and the AMF this core has
// hops of its own, and a message they deliver a little late would be read that
// way, fail its integrity check and be lost - a Registration Complete overtaken
// by the message sent right after it hangs the registration on T3550. So a sqn
// within REPLAY_WINDOW below the highest is first read as that late message,
// the way IPsec's anti-replay window does (RFC 4303 3.4.3), and as the next
// overflow only if that does not verify. The two readings part only after 224
// or more consecutive losses, and the integrity check settles which is meant.
//
// A replay finds its COUNT marked and fails; a forged message verifies with
// neither reading and changes nothing, since only a verified message moves the
// window.
type replayWindow struct {
	highest Counter //the highest COUNT a received message verified with
	seen    bool    //whether any message has verified since the keys were derived
	//bit i marks COUNT highest-1-i as used; highest itself always is
	used uint32
}

// candidates returns the COUNTs a message carrying sqn may have been protected
// with, in the order to try them. late is set only for a COUNT inside the window
// that has not been used.
func (w *replayWindow) candidates(sqn uint8) (ahead Counter, late Counter, hasLate bool) {
	if !w.seen {
		//a context that has received nothing takes the sqn at its overflow
		return newCounter(w.highest.overflow(), sqn), 0, false
	}
	overflow := w.highest.overflow()
	if sqn <= w.highest.sqn() {
		ahead = newCounter(overflow+1, sqn)
		late = newCounter(overflow, sqn)
	} else {
		ahead = newCounter(overflow, sqn)
		if overflow == 0 {
			return //nothing below COUNT 0
		}
		late = newCounter(overflow-1, sqn)
	}
	distance := uint32(w.highest) - uint32(late)
	if distance == 0 || distance > REPLAY_WINDOW || w.used&(1<<(distance-1)) != 0 {
		return ahead, 0, false
	}
	return ahead, late, true
}

// accept records count as used. Called only once its message has verified.
func (w *replayWindow) accept(count Counter) {
	if !w.seen {
		w.highest, w.seen, w.used = count, true, 0
		return
	}
	if count <= w.highest {
		w.used |= 1 << (uint32(w.highest) - uint32(count) - 1)
		return
	}
	shift := uint32(count) - uint32(w.highest)
	//the old highest moves into the window at distance shift; anything pushed
	//past REPLAY_WINDOW is forgotten, and is refused as too old from now on
	if shift > REPLAY_WINDOW {
		w.used = 0
	} else {
		w.used = w.used<<shift | 1<<(shift-1)
	}
	w.highest = count
}

// reset forgets everything: the keys were derived again, and the peer starts
// counting from zero.
func (w *replayWindow) reset() {
	*w = replayWindow{}
}
