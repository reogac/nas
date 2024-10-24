/**generated time: 2024-07-17 15:11:00.944418**/

package nas

/*******************************************************
 * SECURITY MODE COMMAND
 ******************************************************/
type SecurityModeCommand struct {
	MmHeader
	SelectedNasSecurityAlgorithms    SecurityAlgorithms             //V [1]
	Ngksi                            KeySetIdentifier               //V [1/2]
	ReplayedUeSecurityCapabilities   UeSecurityCapability           //LV [3-9]
	ImeisvRequest                    *Uint8                         //TV [E-][1]
	SelectedEpsNasSecurityAlgorithms *Uint8                         //TV [57][2]
	AdditionalSecurityInformation    *AdditionalSecurityInformation //TLV [36][3]
	EapMessage                       *Bytes                         //TLV-E [78][7-1503]
	Abba                             *Bytes                         //TLV [38][4-n]
	ReplayedS1UeSecurityCapabilities *Bytes                         //TLV [19][4-7]
}

func (msg *SecurityModeCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding SecurityModeCommand", err)
		}
	}()
	var buf []byte
	// V[1]
	if buf, err = msg.SelectedNasSecurityAlgorithms.encode(); err != nil {
		err = nasError("encoding SelectedNasSecurityAlgorithms [M V 1]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding SelectedNasSecurityAlgorithms [M V 1]", ErrInvalidSize)
		return
	}
	wire = append(wire, buf[0])

	// V[1/2]
	if buf, err = msg.Ngksi.encode(); err != nil {
		err = nasError("encoding Ngksi [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding Ngksi [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// LV[3-9]
	wire = append(wire, v)

	if buf, err = encodeLV(false, uint16(2), uint16(8), &msg.ReplayedUeSecurityCapabilities); err != nil {
		err = nasError("encoding ReplayedUeSecurityCapabilities [M LV 3-9]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.ImeisvRequest != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0E<<4)|(uint8(*msg.ImeisvRequest)&0x0f))
	}

	if msg.SelectedEpsNasSecurityAlgorithms != nil {
		//TV[2]
		wire = append(wire, []byte{0x57, uint8(*msg.SelectedEpsNasSecurityAlgorithms)}...)
	}

	if msg.AdditionalSecurityInformation != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.AdditionalSecurityInformation); err != nil {
			err = nasError("encoding AdditionalSecurityInformation [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x36), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	if msg.Abba != nil {
		// TLV[4-n]
		if buf, err = encodeLV(false, uint16(2), uint16(0), msg.Abba); err != nil {
			err = nasError("encoding Abba [O TLV 4-n]", err)
			return
		}
		wire = append(append(wire, 0x38), buf...)
	}

	if msg.ReplayedS1UeSecurityCapabilities != nil {
		// TLV[4-7]
		if buf, err = encodeLV(false, uint16(2), uint16(5), msg.ReplayedS1UeSecurityCapabilities); err != nil {
			err = nasError("encoding ReplayedS1UeSecurityCapabilities [O TLV 4-7]", err)
			return
		}
		wire = append(append(wire, 0x19), buf...)
	}

	msg.msgType = SecurityModeCommandMsgType //set message type to SECURITY MODE COMMAND
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *SecurityModeCommand) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding SecurityModeCommand", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1]
	if offset+1 > wireLen {
		err = nasError("decoding SelectedNasSecurityAlgorithms [M V 1]", ErrIncomplete)
		return
	}
	if err = msg.SelectedNasSecurityAlgorithms.decode(wire[offset : offset+1]); err != nil {
		err = nasError("decoding SelectedNasSecurityAlgorithms [M V 1]", err)
		return
	}
	offset++

	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding Ngksi [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.Ngksi.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	// LV[3-9]
	offset++

	if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(8), &msg.ReplayedUeSecurityCapabilities); err != nil {
		err = nasError("decoding ReplayedUeSecurityCapabilities [M LV 3-9]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x0E: //TV[1]
			msg.ImeisvRequest = new(Uint8)
			*msg.ImeisvRequest = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x57: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding SelectedEpsNasSecurityAlgorithms [O TV 2]", ErrIncomplete)
				return
			}
			msg.SelectedEpsNasSecurityAlgorithms = new(Uint8)
			offset++ //consume IEI
			*msg.SelectedEpsNasSecurityAlgorithms = Uint8(wire[offset])
			offset++
		case 0x36: //TLV[3]
			offset++ //consume IEI
			v := &AdditionalSecurityInformation{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalSecurityInformation [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalSecurityInformation = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		case 0x38: //TLV[4-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(0), v); err != nil {
				err = nasError("decoding Abba [O TLV 4-n]", err)
				return
			}
			offset += consumed
			msg.Abba = v
		case 0x19: //TLV[4-7]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(5), v); err != nil {
				err = nasError("decoding ReplayedS1UeSecurityCapabilities [O TLV 4-7]", err)
				return
			}
			offset += consumed
			msg.ReplayedS1UeSecurityCapabilities = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
