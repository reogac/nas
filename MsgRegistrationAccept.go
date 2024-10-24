/**generated time: 2024-07-17 15:11:00.938890**/

package nas

/*******************************************************
 * REGISTRATION ACCEPT
 ******************************************************/
type RegistrationAccept struct {
	MmHeader
	RegistrationResult                                                          RegistrationResult                      //LV [2]
	Guti                                                                        *MobileIdentity                         //TLV-E [77][14]
	EquivalentPlmns                                                             *PlmnList                               //TLV [4A][5-47]
	TaiList                                                                     *TrackingAreaIdentityList               //TLV [54][9-114]
	AllowedNssai                                                                *Nssai                                  //TLV [15][4-74]
	RejectedNssai                                                               *RejectedNssai                          //TLV [11][4-42]
	ConfiguredNssai                                                             *Nssai                                  //TLV [31][4-146]
	NetworkFeatureSupport                                                       *NetworkFeatureSupport                  //TLV [21][3-5]
	PduSessionStatus                                                            *PduSessionStatus                       //TLV [50][4-34]
	PduSessionReactivationResult                                                *PduSessionReactivationResult           //TLV [26][4-34]
	PduSessionReactivationResultErrorCause                                      *PduSessionReactivationResultErrorCause //TLV-E [72][5-515]
	LadnInformation                                                             *LadnInformation                        //TLV-E [79][12-1715]
	MicoIndication                                                              *Uint8                                  //TV [B-][1]
	NetworkSlicingIndication                                                    *Uint8                                  //TV [9-][1]
	ServiceAreaList                                                             *ServiceAreaList                        //TLV [27][6-114]
	T3512Value                                                                  *GprsTimer3                             //TLV [5E][3]
	Non3gppDeRegistrationTimerValue                                             *GprsTimer2                             //TLV [5D][3]
	T3502Value                                                                  *GprsTimer2                             //TLV [16][3]
	EmergencyNumberList                                                         *Bytes                                  //TLV [34][5-50]
	ExtendedEmergencyNumberList                                                 *Bytes                                  //TLV-E [7A][7-65538]
	SorTransparentContainer                                                     *Bytes                                  //TLV-E [73][20-n]
	EapMessage                                                                  *Bytes                                  //TLV-E [78][7-1503]
	NssaiInclusionMode                                                          *Uint8                                  //TV [A-][1]
	OperatorDefinedAccessCategoryDefinitions                                    *Bytes                                  //TLV-E [76][3-8323]
	NegotiatedDrxParameters                                                     *Uint8                                  //TLV [51][3]
	Non3gppNwPolicies                                                           *Uint8                                  //TV [D-][1]
	EpsBearerContextStatus                                                      *Uint16                                 //TLV [60][4]
	NegotiatedExtendedDrxParameters                                             *Bytes                                  //TLV [6E][3-4]
	T3447Value                                                                  *GprsTimer3                             //TLV [6C][3]
	T3448Value                                                                  *GprsTimer2                             //TLV [6B][3]
	T3324Value                                                                  *GprsTimer3                             //TLV [6A][3]
	UeRadioCapabilityId                                                         *Bytes                                  //TLV [67][3-n]
	UeRadioCapabilityIdDeletionIndication                                       *Uint8                                  //TV [E-][1]
	PendingNssai                                                                *Nssai                                  //TLV [39][4-146]
	CipheringKeyData                                                            *Bytes                                  //TLV-E [74][34-n]
	CagInformationList                                                          *Bytes                                  //TLV-E [75][3-n]
	TruncatedSTmsiConfiguration                                                 *Uint8                                  //TLV [1B][3]
	NegotiatedWusAssistanceInformation                                          *Bytes                                  //TLV [1C][3-n]
	NegotiatedNbN1ModeDrxParameters                                             *Uint8                                  //TLV [29][3]
	ExtendedRejectedNssai                                                       *Bytes                                  //TLV [68][5-90]
	ServiceLevelAaContainer                                                     *Bytes                                  //TLV-E [7B][6-n]
	NegotiatedPeipsAssistanceInformation                                        *Bytes                                  //TLV [33][3-n]
	AdditionalRequestResult                                                     *Uint8                                  //TLV [35][3]
	NssrgInformation                                                            *Bytes                                  //TLV-E [70][7-4099]
	DisasterRoamingWaitRange                                                    *Uint16                                 //TLV [14][4]
	DisasterReturnWaitRange                                                     *Uint16                                 //TLV [2C][4]
	ListOfPlmnsToBeUsedInDisasterCondition                                      *Bytes                                  //TLV [13][2-n]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming                    *TrackingAreaIdentityList               //TLV [1D][9-114]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService *TrackingAreaIdentityList               //TLV [1E][9-114]
	ExtendedCagInformationList                                                  *Bytes                                  //TLV-E [71][3-n]
	NsagInformation                                                             *Bytes                                  //TLV-E [7C][9-3143]
}

func (msg *RegistrationAccept) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding RegistrationAccept", err)
		}
	}()
	var buf []byte
	// LV[2]
	if buf, err = encodeLV(false, uint16(1), uint16(1), &msg.RegistrationResult); err != nil {
		err = nasError("encoding RegistrationResult [M LV 2]", err)
		return
	}
	wire = append(wire, buf...)

	if msg.Guti != nil {
		// TLV-E[14]
		if buf, err = encodeLV(true, uint16(11), uint16(11), msg.Guti); err != nil {
			err = nasError("encoding Guti [O TLV-E 14]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	if msg.EquivalentPlmns != nil {
		// TLV[5-47]
		if buf, err = encodeLV(false, uint16(3), uint16(45), msg.EquivalentPlmns); err != nil {
			err = nasError("encoding EquivalentPlmns [O TLV 5-47]", err)
			return
		}
		wire = append(append(wire, 0x4A), buf...)
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

	if msg.RejectedNssai != nil {
		// TLV[4-42]
		if buf, err = encodeLV(false, uint16(2), uint16(40), msg.RejectedNssai); err != nil {
			err = nasError("encoding RejectedNssai [O TLV 4-42]", err)
			return
		}
		wire = append(append(wire, 0x11), buf...)
	}

	if msg.ConfiguredNssai != nil {
		// TLV[4-146]
		if buf, err = encodeLV(false, uint16(2), uint16(144), msg.ConfiguredNssai); err != nil {
			err = nasError("encoding ConfiguredNssai [O TLV 4-146]", err)
			return
		}
		wire = append(append(wire, 0x31), buf...)
	}

	if msg.NetworkFeatureSupport != nil {
		// TLV[3-5]
		if buf, err = encodeLV(false, uint16(1), uint16(3), msg.NetworkFeatureSupport); err != nil {
			err = nasError("encoding NetworkFeatureSupport [O TLV 3-5]", err)
			return
		}
		wire = append(append(wire, 0x21), buf...)
	}

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

	if msg.LadnInformation != nil {
		// TLV-E[12-1715]
		if buf, err = encodeLV(true, uint16(9), uint16(1712), msg.LadnInformation); err != nil {
			err = nasError("encoding LadnInformation [O TLV-E 12-1715]", err)
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

	if msg.ServiceAreaList != nil {
		// TLV[6-114]
		if buf, err = encodeLV(false, uint16(4), uint16(112), msg.ServiceAreaList); err != nil {
			err = nasError("encoding ServiceAreaList [O TLV 6-114]", err)
			return
		}
		wire = append(append(wire, 0x27), buf...)
	}

	if msg.T3512Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3512Value); err != nil {
			err = nasError("encoding T3512Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x5E), buf...)
	}

	if msg.Non3gppDeRegistrationTimerValue != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.Non3gppDeRegistrationTimerValue); err != nil {
			err = nasError("encoding Non3gppDeRegistrationTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x5D), buf...)
	}

	if msg.T3502Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3502Value); err != nil {
			err = nasError("encoding T3502Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x16), buf...)
	}

	if msg.EmergencyNumberList != nil {
		// TLV[5-50]
		if buf, err = encodeLV(false, uint16(3), uint16(48), msg.EmergencyNumberList); err != nil {
			err = nasError("encoding EmergencyNumberList [O TLV 5-50]", err)
			return
		}
		wire = append(append(wire, 0x34), buf...)
	}

	if msg.ExtendedEmergencyNumberList != nil {
		// TLV-E[7-65538]
		if buf, err = encodeLV(true, uint16(4), uint16(0), msg.ExtendedEmergencyNumberList); err != nil {
			err = nasError("encoding ExtendedEmergencyNumberList [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x7A), buf...)
	}

	if msg.SorTransparentContainer != nil {
		// TLV-E[20-n]
		if buf, err = encodeLV(true, uint16(17), uint16(0), msg.SorTransparentContainer); err != nil {
			err = nasError("encoding SorTransparentContainer [O TLV-E 20-n]", err)
			return
		}
		wire = append(append(wire, 0x73), buf...)
	}

	if msg.EapMessage != nil {
		// TLV-E[7-1503]
		if buf, err = encodeLV(true, uint16(4), uint16(1500), msg.EapMessage); err != nil {
			err = nasError("encoding EapMessage [O TLV-E 7-1503]", err)
			return
		}
		wire = append(append(wire, 0x78), buf...)
	}

	if msg.NssaiInclusionMode != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.NssaiInclusionMode)&0x0f))
	}

	if msg.OperatorDefinedAccessCategoryDefinitions != nil {
		// TLV-E[3-8323]
		if buf, err = encodeLV(true, uint16(0), uint16(8320), msg.OperatorDefinedAccessCategoryDefinitions); err != nil {
			err = nasError("encoding OperatorDefinedAccessCategoryDefinitions [O TLV-E 3-8323]", err)
			return
		}
		wire = append(append(wire, 0x76), buf...)
	}

	if msg.NegotiatedDrxParameters != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.NegotiatedDrxParameters); err != nil {
			err = nasError("encoding NegotiatedDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x51), buf...)
	}

	if msg.Non3gppNwPolicies != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0D<<4)|(uint8(*msg.Non3gppNwPolicies)&0x0f))
	}

	if msg.EpsBearerContextStatus != nil {
		// TLV[4]
		if buf, err = encodeLV(false, uint16(2), uint16(2), msg.EpsBearerContextStatus); err != nil {
			err = nasError("encoding EpsBearerContextStatus [O TLV 4]", err)
			return
		}
		wire = append(append(wire, 0x60), buf...)
	}

	if msg.NegotiatedExtendedDrxParameters != nil {
		// TLV[3-4]
		if buf, err = encodeLV(false, uint16(1), uint16(2), msg.NegotiatedExtendedDrxParameters); err != nil {
			err = nasError("encoding NegotiatedExtendedDrxParameters [O TLV 3-4]", err)
			return
		}
		wire = append(append(wire, 0x6E), buf...)
	}

	if msg.T3447Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3447Value); err != nil {
			err = nasError("encoding T3447Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6C), buf...)
	}

	if msg.T3448Value != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3448Value); err != nil {
			err = nasError("encoding T3448Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6B), buf...)
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

	if msg.UeRadioCapabilityIdDeletionIndication != nil {
		// TV[1]
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0E<<4)|(uint8(*msg.UeRadioCapabilityIdDeletionIndication)&0x0f))
	}

	if msg.PendingNssai != nil {
		// TLV[4-146]
		if buf, err = encodeLV(false, uint16(2), uint16(144), msg.PendingNssai); err != nil {
			err = nasError("encoding PendingNssai [O TLV 4-146]", err)
			return
		}
		wire = append(append(wire, 0x39), buf...)
	}

	if msg.CipheringKeyData != nil {
		// TLV-E[34-n]
		if buf, err = encodeLV(true, uint16(31), uint16(0), msg.CipheringKeyData); err != nil {
			err = nasError("encoding CipheringKeyData [O TLV-E 34-n]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
	}

	if msg.CagInformationList != nil {
		// TLV-E[3-n]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.CagInformationList); err != nil {
			err = nasError("encoding CagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x75), buf...)
	}

	if msg.TruncatedSTmsiConfiguration != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.TruncatedSTmsiConfiguration); err != nil {
			err = nasError("encoding TruncatedSTmsiConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1B), buf...)
	}

	if msg.NegotiatedWusAssistanceInformation != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.NegotiatedWusAssistanceInformation); err != nil {
			err = nasError("encoding NegotiatedWusAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x1C), buf...)
	}

	if msg.NegotiatedNbN1ModeDrxParameters != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.NegotiatedNbN1ModeDrxParameters); err != nil {
			err = nasError("encoding NegotiatedNbN1ModeDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
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
		wire = append(append(wire, 0x7B), buf...)
	}

	if msg.NegotiatedPeipsAssistanceInformation != nil {
		// TLV[3-n]
		if buf, err = encodeLV(false, uint16(1), uint16(0), msg.NegotiatedPeipsAssistanceInformation); err != nil {
			err = nasError("encoding NegotiatedPeipsAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x33), buf...)
	}

	if msg.AdditionalRequestResult != nil {
		// TLV[3]
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.AdditionalRequestResult); err != nil {
			err = nasError("encoding AdditionalRequestResult [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x35), buf...)
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

	if msg.ExtendedCagInformationList != nil {
		// TLV-E[3-n]
		if buf, err = encodeLV(true, uint16(0), uint16(0), msg.ExtendedCagInformationList); err != nil {
			err = nasError("encoding ExtendedCagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	if msg.NsagInformation != nil {
		// TLV-E[9-3143]
		if buf, err = encodeLV(true, uint16(6), uint16(3140), msg.NsagInformation); err != nil {
			err = nasError("encoding NsagInformation [O TLV-E 9-3143]", err)
			return
		}
		wire = append(append(wire, 0x7C), buf...)
	}

	msg.msgType = RegistrationAcceptMsgType //set message type to REGISTRATION ACCEPT
	wire = append(msg.headerBytes(), wire...)
	return
}
func (msg *RegistrationAccept) decodeBody(wire []byte) (err error) {
	defer func() {
		if err != nil {
			err = nasError("decoding RegistrationAccept", err)
		}
	}()
	offset := 0
	wireLen := len(wire)
	consumed := 0
	// LV[2]
	if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), &msg.RegistrationResult); err != nil {
		err = nasError("decoding RegistrationResult [M LV 2]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x77: //TLV-E[14]
			offset++ //consume IEI
			v := &MobileIdentity{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(11), uint16(11), v); err != nil {
				err = nasError("decoding Guti [O TLV-E 14]", err)
				return
			}
			offset += consumed
			msg.Guti = v
		case 0x4A: //TLV[5-47]
			offset++ //consume IEI
			v := &PlmnList{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(45), v); err != nil {
				err = nasError("decoding EquivalentPlmns [O TLV 5-47]", err)
				return
			}
			offset += consumed
			msg.EquivalentPlmns = v
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
		case 0x11: //TLV[4-42]
			offset++ //consume IEI
			v := &RejectedNssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(40), v); err != nil {
				err = nasError("decoding RejectedNssai [O TLV 4-42]", err)
				return
			}
			offset += consumed
			msg.RejectedNssai = v
		case 0x31: //TLV[4-146]
			offset++ //consume IEI
			v := &Nssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(144), v); err != nil {
				err = nasError("decoding ConfiguredNssai [O TLV 4-146]", err)
				return
			}
			offset += consumed
			msg.ConfiguredNssai = v
		case 0x21: //TLV[3-5]
			offset++ //consume IEI
			v := &NetworkFeatureSupport{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(3), v); err != nil {
				err = nasError("decoding NetworkFeatureSupport [O TLV 3-5]", err)
				return
			}
			offset += consumed
			msg.NetworkFeatureSupport = v
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
		case 0x79: //TLV-E[12-1715]
			offset++ //consume IEI
			v := &LadnInformation{}
			if consumed, err = decodeLV(wire[offset:], true, uint16(9), uint16(1712), v); err != nil {
				err = nasError("decoding LadnInformation [O TLV-E 12-1715]", err)
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
		case 0x27: //TLV[6-114]
			offset++ //consume IEI
			v := &ServiceAreaList{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(4), uint16(112), v); err != nil {
				err = nasError("decoding ServiceAreaList [O TLV 6-114]", err)
				return
			}
			offset += consumed
			msg.ServiceAreaList = v
		case 0x5E: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3512Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3512Value = v
		case 0x5D: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer2{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding Non3gppDeRegistrationTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.Non3gppDeRegistrationTimerValue = v
		case 0x16: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer2{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3502Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3502Value = v
		case 0x34: //TLV[5-50]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(48), v); err != nil {
				err = nasError("decoding EmergencyNumberList [O TLV 5-50]", err)
				return
			}
			offset += consumed
			msg.EmergencyNumberList = v
		case 0x7A: //TLV-E[7-65538]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedEmergencyNumberList [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedEmergencyNumberList = v
		case 0x73: //TLV-E[20-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(17), uint16(0), v); err != nil {
				err = nasError("decoding SorTransparentContainer [O TLV-E 20-n]", err)
				return
			}
			offset += consumed
			msg.SorTransparentContainer = v
		case 0x78: //TLV-E[7-1503]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = v
		case 0x0A: //TV[1]
			msg.NssaiInclusionMode = new(Uint8)
			*msg.NssaiInclusionMode = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x76: //TLV-E[3-8323]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(8320), v); err != nil {
				err = nasError("decoding OperatorDefinedAccessCategoryDefinitions [O TLV-E 3-8323]", err)
				return
			}
			offset += consumed
			msg.OperatorDefinedAccessCategoryDefinitions = v
		case 0x51: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding NegotiatedDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.NegotiatedDrxParameters = v
		case 0x0D: //TV[1]
			msg.Non3gppNwPolicies = new(Uint8)
			*msg.Non3gppNwPolicies = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
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
				err = nasError("decoding NegotiatedExtendedDrxParameters [O TLV 3-4]", err)
				return
			}
			offset += consumed
			msg.NegotiatedExtendedDrxParameters = v
		case 0x6C: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer3{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3447Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3447Value = v
		case 0x6B: //TLV[3]
			offset++ //consume IEI
			v := &GprsTimer2{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3448Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3448Value = v
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
		case 0x0E: //TV[1]
			msg.UeRadioCapabilityIdDeletionIndication = new(Uint8)
			*msg.UeRadioCapabilityIdDeletionIndication = Uint8(wire[offset] & 0x0f) //take right 1/2
			offset++
		case 0x39: //TLV[4-146]
			offset++ //consume IEI
			v := &Nssai{}
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(144), v); err != nil {
				err = nasError("decoding PendingNssai [O TLV 4-146]", err)
				return
			}
			offset += consumed
			msg.PendingNssai = v
		case 0x74: //TLV-E[34-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(31), uint16(0), v); err != nil {
				err = nasError("decoding CipheringKeyData [O TLV-E 34-n]", err)
				return
			}
			offset += consumed
			msg.CipheringKeyData = v
		case 0x75: //TLV-E[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding CagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.CagInformationList = v
		case 0x1B: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding TruncatedSTmsiConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.TruncatedSTmsiConfiguration = v
		case 0x1C: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NegotiatedWusAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.NegotiatedWusAssistanceInformation = v
		case 0x29: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding NegotiatedNbN1ModeDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.NegotiatedNbN1ModeDrxParameters = v
		case 0x68: //TLV[5-90]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(88), v); err != nil {
				err = nasError("decoding ExtendedRejectedNssai [O TLV 5-90]", err)
				return
			}
			offset += consumed
			msg.ExtendedRejectedNssai = v
		case 0x7B: //TLV-E[6-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = v
		case 0x33: //TLV[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NegotiatedPeipsAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.NegotiatedPeipsAssistanceInformation = v
		case 0x35: //TLV[3]
			offset++ //consume IEI
			v := new(Uint8)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalRequestResult [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalRequestResult = v
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
		case 0x71: //TLV-E[3-n]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedCagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.ExtendedCagInformationList = v
		case 0x7C: //TLV-E[9-3143]
			offset++ //consume IEI
			v := new(Bytes)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(3140), v); err != nil {
				err = nasError("decoding NsagInformation [O TLV-E 9-3143]", err)
				return
			}
			offset += consumed
			msg.NsagInformation = v
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
