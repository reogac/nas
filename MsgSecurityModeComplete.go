/**generated time: 2024-07-17 15:11:00.944726**/

package nas

/*******************************************************
 * SECURITY MODE COMPLETE
 ******************************************************/
type SecurityModeComplete struct {
	MmHeader
	Imeisv              *MobileIdentity //TLV-E [77][12]
	NasMessageContainer *Bytes          //TLV-E [71][4-n]
	NonImeisvPei        *MobileIdentity //TLV-E [78][7-n]
}

func (msg *SecurityModeComplete) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding SecurityModeComplete", err)
		}
	}()
	var buf []byte
	if msg.Imeisv != nil {
		// TLV-E[12]
		if buf, err = encodeLV(true, uint16(9), uint16(9), msg.Imeisv); err != nil {
			err = nasError("encoding Imeisv [O TLV-E 12]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	if msg.NasMessageContainer != nil {
		// TLV-E[4-n]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.NasMessageContainer); err != nil {
			err = nasError("encoding NasMessageContainer [O TLV-E 4-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	if msg.NonImeisvPei != nil {
		// TLV-E[7-n]
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.NonImeisvPei); err != nil {
			err = nasError("encoding NonImeisvPei [O TLV-E 7-n]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	msg.msgType = SecurityModeCompleteMsgType //set message type to SECURITY MODE COMPLETE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *SecurityModeComplete) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding SecurityModeComplete", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x77: //TLV-E[12]
			offset++ //consume IEI
			v := &MobileIdentity{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(9), uint16(9), v); err != nil {
				err = nasError("decoding Imeisv [O TLV-E 12]", err)
				return
			}
			offset += consumed
			msg.Imeisv = v
		case 0x71: //TLV-E[4-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.NasMessageContainer = v
		case 0x78: //TLV-E[7-n]
			offset++ //consume IEI
			v := &MobileIdentity{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding NonImeisvPei [O TLV-E 7-n]", err)
				return
			}
			offset += consumed
			msg.NonImeisvPei = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
