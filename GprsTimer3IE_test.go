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

// The unit was once ignored, so a back-off built with this constructor went out
// with whatever unit bits the multiplier happened to carry.
func Test_NewGprsTimer3(t *testing.T) {
	cases := []struct {
		u, v uint8
		want uint8
	}{
		{0b011, 4, 0x64},   // 4 x 2 seconds
		{0b000, 5, 0x05},   // 5 x 10 minutes
		{0b001, 1, 0x21},   // 1 hour
		{0b111, 0, 0xe0},   // deactivated
		{0xff, 0xff, 0xff}, // bits above each field are dropped
	}
	for _, c := range cases {
		ie := NewGprsTimer3(c.u, c.v)
		if ie.Value != c.want {
			t.Errorf("NewGprsTimer3(%03b, %d) = %#02x, want %#02x", c.u, c.v, ie.Value, c.want)
		}
		wire, err := ie.encode()
		if err != nil {
			t.Fatalf("encode: %v", err)
		}
		var back GprsTimer3
		if err = back.decode(wire); err != nil || back.Value != c.want {
			t.Errorf("round trip of %#02x = %#02x, %v", c.want, back.Value, err)
		}
	}
}
