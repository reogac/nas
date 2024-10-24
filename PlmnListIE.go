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
