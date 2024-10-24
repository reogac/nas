package nas

import (
	//	"fmt"
	"testing"
)

func Test_PlmnId(t *testing.T) {
	var plmnId PlmnId
	var err error
	mcc := "108"
	mnc := "90"
	if err = plmnId.Set(mcc, mnc); err != nil {
		t.Errorf("Fail to set PlmnId: %+v", err)
	}
	rMcc, rMnc := plmnId.Get()
	if mcc != rMcc || mnc != rMnc {
		t.Errorf("Wrong decoded  PlmnId")
	}
}
