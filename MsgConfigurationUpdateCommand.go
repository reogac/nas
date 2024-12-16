/**generated time: 2024-12-16 16:36:18.694005**/

package nas

/*******************************************************
 * CONFIGURATION UPDATE COMMAND
 ******************************************************/
type ConfigurationUpdateCommand struct {
	MmHeader
	ConfigurationUpdateIndication            *uint8                    //O: TV [D-][1]
	Guti                                     *MobileIdentity           //O: TLV-E [77][14]
	TaiList                                  *TrackingAreaIdentityList //O: TLV [54][9-114]
	AllowedNssai                             *Nssai                    //O: TLV [15][4-74]
	ServiceAreaList                          *ServiceAreaList          //O: TLV [27][6-114]
	FullNameForNetwork                       *NetworkName              //O: TLV [43][3-n]
	ShortNameForNetwork                      *NetworkName              //O: TLV [45][3-n]
	LocalTimeZone                            *TimeZone                 //O: TV [46][2]
	UniversalTimeAndLocalTimeZone            []byte                    //O: TV [47][8]
	NetworkDaylightSavingTime                *uint8                    //O: TLV [49][3]
	LadnInformation                          *LadnInformation          //O: TLV-E [79][3-1715]
	MicoIndication                           *uint8                    //O: TV [B-][1]
	NetworkSlicingIndication                 *uint8                    //O: TV [9-][1]
	ConfiguredNssai                          *Nssai                    //O: TLV [31][4-146]
	RejectedNssai                            *RejectedNssai            //O: TLV [11][4-42]
	OperatorDefinedAccessCategoryDefinitions []byte                    //O: TLV-E [76][3-8323]
	SmsIndication                            *uint8                    //O: TV [F-][1]
	T3447Value                               *GprsTimer3               //O: TLV [6C][3]
	CagInformationList                       []byte                    //O: TLV-E [75][3-n]
	UeRadioCapabilityId                      []byte                    //O: TLV [67][3-n]
	UeRadioCapabilityIdDeletionIndication    *uint8                    //O: TV [A-][1]
	RegistrationResult                       *RegistrationResult       //O: TLV [44][3]
	TruncatedSTmsiConfiguration              *uint8                    //O: TLV [1B][3]
	AdditionalConfigurationIndication        *uint8                    //O: TV [C-][1]
	ExtendedRejectedNssai                    []byte                    //O: TLV [68][5-90]
	ServiceLevelAaContainer                  []byte                    //O: TLV-E [72][6-n]
	NssrgInformation                         []byte                    //O: TLV-E [70][7-4099]
	DisasterRoamingWaitRange                 *uint16                   //O: TLV [14][4]
	DisasterReturnWaitRange                  *uint16                   //O: TLV [2C][4]
	ListOfPlmnsToBeUsedInDisasterCondition   []byte                    //O: TLV [13][2-n]
	ExtendedCagInformationList               []byte                    //O: TLV-E [71][3-n]
	UpdatedPeipsAssistanceInformation        []byte                    //O: TLV [1F][3-n]
	NsagInformation                          []byte                    //O: TLV-E [73][9-3143]
	PriorityIndicator                        *uint8                    //O: TV [E-][1]
}

func (msg *ConfigurationUpdateCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding ConfigurationUpdateCommand", err)
		}
	}()
	var buf []byte
	// O: TV[1]
	if msg.ConfigurationUpdateIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0D<<4)|(uint8(*msg.ConfigurationUpdateIndication)&0x0f))
	}

	// O: TLV-E[14]
	if msg.Guti != nil {
		if buf, err = encodeLV(true, uint16(11), uint16(11), msg.Guti); err != nil {
			err = nasError("encoding Guti [O TLV-E 14]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	// O: TLV[9-114]
	if msg.TaiList != nil {
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.TaiList); err != nil {
			err = nasError("encoding TaiList [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x54), buf...)
	}

	// O: TLV[4-74]
	if msg.AllowedNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(72), msg.AllowedNssai); err != nil {
			err = nasError("encoding AllowedNssai [O TLV 4-74]", err)
			return
		}
		wire = append(append(wire, 0x15), buf...)
	}

	// O: TLV[6-114]
	if msg.ServiceAreaList != nil {
		if buf, err = encodeLV(false, uint16(4), uint16(112), msg.ServiceAreaList); err != nil {
			err = nasError("encoding ServiceAreaList [O TLV 6-114]", err)
			return
		}
		wire = append(append(wire, 0x27), buf...)
	}

	// O: TLV[3-n]
	if msg.FullNameForNetwork != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.FullNameForNetwork); err != nil {
			err = nasError("encoding FullNameForNetwork [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x43), buf...)
	}

	// O: TLV[3-n]
	if msg.ShortNameForNetwork != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.ShortNameForNetwork); err != nil {
			err = nasError("encoding ShortNameForNetwork [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x45), buf...)
	}

	// O: TV[2]
	if msg.LocalTimeZone != nil {
		if buf, err = msg.LocalTimeZone.encode(); err != nil {
			err = nasError("encoding LocalTimeZone [O TV 2]", err)
			return
		}
		if len(buf) != 1 {
			err = nasError("encoding LocalTimeZone [O TV 2]", ErrInvalidSize)
			return
		}
		wire = append(wire, []byte{0x46, buf[0]}...)
	}

	// O: TV[8]
	if len(msg.UniversalTimeAndLocalTimeZone) > 0 {
		tmp := newBytesEncoder(msg.UniversalTimeAndLocalTimeZone)
		if buf, err = tmp.encode(); err != nil {
			err = nasError("encoding UniversalTimeAndLocalTimeZone [O TV 8]", err)
			return
		}
		if len(buf) != 7 {
			err = nasError("encoding UniversalTimeAndLocalTimeZone [O TV 8]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x47), buf...)
	}

	// O: TLV[3]
	if msg.NetworkDaylightSavingTime != nil {
		tmp := newUint8Encoder(*msg.NetworkDaylightSavingTime)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding NetworkDaylightSavingTime [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x49), buf...)
	}

	// O: TLV-E[3-1715]
	if msg.LadnInformation != nil {
		if buf, err = encodeLV(true, uint16(0), uint16(1712), msg.LadnInformation); err != nil {
			err = nasError("encoding LadnInformation [O TLV-E 3-1715]", err)
			return
		}
		wire = append(append(wire, 0x79), buf...)
	}

	// O: TV[1]
	if msg.MicoIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0B<<4)|(uint8(*msg.MicoIndication)&0x0f))
	}

	// O: TV[1]
	if msg.NetworkSlicingIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x09<<4)|(uint8(*msg.NetworkSlicingIndication)&0x0f))
	}

	// O: TLV[4-146]
	if msg.ConfiguredNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(144), msg.ConfiguredNssai); err != nil {
			err = nasError("encoding ConfiguredNssai [O TLV 4-146]", err)
			return
		}
		wire = append(append(wire, 0x31), buf...)
	}

	// O: TLV[4-42]
	if msg.RejectedNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(40), msg.RejectedNssai); err != nil {
			err = nasError("encoding RejectedNssai [O TLV 4-42]", err)
			return
		}
		wire = append(append(wire, 0x11), buf...)
	}

	// O: TLV-E[3-8323]
	if len(msg.OperatorDefinedAccessCategoryDefinitions) > 0 {
		tmp := newBytesEncoder(msg.OperatorDefinedAccessCategoryDefinitions)
		if buf, err = encodeLV(true, uint16(0), uint16(8320), tmp); err != nil {
			err = nasError("encoding OperatorDefinedAccessCategoryDefinitions [O TLV-E 3-8323]", err)
			return
		}
		wire = append(append(wire, 0x76), buf...)
	}

	// O: TV[1]
	if msg.SmsIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0F<<4)|(uint8(*msg.SmsIndication)&0x0f))
	}

	// O: TLV[3]
	if msg.T3447Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3447Value); err != nil {
			err = nasError("encoding T3447Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6C), buf...)
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

	// O: TLV[3-n]
	if len(msg.UeRadioCapabilityId) > 0 {
		tmp := newBytesEncoder(msg.UeRadioCapabilityId)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding UeRadioCapabilityId [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x67), buf...)
	}

	// O: TV[1]
	if msg.UeRadioCapabilityIdDeletionIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.UeRadioCapabilityIdDeletionIndication)&0x0f))
	}

	// O: TLV[3]
	if msg.RegistrationResult != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.RegistrationResult); err != nil {
			err = nasError("encoding RegistrationResult [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x44), buf...)
	}

	// O: TLV[3]
	if msg.TruncatedSTmsiConfiguration != nil {
		tmp := newUint8Encoder(*msg.TruncatedSTmsiConfiguration)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding TruncatedSTmsiConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1B), buf...)
	}

	// O: TV[1]
	if msg.AdditionalConfigurationIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0C<<4)|(uint8(*msg.AdditionalConfigurationIndication)&0x0f))
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

	// O: TLV-E[6-n]
	if len(msg.ServiceLevelAaContainer) > 0 {
		tmp := newBytesEncoder(msg.ServiceLevelAaContainer)
		if buf, err = encodeLV(true, uint16(3), uint16(0), tmp); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	// O: TLV-E[7-4099]
	if len(msg.NssrgInformation) > 0 {
		tmp := newBytesEncoder(msg.NssrgInformation)
		if buf, err = encodeLV(true, uint16(4), uint16(4096), tmp); err != nil {
			err = nasError("encoding NssrgInformation [O TLV-E 7-4099]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
	}

	// O: TLV[4]
	if msg.DisasterRoamingWaitRange != nil {
		tmp := newUint16Encoder(*msg.DisasterRoamingWaitRange)
		if buf, err = encodeLV(false, uint16(2), uint16(2), tmp); err != nil {
			err = nasError("encoding DisasterRoamingWaitRange [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x14), buf...)
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

	// O: TLV[2-n]
	if len(msg.ListOfPlmnsToBeUsedInDisasterCondition) > 0 {
		tmp := newBytesEncoder(msg.ListOfPlmnsToBeUsedInDisasterCondition)
		if buf, err = encodeLV(false, uint16(0), uint16(0), tmp); err != nil {
			err = nasError("encoding ListOfPlmnsToBeUsedInDisasterCondition [O TLV 2-n]", err)
			return
		}
		wire = append(append(wire, 0x13), buf...)
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

	// O: TLV[3-n]
	if len(msg.UpdatedPeipsAssistanceInformation) > 0 {
		tmp := newBytesEncoder(msg.UpdatedPeipsAssistanceInformation)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding UpdatedPeipsAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x1F), buf...)
	}

	// O: TLV-E[9-3143]
	if len(msg.NsagInformation) > 0 {
		tmp := newBytesEncoder(msg.NsagInformation)
		if buf, err = encodeLV(true, uint16(6), uint16(3140), tmp); err != nil {
			err = nasError("encoding NsagInformation [O TLV-E 9-3143]", err)
			return
		}
		wire = append(append(wire, 0x73), buf...)
	}

	// O: TV[1]
	if msg.PriorityIndicator != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0E<<4)|(uint8(*msg.PriorityIndicator)&0x0f))
	}

	msg.msgType = ConfigurationUpdateCommandMsgType //set message type to CONFIGURATION UPDATE COMMAND
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *ConfigurationUpdateCommand) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding ConfigurationUpdateCommand", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x0D: //O: TV[1]
			msg.ConfigurationUpdateIndication = new(uint8)
			*msg.ConfigurationUpdateIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x77: //O: TLV-E[14]
			offset++ //consume IEI
			v := new(MobileIdentity)
			if consumed, err = decodeLV(wire[offset:], true, uint16(11), uint16(11), v); err != nil {
				err = nasError("decoding Guti [O TLV-E 14]", err)
				return
			}
			offset += consumed
			msg.Guti = v
		case 0x54: //O: TLV[9-114]
			offset++ //consume IEI
			v := new(TrackingAreaIdentityList)
			if consumed, err = decodeLV(wire[offset:], false, uint16(7), uint16(112), v); err != nil {
				err = nasError("decoding TaiList [O TLV 9-114]", err)
				return
			}
			offset += consumed
			msg.TaiList = v
		case 0x15: //O: TLV[4-74]
			offset++ //consume IEI
			v := new(Nssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(72), v); err != nil {
				err = nasError("decoding AllowedNssai [O TLV 4-74]", err)
				return
			}
			offset += consumed
			msg.AllowedNssai = v
		case 0x27: //O: TLV[6-114]
			offset++ //consume IEI
			v := new(ServiceAreaList)
			if consumed, err = decodeLV(wire[offset:], false, uint16(4), uint16(112), v); err != nil {
				err = nasError("decoding ServiceAreaList [O TLV 6-114]", err)
				return
			}
			offset += consumed
			msg.ServiceAreaList = v
		case 0x43: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(NetworkName)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding FullNameForNetwork [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.FullNameForNetwork = v
		case 0x45: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(NetworkName)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ShortNameForNetwork [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.ShortNameForNetwork = v
		case 0x46: //O: TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding LocalTimeZone [O TV 2]", ErrIncomplete)
				return
			}
			v := new(TimeZone)
			offset++ //consume IEI
			if err = v.decode(wire[offset : offset+1]); err != nil {
				err = nasError("decoding LocalTimeZone [O TV 2]", err)
				return
			}
			msg.LocalTimeZone = v
			offset++
		case 0x47: //O: TV[8]
			if offset+8 > wireLen {
				err = nasError("decoding UniversalTimeAndLocalTimeZone [O TV 8]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(bytesDecoder)
			if err = v.decode(wire[offset : offset+7]); err != nil {
				err = nasError("decoding UniversalTimeAndLocalTimeZone [O TV 8]", err)
				return
			}
			msg.UniversalTimeAndLocalTimeZone = []byte(*v)
			offset += 7

		case 0x49: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding NetworkDaylightSavingTime [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.NetworkDaylightSavingTime = (*uint8)(v)
		case 0x79: //O: TLV-E[3-1715]
			offset++ //consume IEI
			v := new(LadnInformation)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(1712), v); err != nil {
				err = nasError("decoding LadnInformation [O TLV-E 3-1715]", err)
				return
			}
			offset += consumed
			msg.LadnInformation = v
		case 0x0B: //O: TV[1]
			msg.MicoIndication = new(uint8)
			*msg.MicoIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x09: //O: TV[1]
			msg.NetworkSlicingIndication = new(uint8)
			*msg.NetworkSlicingIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x31: //O: TLV[4-146]
			offset++ //consume IEI
			v := new(Nssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(144), v); err != nil {
				err = nasError("decoding ConfiguredNssai [O TLV 4-146]", err)
				return
			}
			offset += consumed
			msg.ConfiguredNssai = v
		case 0x11: //O: TLV[4-42]
			offset++ //consume IEI
			v := new(RejectedNssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(40), v); err != nil {
				err = nasError("decoding RejectedNssai [O TLV 4-42]", err)
				return
			}
			offset += consumed
			msg.RejectedNssai = v
		case 0x76: //O: TLV-E[3-8323]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(8320), v); err != nil {
				err = nasError("decoding OperatorDefinedAccessCategoryDefinitions [O TLV-E 3-8323]", err)
				return
			}
			offset += consumed
			msg.OperatorDefinedAccessCategoryDefinitions = []byte(*v)
		case 0x0F: //O: TV[1]
			msg.SmsIndication = new(uint8)
			*msg.SmsIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x6C: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3447Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3447Value = v
		case 0x75: //O: TLV-E[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding CagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.CagInformationList = []byte(*v)
		case 0x67: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding UeRadioCapabilityId [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.UeRadioCapabilityId = []byte(*v)
		case 0x0A: //O: TV[1]
			msg.UeRadioCapabilityIdDeletionIndication = new(uint8)
			*msg.UeRadioCapabilityIdDeletionIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x44: //O: TLV[3]
			offset++ //consume IEI
			v := new(RegistrationResult)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding RegistrationResult [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.RegistrationResult = v
		case 0x1B: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding TruncatedSTmsiConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.TruncatedSTmsiConfiguration = (*uint8)(v)
		case 0x0C: //O: TV[1]
			msg.AdditionalConfigurationIndication = new(uint8)
			*msg.AdditionalConfigurationIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x68: //O: TLV[5-90]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(88), v); err != nil {
				err = nasError("decoding ExtendedRejectedNssai [O TLV 5-90]", err)
				return
			}
			offset += consumed
			msg.ExtendedRejectedNssai = []byte(*v)
		case 0x72: //O: TLV-E[6-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = []byte(*v)
		case 0x70: //O: TLV-E[7-4099]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(4096), v); err != nil {
				err = nasError("decoding NssrgInformation [O TLV-E 7-4099]", err)
				return
			}
			offset += consumed
			msg.NssrgInformation = []byte(*v)
		case 0x14: //O: TLV[4]
			offset++ //consume IEI
			v := new(uint16Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding DisasterRoamingWaitRange [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.DisasterRoamingWaitRange = (*uint16)(v)
		case 0x2C: //O: TLV[4]
			offset++ //consume IEI
			v := new(uint16Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding DisasterReturnWaitRange [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.DisasterReturnWaitRange = (*uint16)(v)
		case 0x13: //O: TLV[2-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ListOfPlmnsToBeUsedInDisasterCondition [O TLV 2-n]", err)
				return
			}
			offset += consumed
			msg.ListOfPlmnsToBeUsedInDisasterCondition = []byte(*v)
		case 0x71: //O: TLV-E[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedCagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.ExtendedCagInformationList = []byte(*v)
		case 0x1F: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding UpdatedPeipsAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.UpdatedPeipsAssistanceInformation = []byte(*v)
		case 0x73: //O: TLV-E[9-3143]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(3140), v); err != nil {
				err = nasError("decoding NsagInformation [O TLV-E 9-3143]", err)
				return
			}
			offset += consumed
			msg.NsagInformation = []byte(*v)
		case 0x0E: //O: TV[1]
			msg.PriorityIndicator = new(uint8)
			*msg.PriorityIndicator = wire[offset] & 0x0f //take right 1/2
			offset++
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
