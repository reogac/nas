/**generated time: 2024-07-17 15:11:00.941613**/

package nas

/*******************************************************
 * SERVICE REQUEST
 ******************************************************/
type ServiceRequest struct {
	MmHeader
	Ngksi                   KeySetIdentifier         //V [1/2]
	ServiceType             Uint8                    //V [1/2]
	STmsi                   MobileIdentity           //LV-E [9]
	UplinkDataStatus        *UplinkDataStatus        //TLV [40][4-34]
	PduSessionStatus        *PduSessionStatus        //TLV [50][4-34]
	AllowedPduSessionStatus *AllowedPduSessionStatus //TLV [25][4-34]
	NasMessageContainer     *Bytes                   //TLV-E [71][4-n]
	UeRequestType           *Uint8                   //TLV [29][3]
	PagingRestriction       *Bytes                   //TLV [28][3-35]
}

func (msg *ServiceRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding ServiceRequest", err)
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
	// V[1/2]
	v |= (uint8(msg.ServiceType) & 0x0f) << 4 //fill lefthalf
	wire = append(wire, v)

	// LV-E[9]
	if buf, err = encodeLV(true, uint16(7), uint16(7), &msg.STmsi); err != nil {
		err = nasError("encoding STmsi [M LV-E 9]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.UplinkDataStatus != nil {
		// TLV[4-34]
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.UplinkDataStatus); err != nil {
			err = nasError("encoding UplinkDataStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x40), buf...)
	}

	if msg.PduSessionStatus != nil {
		// TLV[4-34]
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionStatus); err != nil {
			err = nasError("encoding PduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x50), buf...)
	}

	if msg.AllowedPduSessionStatus != nil {
		// TLV[4-34]
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.AllowedPduSessionStatus); err != nil {
			err = nasError("encoding AllowedPduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
	}

	if msg.NasMessageContainer != nil {
		// TLV-E[4-n]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.NasMessageContainer); err != nil {
			err = nasError("encoding NasMessageContainer [O TLV-E 4-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	if msg.UeRequestType != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.UeRequestType); err != nil {
			err = nasError("encoding UeRequestType [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
	}

	if msg.PagingRestriction != nil {
		// TLV[3-35]
		if buf, err = encodeLV(false, uint16(1), uint16(33), msg.PagingRestriction); err != nil {
			err = nasError("encoding PagingRestriction [O TLV 3-35]", err)
			return
		}
		wire = append(append(wire, 0x28), buf...)
	}

	msg.msgType = ServiceRequestMsgType //set message type to SERVICE REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *ServiceRequest) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding ServiceRequest", err)
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
	// V[1/2]
	msg.ServiceType = Uint8((0xf0 & wire[offset]) >> 4) //lefthalf
	offset++

	// LV-E[9]
	if consumed, err = decodeLV(wire[offset:], true, uint16(7), uint16(7), &msg.STmsi); err != nil {
		err = nasError("decoding STmsi [M LV-E 9]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x40: //TLV[4-34]
			offset++ //consume IEI
			v := &UplinkDataStatus{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding UplinkDataStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.UplinkDataStatus = v
		case 0x50: //TLV[4-34]
			offset++ //consume IEI
			v := &PduSessionStatus{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionStatus = v
		case 0x25: //TLV[4-34]
			offset++ //consume IEI
			v := &AllowedPduSessionStatus{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding AllowedPduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.AllowedPduSessionStatus = v
		case 0x71: //TLV-E[4-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.NasMessageContainer = v
		case 0x29: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UeRequestType [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UeRequestType = v
		case 0x28: //TLV[3-35]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(33), v); err != nil {
				err = nasError("decoding PagingRestriction [O TLV 3-35]", err)
				return
			}
			offset += consumed
			msg.PagingRestriction = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
