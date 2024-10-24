/**generated time: 2024-07-17 15:11:00.945448**/

package nas

/*******************************************************
 * DL NAS TRANSPORT
 ******************************************************/
type DlNasTransport struct {
	MmHeader
	PayloadContainerType  Uint8       //V [1/2]
	PayloadContainer      Bytes       //LV-E [3-65537]
	PduSessionId          *Uint8      //TV [12][2]
	AdditionalInformation *Bytes      //TLV [24][3-n]
	GmmCause              *Uint8      //TV [58][2]
	BackOffTimerValue     *GprsTimer3 //TLV [37][3]
	LowerBoundTimerValue  *GprsTimer3 //TLV [3A][3]
}

func (msg *DlNasTransport) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding DlNasTransport", err)
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

	if msg.AdditionalInformation != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.AdditionalInformation); err != nil {
			err = nasError("encoding AdditionalInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x24), buf...)
	}

	if msg.GmmCause != nil {
		//TV[2]
		wire = append(wire, []byte{0x58, uint8(*msg.GmmCause)}...)
	}

	if msg.BackOffTimerValue != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.BackOffTimerValue); err != nil {
			err = nasError("encoding BackOffTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x37), buf...)
	}

	if msg.LowerBoundTimerValue != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.LowerBoundTimerValue); err != nil {
			err = nasError("encoding LowerBoundTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x3A), buf...)
	}

	msg.msgType = DlNasTransportMsgType //set message type to DL NAS TRANSPORT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *DlNasTransport) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding DlNasTransport", err)
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
		case 0x24: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding AdditionalInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.AdditionalInformation = v
		case 0x58: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GmmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GmmCause = new(Uint8)
			offset++ //consume IEI
			*msg.GmmCause = Uint8(wire[offset])
			offset++
		case 0x37: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding BackOffTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.BackOffTimerValue = v
		case 0x3A: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding LowerBoundTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.LowerBoundTimerValue = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
