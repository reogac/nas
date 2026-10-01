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

import "testing"

// A protected message cut short anywhere is refused, not read past its end. Six
// octets - a header and a MAC with no sequence number - used to panic.
func TestATruncatedProtectedMessageIsRefused(t *testing.T) {
	full := []byte{EPD_5GMM, NasSecIntegrity, 1, 2, 3, 4, 0, EPD_5GMM, NasSecNone, RegistrationCompleteMsgType}
	for _, ctx := range []*NasContext{nil, NewNasContext(true)} {
		for n := 2; n < 7; n++ {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("%d octets panicked: %v", n, r)
					}
				}()
				if _, err := Decode(ctx, full[:n], true); err == nil {
					t.Errorf("%d octets were accepted", n)
				}
			}()
		}
	}
}
