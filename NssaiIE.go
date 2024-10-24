package nas

import (
	"fmt"
)

// 9.11.3.37 NSSAI [O TLV 4-74]
type Nssai struct {
	List []SNssai
}

func (ie *Nssai) Add(slice *SNssai) {
	ie.List = append(ie.List, *slice)
}

func (ie *Nssai) encode() (wire []byte, err error) {
	var buf []byte
	for i, s := range ie.List {
		if i >= 16 { //no more than 16 slices
			break
		}
		if buf, err = s.encode(); err != nil {
			err = fmt.Errorf("encode element %d fails: %+v", i, err)
			return
		}
		wire = append(wire, uint8(len(buf)))
		wire = append(wire, buf...)
	}

	return
}

func (ie *Nssai) decode(wire []byte) (err error) {
	offset := 0
	wireLen := len(wire)
	i := 0
	for offset < wireLen {
		if i >= 16 {
			break //no more than 16 slices
		}

		sliceSize := int(wire[offset]) + 1 //including length byte
		if offset+sliceSize > wireLen {
			return ErrIncomplete
		}
		slice := new(SNssai)
		if err = slice.decode(wire[offset+1 : offset+sliceSize]); err != nil {
			err = fmt.Errorf("decode element %d fails: %+v", i, err)
			return
		}
		ie.List = append(ie.List, *slice)
		i++
		offset += sliceSize
	}
	return
}
