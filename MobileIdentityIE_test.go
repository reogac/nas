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
	"encoding/hex"
	"testing"
)

func Test_MobileIdentity(t *testing.T) {
	log.Infof("Test mobile identity encoding/decoding")
	testCases := make(map[string]MobileIdentityInf)
	//add suci ids
	suciList := make(map[string]string)
	suciList["gci-"] = "gci-"
	suciList["nai-324024243243"] = "nai-324024243243"
	suciList["suci-12929-6783198712"] = "imsi-12929-6783198712"
	suciList["imsi-129290-0000000001"] = "imsi-129290-0000000001"
	suciList["suci-129-29-0001-1-1-6783198712"] = "suci-129-29-0001-1-1-6783198712"
	suciList["suci-129-29-0001-0-1-6783198712"] = "imsi-12929-6783198712"
	for in, out := range suciList {
		suci := new(Suci)
		suci.Parse(in)
		testCases[out] = suci
	}
	//add imei id
	imeiStr := "000111222333"
	imei := new(Imei)
	imei.Parse(imeiStr)
	testCases[imeiStr] = imei

	// add guti ids
	gutiList := []string{
		"1209011111100000001",
		"12090911111180000001",
		"12090911111180000001",
		"1209011111109000001",
	}
	for _, gutiStr := range gutiList {
		guti := new(Guti)
		guti.Parse(gutiStr)
		testCases[gutiStr] = guti
	}
	// add tmsi5gs ids
	tmsi5gsList := []string{
		"121100000001",
		"131180000002",
		"211180000003",
		"111109000004",
	}
	for _, tmsi5gsStr := range tmsi5gsList {
		tmsi5gs := new(Tmsi5Gs)
		tmsi5gs.Parse(tmsi5gsStr)
		testCases[tmsi5gsStr] = tmsi5gs
	}

	//add Mac id
	mac := &MacIdentity{}
	copy(mac.Bytes[:], []byte{1, 1, 1, 1, 1, 1})
	macStr := hex.EncodeToString(mac.Bytes[:])
	testCases[macStr] = mac

	//add Eui64 id
	eui64 := &Eui64Identity{}
	copy(eui64.Bytes[:], []byte{1, 1, 1, 1, 1, 1, 2, 2})

	eui64Str := hex.EncodeToString(eui64.Bytes[:])
	testCases[eui64Str] = eui64

	var wire []byte
	var err error
	for idStr, idContent := range testCases {
		id := &MobileIdentity{
			Id: idContent,
		}
		if wire, err = id.encode(); err != nil {
			t.Errorf("Encode mobile id fails: %+v", err)
			continue
		}
		newId := new(MobileIdentity)
		if err = newId.decode(wire); err != nil {
			t.Errorf("Decode mobile id fails: %+v", err)
			continue
		}
		if newId.Id.String() != idStr {
			t.Errorf("Not equal: %s != %s", idStr, newId.Id.String())
			continue
		}
	}

}
