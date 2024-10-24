/**generated time: 2024-07-17 15:11:00.945119**/

package nas

/*******************************************************
 * UL NAS TRANSPORT
 ******************************************************/
type UlNasTransport struct {
	MmHeader
	PayloadContainerType        Uint8   //V [1/2]
	PayloadContainer            Bytes   //LV-E [3-65537]
	PduSessionId                *Uint8  //TV [12][2]
	OldPduSessionId             *Uint8  //TV [59][2]
	RequestType                 *Uint8  //TV [8-][1]
	SNssai                      *SNssai //TLV [22][3-10]
	Dnn                         *Dnn    //TLV [25][3-102]
	AdditionalInformation       *Bytes  //TLV [24][3-n]
	MaPduSessionInformation     *Uint8  //TV [A-][1]
	ReleaseAssistanceIndication *Uint8  //TV [F-][1]
}

func (msg *UlNasTransport) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding UlNasTransport", err)
		}
	}()
	var buf []byte
	//V[1/2]
	v := (uint8(msg.PayloadContainerType) & 0x0f) //fill righthalf
	// LV-E[3-65537]
	wire = append(wire, v)

	if buf, err = encodeLV(true, uint16(1), uint16(0), &msg.PayloadContainer); err != nil {
		err = nasError("encoding PayloadContainer [M LV-E 3-65537]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.PduSessionId != nil {
		//TV[2]
		wire = append(wire, []byte{0x12, uint8(*msg.PduSessionId)}...)
	}

	if msg.OldPduSessionId != nil {
		//TV[2]
		wire = append(wire, []byte{0x59, uint8(*msg.OldPduSessionId)}...)
	}

	if msg.RequestType != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.RequestType)&0x0f))
	}

	if msg.SNssai != nil {
		// TLV[3-10]
		if buf, err = encodeLV(false, uint16(1), uint16(8), msg.SNssai); err != nil {
			err = nasError("encoding SNssai [O TLV 3-10]", err)
			return
		}
		wire = append(append(wire, 0x22), buf...)
	}

	if msg.Dnn != nil {
		// TLV[3-102]
		if buf, err = encodeLV(false, uint16(1), uint16(100), msg.Dnn); err != nil {
			err = nasError("encoding Dnn [O TLV 3-102]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
	}

	if msg.AdditionalInformation != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.AdditionalInformation); err != nil {
			err = nasError("encoding AdditionalInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x24), buf...)
	}

	if msg.MaPduSessionInformation != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.MaPduSessionInformation)&0x0f))
	}

	if msg.ReleaseAssistanceIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0F<<4)|(uint8(*msg.ReleaseAssistanceIndication)&0x0f))
	}

	msg.msgType = UlNasTransportMsgType //set message type to UL NAS TRANSPORT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *UlNasTransport) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding UlNasTransport", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding PayloadContainerType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.PayloadContainerType = Uint8(0x0f & wire[offset]) //righthalf
	// LV-E[3-65537]
	offset++

	if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), &msg.PayloadContainer); err != nil {
		err = nasError("decoding PayloadContainer [M LV-E 3-65537]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x12: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding PduSessionId [C TV 2]", ErrIncomplete)
				return
			}
			msg.PduSessionId = new(Uint8)
			offset++ //consume IEI
			*msg.PduSessionId = Uint8(wire[offset])
			offset++
		case 0x59: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding OldPduSessionId [O TV 2]", ErrIncomplete)
				return
			}
			msg.OldPduSessionId = new(Uint8)
			offset++ //consume IEI
			*msg.OldPduSessionId = Uint8(wire[offset])
			offset++
		case 0x08: //TV[1]
			msg.RequestType = new(Uint8)
			*msg.RequestType = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x22: //TLV[3-10]
			offset++ //consume IEI
			v := &SNssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(8), v); err != nil {
				err = nasError("decoding SNssai [O TLV 3-10]", err)
				return
			}
			offset += consumed
			msg.SNssai = v
		case 0x25: //TLV[3-102]
			offset++ //consume IEI
			v := &Dnn{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(100), v); err != nil {
				err = nasError("decoding Dnn [O TLV 3-102]", err)
				return
			}
			offset += consumed
			msg.Dnn = v
		case 0x24: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding AdditionalInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.AdditionalInformation = v
		case 0x0A: //TV[1]
			msg.MaPduSessionInformation = new(Uint8)
			*msg.MaPduSessionInformation = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x0F: //TV[1]
			msg.ReleaseAssistanceIndication = new(Uint8)
			*msg.ReleaseAssistanceIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
