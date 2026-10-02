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

// Each flag reads back what was set, through encoding and decoding. IMS VoPS
// over 3GPP access and EMCN read inverted: a network that set neither was read
// as supporting both.
func Test_NetworkFeatureSupportFlagsReadBack(t *testing.T) {
	flags := []struct {
		name string
		set  func(*NetworkFeatureSupport, bool)
		get  func(*NetworkFeatureSupport) bool
	}{
		{"IMSVoPS3GPP", (*NetworkFeatureSupport).SetIMSVoPS3GPP, (*NetworkFeatureSupport).GetIMSVoPS3GPP},
		{"IMSVoPSN3GPP", (*NetworkFeatureSupport).SetIMSVoPSN3GPP, (*NetworkFeatureSupport).GetIMSVoPSN3GPP},
		{"IWKN26", (*NetworkFeatureSupport).SetIWKN26, (*NetworkFeatureSupport).GetIWKN26},
		{"MPSI", (*NetworkFeatureSupport).SetMPSI, (*NetworkFeatureSupport).GetMPSI},
		{"EMCN", (*NetworkFeatureSupport).SetEMCN, (*NetworkFeatureSupport).GetEMCN},
		{"MCSI", (*NetworkFeatureSupport).SetMCSI, (*NetworkFeatureSupport).GetMCSI},
	}
	for _, f := range flags {
		for _, want := range []bool{false, true} {
			ie := new(NetworkFeatureSupport)
			//the second octet present either way, so a flag of it is read
			ie.SetMCSI(false)
			f.set(ie, want)
			wire, err := ie.encode()
			if err != nil {
				t.Fatalf("%s: encode: %+v", f.name, err)
			}
			back := new(NetworkFeatureSupport)
			if err = back.decode(wire); err != nil {
				t.Fatalf("%s: decode: %+v", f.name, err)
			}
			if got := f.get(back); got != want {
				t.Errorf("%s set %v reads back %v", f.name, want, got)
			}
		}
	}
	//nothing set reads as nothing supported
	none := new(NetworkFeatureSupport)
	if none.GetIMSVoPS3GPP() || none.GetIMSVoPSN3GPP() || none.GetEMCN() || none.GetEMC() != 0 || none.GetEMF() != 0 {
		t.Errorf("an empty IE reads as supporting something: %+v", none)
	}
}
