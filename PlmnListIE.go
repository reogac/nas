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
	"fmt"
)

// 9.11.3.45 PLMN list [O TLV 5-47]
type PlmnList struct {
	plmnIdList []PlmnId
}

func (ie *PlmnList) encode() (wire []byte, err error) {
	var buf []byte
	for _, plmnId := range ie.plmnIdList {
		if buf, err = plmnId.encode(); err != nil {
			return
		}
		wire = append(wire, buf...)
	}
	return
}

func (ie *PlmnList) decode(wire []byte) (err error) {
	l := len(wire)
	if l%3 != 0 {
		return fmt.Errorf("Invalid Plmn Id list size")
	}
	offset := 0
	for offset < l {
		plmnId := new(PlmnId)
		if err = plmnId.decode(wire[offset : offset+3]); err != nil {
			return nasError("decode PlmnId list", err)
		}
		ie.plmnIdList = append(ie.plmnIdList, *plmnId)
		offset += 3
	}

	return
}

func NewPlmnList(list []PlmnId) *PlmnList {
	ie := &PlmnList{
		plmnIdList: list,
	}
	return ie
}
func (ie *PlmnList) GetList() []PlmnId {
	return ie.plmnIdList
}
