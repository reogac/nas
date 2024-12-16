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
	"math"
	"testing"
)

func Test_Counter(t *testing.T) {
	c := newCounter(100, 10)
	if c.overflow() != 100 {
		t.Errorf("wrong overflow")
	}

	if c.sqn() != 10 {
		t.Errorf("wrong sqn")
	}
	c.inc()

	if c.sqn() != 11 {
		t.Errorf("wrong inc")
	}
	c.setSqn(math.MaxUint8)
	c.inc()
	if c.sqn() != 0 || c.overflow() != 101 {
		t.Errorf("fail increase sqn with overflow")
	}
}
