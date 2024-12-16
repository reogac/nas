/**generated time: 2024-12-16 16:36:18.690460**/

package nas

/*******************************************************
 * REGISTRATION REQUEST
 ******************************************************/
type RegistrationRequest struct {
	MmHeader
	RegistrationType                      RegistrationType         //M: V [1/2]
	Ngksi                                 KeySetIdentifier         //M: V [1/2]
	MobileIdentity                        MobileIdentity           //M: LV-E [6-n]
	NonCurrentNativeNasKeySetIdentifier   *KeySetIdentifier        //O: TV [C-][1]
	GmmCapability                         *GmmCapability           //O: TLV [10][3-15]
	UeSecurityCapability                  *UeSecurityCapability    //O: TLV [2E][4-10]
	RequestedNssai                        *Nssai                   //O: TLV [2F][4-74]
	LastVisitedRegisteredTai              *TrackingAreaIdentity    //O: TV [52][7]
	S1UeNetworkCapability                 []byte                   //O: TLV [17][4-15]
	UplinkDataStatus                      *UplinkDataStatus        //O: TLV [40][4-34]
	PduSessionStatus                      *PduSessionStatus        //O: TLV [50][4-34]
	MicoIndication                        *uint8                   //O: TV [B-][1]
	UeStatus                              *uint8                   //O: TLV [2B][3]
	AdditionalGuti                        *MobileIdentity          //O: TLV-E [77][14]
	AllowedPduSessionStatus               *AllowedPduSessionStatus //O: TLV [25][4-34]
	UeUsageSetting                        *uint8                   //O: TLV [18][3]
	RequestedDrxParameters                *uint8                   //O: TLV [51][3]
	EpsNasMessageContainer                []byte                   //O: TLV-E [70][4-n]
	LadnIndication                        []byte                   //O: TLV-E [74][3-811]
	PayloadContainerType                  *uint8                   //O: TV [8-][1]
	PayloadContainer                      []byte                   //O: TLV-E [7B][4-65538]
	NetworkSlicingIndication              *uint8                   //O: TV [9-][1]
	UpdateType                            *uint8                   //O: TLV [53][3]
	MobileStationClassmark2               []byte                   //O: TLV [41][5]
	SupportedCodecs                       []byte                   //O: TLV [42][5-n]
	NasMessageContainer                   []byte                   //O: TLV-E [71][4-n]
	EpsBearerContextStatus                *uint16                  //O: TLV [60][4]
	RequestedExtendedDrxParameters        []byte                   //O: TLV [6E][3-4]
	T3324Value                            *GprsTimer3              //O: TLV [6A][3]
	UeRadioCapabilityId                   []byte                   //O: TLV [67][3-n]
	RequestedMappedNssai                  *MappedNssai             //O: TLV [35][3-42]
	AdditionalInformationRequested        *uint8                   //O: TLV [48][3]
	RequestedWusAssistanceInformation     []byte                   //O: TLV [1A][3-n]
	N5gcIndication                        *uint8                   //O: TV [A-][1]
	RequestedNbN1ModeDrxParameters        *uint8                   //O: TLV [30][3]
	UeRequestType                         *uint8                   //O: TLV [29][3]
	PagingRestriction                     []byte                   //O: TLV [28][3-35]
	ServiceLevelAaContainer               []byte                   //O: TLV-E [72][6-n]
	Nid                                   []byte                   //O: TLV [32][8]
	MsDeterminedPlmnWithDisasterCondition []byte                   //O: TLV [16][5]
	RequestedPeipsAssistanceInformation   []byte                   //O: TLV [2A][3-n]
	RequestedT3512Value                   *GprsTimer3              //O: TLV [3B][3]
}

func (msg *RegistrationRequest) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding RegistrationRequest", err)
		}
	}()
	var buf []byte
	// M: V[1/2]
	if buf, err = msg.RegistrationType.encode(); err != nil {
		err = nasError("encoding RegistrationType [M V 1/2]", err)
		return
	}
	if len(buf) != 1 {
		err = nasError("encoding RegistrationType [M V 1/2]", ErrInvalidSize)
		return
	}
	v := (buf[0] & 0x0f) //fill righthalf
	// M: V[1/2]
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

	// M: LV-E[6-n]
	if buf, err = encodeLV(true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("encoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TV[1]
	if msg.NonCurrentNativeNasKeySetIdentifier != nil {
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

	// O: TLV[3-15]
	if msg.GmmCapability != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(13), msg.GmmCapability); err != nil {
			err = nasError("encoding GmmCapability [O TLV 3-15]", err)
			return
		}
		wire = append(append(wire, 0x10), buf...)
	}

	// O: TLV[4-10]
	if msg.UeSecurityCapability != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(8), msg.UeSecurityCapability); err != nil {
			err = nasError("encoding UeSecurityCapability [O TLV 4-10]", err)
			return
		}
		wire = append(append(wire, 0x2E), buf...)
	}

	// O: TLV[4-74]
	if msg.RequestedNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(72), msg.RequestedNssai); err != nil {
			err = nasError("encoding RequestedNssai [O TLV 4-74]", err)
			return
		}
		wire = append(append(wire, 0x2F), buf...)
	}

	// O: TV[7]
	if msg.LastVisitedRegisteredTai != nil {
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

	// O: TLV[4-15]
	if len(msg.S1UeNetworkCapability) > 0 {
		tmp := newBytesEncoder(msg.S1UeNetworkCapability)
		if buf, err = encodeLV(false, uint16(2), uint16(13), tmp); err != nil {
			err = nasError("encoding S1UeNetworkCapability [O TLV 4-15]", err)
			return
		}
		wire = append(append(wire, 0x17), buf...)
	}

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

	// O: TV[1]
	if msg.MicoIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0B<<4)|(uint8(*msg.MicoIndication)&0x0f))
	}

	// O: TLV[3]
	if msg.UeStatus != nil {
		tmp := newUint8Encoder(*msg.UeStatus)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding UeStatus [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x2B), buf...)
	}

	// O: TLV-E[14]
	if msg.AdditionalGuti != nil {
		if buf, err = encodeLV(true, uint16(11), uint16(11), msg.AdditionalGuti); err != nil {
			err = nasError("encoding AdditionalGuti [O TLV-E 14]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	// O: TLV[4-34]
	if msg.AllowedPduSessionStatus != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.AllowedPduSessionStatus); err != nil {
			err = nasError("encoding AllowedPduSessionStatus [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x25), buf...)
	}

	// O: TLV[3]
	if msg.UeUsageSetting != nil {
		tmp := newUint8Encoder(*msg.UeUsageSetting)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding UeUsageSetting [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x18), buf...)
	}

	// O: TLV[3]
	if msg.RequestedDrxParameters != nil {
		tmp := newUint8Encoder(*msg.RequestedDrxParameters)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding RequestedDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x51), buf...)
	}

	// O: TLV-E[4-n]
	if len(msg.EpsNasMessageContainer) > 0 {
		tmp := newBytesEncoder(msg.EpsNasMessageContainer)
		if buf, err = encodeLV(true, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding EpsNasMessageContainer [O TLV-E 4-n]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
	}

	// O: TLV-E[3-811]
	if len(msg.LadnIndication) > 0 {
		tmp := newBytesEncoder(msg.LadnIndication)
		if buf, err = encodeLV(true, uint16(0), uint16(808), tmp); err != nil {
			err = nasError("encoding LadnIndication [O TLV-E 3-811]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
	}

	// O: TV[1]
	if msg.PayloadContainerType != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x08<<4)|(uint8(*msg.PayloadContainerType)&0x0f))
	}

	// O: TLV-E[4-65538]
	if len(msg.PayloadContainer) > 0 {
		tmp := newBytesEncoder(msg.PayloadContainer)
		if buf, err = encodeLV(true, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding PayloadContainer [O TLV-E 4-65538]", err)
			return
		}
		wire = append(append(wire, 0x7B), buf...)
	}

	// O: TV[1]
	if msg.NetworkSlicingIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x09<<4)|(uint8(*msg.NetworkSlicingIndication)&0x0f))
	}

	// O: TLV[3]
	if msg.UpdateType != nil {
		tmp := newUint8Encoder(*msg.UpdateType)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding UpdateType [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x53), buf...)
	}

	// O: TLV[5]
	if len(msg.MobileStationClassmark2) > 0 {
		tmp := newBytesEncoder(msg.MobileStationClassmark2)
		if buf, err = encodeLV(false, uint16(3), uint16(3), tmp); err != nil {
			err = nasError("encoding MobileStationClassmark2 [O TLV 5]", err)
			return
		}
		wire = append(append(wire, 0x41), buf...)
	}

	// O: TLV[5-n]
	if len(msg.SupportedCodecs) > 0 {
		tmp := newBytesEncoder(msg.SupportedCodecs)
		if buf, err = encodeLV(false, uint16(3), uint16(0), tmp); err != nil {
			err = nasError("encoding SupportedCodecs [O TLV 5-n]", err)
			return
		}
		wire = append(append(wire, 0x42), buf...)
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

	// O: TLV[4]
	if msg.EpsBearerContextStatus != nil {
		tmp := newUint16Encoder(*msg.EpsBearerContextStatus)
		if buf, err = encodeLV(false, uint16(2), uint16(2), tmp); err != nil {
			err = nasError("encoding EpsBearerContextStatus [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x60), buf...)
	}

	// O: TLV[3-4]
	if len(msg.RequestedExtendedDrxParameters) > 0 {
		tmp := newBytesEncoder(msg.RequestedExtendedDrxParameters)
		if buf, err = encodeLV(false, uint16(1), uint16(2), tmp); err != nil {
			err = nasError("encoding RequestedExtendedDrxParameters [O TLV 3-4]", err)
			return
		}
		wire = append(append(wire, 0x6E), buf...)
	}

	// O: TLV[3]
	if msg.T3324Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3324Value); err != nil {
			err = nasError("encoding T3324Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6A), buf...)
	}

	// O: TLV[3-n]
	if len(msg.UeRadioCapabilityId) > 0 {
		tmp := newBytesEncoder(msg.UeRadioCapabilityId)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding UeRadioCapabilityId [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x67), buf...)
	}

	// O: TLV[3-42]
	if msg.RequestedMappedNssai != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(40), msg.RequestedMappedNssai); err != nil {
			err = nasError("encoding RequestedMappedNssai [O TLV 3-42]", err)
			return
		}
		wire = append(append(wire, 0x35), buf...)
	}

	// O: TLV[3]
	if msg.AdditionalInformationRequested != nil {
		tmp := newUint8Encoder(*msg.AdditionalInformationRequested)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding AdditionalInformationRequested [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x48), buf...)
	}

	// O: TLV[3-n]
	if len(msg.RequestedWusAssistanceInformation) > 0 {
		tmp := newBytesEncoder(msg.RequestedWusAssistanceInformation)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding RequestedWusAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x1A), buf...)
	}

	// O: TV[1]
	if msg.N5gcIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.N5gcIndication)&0x0f))
	}

	// O: TLV[3]
	if msg.RequestedNbN1ModeDrxParameters != nil {
		tmp := newUint8Encoder(*msg.RequestedNbN1ModeDrxParameters)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding RequestedNbN1ModeDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x30), buf...)
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

	// O: TLV-E[6-n]
	if len(msg.ServiceLevelAaContainer) > 0 {
		tmp := newBytesEncoder(msg.ServiceLevelAaContainer)
		if buf, err = encodeLV(true, uint16(3), uint16(0), tmp); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	// O: TLV[8]
	if len(msg.Nid) > 0 {
		tmp := newBytesEncoder(msg.Nid)
		if buf, err = encodeLV(false, uint16(6), uint16(6), tmp); err != nil {
			err = nasError("encoding Nid [O TLV 8]", err)
			return
		}
		wire = append(append(wire, 0x32), buf...)
	}

	// O: TLV[5]
	if len(msg.MsDeterminedPlmnWithDisasterCondition) > 0 {
		tmp := newBytesEncoder(msg.MsDeterminedPlmnWithDisasterCondition)
		if buf, err = encodeLV(false, uint16(3), uint16(3), tmp); err != nil {
			err = nasError("encoding MsDeterminedPlmnWithDisasterCondition [O TLV 5]", err)
			return
		}
		wire = append(append(wire, 0x16), buf...)
	}

	// O: TLV[3-n]
	if len(msg.RequestedPeipsAssistanceInformation) > 0 {
		tmp := newBytesEncoder(msg.RequestedPeipsAssistanceInformation)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding RequestedPeipsAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x2A), buf...)
	}

	// O: TLV[3]
	if msg.RequestedT3512Value != nil {
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
	// M V[1/2]
	if offset+1 > wireLen {
		err = nasError("decoding RegistrationType [M V 1/2]", ErrIncomplete)
		return
	}
	if err = msg.RegistrationType.decode([]byte{0x0f & wire[offset] /*righthalf*/}); err != nil {
		err = nasError("decoding RegistrationType [M V 1/2]", err)
		return
	}
	// M V[1/2]
	if err = msg.Ngksi.decode([]byte{(0xf0 & wire[offset]) >> 4 /*lefthalf*/}); err != nil {
		err = nasError("decoding Ngksi [M V 1/2]", err)
		return
	}
	offset++

	// M LV-E[6-n]
	if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), &msg.MobileIdentity); err != nil {
		err = nasError("decoding MobileIdentity [M LV-E 6-n]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x0C: //O: TV[1]
			v := new(KeySetIdentifier)
			if err = v.decode([]byte{wire[offset] & 0x0f /*righthalf*/}); err != nil {
				err = nasError("decoding NonCurrentNativeNasKeySetIdentifier [O TV 1]", err)
				return
			}
			msg.NonCurrentNativeNasKeySetIdentifier = v
			offset++
		case 0x10: //O: TLV[3-15]
			offset++ //consume IEI
			v := new(GmmCapability)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(13), v); err != nil {
				err = nasError("decoding GmmCapability [O TLV 3-15]", err)
				return
			}
			offset += consumed
			msg.GmmCapability = v
		case 0x2E: //O: TLV[4-10]
			offset++ //consume IEI
			v := new(UeSecurityCapability)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(8), v); err != nil {
				err = nasError("decoding UeSecurityCapability [O TLV 4-10]", err)
				return
			}
			offset += consumed
			msg.UeSecurityCapability = v
		case 0x2F: //O: TLV[4-74]
			offset++ //consume IEI
			v := new(Nssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(72), v); err != nil {
				err = nasError("decoding RequestedNssai [O TLV 4-74]", err)
				return
			}
			offset += consumed
			msg.RequestedNssai = v
		case 0x52: //O: TV[7]
			if offset+7 > wireLen {
				err = nasError("decoding LastVisitedRegisteredTai [O TV 7]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(TrackingAreaIdentity)
			if err = v.decode(wire[offset : offset+6]); err != nil {
				err = nasError("decoding LastVisitedRegisteredTai [O TV 7]", err)
				return
			}
			msg.LastVisitedRegisteredTai = v
			offset += 6

		case 0x17: //O: TLV[4-15]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(13), v); err != nil {
				err = nasError("decoding S1UeNetworkCapability [O TLV 4-15]", err)
				return
			}
			offset += consumed
			msg.S1UeNetworkCapability = []byte(*v)
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
		case 0x0B: //O: TV[1]
			msg.MicoIndication = new(uint8)
			*msg.MicoIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x2B: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UeStatus [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UeStatus = (*uint8)(v)
		case 0x77: //O: TLV-E[14]
			offset++ //consume IEI
			v := new(MobileIdentity)
			if consumed, err = decodeLV(wire[offset:], true, uint16(11), uint16(11), v); err != nil {
				err = nasError("decoding AdditionalGuti [O TLV-E 14]", err)
				return
			}
			offset += consumed
			msg.AdditionalGuti = v
		case 0x25: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(AllowedPduSessionStatus)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding AllowedPduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.AllowedPduSessionStatus = v
		case 0x18: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UeUsageSetting [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UeUsageSetting = (*uint8)(v)
		case 0x51: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding RequestedDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.RequestedDrxParameters = (*uint8)(v)
		case 0x70: //O: TLV-E[4-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding EpsNasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.EpsNasMessageContainer = []byte(*v)
		case 0x74: //O: TLV-E[3-811]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(808), v); err != nil {
				err = nasError("decoding LadnIndication [O TLV-E 3-811]", err)
				return
			}
			offset += consumed
			msg.LadnIndication = []byte(*v)
		case 0x08: //O: TV[1]
			msg.PayloadContainerType = new(uint8)
			*msg.PayloadContainerType = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x7B: //O: TLV-E[4-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding PayloadContainer [O TLV-E 4-65538]", err)
				return
			}
			offset += consumed
			msg.PayloadContainer = []byte(*v)
		case 0x09: //O: TV[1]
			msg.NetworkSlicingIndication = new(uint8)
			*msg.NetworkSlicingIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x53: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding UpdateType [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.UpdateType = (*uint8)(v)
		case 0x41: //O: TLV[5]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(3), v); err != nil {
				err = nasError("decoding MobileStationClassmark2 [O TLV 5]", err)
				return
			}
			offset += consumed
			msg.MobileStationClassmark2 = []byte(*v)
		case 0x42: //O: TLV[5-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding SupportedCodecs [O TLV 5-n]", err)
				return
			}
			offset += consumed
			msg.SupportedCodecs = []byte(*v)
		case 0x71: //O: TLV-E[4-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NasMessageContainer [O TLV-E 4-n]", err)
				return
			}
			offset += consumed
			msg.NasMessageContainer = []byte(*v)
		case 0x60: //O: TLV[4]
			offset++ //consume IEI
			v := new(uint16Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding EpsBearerContextStatus [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.EpsBearerContextStatus = (*uint16)(v)
		case 0x6E: //O: TLV[3-4]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(2), v); err != nil {
				err = nasError("decoding RequestedExtendedDrxParameters [O TLV 3-4]", err)
				return
			}
			offset += consumed
			msg.RequestedExtendedDrxParameters = []byte(*v)
		case 0x6A: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3324Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3324Value = v
		case 0x67: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding UeRadioCapabilityId [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.UeRadioCapabilityId = []byte(*v)
		case 0x35: //O: TLV[3-42]
			offset++ //consume IEI
			v := new(MappedNssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(40), v); err != nil {
				err = nasError("decoding RequestedMappedNssai [O TLV 3-42]", err)
				return
			}
			offset += consumed
			msg.RequestedMappedNssai = v
		case 0x48: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalInformationRequested [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalInformationRequested = (*uint8)(v)
		case 0x1A: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding RequestedWusAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.RequestedWusAssistanceInformation = []byte(*v)
		case 0x0A: //O: TV[1]
			msg.N5gcIndication = new(uint8)
			*msg.N5gcIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x30: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding RequestedNbN1ModeDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.RequestedNbN1ModeDrxParameters = (*uint8)(v)
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
		case 0x72: //O: TLV-E[6-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = []byte(*v)
		case 0x32: //O: TLV[8]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(6), uint16(6), v); err != nil {
				err = nasError("decoding Nid [O TLV 8]", err)
				return
			}
			offset += consumed
			msg.Nid = []byte(*v)
		case 0x16: //O: TLV[5]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(3), v); err != nil {
				err = nasError("decoding MsDeterminedPlmnWithDisasterCondition [O TLV 5]", err)
				return
			}
			offset += consumed
			msg.MsDeterminedPlmnWithDisasterCondition = []byte(*v)
		case 0x2A: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding RequestedPeipsAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.RequestedPeipsAssistanceInformation = []byte(*v)
		case 0x3B: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
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
