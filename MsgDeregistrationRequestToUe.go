/**generated time: 2024-12-16 16:36:18.693017**/

package nas

/*******************************************************
 * DEREGISTRATION REQUEST TO UE
 ******************************************************/
type DeregistrationRequestToUe struct {
	MmHeader
	DeRegistrationType                                                          DeRegistrationType        //M: V [1/2]
	GmmCause                                                                    *uint8                    //O: TV [58][2]
	T3346Value                                                                  *GprsTimer2               //O: TLV [5F][3]
	RejectedNssai                                                               *RejectedNssai            //O: TLV [6D][4-42]
	CagInformationList                                                          []byte                    //O: TLV-E [75][3-n]
	ExtendedRejectedNssai                                                       []byte                    //O: TLV [68][5-90]
	DisasterReturnWaitRange                                                     *uint16                   //O: TLV [2C][4]
	ExtendedCagInformationList                                                  []byte                    //O: TLV-E [71][3-n]
	LowerBoundTimerValue                                                        *GprsTimer3               //O: TLV [3A][3]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming                    *TrackingAreaIdentityList //O: TLV [1D][9-114]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService *TrackingAreaIdentityList //O: TLV [1E][9-114]
}

func (msg *DeregistrationRequestToUe) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding DeregistrationRequestToUe", err)
		}
	}()
	var buf []byte
	// M: V[1/2]
	if buf, err = msg.DeRegistrationType.encode(); err != nil {
		err = nasError("encoding DeRegistrationType [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding DeRegistrationType [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	wire = append(wire, v)

	// O: TV[2]
	if msg.GmmCause != nil {
		wire = append(wire, []byte{0x58, uint8(*msg.GmmCause)}...)
	}

	// O: TLV[3]
	if msg.T3346Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3346Value); err != nil {
			err = nasError("encoding T3346Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x5F), buf...)
	}

	// O: TLV[4-42]
	if msg.RejectedNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(40), msg.RejectedNssai); err != nil {
			err = nasError("encoding RejectedNssai [O TLV 4-42]", err)
			return
		}
		wire = append(append(wire, 0x6D), buf...)
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

	// O: TLV[5-90]
	if len(msg.ExtendedRejectedNssai) > 0 {
		tmp := newBytesEncoder(msg.ExtendedRejectedNssai)
		if buf, err = encodeLV(false, uint16(3), uint16(88), tmp); err != nil {
			err = nasError("encoding ExtendedRejectedNssai [O TLV 5-90]", err)
			return
		}
		wire = append(append(wire, 0x68), buf...)
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

	msg.msgType = DeregistrationRequestToUeMsgType //set message type to DEREGISTRATION REQUEST TO UE
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *DeregistrationRequestToUe) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding DeregistrationRequestToUe", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding DeRegistrationType [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.DeRegistrationType.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding DeRegistrationType [M V 1/2]", err)
		return
	}
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x58: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding GmmCause [O TV 2]", ErrIncomplete)
				return
			}
			msg.GmmCause = new(uint8)
			offset++ //consume IEI
			*msg.GmmCause = wire[offset]
			offset++
		case 0x5F: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer2)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3346Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3346Value = v
		case 0x6D: //O: TLV[4-42]
			offset++ //consume IEI
			v := new(RejectedNssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(40), v); err != nil {
				err = nasError("decoding RejectedNssai [O TLV 4-42]", err)
				return
			}
			offset += consumed
			msg.RejectedNssai = v
		case 0x75: //O: TLV-E[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding CagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.CagInformationList = []byte(*v)
		case 0x68: //O: TLV[5-90]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(88), v); err != nil {
				err = nasError("decoding ExtendedRejectedNssai [O TLV 5-90]", err)
				return
			}
			offset += consumed
			msg.ExtendedRejectedNssai = []byte(*v)
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
