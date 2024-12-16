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

func Test_NasEncoding(t *testing.T) {
	log.Infof("Test nas message encoding/decoding")
	guti := new(Guti)
	guti.PlmnId.Set("100", "20")
	guti.AmfId.Set(10, 100, 50)

	msg := &RegistrationRequest{
		RegistrationType: NewRegistrationType(true, RegistrationType5GSInitialRegistration),
		MobileIdentity: MobileIdentity{
			Id: guti,
		},
		Ngksi: KeySetIdentifier{
			Tsc: 1,
			Id:  3,
		},
		GmmCapability: &GmmCapability{
			Bytes: []byte{1, 2, 3},
		},
		//LastVisitedRegisteredTai: NewBytes([]byte{1, 2, 3, 4, 5, 6}), //6 bytes
		NonCurrentNativeNasKeySetIdentifier: &KeySetIdentifier{
			Tsc: 0,
			Id:  4,
		},
		PayloadContainerType:   new(uint8),
		NasMessageContainer:    []byte{1}, //not-empty
		EpsBearerContextStatus: new(uint16),
	}
	*msg.PayloadContainerType = 1
	*msg.EpsBearerContextStatus = 100
	msg.SetSecurityHeader(0)
	if buf, err := EncodeMm(nil, msg); err != nil {
		t.Errorf("Encode fail: %+v", err)
		return
	} else {
		var nasMsg NasMessage
		if nasMsg, err = Decode(nil, buf); err != nil {
			t.Errorf("Decode fail: %s", err.Error())
			return
		}
		var newMsg *RegistrationRequest
		newMsg = nasMsg.Gmm.RegistrationRequest
		if newMsg.RegistrationType.value != msg.RegistrationType.value {
			t.Errorf("Decode fail: mismatched registration type %d:%d", msg.RegistrationType.value, newMsg.RegistrationType.value)
		}
		newGuti := newMsg.MobileIdentity.Id.(*Guti)
		mcc, mnc := newGuti.PlmnId.Get()
		if mcc != "100" || mnc != "20" {
			t.Errorf("Decode fail: mismatched guti")
		}
		if uint8(*newMsg.PayloadContainerType) != uint8(*msg.PayloadContainerType) {
			t.Errorf("Decode fail: mismatched payload container type")
		}
		if *newMsg.EpsBearerContextStatus != *msg.EpsBearerContextStatus {
			t.Errorf("Decode fail: mismatched eps bearer context status")
		}
		t.Logf("EpsBearerContextStatus=%d\n", *newMsg.EpsBearerContextStatus)
	}
}
