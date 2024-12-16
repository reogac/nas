/** this file was generated at 2024-12-16 17:55:27.326048 by tqtung@etri.re.kr **/

package nas

/*******************************************************
 * SERVICE REQUEST
 ******************************************************/
type ServiceRequest struct {
	MmHeader
	Ngksi                   KeySetIdentifier         //M: V [1/2]
	ServiceType             uint8                    //M: V [1/2]
	STmsi                   MobileIdentity           //M: LV-E [9]
	UplinkDataStatus        *UplinkDataStatus        //O: TLV [40][4-34]
	PduSessionStatus        *PduSessionStatus        //O: TLV [50][4-34]
	AllowedPduSessionStatus *AllowedPduSessionStatus //O: TLV [25][4-34]
	NasMessageContainer     []byte                   //O: TLV-E [71][4-n]
	UeRequestType           *uint8                   //O: TLV [29][3]
	PagingRestriction       []byte                   //O: TLV [28][3-35]
}

func (msg *ServiceRequest) encode() (wire []byte, err error) {
	var buf []byte
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
	// M: V[1/2]
	v |= (uint8(msg.ServiceType) & 0x0f) << 4 //fill lefthalf
	wire = append(wire, v)

	// M: LV-E[9]
	if buf, err = encodeLV(true, uint16(7), uint16(7), &msg.STmsi); err != nil {
		err = nasError("encoding STmsi [M LV-E 9]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TLV[4-34]
	if msg.UplinkDataStatus != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.UplinkDataStatus); err != nil {
			err = nasError("encoding UplinkDataStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x40), buf...)
	}

	// O: TLV[4-34]
	if msg.PduSessionStatus != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionStatus); err != nil {
			err = nasError("encoding PduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x50), buf...)
	}

	// O: TLV[4-34]
	if msg.AllowedPduSessionStatus != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.AllowedPduSessionStatus); err != nil {
			err = nasError("encoding AllowedPduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
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

	// O: TLV[3]
	if msg.UeRequestType != nil {
		tmp := newUint8Encoder(*msg.UeRequestType)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding UeRequestType [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
	}

	// O: TLV[3-35]
	if len(msg.PagingRestriction) > 0 {
		tmp := newBytesEncoder(msg.PagingRestriction)
		if buf, err = encodeLV(false, uint16(1), uint16(33), tmp); err != nil {
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
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding Ngksi [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.Ngksi.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	// M V[1/2]
	msg.ServiceType = (0xf0 & wire[offset]) >> 4 //lefthalf
	offset++

	// M LV-E[9]
	if consumed, err = decodeLV(wire[offset:], true, uint16(7), uint16(7), &msg.STmsi); err != nil {
		err = nasError("decoding STmsi [M LV-E 9]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x40: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(UplinkDataStatus)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding UplinkDataStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.UplinkDataStatus = v
		case 0x50: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(PduSessionStatus)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionStatus = v
		case 0x25: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(AllowedPduSessionStatus)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding AllowedPduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.AllowedPduSessionStatus = v
		case 0x71: //O: TLV-E[4-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.NasMessageContainer = []byte(*v)
		case 0x29: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UeRequestType [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UeRequestType = (*uint8)(v)
		case 0x28: //O: TLV[3-35]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(33), v); err != nil {
				err = nasError("decoding PagingRestriction [O TLV 3-35]", err)
				return
			}
			offset += consumed
			msg.PagingRestriction = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
