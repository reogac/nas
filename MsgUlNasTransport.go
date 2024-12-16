/**generated time: 2024-12-16 16:36:18.695815**/

package nas

/*******************************************************
 * UL NAS TRANSPORT
 ******************************************************/
type UlNasTransport struct {
	MmHeader
	PayloadContainerType        uint8   //M: V [1/2]
	PayloadContainer            []byte  //M: LV-E [3-65537]
	PduSessionId                *uint8  //O: TV [12][2]
	OldPduSessionId             *uint8  //O: TV [59][2]
	RequestType                 *uint8  //O: TV [8-][1]
	SNssai                      *SNssai //O: TLV [22][3-10]
	Dnn                         *Dnn    //O: TLV [25][3-102]
	AdditionalInformation       []byte  //O: TLV [24][3-n]
	MaPduSessionInformation     *uint8  //O: TV [A-][1]
	ReleaseAssistanceIndication *uint8  //O: TV [F-][1]
}

func (msg *UlNasTransport) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding UlNasTransport", err)
		}
	}()
	var buf []byte
	//M: V[1/2]
	v := (uint8(msg.PayloadContainerType) & 0x0f) //fill righthalf
	// M: LV-E[3-65537]
	wire = append(wire, v)

	tmp := newBytesEncoder(msg.PayloadContainer)
	if buf, err = encodeLV(true, uint16(1), uint16(0), tmp); err != nil {
		err = nasError("encoding PayloadContainer [M LV-E 3-65537]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TV[2]
	if msg.PduSessionId != nil {
		wire = append(wire, []byte{0x12, uint8(*msg.PduSessionId)}...)
	}

	// O: TV[2]
	if msg.OldPduSessionId != nil {
		wire = append(wire, []byte{0x59, uint8(*msg.OldPduSessionId)}...)
	}

	// O: TV[1]
	if msg.RequestType != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.RequestType)&0x0f))
	}

	// O: TLV[3-10]
	if msg.SNssai != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(8), msg.SNssai); err != nil {
			err = nasError("encoding SNssai [O TLV 3-10]", err)
			return
		}
		wire = append(append(wire, 0x22), buf...)
	}

	// O: TLV[3-102]
	if msg.Dnn != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(100), msg.Dnn); err != nil {
			err = nasError("encoding Dnn [O TLV 3-102]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
	}

	// O: TLV[3-n]
	if len(msg.AdditionalInformation) > 0 {
		tmp := newBytesEncoder(msg.AdditionalInformation)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding AdditionalInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x24), buf...)
	}

	// O: TV[1]
	if msg.MaPduSessionInformation != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.MaPduSessionInformation)&0x0f))
	}

	// O: TV[1]
	if msg.ReleaseAssistanceIndication != nil {
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
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding PayloadContainerType [M V 1/2]", ErrIncomplete)
		return
	}
	msg.PayloadContainerType = 0x0f & wire[offset] //righthalf
	// M LV-E[3-65537]
	offset++

	v := new(bytesDecoder)
	if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
		err = nasError("decoding PayloadContainer [M LV-E 3-65537]", err)
		return
	}
	offset += consumed
	msg.PayloadContainer = []byte(*v)
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x12: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding PduSessionId [C TV 2]", ErrIncomplete)
				return
			}
			msg.PduSessionId = new(uint8)
			offset++ //consume IEI
			*msg.PduSessionId = wire[offset]
			offset++
		case 0x59: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding OldPduSessionId [O TV 2]", ErrIncomplete)
				return
			}
			msg.OldPduSessionId = new(uint8)
			offset++ //consume IEI
			*msg.OldPduSessionId = wire[offset]
			offset++
		case 0x08: //O: TV[1]
			msg.RequestType = new(uint8)
			*msg.RequestType = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x22: //O: TLV[3-10]
			offset++ //consume IEI
			v := new(SNssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(8), v); err != nil {
				err = nasError("decoding SNssai [O TLV 3-10]", err)
				return
			}
			offset += consumed
			msg.SNssai = v
		case 0x25: //O: TLV[3-102]
			offset++ //consume IEI
			v := new(Dnn)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(100), v); err != nil {
				err = nasError("decoding Dnn [O TLV 3-102]", err)
				return
			}
			offset += consumed
			msg.Dnn = v
		case 0x24: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding AdditionalInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.AdditionalInformation = []byte(*v)
		case 0x0A: //O: TV[1]
			msg.MaPduSessionInformation = new(uint8)
			*msg.MaPduSessionInformation = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x0F: //O: TV[1]
			msg.ReleaseAssistanceIndication = new(uint8)
			*msg.ReleaseAssistanceIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
