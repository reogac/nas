/**generated time: 2024-07-17 15:11:00.943708**/

package nas

/*******************************************************
 * AUTHENTICATION REQUEST
 ******************************************************/
type AuthenticationRequest struct {
	MmHeader
	Ngksi                       KeySetIdentifier //V [1/2]
	Abba                        Bytes            //LV [3-n]
	AuthenticationParameterRand *Bytes           //TV [21][17]
	AuthenticationParameterAutn *Bytes           //TLV [20][18]
	EapMessage                  *Bytes           //TLV-E [78][7-1503]
}

func (msg *AuthenticationRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding AuthenticationRequest", err)
		}
	}()
	var buf []byte
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
	// LV[3-n]
	wire = append(wire, v)

	if buf, err = encodeLV(false, uint16(2), uint16(0), &msg.Abba); err != nil {
		err = nasError("encoding Abba [M LV 3-n]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.AuthenticationParameterRand != nil {
		// TV[17]
		if buf, err = msg.AuthenticationParameterRand.encode(); err != nil {
			err = nasError("encoding AuthenticationParameterRand [O TV 17]", err)
			return
		}
		if len(buf) != 16 {
			err = nasError("encoding AuthenticationParameterRand [O TV 17]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x21), buf...)
	}

	if msg.AuthenticationParameterAutn != nil {
		// TLV[18]
		if buf, err = encodeLV(false, uint16(16), uint16(16), msg.AuthenticationParameterAutn); err != nil {
			err = nasError("encoding AuthenticationParameterAutn [O TLV 18]", err)
			return
		}
		wire = append(append(wire, 0x20), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	msg.msgType = AuthenticationRequestMsgType //set message type to AUTHENTICATION REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *AuthenticationRequest) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding AuthenticationRequest", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding Ngksi [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.Ngksi.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	// LV[3-n]
	offset++

	if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(0), &msg.Abba); err != nil {
		err = nasError("decoding Abba [M LV 3-n]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x21: //TV[17]
			if offset+17 > wireLen {
				err = nasError("decoding AuthenticationParameterRand [O TV 17]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(Bytes)
			if err = v.decode(wire[offset : offset+16]); err != nil {
				err = nasError("decoding AuthenticationParameterRand [O TV 17]", err)
				return
			}
			msg.AuthenticationParameterRand = v
			offset += 16

		case 0x20: //TLV[18]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(16), uint16(16), v); err != nil {
				err = nasError("decoding AuthenticationParameterAutn [O TLV 18]", err)
				return
			}
			offset += consumed
			msg.AuthenticationParameterAutn = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
