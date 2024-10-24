package nas

import (
	"fmt"
)

// 9.11.3.43 PDU session reactivation result error cause [O TLV-E 5-515]
type PduSessionReactivationResultErrorCause struct {
	bytes []byte
}

func (ie *PduSessionReactivationResultErrorCause) encode() (wire []byte, err error) {
	wire = ie.bytes
	return
}

func (ie *PduSessionReactivationResultErrorCause) decode(wire []byte) (err error) {
	if len(wire)%2 == 1 {
		return fmt.Errorf("Expect even number of bytes")
	}
	ie.bytes = make([]byte, len(wire))
	copy(ie.bytes, wire)
	return
}

func (ie *PduSessionReactivationResultErrorCause) Set(idList, causeList []uint8) (err error) {
	if len(idList) != len(causeList) {
		return fmt.Errorf("Mismatched id list and cause list")
	}
	ie.bytes = make([]byte, 2*len(idList))
	for i, id := range idList {
		ie.bytes[2*i] = id
		ie.bytes[2*i+1] = causeList[i]
	}
	return
}
