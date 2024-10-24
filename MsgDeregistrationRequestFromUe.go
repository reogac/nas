/**generated time: 2024-07-17 15:11:00.941103**/

package nas

/*******************************************************
 * DEREGISTRATION REQUEST FROM UE
 ******************************************************/
type DeregistrationRequestFromUe struct {
	MmHeader
	DeRegistrationType DeRegistrationType //V [1/2]
	Ngksi              KeySetIdentifier   //V [1/2]
	MobileIdentity     MobileIdentity     //LV-E [6-n]
}

func (msg *DeregistrationRequestFromUe) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding DeregistrationRequestFromUe", err)
		}
	}()
	var buf []byte
	// V[1/2]
	if buf, err = msg.DeRegistrationType.encode(); err != nil {
		err = nasError("encoding DeRegistrationType [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding DeRegistrationType [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// V[1/2]
	if buf, err = msg.Ngksi.encode(); err != nil {
		err = nasError("encoding Ngksi [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding Ngksi [M V 1/2]", ErrInvalidSize)
		return
	}
	v |= (buf[0] & 0x0f) << 4 //fill lefthalf
	wire = append(wire, v)

	// LV-E[6-n]
	if buf, err = encodeLV(true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("encoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	wire = append(wire, buf...)

	msg.msgType = DeregistrationRequestFromUeMsgType //set message type to DEREGISTRATION REQUEST FROM UE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *DeregistrationRequestFromUe) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding DeregistrationRequestFromUe", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding DeRegistrationType [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.DeRegistrationType.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding DeRegistrationType [M V 1/2]", err)
		return
	}
	// V[1/2]
	if err = msg.Ngksi.decode([]byte{(0xf0 & wire[offset]) >> 4 /*lefthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	offset++

	// LV-E[6-n]
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("decoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	offset += consumed
	return
}
