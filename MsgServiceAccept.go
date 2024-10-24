/**generated time: 2024-07-17 15:11:00.942305**/

package nas

/*******************************************************
 * SERVICE ACCEPT
 ******************************************************/
type ServiceAccept struct {
	MmHeader
	PduSessionStatus                                                            *PduSessionStatus                       //TLV [50][4-34]
	PduSessionReactivationResult                                                *PduSessionReactivationResult           //TLV [26][4-34]
	PduSessionReactivationResultErrorCause                                      *PduSessionReactivationResultErrorCause //TLV-E [72][5-515]
	EapMessage                                                                  *Bytes                                  //TLV-E [78][7-1503]
	T3448Value                                                                  *GprsTimer2                             //TLV [6B][3]
	AdditionalRequestResult                                                     *Uint8                                  //TLV [34][3]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming                    *TrackingAreaIdentityList               //TLV [1D][9-114]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService *TrackingAreaIdentityList               //TLV [1E][9-114]
}

func (msg *ServiceAccept) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding ServiceAccept", err)
		}
	}()
	var buf []byte
	if msg.PduSessionStatus != nil {
		// TLV[4-34]
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionStatus); err != nil {
			err = nasError("encoding PduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x50), buf...)
	}

	if msg.PduSessionReactivationResult != nil {
		// TLV[4-34]
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionReactivationResult); err != nil {
			err = nasError("encoding PduSessionReactivationResult [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x26), buf...)
	}

	if msg.PduSessionReactivationResultErrorCause != nil {
		// TLV-E[5-515]
		if buf, err = encodeLV(true, uint16(2), uint16(512), msg.PduSessionReactivationResultErrorCause); err != nil {
			err = nasError("encoding PduSessionReactivationResultErrorCause [O TLV-E 5-515]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	if msg.T3448Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3448Value); err != nil {
			err = nasError("encoding T3448Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6B), buf...)
	}

	if msg.AdditionalRequestResult != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.AdditionalRequestResult); err != nil {
			err = nasError("encoding AdditionalRequestResult [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x34), buf...)
	}

	if msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming != nil {
		// TLV[9-114]
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming); err != nil {
			err = nasError("encoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x1D), buf...)
	}

	if msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService != nil {
		// TLV[9-114]
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService); err != nil {
			err = nasError("encoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x1E), buf...)
	}

	msg.msgType = ServiceAcceptMsgType //set message type to SERVICE ACCEPT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *ServiceAccept) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding ServiceAccept", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x50: //TLV[4-34]
			offset++ //consume IEI
			v := &PduSessionStatus{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionStatus = v
		case 0x26: //TLV[4-34]
			offset++ //consume IEI
			v := &PduSessionReactivationResult{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionReactivationResult [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionReactivationResult = v
		case 0x72: //TLV-E[5-515]
			offset++ //consume IEI
			v := &PduSessionReactivationResultErrorCause{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(2), uint16(512), v); err != nil {
				err = nasError("decoding PduSessionReactivationResultErrorCause [O TLV-E 5-515]", err)
				return
			}
			offset += consumed
			msg.PduSessionReactivationResultErrorCause = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		case 0x6B: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer2{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3448Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3448Value = v
		case 0x34: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalRequestResult [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalRequestResult = v
		case 0x1D: //TLV[9-114]
			offset++ //consume IEI
			v := &TrackingAreaIdentityList{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(7), uint16(112), v); err != nil {
				err = nasError("decoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming [O TLV 9-114]", err)
				return
			}
			offset += consumed
			msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming = v
		case 0x1E: //TLV[9-114]
			offset++ //consume IEI
			v := &TrackingAreaIdentityList{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(7), uint16(112), v); err != nil {
				err = nasError("decoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService [O TLV 9-114]", err)
				return
			}
			offset += consumed
			msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
