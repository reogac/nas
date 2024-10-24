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
