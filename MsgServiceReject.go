/**generated time: 2024-12-16 16:36:18.693517**/

package nas

/*******************************************************
 * SERVICE REJECT
 ******************************************************/
type ServiceReject struct {
	MmHeader
	GmmCause                                                                    uint8                     //M: V [1]
	PduSessionStatus                                                            *PduSessionStatus         //O: TLV [50][4-34]
	T3346Value                                                                  *GprsTimer2               //O: TLV [5F][3]
	EapMessage                                                                  []byte                    //O: TLV-E [78][7-1503]
	T3448Value                                                                  *GprsTimer2               //O: TLV [6B][3]
	CagInformationList                                                          []byte                    //O: TLV-E [75][3-n]
	DisasterReturnWaitRange                                                     *uint16                   //O: TLV [2C][4]
	ExtendedCagInformationList                                                  []byte                    //O: TLV-E [71][3-n]
	LowerBoundTimerValue                                                        *GprsTimer3               //O: TLV [3A][3]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming                    *TrackingAreaIdentityList //O: TLV [1D][9-114]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService *TrackingAreaIdentityList //O: TLV [1E][9-114]
}

func (msg *ServiceReject) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding ServiceReject", err)
		}
	}()
	var buf []byte
	// M: V[1]
	wire = append(wire, uint8(msg.GmmCause))

	// O: TLV[4-34]
	if msg.PduSessionStatus != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionStatus); err != nil {
			err = nasError("encoding PduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x50), buf...)
	}

	// O: TLV[3]
	if msg.T3346Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3346Value); err != nil {
			err = nasError("encoding T3346Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x5F), buf...)
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

	// O: TLV[3]
	if msg.T3448Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3448Value); err != nil {
			err = nasError("encoding T3448Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6B), buf...)
	}

	// O: TLV-E[3-n]
	if len(msg.CagInformationList) > 0 {
		tmp := newBytesEncoder(msg.CagInformationList)
		if buf, err = encodeLV(true, uint16(0), uint16(0), tmp); err != nil {
			err = nasError("encoding CagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
	}

	// O: TLV[4]
	if msg.DisasterReturnWaitRange != nil {
		tmp := newUint16Encoder(*msg.DisasterReturnWaitRange)
		if buf, err = encodeLV(false, uint16(2), uint16(2), tmp); err != nil {
			err = nasError("encoding DisasterReturnWaitRange [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x2C), buf...)
	}

	// O: TLV-E[3-n]
	if len(msg.ExtendedCagInformationList) > 0 {
		tmp := newBytesEncoder(msg.ExtendedCagInformationList)
		if buf, err = encodeLV(true, uint16(0), uint16(0), tmp); err != nil {
			err = nasError("encoding ExtendedCagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	// O: TLV[3]
	if msg.LowerBoundTimerValue != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.LowerBoundTimerValue); err != nil {
			err = nasError("encoding LowerBoundTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x3A), buf...)
	}

	// O: TLV[9-114]
	if msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming != nil {
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming); err != nil {
			err = nasError("encoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x1D), buf...)
	}

	// O: TLV[9-114]
	if msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService != nil {
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService); err != nil {
			err = nasError("encoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x1E), buf...)
	}

	msg.msgType = ServiceRejectMsgType //set message type to SERVICE REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *ServiceReject) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding ServiceReject", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = wire[offset]
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x50: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(PduSessionStatus)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionStatus = v
		case 0x5F: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer2)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3346Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3346Value = v
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
		case 0x6B: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer2)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3448Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3448Value = v
		case 0x75: //O: TLV-E[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding CagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.CagInformationList = []byte(*v)
		case 0x2C: //O: TLV[4]
			offset++ //consume IEI
			v := new(uint16Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding DisasterReturnWaitRange [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.DisasterReturnWaitRange = (*uint16)(v)
		case 0x71: //O: TLV-E[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedCagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.ExtendedCagInformationList = []byte(*v)
		case 0x3A: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding LowerBoundTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.LowerBoundTimerValue = v
		case 0x1D: //O: TLV[9-114]
			offset++ //consume IEI
			v := new(TrackingAreaIdentityList)
			if consumed, err = decodeLV(wire[offset:], false, uint16(7), uint16(112), v); err != nil {
				err = nasError("decoding ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming [O TLV 9-114]", err)
				return
			}
			offset += consumed
			msg.ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming = v
		case 0x1E: //O: TLV[9-114]
			offset++ //consume IEI
			v := new(TrackingAreaIdentityList)
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
