/** this file was generated at 2024-12-16 17:55:27.328245 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * SECURITY MODE COMPLETE
 ******************************************************/
type SecurityModeComplete struct {
	MmHeader
	Imeisv              *MobileIdentity //O: TLV-E [77][12]
	NasMessageContainer []byte          //O: TLV-E [71][4-n]
	NonImeisvPei        *MobileIdentity //O: TLV-E [78][7-n]
}

func (msg *SecurityModeComplete) encode() (wire []byte, err error) {
	var buf []byte
	// O: TLV-E[12]
	if msg.Imeisv != nil {
		if buf, err = encodeLV(true, uint16(9), uint16(9), msg.Imeisv); err != nil {
			err = nasError("encoding Imeisv [O TLV-E 12]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	// O: TLV-E[4-n]
	if len(msg.NasMessageContainer) > 0 {
		tmp := newBytesEncoder(msg.NasMessageContainer)
		if buf, err = encodeLV(true, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding NasMessageContainer [O TLV-E 4-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	// O: TLV-E[7-n]
	if msg.NonImeisvPei != nil {
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
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x77: //O: TLV-E[12]
			offset++ //consume IEI
			v := new(MobileIdentity)
			if consumed, err = decodeLV(wire[offset:], true, uint16(9), uint16(9), v); err != nil {
				err = nasError("decoding Imeisv [O TLV-E 12]", err)
				return
			}
			offset += consumed
			msg.Imeisv = v
		case 0x71: //O: TLV-E[4-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.NasMessageContainer = []byte(*v)
		case 0x78: //O: TLV-E[7-n]
			offset++ //consume IEI
			v := new(MobileIdentity)
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
