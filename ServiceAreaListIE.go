package nas

import (
	"fmt"
)

// 9.11.3.49 Service area list [O TLV 6-114]
type ServiceAreaList struct {
	Lists []TaiListInf
}

func (l *ServiceAreaList) encode() (wire []byte, err error) {
	for _, inf := range l.Lists {
		wire = append(wire, inf.bytes()...)
	}
	return
}

func (l *ServiceAreaList) decode(wire []byte) (err error) {
	wireLen := len(wire)
	offset := 0
	var listType, numElems uint8
	var listSize int
	var listInf TaiListInf
	for offset < wireLen {
		listType = wire[offset] & 0x60 >> 5 //bit 6,5
		numElems = wire[offset]&0x1f + 1
		if numElems > 16 {
			err = fmt.Errorf("TAI list must have less than 16 elements")
			return
		}
		//log.Infof("num elems = %d", numElems)
		switch listType {
		case 0:
			listSize = 4 + 3*int(numElems)
			//log.Infof("List 0 len=%d", listSize)
			listInf = new(TaiListType0)
		case 1:
			listSize = 7
			//log.Infof("List 1 len=%d", listSize)
			listInf = new(TaiListType1)
		case 2:
			listSize = 1 + 6*int(numElems)
			//log.Infof("List 2 len=%d", listSize)
			listInf = new(TaiListType2)
		case 3:
			listSize = 4
			//log.Infof("List 3 len=%d", listSize)
			listInf = new(TaiListType3)

		default:
			err = fmt.Errorf("Unknown list type")
		}
		if err != nil {
			return
		}
		if offset+listSize > wireLen {
			err = ErrIncomplete
			return
		} else {
			listInf.readBytes(wire[offset : offset+listSize])
			l.Lists = append(l.Lists, listInf)
			offset += listSize
		}
	}
	return
}
