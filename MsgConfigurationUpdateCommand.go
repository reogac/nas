/**generated time: 2024-07-17 15:11:00.942610**/

package nas

/*******************************************************
 * CONFIGURATION UPDATE COMMAND
 ******************************************************/
type ConfigurationUpdateCommand struct {
	MmHeader
	ConfigurationUpdateIndication            *Uint8                    //TV [D-][1]
	Guti                                     *MobileIdentity           //TLV-E [77][14]
	TaiList                                  *TrackingAreaIdentityList //TLV [54][9-114]
	AllowedNssai                             *Nssai                    //TLV [15][4-74]
	ServiceAreaList                          *ServiceAreaList          //TLV [27][6-114]
	FullNameForNetwork                       *NetworkName              //TLV [43][3-n]
	ShortNameForNetwork                      *NetworkName              //TLV [45][3-n]
	LocalTimeZone                            *TimeZone                 //TV [46][2]
	UniversalTimeAndLocalTimeZone            *Bytes                    //TV [47][8]
	NetworkDaylightSavingTime                *Uint8                    //TLV [49][3]
	LadnInformation                          *LadnInformation          //TLV-E [79][3-1715]
	MicoIndication                           *Uint8                    //TV [B-][1]
	NetworkSlicingIndication                 *Uint8                    //TV [9-][1]
	ConfiguredNssai                          *Nssai                    //TLV [31][4-146]
	RejectedNssai                            *RejectedNssai            //TLV [11][4-42]
	OperatorDefinedAccessCategoryDefinitions *Bytes                    //TLV-E [76][3-8323]
	SmsIndication                            *Uint8                    //TV [F-][1]
	T3447Value                               *GprsTimer3               //TLV [6C][3]
	CagInformationList                       *Bytes                    //TLV-E [75][3-n]
	UeRadioCapabilityId                      *Bytes                    //TLV [67][3-n]
	UeRadioCapabilityIdDeletionIndication    *Uint8                    //TV [A-][1]
	RegistrationResult                       *RegistrationResult       //TLV [44][3]
	TruncatedSTmsiConfiguration              *Uint8                    //TLV [1B][3]
	AdditionalConfigurationIndication        *Uint8                    //TV [C-][1]
	ExtendedRejectedNssai                    *Bytes                    //TLV [68][5-90]
	ServiceLevelAaContainer                  *Bytes                    //TLV-E [72][6-n]
	NssrgInformation                         *Bytes                    //TLV-E [70][7-4099]
	DisasterRoamingWaitRange                 *Uint16                   //TLV [14][4]
	DisasterReturnWaitRange                  *Uint16                   //TLV [2C][4]
	ListOfPlmnsToBeUsedInDisasterCondition   *Bytes                    //TLV [13][2-n]
	ExtendedCagInformationList               *Bytes                    //TLV-E [71][3-n]
	UpdatedPeipsAssistanceInformation        *Bytes                    //TLV [1F][3-n]
	NsagInformation                          *Bytes                    //TLV-E [73][9-3143]
	PriorityIndicator                        *Uint8                    //TV [E-][1]
}

func (msg *ConfigurationUpdateCommand) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding ConfigurationUpdateCommand", err)
		}
	}()
	var buf []byte
	if msg.ConfigurationUpdateIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0D<<4)|(uint8(*msg.ConfigurationUpdateIndication)&0x0f))
	}

	if msg.Guti != nil {
		// TLV-E[14]
		if buf, err = encodeLV(true, uint16(11), uint16(11), msg.Guti); err != nil {
			err = nasError("encoding Guti [O TLV-E 14]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	if msg.TaiList != nil {
		// TLV[9-114]
		if buf, err = encodeLV(false, uint16(7), uint16(112), msg.TaiList); err != nil {
			err = nasError("encoding TaiList [O TLV 9-114]", err)
			return
		}
		wire = append(append(wire, 0x54), buf...)
	}

	if msg.AllowedNssai != nil {
		// TLV[4-74]
		if buf, err = encodeLV(false, uint16(2), uint16(72), msg.AllowedNssai); err != nil {
			err = nasError("encoding AllowedNssai [O TLV 4-74]", err)
			return
		}
		wire = append(append(wire, 0x15), buf...)
	}

	if msg.ServiceAreaList != nil {
		// TLV[6-114]
		if buf, err = encodeLV(false, uint16(4), uint16(112), msg.ServiceAreaList); err != nil {
			err = nasError("encoding ServiceAreaList [O TLV 6-114]", err)
			return
		}
		wire = append(append(wire, 0x27), buf...)
	}

	if msg.FullNameForNetwork != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.FullNameForNetwork); err != nil {
			err = nasError("encoding FullNameForNetwork [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x43), buf...)
	}

	if msg.ShortNameForNetwork != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.ShortNameForNetwork); err != nil {
			err = nasError("encoding ShortNameForNetwork [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x45), buf...)
	}

	if msg.LocalTimeZone != nil {
		// TV[2]
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

	if msg.UniversalTimeAndLocalTimeZone != nil {
		// TV[8]
		if buf, err = msg.UniversalTimeAndLocalTimeZone.encode(); err != nil {
			err = nasError("encoding UniversalTimeAndLocalTimeZone [O TV 8]", err)
			return
		}
		if len(buf) != 7 {
			err = nasError("encoding UniversalTimeAndLocalTimeZone [O TV 8]", ErrInvalidSize)
			return
		}
		wire = append(append(wire, 0x47), buf...)
	}

	if msg.NetworkDaylightSavingTime != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.NetworkDaylightSavingTime); err != nil {
			err = nasError("encoding NetworkDaylightSavingTime [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x49), buf...)
	}

	if msg.LadnInformation != nil {
		// TLV-E[3-1715]
		if buf, err = encodeLV(true, uint16(0), uint16(1712), msg.LadnInformation); err != nil {
			err = nasError("encoding LadnInformation [O TLV-E 3-1715]", err)
			return
		}
		wire = append(append(wire, 0x79), buf...)
	}

	if msg.MicoIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0B<<4)|(uint8(*msg.MicoIndication)&0x0f))
	}

	if msg.NetworkSlicingIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x09<<4)|(uint8(*msg.NetworkSlicingIndication)&0x0f))
	}

	if msg.ConfiguredNssai != nil {
		// TLV[4-146]
		if buf, err = encodeLV(false, uint16(2), uint16(144), msg.ConfiguredNssai); err != nil {
			err = nasError("encoding ConfiguredNssai [O TLV 4-146]", err)
			return
		}
		wire = append(append(wire, 0x31), buf...)
	}

	if msg.RejectedNssai != nil {
		// TLV[4-42]
		if buf, err = encodeLV(false, uint16(2), uint16(40), msg.RejectedNssai); err != nil {
			err = nasError("encoding RejectedNssai [O TLV 4-42]", err)
			return
		}
		wire = append(append(wire, 0x11), buf...)
	}

	if msg.OperatorDefinedAccessCategoryDefinitions != nil {
		// TLV-E[3-8323]
		if buf, err = encodeLV(true, uint16(0), uint16(8320), msg.OperatorDefinedAccessCategoryDefinitions); err != nil {
			err = nasError("encoding OperatorDefinedAccessCategoryDefinitions [O TLV-E 3-8323]", err)
			return
		}
		wire = append(append(wire, 0x76), buf...)
	}

	if msg.SmsIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0F<<4)|(uint8(*msg.SmsIndication)&0x0f))
	}

	if msg.T3447Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3447Value); err != nil {
			err = nasError("encoding T3447Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6C), buf...)
	}

	if msg.CagInformationList != nil {
		// TLV-E[3-n]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.CagInformationList); err != nil {
			err = nasError("encoding CagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
	}

	if msg.UeRadioCapabilityId != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.UeRadioCapabilityId); err != nil {
			err = nasError("encoding UeRadioCapabilityId [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x67), buf...)
	}

	if msg.UeRadioCapabilityIdDeletionIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.UeRadioCapabilityIdDeletionIndication)&0x0f))
	}

	if msg.RegistrationResult != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.RegistrationResult); err != nil {
			err = nasError("encoding RegistrationResult [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x44), buf...)
	}

	if msg.TruncatedSTmsiConfiguration != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.TruncatedSTmsiConfiguration); err != nil {
			err = nasError("encoding TruncatedSTmsiConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1B), buf...)
	}

	if msg.AdditionalConfigurationIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0C<<4)|(uint8(*msg.AdditionalConfigurationIndication)&0x0f))
	}

	if msg.ExtendedRejectedNssai != nil {
		// TLV[5-90]
		if buf, err = encodeLV(false, uint16(3), uint16(88), msg.ExtendedRejectedNssai); err != nil {
			err = nasError("encoding ExtendedRejectedNssai [O TLV 5-90]", err)
			return
		}
		wire = append(append(wire, 0x68), buf...)
	}

	if msg.ServiceLevelAaContainer != nil {
		// TLV-E[6-n]
		if buf, err = encodeLV(true, uint16(3), uint16(0), msg.ServiceLevelAaContainer); err != nil {
			err = nasError("encoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	if msg.NssrgInformation != nil {
		// TLV-E[7-4099]
		if buf, err = encodeLV(true, uint16(4), uint16(4096), msg.NssrgInformation); err != nil {
			err = nasError("encoding NssrgInformation [O TLV-E 7-4099]", err)
			return
		}
		wire = append(append(wire, 0x70), buf...)
	}

	if msg.DisasterRoamingWaitRange != nil {
		// TLV[4]
		if buf, err = encodeLV(false, uint16(2), uint16(2), msg.DisasterRoamingWaitRange); err != nil {
			err = nasError("encoding DisasterRoamingWaitRange [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x14), buf...)
	}

	if msg.DisasterReturnWaitRange != nil {
		// TLV[4]
		if buf, err = encodeLV(false, uint16(2), uint16(2), msg.DisasterReturnWaitRange); err != nil {
			err = nasError("encoding DisasterReturnWaitRange [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x2C), buf...)
	}

	if msg.ListOfPlmnsToBeUsedInDisasterCondition != nil {
		// TLV[2-n]
		if buf, err = encodeLV(false, uint16(0), uint16(0), msg.ListOfPlmnsToBeUsedInDisasterCondition); err != nil {
			err = nasError("encoding ListOfPlmnsToBeUsedInDisasterCondition [O TLV 2-n]", err)
			return
		}
		wire = append(append(wire, 0x13), buf...)
	}

	if msg.ExtendedCagInformationList != nil {
		// TLV-E[3-n]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.ExtendedCagInformationList); err != nil {
			err = nasError("encoding ExtendedCagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	if msg.UpdatedPeipsAssistanceInformation != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.UpdatedPeipsAssistanceInformation); err != nil {
			err = nasError("encoding UpdatedPeipsAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x1F), buf...)
	}

	if msg.NsagInformation != nil {
		// TLV-E[9-3143]
		if buf, err = encodeLV(true, uint16(6), uint16(3140), msg.NsagInformation); err != nil {
			err = nasError("encoding NsagInformation [O TLV-E 9-3143]", err)
			return
		}
		wire = append(append(wire, 0x73), buf...)
	}

	if msg.PriorityIndicator != nil {
		// TV[1]
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
		case 0x0D: //TV[1]
			msg.ConfigurationUpdateIndication = new(Uint8)
			*msg.ConfigurationUpdateIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x77: //TLV-E[14]
			offset++ //consume IEI
			v := &MobileIdentity{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(11), uint16(11), v); err != nil {
				err = nasError("decoding Guti [O TLV-E 14]", err)
				return
			}
			offset += consumed
			msg.Guti = v
		case 0x54: //TLV[9-114]
			offset++ //consume IEI
			v := &TrackingAreaIdentityList{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(7), uint16(112), v); err != nil {
				err = nasError("decoding TaiList [O TLV 9-114]", err)
				return
			}
			offset += consumed
			msg.TaiList = v
		case 0x15: //TLV[4-74]
			offset++ //consume IEI
			v := &Nssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(72), v); err != nil {
				err = nasError("decoding AllowedNssai [O TLV 4-74]", err)
				return
			}
			offset += consumed
			msg.AllowedNssai = v
		case 0x27: //TLV[6-114]
			offset++ //consume IEI
			v := &ServiceAreaList{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(4), uint16(112), v); err != nil {
				err = nasError("decoding ServiceAreaList [O TLV 6-114]", err)
				return
			}
			offset += consumed
			msg.ServiceAreaList = v
		case 0x43: //TLV[3-n]
			offset++ //consume IEI
			v := &NetworkName{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding FullNameForNetwork [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.FullNameForNetwork = v
		case 0x45: //TLV[3-n]
			offset++ //consume IEI
			v := &NetworkName{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding ShortNameForNetwork [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.ShortNameForNetwork = v
		case 0x46: //TV[2]
			if offset+2 > wireLen {
				err = nasError("decoding LocalTimeZone [O TV 2]", ErrIncomplete)
				return
			}
			v := &TimeZone{}
			offset++ //consume IEI
			if err = v.decode(wire[offset : offset+1]); err != nil {
				err = nasError("decoding LocalTimeZone [O TV 2]", err)
				return
			}
			msg.LocalTimeZone = v
			offset++
		case 0x47: //TV[8]
			if offset+8 > wireLen {
				err = nasError("decoding UniversalTimeAndLocalTimeZone [O TV 8]", ErrIncomplete)
				return
			}
			offset++ //consume IEI
			v := new(Bytes)
			if err = v.decode(wire[offset : offset+7]); err != nil {
				err = nasError("decoding UniversalTimeAndLocalTimeZone [O TV 8]", err)
				return
			}
			msg.UniversalTimeAndLocalTimeZone = v
			offset += 7

		case 0x49: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding NetworkDaylightSavingTime [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.NetworkDaylightSavingTime = v
		case 0x79: //TLV-E[3-1715]
			offset++ //consume IEI
			v := &LadnInformation{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(1712), v); err != nil {
				err = nasError("decoding LadnInformation [O TLV-E 3-1715]", err)
				return
			}
			offset += consumed
			msg.LadnInformation = v
		case 0x0B: //TV[1]
			msg.MicoIndication = new(Uint8)
			*msg.MicoIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x09: //TV[1]
			msg.NetworkSlicingIndication = new(Uint8)
			*msg.NetworkSlicingIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x31: //TLV[4-146]
			offset++ //consume IEI
			v := &Nssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(144), v); err != nil {
				err = nasError("decoding ConfiguredNssai [O TLV 4-146]", err)
				return
			}
			offset += consumed
			msg.ConfiguredNssai = v
		case 0x11: //TLV[4-42]
			offset++ //consume IEI
			v := &RejectedNssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(40), v); err != nil {
				err = nasError("decoding RejectedNssai [O TLV 4-42]", err)
				return
			}
			offset += consumed
			msg.RejectedNssai = v
		case 0x76: //TLV-E[3-8323]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(8320), v); err != nil {
				err = nasError("decoding OperatorDefinedAccessCategoryDefinitions [O TLV-E 3-8323]", err)
				return
			}
			offset += consumed
			msg.OperatorDefinedAccessCategoryDefinitions = v
		case 0x0F: //TV[1]
			msg.SmsIndication = new(Uint8)
			*msg.SmsIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x6C: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3447Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3447Value = v
		case 0x75: //TLV-E[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding CagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.CagInformationList = v
		case 0x67: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding UeRadioCapabilityId [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.UeRadioCapabilityId = v
		case 0x0A: //TV[1]
			msg.UeRadioCapabilityIdDeletionIndication = new(Uint8)
			*msg.UeRadioCapabilityIdDeletionIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x44: //TLV[3]
			offset++ //consume IEI
			v := &RegistrationResult{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding RegistrationResult [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.RegistrationResult = v
		case 0x1B: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding TruncatedSTmsiConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.TruncatedSTmsiConfiguration = v
		case 0x0C: //TV[1]
			msg.AdditionalConfigurationIndication = new(Uint8)
			*msg.AdditionalConfigurationIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x68: //TLV[5-90]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(88), v); err != nil {
				err = nasError("decoding ExtendedRejectedNssai [O TLV 5-90]", err)
				return
			}
			offset += consumed
			msg.ExtendedRejectedNssai = v
		case 0x72: //TLV-E[6-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = v
		case 0x70: //TLV-E[7-4099]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(4096), v); err != nil {
				err = nasError("decoding NssrgInformation [O TLV-E 7-4099]", err)
				return
			}
			offset += consumed
			msg.NssrgInformation = v
		case 0x14: //TLV[4]
			offset++ //consume IEI
			v := new(Uint16)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding DisasterRoamingWaitRange [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.DisasterRoamingWaitRange = v
		case 0x2C: //TLV[4]
			offset++ //consume IEI
			v := new(Uint16)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(2), v); err != nil {
				err = nasError("decoding DisasterReturnWaitRange [O TLV 4]", err)
				return
			}
			offset += consumed
			msg.DisasterReturnWaitRange = v
		case 0x13: //TLV[2-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ListOfPlmnsToBeUsedInDisasterCondition [O TLV 2-n]", err)
				return
			}
			offset += consumed
			msg.ListOfPlmnsToBeUsedInDisasterCondition = v
		case 0x71: //TLV-E[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedCagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.ExtendedCagInformationList = v
		case 0x1F: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding UpdatedPeipsAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.UpdatedPeipsAssistanceInformation = v
		case 0x73: //TLV-E[9-3143]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(3140), v); err != nil {
				err = nasError("decoding NsagInformation [O TLV-E 9-3143]", err)
				return
			}
			offset += consumed
			msg.NsagInformation = v
		case 0x0E: //TV[1]
			msg.PriorityIndicator = new(Uint8)
			*msg.PriorityIndicator = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
