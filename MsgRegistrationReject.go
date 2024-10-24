/**generated time: 2024-07-17 15:11:00.940682**/

package nas

/*******************************************************
 * REGISTRATION REJECT
 ******************************************************/
type RegistrationReject struct {
	MmHeader
	GmmCause                                                                    Uint8                     //V [1]
	T3346Value                                                                  *GprsTimer2               //TLV [5F][3]
	T3502Value                                                                  *GprsTimer2               //TLV [16][3]
	EapMessage                                                                  *Bytes                    //TLV-E [78][7-1503]
	RejectedNssai                                                               *RejectedNssai            //TLV [69][4-42]
	CagInformationList                                                          *Bytes                    //TLV-E [75][3-n]
	ExtendedRejectedNssai                                                       *Bytes                    //TLV [68][5-90]
	DisasterReturnWaitRange                                                     *Uint16                   //TLV [2C][4]
	ExtendedCagInformationList                                                  *Bytes                    //TLV-E [71][3-n]
	LowerBoundTimerValue                                                        *GprsTimer3               //TLV [3A][3]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming                    *TrackingAreaIdentityList //TLV [1D][9-114]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService *TrackingAreaIdentityList //TLV [1E][9-114]
}

func (msg *RegistrationReject) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding RegistrationReject", err)
		}
	}()
	var buf []byte
	// V[1]
	wire = append(wire, uint8(msg.GmmCause))

	if msg.T3346Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3346Value); err != nil {
			err = nasError("encoding T3346Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x5F), buf...)
	}

	if msg.T3502Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3502Value); err != nil {
			err = nasError("encoding T3502Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x16), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	if msg.RejectedNssai != nil {
		// TLV[4-42]
		if buf, err = encodeLV(false, uint16(2), uint16(40), msg.RejectedNssai); err != nil {
			err = nasError("encoding RejectedNssai [O TLV 4-42]", err)
			return
		}
		wire = append(append(wire, 0x69), buf...)
	}

	if msg.CagInformationList != nil {
		// TLV-E[3-n]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.CagInformationList); err != nil {
			err = nasError("encoding CagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
	}

	if msg.ExtendedRejectedNssai != nil {
		// TLV[5-90]
		if buf, err = encodeLV(false, uint16(3), uint16(88), msg.ExtendedRejectedNssai); err != nil {
			err = nasError("encoding ExtendedRejectedNssai [O TLV 5-90]", err)
			return
		}
		wire = append(append(wire, 0x68), buf...)
	}

	if msg.DisasterReturnWaitRange != nil {
		// TLV[4]
		if buf, err = encodeLV(false, uint16(2), uint16(2), msg.DisasterReturnWaitRange); err != nil {
			err = nasError("encoding DisasterReturnWaitRange [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x2C), buf...)
	}

	if msg.ExtendedCagInformationList != nil {
		// TLV-E[3-n]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.ExtendedCagInformationList); err != nil {
			err = nasError("encoding ExtendedCagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	if msg.LowerBoundTimerValue != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.LowerBoundTimerValue); err != nil {
			err = nasError("encoding LowerBoundTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x3A), buf...)
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

	msg.msgType = RegistrationRejectMsgType //set message type to REGISTRATION REJECT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *RegistrationReject) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding RegistrationReject", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1]
	if offset+1 > wireLen {
		err = nasError("decoding GmmCause [M V 1]", ErrIncomplete)
		return
	}
	msg.GmmCause = Uint8(wire[offset])
	offset++

	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x5F: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer2{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3346Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3346Value = v
		case 0x16: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer2{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3502Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3502Value = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		case 0x69: //TLV[4-42]
			offset++ //consume IEI
			v := &RejectedNssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(40), v); err != nil {
				err = nasError("decoding RejectedNssai [O TLV 4-42]", err)
				return
			}
			offset += consumed
			msg.RejectedNssai = v
		case 0x75: //TLV-E[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding CagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.CagInformationList = v
		case 0x68: //TLV[5-90]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(88), v); err != nil {
				err = nasError("decoding ExtendedRejectedNssai [O TLV 5-90]", err)
				return
			}
			offset += consumed
			msg.ExtendedRejectedNssai = v
		case 0x2C: //TLV[4]
			offset++ //consume IEI
			v := new(Uint16)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding DisasterReturnWaitRange [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.DisasterReturnWaitRange = v
		case 0x71: //TLV-E[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedCagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.ExtendedCagInformationList = v
		case 0x3A: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding LowerBoundTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.LowerBoundTimerValue = v
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
