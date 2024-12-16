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

// 9.11.2.5 GPRS timer 3 [O TLV 3]
type GprsTimer3 struct {
	Value uint8
}

func (ie *GprsTimer3) encode() (wire []byte, err error) {
	wire = []byte{ie.Value}
	return
}

func (ie *GprsTimer3) decode(wire []byte) (err error) {
	if len(wire) != 1 {
		err = ErrInvalidSize
		return
	}
	ie.Value = wire[0]
	return
}

// just for testing, need to correct this
func NewGprsTimer3(u uint8, v uint8) *GprsTimer3 {
	//TODO:
	return &GprsTimer3{
		Value: v,
	}
}
