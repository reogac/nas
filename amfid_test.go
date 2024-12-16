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
	//	"fmt"
	"testing"
)

func Test_AmfId(t *testing.T) {
	var id AmfId
	var err error
	id.SetRegion(3)
	id.SetSet(256)
	id.SetPointer(37)
	var newId *AmfId
	if newId, err = ParseAmfId(id.String()); err != nil {
		t.Errorf("Fail to parse AmfId: %+v", err)
	}
	if newId.GetRegion() != 3 || newId.GetSet() != 256 || newId.GetPointer() != 37 {
		t.Errorf("Invalid parsed AmfId: %d-%d-%d", newId.GetRegion(), newId.GetSet(), newId.GetPointer())
	}
}
