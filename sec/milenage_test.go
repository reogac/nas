package sec

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type MilenageTestCase struct {
	K       string
	RAND    string
	SQN     string
	AMF     string
	OP      string
	eOPC    string
	f1      string
	f1star  string
	eRES    string
	eAK     string
	eCK     string
	eIK     string
	eAKstar string
}

func TestAll(t *testing.T) {
	table := []MilenageTestCase{
		{
			K:       "465b5ce8b199b49faa5f0a2ee238a6bc",
			RAND:    "23553cbe9637a89d218ae64dae47bf35",
			SQN:     "ff9bb4d0b607",
			AMF:     "b9b9",
			OP:      "cdc202d5123e20f62b6d676ac72cb318",
			eOPC:    "cd63cb71954a9f4e48a5994e37a02baf",
			f1:      "4a9ffac354dfafb3",
			f1star:  "01cfaf9ec4e871e9",
			eRES:    "a54211d5e3ba50bf",
			eAK:     "aa689c648370",
			eCK:     "b40ba9a3c58b2a05bbf0d987b21bf8cb",
			eIK:     "f769bcd751044604127672711c6d3441",
			eAKstar: "451e8beca43b",
		},
		{
			K:       "0396eb317b6d1c36f19c1c84cd6ffd16",
			RAND:    "c00d603103dcee52c4478119494202e8",
			SQN:     "fd8eef40df7d",
			AMF:     "af17",
			OP:      "ff53bade17df5d4e793073ce9d7579fa",
			eOPC:    "53c15671c60a4b731c55b4a441c0bde2",
			f1:      "5df5b31807e258b0",
			f1star:  "a8c016e51ef4a343",
			eRES:    "d3a628ed988620f0",
			eAK:     "c47783995f72",
			eCK:     "58c433ff7a7082acd424220f2b67c556",
			eIK:     "21a8c1f929702adb3e738488b9f5c5da",
			eAKstar: "30f1197061c1",
		},
		{
			K:       "fec86ba6eb707ed08905757b1bb44b8f",
			RAND:    "9f7c8d021accf4db213ccff0c7f71a6a",
			SQN:     "9d0277595ffc",
			AMF:     "725c",
			OP:      "dbc59adcb6f9a0ef735477b7fadf8374",
			eOPC:    "1006020f0a478bf6b699f15c062e42b3",
			f1:      "9cabc3e99baf7281",
			f1star:  "95814ba2b3044324",
			eRES:    "8011c48c0c214ed2",
			eAK:     "33484dc2136b",
			eCK:     "5dbdbb2954e8f3cde665b046179a5098",
			eIK:     "59a92d3b476a0443487055cf88b2307b",
			eAKstar: "deacdd848cc6",
		},
		{
			K:       "9e5944aea94b81165c82fbf9f32db751",
			RAND:    "ce83dbc54ac0274a157c17f80d017bd6",
			SQN:     "0b604a81eca8",
			AMF:     "9e09",
			OP:      "223014c5806694c007ca1eeef57f004f",
			eOPC:    "a64a507ae1a2a98bb88eb4210135dc87",
			f1:      "74a58220cba84c49",
			f1star:  "ac2cc74a96871837",
			eRES:    "f365cd683cd92e96",
			eAK:     "f0b9c08ad02e",
			eCK:     "e203edb3971574f5a94b0d61b816345d",
			eIK:     "0c4524adeac041c4dd830d20854fc46b",
			eAKstar: "6085a86c6f63",
		},
		{
			K:       "4ab1deb05ca6ceb051fc98e77d026a84",
			RAND:    "74b0cd6031a1c8339b2b6ce2b8c4a186",
			SQN:     "e880a1b580b6",
			AMF:     "9f07",
			OP:      "2d16c5cd1fdf6b22383584e3bef2a8d8",
			eOPC:    "dcf07cbd51855290b92a07a9891e523e",
			f1:      "49e785dd12626ef2",
			f1star:  "9e85790336bb3fa2",
			eRES:    "5860fc1bce351e7e",
			eAK:     "31e11a609118",
			eCK:     "7657766b373d1c2138f307e3de9242f9",
			eIK:     "1c42e960d89b8fa99f2744e0708ccb53",
			eAKstar: "fe2555e54aa9",
		}, /*
			{
				K:      "6c38a116ac280c454f59332ee35c8c4f",
				RAND:   "ee6466bc96202c5a557abbeff8babf63",
				SQN:    "414b98222181",
				AMF:    "4464",
				OP:     "1ba00a1a7c6700ac8c3ff3e96ad08725",
				eOPC:   "3803ef5363b947c6aaa225e58fae3934",
				f1:     "078adfb488241a57",
				f1star: "80246b8d0186bcf1",
			eIK:     "a7466cc1e6b2a1337d49d3b66e95d7b4",
			eAKStar: "1f53cd2b1113",
			},
		*/
	}
	for _, testcase := range table {
		K, err := hex.DecodeString(strings.Repeat(testcase.K, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		RAND, err := hex.DecodeString(strings.Repeat(testcase.RAND, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		SQN, err := hex.DecodeString(strings.Repeat(testcase.SQN, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		AMF, err := hex.DecodeString(strings.Repeat(testcase.AMF, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		OP, err := hex.DecodeString(strings.Repeat(testcase.OP, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		eOPC, err := hex.DecodeString(strings.Repeat(testcase.eOPC, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		f1, err := hex.DecodeString(strings.Repeat(testcase.f1, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		f1star, err := hex.DecodeString(strings.Repeat(testcase.f1star, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		eRES, err := hex.DecodeString(strings.Repeat(testcase.eRES, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		eCK, err := hex.DecodeString(strings.Repeat(testcase.eCK, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		eAK, err := hex.DecodeString(strings.Repeat(testcase.eAK, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		eIK, err := hex.DecodeString(strings.Repeat(testcase.eIK, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}

		eAKstar, err := hex.DecodeString(strings.Repeat(testcase.eAKstar, 1))
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		var m *Milenage
		if m, err = NewMilenage(K, OP, false); err != nil {
			t.Errorf("Failed to create a Milenage object")
		}
		if err := m.SetRand(RAND); err != nil {
			t.Errorf("err: %+v\n", err)
		}
		a, b, err := m.F1(SQN, AMF)
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
		res, ak := m.F2F5()
		ck := m.F3()
		ik := m.F4()
		akstar := m.F5star()

		if !reflect.DeepEqual(m.opc[:], eOPC) {
			t.Errorf("test generate OPC failed")
		}
		if !reflect.DeepEqual(a, f1) {
			t.Errorf("test F1 failed")
		}

		if !reflect.DeepEqual(b, f1star) {
			t.Errorf("test F1Star failed")
		}

		if !reflect.DeepEqual(res, eRES) {
			t.Errorf("test RES failed")
		}

		if !reflect.DeepEqual(ak, eAK) {
			t.Errorf("test AK failed")
		}
		if !reflect.DeepEqual(ck, eCK) {
			t.Errorf("test CK failed")
		}

		if !reflect.DeepEqual(ik, eIK) {
			t.Errorf("test RES failed")
		}

		if !reflect.DeepEqual(akstar, eAKstar) {
			t.Errorf("test AKstar failed")
		}

	}
	extra()
	fmt.Printf("TestMilenage is done\n")
}

func extra() {
	K, _ := hex.DecodeString("8BAF473F2F8FD09487CCCBD7097C6862")
	OPC, _ := hex.DecodeString("8E27B6AF0E692E750F32667A3B14605D")
	RAND, _ := hex.DecodeString("FEE14720A39CC164BCC0E3628452847C")
	//RAND, _ := hex.DecodeString("3D0029B53AF6BCA2CDC2276593B319F")
	SQN, _ := hex.DecodeString("0000000000af")
	AMF, _ := hex.DecodeString("8000")
	var m *Milenage
	var err error
	if m, err = NewMilenage(K, OPC, true); err != nil {
		log.Errorf("Failed to create a Milenage object")
	}
	if err := m.SetRand(RAND); err != nil {
		log.Errorf("err: %+v\n", err)
	}
	/*
		a, b, err := m.F1(SQN, AMF)
		if err != nil {
			t.Errorf("err: %+v\n", err)
		}
	*/
	_, ak := m.F2F5()
	_, b, _ := m.F1(SQN, AMF)
	log.Infof("ak=%x, macs=%x", ak, b)
}
