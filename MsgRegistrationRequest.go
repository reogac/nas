/**generated time: 2024-07-17 15:11:00.937546**/

package nas

/*******************************************************
 * REGISTRATION REQUEST
 ******************************************************/
type RegistrationRequest struct {
	MmHeader
	RegistrationType                      RegistrationType         //V [1/2]
	Ngksi                                 KeySetIdentifier         //V [1/2]
	MobileIdentity                        MobileIdentity           //LV-E [6-n]
	NonCurrentNativeNasKeySetIdentifier   *KeySetIdentifier        //TV [C-][1]
	GmmCapability                         *GmmCapability           //TLV [10][3-15]
	UeSecurityCapability                  *UeSecurityCapability    //TLV [2E][4-10]
	RequestedNssai                        *Nssai                   //TLV [2F][4-74]
	LastVisitedRegisteredTai              *TrackingAreaIdentity    //TV [52][7]
	S1UeNetworkCapability                 *Bytes                   //TLV [17][4-15]
	UplinkDataStatus                      *UplinkDataStatus        //TLV [40][4-34]
	PduSessionStatus                      *PduSessionStatus        //TLV [50][4-34]
	MicoIndication                        *Uint8                   //TV [B-][1]
	UeStatus                              *Uint8                   //TLV [2B][3]
	AdditionalGuti                        *MobileIdentity          //TLV-E [77][14]
	AllowedPduSessionStatus               *AllowedPduSessionStatus //TLV [25][4-34]
	UeUsageSetting                        *Uint8                   //TLV [18][3]
	RequestedDrxParameters                *Uint8                   //TLV [51][3]
	EpsNasMessageContainer                *Bytes                   //TLV-E [70][4-n]
	LadnIndication                        *Bytes                   //TLV-E [74][3-811]
	PayloadContainerType                  *Uint8                   //TV [8-][1]
	PayloadContainer                      *Bytes                   //TLV-E [7B][4-65538]
	NetworkSlicingIndication              *Uint8                   //TV [9-][1]
	UpdateType                            *Uint8                   //TLV [53][3]
	MobileStationClassmark2               *Bytes                   //TLV [41][5]
	SupportedCodecs                       *Bytes                   //TLV [42][5-n]
	NasMessageContainer                   *Bytes                   //TLV-E [71][4-n]
	EpsBearerContextStatus                *Uint16                  //TLV [60][4]
	RequestedExtendedDrxParameters        *Bytes                   //TLV [6E][3-4]
	T3324Value                            *GprsTimer3              //TLV [6A][3]
	UeRadioCapabilityId                   *Bytes                   //TLV [67][3-n]
	RequestedMappedNssai                  *MappedNssai             //TLV [35][3-42]
	AdditionalInformationRequested        *Uint8                   //TLV [48][3]
	RequestedWusAssistanceInformation     *Bytes                   //TLV [1A][3-n]
	N5gcIndication                        *Uint8                   //TV [A-][1]
	RequestedNbN1ModeDrxParameters        *Uint8                   //TLV [30][3]
	UeRequestType                         *Uint8                   //TLV [29][3]
	PagingRestriction                     *Bytes                   //TLV [28][3-35]
	ServiceLevelAaContainer               *Bytes                   //TLV-E [72][6-n]
	Nid                                   *Bytes                   //TLV [32][8]
	MsDeterminedPlmnWithDisasterCondition *Bytes                   //TLV [16][5]
	RequestedPeipsAssistanceInformation   *Bytes                   //TLV [2A][3-n]
	RequestedT3512Value                   *GprsTimer3              //TLV [3B][3]
}

func (msg *RegistrationRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding RegistrationRequest", err)
		}
	}()
	var buf []byte
	// V[1/2]
	if buf, err = msg.RegistrationType.encode(); err != nil {
		err = nasError("encoding RegistrationType [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding RegistrationType [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// V[1/2]
	if buf, err = msg.Ngksi.encode(); err != nil {
		err = nasError("encoding Ngksi [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding Ngksi [M V 1/2]", ErrInvalidSize)
		return
	}
	v |= (buf[0] & 0x0f) << 4 //fill lefthalf
	wire = append(wire, v)

	// LV-E[6-n]
	if buf, err = encodeLV(true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("encoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.NonCurrentNativeNasKeySetIdentifier != nil {
		// TV[1]
		if buf, err = msg.NonCurrentNativeNasKeySetIdentifier.encode(); err != nil {
			err = nasError("encoding NonCurrentNativeNasKeySetIdentifier [O TV 1]", err)
			return
		}
		if len(buf) != 1 {
			err = nasError("encoding NonCurrentNativeNasKeySetIdentifier [O TV 1]", ErrInvalidSize)
			return
		}
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0C<<4)|(buf[0]&0x0f))
	}

	if msg.GmmCapability != nil {
		// TLV[3-15]
		if buf, err = encodeLV(false, uint16(1), uint16(13), msg.GmmCapability); err != nil {
			err = nasError("encoding GmmCapability [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x10), buf...)
	}

	if msg.UeSecurityCapability != nil {
		// TLV[4-10]
		if buf, err = encodeLV(false, uint16(2), uint16(8), msg.UeSecurityCapability); err != nil {
			err = nasError("encoding UeSecurityCapability [O TLV 4-10]", err)
			return
		}
		wire = append(append(wire, 0x2E), buf...)
	}

	if msg.RequestedNssai != nil {
		// TLV[4-74]
		if buf, err = encodeLV(false, uint16(2), uint16(72), msg.RequestedNssai); err != nil {
			err = nasError("encoding RequestedNssai [O TLV 4-74]", err)
			return
		}
		wire = append(append(wire, 0x2F), buf...)
	}

	if msg.LastVisitedRegisteredTai != nil {
		// TV[7]
		if buf, err = msg.LastVisitedRegisteredTai.encode(); err != nil {
			err = nasError("encoding LastVisitedRegisteredTai [O TV 7]", err)
			return
		}
		if len(buf) != 6 {
			err = nasError("encoding LastVisitedRegisteredTai [O TV 7]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x52), buf...)
	}

	if msg.S1UeNetworkCapability != nil {
		// TLV[4-15]
		if buf, err = encodeLV(false, uint16(2), uint16(13), msg.S1UeNetworkCapability); err != nil {
			err = nasError("encoding S1UeNetworkCapability [O TLV 4-15]", err)
			return
		}
		wire = append(append(wire, 0x17), buf...)
	}

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

	if msg.MicoIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0B<<4)|(uint8(*msg.MicoIndication)&0x0f))
	}

	if msg.UeStatus != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.UeStatus); err != nil {
			err = nasError("encoding UeStatus [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x2B), buf...)
	}

	if msg.AdditionalGuti != nil {
		// TLV-E[14]
		if buf, err = encodeLV(true, uint16(11), uint16(11), msg.AdditionalGuti); err != nil {
			err = nasError("encoding AdditionalGuti [O TLV-E 14]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	if msg.AllowedPduSessionStatus != nil {
		// TLV[4-34]
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.AllowedPduSessionStatus); err != nil {
			err = nasError("encoding AllowedPduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
	}

	if msg.UeUsageSetting != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.UeUsageSetting); err != nil {
			err = nasError("encoding UeUsageSetting [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x18), buf...)
	}

	if msg.RequestedDrxParameters != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.RequestedDrxParameters); err != nil {
			err = nasError("encoding RequestedDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x51), buf...)
	}

	if msg.EpsNasMessageContainer != nil {
		// TLV-E[4-n]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.EpsNasMessageContainer); err != nil {
			err = nasError("encoding EpsNasMessageContainer [O TLV-E 4-n]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
	}

	if msg.LadnIndication != nil {
		// TLV-E[3-811]
		if buf, err = encodeLV(true, uint16(0), uint16(808), msg.LadnIndication); err != nil {
			err = nasError("encoding LadnIndication [O TLV-E 3-811]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
	}

	if msg.PayloadContainerType != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.PayloadContainerType)&0x0f))
	}

	if msg.PayloadContainer != nil {
		// TLV-E[4-65538]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.PayloadContainer); err != nil {
			err = nasError("encoding PayloadContainer [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	if msg.NetworkSlicingIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x09<<4)|(uint8(*msg.NetworkSlicingIndication)&0x0f))
	}

	if msg.UpdateType != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.UpdateType); err != nil {
			err = nasError("encoding UpdateType [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x53), buf...)
	}

	if msg.MobileStationClassmark2 != nil {
		// TLV[5]
		if buf, err = encodeLV(false, uint16(3), uint16(3), msg.MobileStationClassmark2); err != nil {
			err = nasError("encoding MobileStationClassmark2 [O TLV 5]", err)
			return
		}
		wire = append(append(wire, 0x41), buf...)
	}

	if msg.SupportedCodecs != nil {
		// TLV[5-n]
		if buf, err = encodeLV(false, uint16(3), uint16(0), msg.SupportedCodecs); err != nil {
			err = nasError("encoding SupportedCodecs [O TLV 5-n]", err)
			return
		}
		wire = append(append(wire, 0x42), buf...)
	}

	if msg.NasMessageContainer != nil {
		// TLV-E[4-n]
		if buf, err = encodeLV(true, uint16(1), uint16(0), msg.NasMessageContainer); err != nil {
			err = nasError("encoding NasMessageContainer [O TLV-E 4-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	if msg.EpsBearerContextStatus != nil {
		// TLV[4]
		if buf, err = encodeLV(false, uint16(2), uint16(2), msg.EpsBearerContextStatus); err != nil {
			err = nasError("encoding EpsBearerContextStatus [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x60), buf...)
	}

	if msg.RequestedExtendedDrxParameters != nil {
		// TLV[3-4]
		if buf, err = encodeLV(false, uint16(1), uint16(2), msg.RequestedExtendedDrxParameters); err != nil {
			err = nasError("encoding RequestedExtendedDrxParameters [O TLV 3-4]", err)
			return
		}
		wire = append(append(wire, 0x6E), buf...)
	}

	if msg.T3324Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3324Value); err != nil {
			err = nasError("encoding T3324Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6A), buf...)
	}

	if msg.UeRadioCapabilityId != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.UeRadioCapabilityId); err != nil {
			err = nasError("encoding UeRadioCapabilityId [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x67), buf...)
	}

	if msg.RequestedMappedNssai != nil {
		// TLV[3-42]
		if buf, err = encodeLV(false, uint16(1), uint16(40), msg.RequestedMappedNssai); err != nil {
			err = nasError("encoding RequestedMappedNssai [O TLV 3-42]", err)
			return
		}
		wire = append(append(wire, 0x35), buf...)
	}

	if msg.AdditionalInformationRequested != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.AdditionalInformationRequested); err != nil {
			err = nasError("encoding AdditionalInformationRequested [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x48), buf...)
	}

	if msg.RequestedWusAssistanceInformation != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.RequestedWusAssistanceInformation); err != nil {
			err = nasError("encoding RequestedWusAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x1A), buf...)
	}

	if msg.N5gcIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.N5gcIndication)&0x0f))
	}

	if msg.RequestedNbN1ModeDrxParameters != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.RequestedNbN1ModeDrxParameters); err != nil {
			err = nasError("encoding RequestedNbN1ModeDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x30), buf...)
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

	if msg.ServiceLevelAaContainer != nil {
		// TLV-E[6-n]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.ServiceLevelAaContainer); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	if msg.Nid != nil {
		// TLV[8]
		if buf, err = encodeLV(false, uint16(6), uint16(6), msg.Nid); err != nil {
			err = nasError("encoding Nid [O TLV 8]", err)
			return
		}
		wire = append(append(wire, 0x32), buf...)
	}

	if msg.MsDeterminedPlmnWithDisasterCondition != nil {
		// TLV[5]
		if buf, err = encodeLV(false, uint16(3), uint16(3), msg.MsDeterminedPlmnWithDisasterCondition); err != nil {
			err = nasError("encoding MsDeterminedPlmnWithDisasterCondition [O TLV 5]", err)
			return
		}
		wire = append(append(wire, 0x16), buf...)
	}

	if msg.RequestedPeipsAssistanceInformation != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.RequestedPeipsAssistanceInformation); err != nil {
			err = nasError("encoding RequestedPeipsAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x2A), buf...)
	}

	if msg.RequestedT3512Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.RequestedT3512Value); err != nil {
			err = nasError("encoding RequestedT3512Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x3B), buf...)
	}

	msg.msgType = RegistrationRequestMsgType //set message type to REGISTRATION REQUEST
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *RegistrationRequest) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding RegistrationRequest", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding RegistrationType [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.RegistrationType.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding RegistrationType [M V 1/2]", err)
		return
	}
	// V[1/2]
	if err = msg.Ngksi.decode([]byte{(0xf0 & wire[offset]) >> 4 /*lefthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	offset++

	// LV-E[6-n]
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("decoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x0C: //TV[1]
			v := &KeySetIdentifier{}
			if err = v.decode([]byte{wire[offset] & 0x0f /*righthalf*/}); err != nil {
				err = nasError("decoding NonCurrentNativeNasKeySetIdentifier [O TV 1]", err)
				return
			}
			msg.NonCurrentNativeNasKeySetIdentifier = v
			offset++
		case 0x10: //TLV[3-15]
			offset++ //consume IEI
			v := &GmmCapability{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GmmCapability [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GmmCapability = v
		case 0x2E: //TLV[4-10]
			offset++ //consume IEI
			v := &UeSecurityCapability{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(8), v); err != nil {
				err = nasError("decoding UeSecurityCapability [O TLV 4-10]", err)
				return
			}
			offset += consumed
			msg.UeSecurityCapability = v
		case 0x2F: //TLV[4-74]
			offset++ //consume IEI
			v := &Nssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(72), v); err != nil {
				err = nasError("decoding RequestedNssai [O TLV 4-74]", err)
				return
			}
			offset += consumed
			msg.RequestedNssai = v
		case 0x52: //TV[7]
			if offset+7 > wireLen {
				err = nasError("decoding LastVisitedRegisteredTai [O TV 7]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := &TrackingAreaIdentity{}
			if err = v.decode(wire[offset : offset+6]); err != nil {
				err = nasError("decoding LastVisitedRegisteredTai [O TV 7]", err)
				return
			}
			msg.LastVisitedRegisteredTai = v
			offset += 6

		case 0x17: //TLV[4-15]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(13), v); err != nil {
				err = nasError("decoding S1UeNetworkCapability [O TLV 4-15]", err)
				return
			}
			offset += consumed
			msg.S1UeNetworkCapability = v
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
		case 0x0B: //TV[1]
			msg.MicoIndication = new(Uint8)
			*msg.MicoIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x2B: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UeStatus [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UeStatus = v
		case 0x77: //TLV-E[14]
			offset++ //consume IEI
			v := &MobileIdentity{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(11), uint16(11), v); err != nil {
				err = nasError("decoding AdditionalGuti [O TLV-E 14]", err)
				return
			}
			offset += consumed
			msg.AdditionalGuti = v
		case 0x25: //TLV[4-34]
			offset++ //consume IEI
			v := &AllowedPduSessionStatus{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding AllowedPduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.AllowedPduSessionStatus = v
		case 0x18: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UeUsageSetting [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UeUsageSetting = v
		case 0x51: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding RequestedDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.RequestedDrxParameters = v
		case 0x70: //TLV-E[4-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding EpsNasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.EpsNasMessageContainer = v
		case 0x74: //TLV-E[3-811]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(808), v); err != nil {
				err = nasError("decoding LadnIndication [O TLV-E 3-811]", err)
				return
			}
			offset += consumed
			msg.LadnIndication = v
		case 0x08: //TV[1]
			msg.PayloadContainerType = new(Uint8)
			*msg.PayloadContainerType = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x7B: //TLV-E[4-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PayloadContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PayloadContainer = v
		case 0x09: //TV[1]
			msg.NetworkSlicingIndication = new(Uint8)
			*msg.NetworkSlicingIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x53: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UpdateType [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UpdateType = v
		case 0x41: //TLV[5]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(3), v); err != nil {
				err = nasError("decoding MobileStationClassmark2 [O TLV 5]", err)
				return
			}
			offset += consumed
			msg.MobileStationClassmark2 = v
		case 0x42: //TLV[5-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding SupportedCodecs [O TLV 5-n]", err)
				return
			}
			offset += consumed
			msg.SupportedCodecs = v
		case 0x71: //TLV-E[4-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.NasMessageContainer = v
		case 0x60: //TLV[4]
			offset++ //consume IEI
			v := new(Uint16)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding EpsBearerContextStatus [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.EpsBearerContextStatus = v
		case 0x6E: //TLV[3-4]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(2), v); err != nil {
				err = nasError("decoding RequestedExtendedDrxParameters [O TLV 3-4]", err)
				return
			}
			offset += consumed
			msg.RequestedExtendedDrxParameters = v
		case 0x6A: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3324Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3324Value = v
		case 0x67: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding UeRadioCapabilityId [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.UeRadioCapabilityId = v
		case 0x35: //TLV[3-42]
			offset++ //consume IEI
			v := &MappedNssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(40), v); err != nil {
				err = nasError("decoding RequestedMappedNssai [O TLV 3-42]", err)
				return
			}
			offset += consumed
			msg.RequestedMappedNssai = v
		case 0x48: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalInformationRequested [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalInformationRequested = v
		case 0x1A: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding RequestedWusAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.RequestedWusAssistanceInformation = v
		case 0x0A: //TV[1]
			msg.N5gcIndication = new(Uint8)
			*msg.N5gcIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x30: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding RequestedNbN1ModeDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.RequestedNbN1ModeDrxParameters = v
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
		case 0x72: //TLV-E[6-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = v
		case 0x32: //TLV[8]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), v); err != nil {
				err = nasError("decoding Nid [O TLV 8]", err)
				return
			}
			offset += consumed
			msg.Nid = v
		case 0x16: //TLV[5]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(3), v); err != nil {
				err = nasError("decoding MsDeterminedPlmnWithDisasterCondition [O TLV 5]", err)
				return
			}
			offset += consumed
			msg.MsDeterminedPlmnWithDisasterCondition = v
		case 0x2A: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding RequestedPeipsAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.RequestedPeipsAssistanceInformation = v
		case 0x3B: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding RequestedT3512Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.RequestedT3512Value = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
