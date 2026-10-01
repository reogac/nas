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
	"encoding/binary"
)

const (
	NAS_COUNT_WINDOW uint8 = 20
)

/*
TS 33.501 6.4.3.1

	COUNT (32 bits) := 0x00 || NAS COUNT (24 bits)
	NAS COUNT (24 bits) := NAS OVERFLOW (16 bits) || NAS SQN (8 bits)
*/

type Counter uint32

func newCounter(overflow uint16, sqn uint8) (c Counter) {
	c.set(overflow, sqn)
	return
}

func (c *Counter) bytes() (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(*c))
	return
}

func (c *Counter) set(overflow uint16, sqn uint8) {
	v := uint32(sqn)           //initialize with sqn
	v |= uint32(overflow) << 8 //fill in overflow
	*c = Counter(v)
}

func (c *Counter) setSqn(sqn uint8) {
	*c = Counter((uint32(*c) & 0xffffff00) | uint32(sqn))
}

// NAS_COUNT_MAX is the last COUNT a message may be protected with. One more
// would wrap to 0 and repeat a COUNT already used under the same keys (TS 24.501
// 4.4.3.5), so a counter past it is spent and protects nothing.
const NAS_COUNT_MAX Counter = 0x00ffffff

// NAS_COUNT_MARGIN is how close to NAS_COUNT_MAX a COUNT is close to the limit:
// the AMF then takes new keys into use before the limit is reached (TS 24.501
// 4.4.3.5), with that many messages left to do it in.
const NAS_COUNT_MARGIN Counter = 1 << 16

// inc moves to the next COUNT. Past NAS_COUNT_MAX it stops, spent, rather than
// wrapping to 0.
func (c *Counter) inc() {
	if *c <= NAS_COUNT_MAX {
		*c++
	}
}

// spent reports a counter past the last COUNT it may use.
func (c Counter) spent() bool { return c > NAS_COUNT_MAX }

func (c *Counter) sqn() uint8 {
	return uint8(*c & 0x000000ff)
}

func (c *Counter) overflow() uint16 {
	return uint16((*c & 0x00ffff00) >> 8)
}

func (c *Counter) setOverflow(overflow uint16) {
	*c = Counter((uint32(*c) & 0xff0000ff) | (uint32(overflow) << 8))
}
