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

func Test_SNssai(t *testing.T) {
	sst1 := uint8(11)
	sst2 := uint8(22)
	sd1 := "000001"
	sd2 := "000002"
	var list [5]SNssai
	list[0].Set(sst1, "")
	list[1].Set(sst1, "")
	list[1].SetMapped(sst2, "")
	list[2].Set(sst1, sd1)
	list[3].Set(sst1, sd1)
	list[3].SetMapped(sst2, "")
	list[4].Set(sst1, sd1)
	list[4].SetMapped(sst2, sd2)

	for i := 0; i < 5; i++ {
		buf, _ := list[i].encode()
		snssai := new(SNssai)
		snssai.decode(buf)
		if !compareSnssai(&list[i], snssai) {
			t.Errorf("check equality fails on item %d", i)
		}
	}
}

func compareSnssai(s1, s2 *SNssai) bool {
	if s1.Sst != s2.Sst {
		return false
	}
	if !bytes.Equal(s1.Sd, s2.Sd) {
		return false
	}
	if s1.Mapped == nil && s2.Mapped == nil {
		return true
	}
	if s1.Mapped != nil && s2.Mapped == nil {
		return false
	}

	if s1.Mapped == nil && s2.Mapped != nil {
		return false
	}
	if s1.Mapped.Sst != s2.Mapped.Sst {
		return false
	}
	return bytes.Equal(s1.Mapped.Sd, s2.Mapped.Sd)
}
