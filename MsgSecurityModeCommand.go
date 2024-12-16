/**generated time: 2024-12-16 16:36:18.695300**/

package nas

/*******************************************************
 * SECURITY MODE COMMAND
 ******************************************************/
type SecurityModeCommand struct {
	MmHeader
	SelectedNasSecurityAlgorithms    SecurityAlgorithms             //M: V [1]
	Ngksi                            KeySetIdentifier               //M: V [1/2]
	ReplayedUeSecurityCapabilities   UeSecurityCapability           //M: LV [3-9]
	ImeisvRequest                    *uint8                         //O: TV [E-][1]
	SelectedEpsNasSecurityAlgorithms *uint8                         //O: TV [57][2]
	AdditionalSecurityInformation    *AdditionalSecurityInformation //O: TLV [36][3]
	EapMessage                       []byte                         //O: TLV-E [78][7-1503]
	Abba                             []byte                         //O: TLV [38][4-n]
	ReplayedS1UeSecurityCapabilities []byte                         //O: TLV [19][4-7]
}

func (msg *SecurityModeCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding SecurityModeCommand", err)
		}
	}()
	var buf []byte
	// M: V[1]
	if buf, err = msg.SelectedNasSecurityAlgorithms.encode(); err != nil {
		err = nasError("encoding SelectedNasSecurityAlgorithms [M V 1]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding SelectedNasSecurityAlgorithms [M V 1]", ErrInvalidSize)
		return
	}
	wire = append(wire, buf[0])

	// M: V[1/2]
	if buf, err = msg.Ngksi.encode(); err != nil {
		err = nasError("encoding Ngksi [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding Ngksi [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// M: LV[3-9]
	wire = append(wire, v)

	if buf, err = encodeLV(false, uint16(2), uint16(8), &msg.ReplayedUeSecurityCapabilities); err != nil {
		err = nasError("encoding ReplayedUeSecurityCapabilities [M LV 3-9]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TV[1]
	if msg.ImeisvRequest != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0E<<4)|(uint8(*msg.ImeisvRequest)&0x0f))
	}

	// O: TV[2]
	if msg.SelectedEpsNasSecurityAlgorithms != nil {
		wire = append(wire, []byte{0x57, uint8(*msg.SelectedEpsNasSecurityAlgorithms)}...)
	}

	// O: TLV[3]
	if msg.AdditionalSecurityInformation != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.AdditionalSecurityInformation); err != nil {
			err = nasError("encoding AdditionalSecurityInformation [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x36), buf...)
	}

	// O: TLV-E[7-1503]
	if len(msg.EapMessage) > 0 {
		tmp := newBytesEncoder(msg.EapMessage)
		if buf, err = encodeLV(true, uint16(4), uint16(1500), tmp); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	// O: TLV[4-n]
	if len(msg.Abba) > 0 {
		tmp := newBytesEncoder(msg.Abba)
		if buf, err = encodeLV(false, uint16(2), uint16(0), tmp); err != nil {
			err = nasError("encoding Abba [O TLV 4-n]", err)
			return
		}
		wire = append(append(wire, 0x38), buf...)
	}

	// O: TLV[4-7]
	if len(msg.ReplayedS1UeSecurityCapabilities) > 0 {
		tmp := newBytesEncoder(msg.ReplayedS1UeSecurityCapabilities)
		if buf, err = encodeLV(false, uint16(2), uint16(5), tmp); err != nil {
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
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding SelectedNasSecurityAlgorithms [M V 1]", ErrIncomplete)
		return
	}
	if err = msg.SelectedNasSecurityAlgorithms.decode(wire[offset : offset+1]); err != nil {
		err = nasError("decoding SelectedNasSecurityAlgorithms [M V 1]", err)
		return
	}
	offset++

	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding Ngksi [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.Ngksi.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	// M LV[3-9]
	offset++

	if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(8), &msg.ReplayedUeSecurityCapabilities); err != nil {
		err = nasError("decoding ReplayedUeSecurityCapabilities [M LV 3-9]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x0E: //O: TV[1]
			msg.ImeisvRequest = new(uint8)
			*msg.ImeisvRequest = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x57: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding SelectedEpsNasSecurityAlgorithms [O TV 2]", ErrIncomplete)
				return
			}
			msg.SelectedEpsNasSecurityAlgorithms = new(uint8)
			offset++ //consume IEI
			*msg.SelectedEpsNasSecurityAlgorithms = wire[offset]
			offset++
		case 0x36: //O: TLV[3]
			offset++ //consume IEI
			v := new(AdditionalSecurityInformation)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalSecurityInformation [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalSecurityInformation = v
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
		case 0x38: //O: TLV[4-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(0), v); err != nil {
				err = nasError("decoding Abba [O TLV 4-n]", err)
				return
			}
			offset += consumed
			msg.Abba = []byte(*v)
		case 0x19: //O: TLV[4-7]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(5), v); err != nil {
				err = nasError("decoding ReplayedS1UeSecurityCapabilities [O TLV 4-7]", err)
				return
			}
			offset += consumed
			msg.ReplayedS1UeSecurityCapabilities = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
