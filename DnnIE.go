package nas

import (
	"fmt"
	"strings"
)

// 9.11.2.1B DNN [O TLV 3-102]
type Dnn struct {
	fqdn string
}

func NewDnn(fqdn string) *Dnn {
	return &Dnn{
		fqdn: fqdn,
	}
}
func (ie *Dnn) encode() (wire []byte, err error) {
	parts := strings.Split(ie.fqdn, ".")

	for _, part := range parts {
		// In RFC 1035 max length is 63, but in TS 23.003
		// including length octet
		if len(part) > 62 {
			err = fmt.Errorf("DNN limit the label to 62 octets or less")
			return
		}
		wire = append(wire, uint8(len(part)))
		wire = append(wire, []byte(part)...)
	}

	if len(wire) > 100 {
		err = fmt.Errorf("DNN should less then 100 octet")
	}
	return
}

func (ie *Dnn) decode(wire []byte) (err error) {
	parts := []string{}
	offset := 0
	wireLen := len(wire)
	for offset < wireLen {
		partLen := int(wire[offset])
		offset++
		if offset+partLen > wireLen {
			return ErrIncomplete
		}
		parts = append(parts, string(wire[offset:offset+partLen]))
		offset += partLen
	}
	ie.fqdn = strings.Join(parts, ".")
	return
}

func (ie *Dnn) String() string {
	return ie.fqdn
}
