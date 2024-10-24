package nas

import (
	"testing"
)

func Test_Tailist(t *testing.T) {
	log.Infof("Test tailist encoding/decoding")
	var plmnId PlmnId
	plmnId.Parse("10020")
	list1 := &TaiListType0{
		PlmnId: plmnId,
	}
	list2 := &TaiListType1{
		PlmnId: plmnId,
		Size:   30,
		Start:  20,
	}
	list3 := &TaiListType2{}

	for i := 0; i < 20; i++ {
		list3.List = append(list3.List, TrackingAreaIdentity{
			PlmnId: plmnId,
			Tac:    uint32(i + 1),
		})
		list1.TacList = append(list1.TacList, uint32(i+1))
	}

	//test tracking area identity list
	lists := TrackingAreaIdentityList{
		Lists: []TaiListInf{
			list1,
			list2,
			list3,
		},
	}
	var wire []byte
	var err error
	if wire, err = lists.encode(); err != nil {
		t.Errorf("fail to encode list of tai-lists: %+v", err)
		return
	}
	var newLists TrackingAreaIdentityList
	if err = newLists.decode(wire); err != nil {
		t.Errorf("fail to decode list of tai-lists: %+v", err)
		return
	}
	if len(newLists.Lists) != 3 {
		t.Errorf("wrong number of lists")
		return
	}
	if newLists.Lists[0].GetType() != uint8(0) || newLists.Lists[1].GetType() != uint8(1) || newLists.Lists[2].GetType() != uint8(2) {
		t.Errorf("wrong type of list")
		return
	}

	//test service area list
	list4 := &TaiListType3{
		PlmnId:  plmnId,
		Allowed: true,
	}
	list1.Allowed = false
	list2.Allowed = true
	list3.Allowed = false

	sLists := ServiceAreaList{
		Lists: []TaiListInf{
			list1,
			list2,
			list3,
			list4,
		},
	}
	if wire, err = sLists.encode(); err != nil {
		t.Errorf("fail to encode list of tai-lists: %+v", err)
		return
	}
	var newSLists ServiceAreaList
	if err = newSLists.decode(wire); err != nil {
		t.Errorf("fail to decode list of tai-lists: %+v", err)
		return
	}
	if len(newSLists.Lists) != 4 {
		t.Errorf("wrong number of service are lists")
		return
	}
	if newSLists.Lists[0].GetType() != uint8(0) || newSLists.Lists[1].GetType() != uint8(1) || newSLists.Lists[2].GetType() != uint8(2) || newSLists.Lists[3].GetType() != uint8(3) {
		t.Errorf("wrong type of service are list")
		return
	}
	newList1 := newSLists.Lists[0].(*TaiListType0)
	newList2 := newSLists.Lists[1].(*TaiListType1)
	newList3 := newSLists.Lists[2].(*TaiListType2)
	newList4 := newSLists.Lists[3].(*TaiListType3)
	if newList1.Allowed || !newList2.Allowed && newList3.Allowed && !newList4.Allowed {
		t.Errorf("wrong value of allowance bit")
	}
}
