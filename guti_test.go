package nas

import (
	//	"fmt"
	"testing"
)

func Test_Guti(t *testing.T) {
	gutis := []string{
		"1209011111100000001",
		"12090911111180000001",
		"12090911111180000001",
		"1209011111109000001",
	}
	var err error
	var wire []byte
	for _, gutiStr := range gutis {
		guti := new(Guti)
		if err = guti.Parse(gutiStr); err != nil {
			t.Errorf("Parse GUTI fail: %+v", err)
			continue
		}
		if wire, err = guti.encode(); err != nil {
			t.Errorf("encode GUTI fail: %+v", err)
			continue
		}
		newGuti := new(Guti)
		if err = newGuti.decode(wire); err != nil {
			t.Errorf("decode GUTI fail: %+v", err)
			continue
		}
		if newGuti.String() != gutiStr {
			t.Errorf("not equal %s != %s", gutiStr, newGuti.String())
			continue
		}
		tmsi5gsStr := guti.Tmsi5GS()
		tmsi5gs := new(Tmsi5Gs)
		if err = tmsi5gs.Parse(tmsi5gsStr); err != nil {
			t.Errorf("Fail to parse Tmsi5Gs[%s] :%+v", tmsi5gsStr, err)
			continue
		}
		if newTmsi5gsStr := tmsi5gs.String(); newTmsi5gsStr != tmsi5gsStr {
			t.Errorf("tmsi5Gs not equal: %s != %s", tmsi5gsStr, newTmsi5gsStr)
			continue
		}
	}
}
