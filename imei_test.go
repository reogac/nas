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

func Test_Imei(t *testing.T) {
	imeiStr := "123203209090402"
	imei := new(Imei)
	var err error
	if err = imei.Parse(imeiStr); err != nil {
		t.Errorf("Parse IMEI fails: %+v", err)
	}
	var wire []byte
	if wire, err = imei.encode(); err != nil {
		t.Errorf("encode IMEI fails: %+v", err)
	}
	var newImei Imei
	if err = newImei.decode(wire); err != nil {
		t.Errorf("decode IMEI fails: %+v", err)
	}
	if imei.String() != newImei.String() {
		t.Errorf("not equal: %s != %s", imei.String(), newImei.String())
	}
}
