/**generated time: 2024-12-16 16:36:18.691414**/

package nas

/*******************************************************
 * REGISTRATION ACCEPT
 ******************************************************/
type RegistrationAccept struct {
	MmHeader
	RegistrationResult                                                          RegistrationResult                      //M: LV [2]
	Guti                                                                        *MobileIdentity                         //O: TLV-E [77][14]
	EquivalentPlmns                                                             *PlmnList                               //O: TLV [4A][5-47]
	TaiList                                                                     *TrackingAreaIdentityList               //O: TLV [54][9-114]
	AllowedNssai                                                                *Nssai                                  //O: TLV [15][4-74]
	RejectedNssai                                                               *RejectedNssai                          //O: TLV [11][4-42]
	ConfiguredNssai                                                             *Nssai                                  //O: TLV [31][4-146]
	NetworkFeatureSupport                                                       *NetworkFeatureSupport                  //O: TLV [21][3-5]
	PduSessionStatus                                                            *PduSessionStatus                       //O: TLV [50][4-34]
	PduSessionReactivationResult                                                *PduSessionReactivationResult           //O: TLV [26][4-34]
	PduSessionReactivationResultErrorCause                                      *PduSessionReactivationResultErrorCause //O: TLV-E [72][5-515]
	LadnInformation                                                             *LadnInformation                        //O: TLV-E [79][12-1715]
	MicoIndication                                                              *uint8                                  //O: TV [B-][1]
	NetworkSlicingIndication                                                    *uint8                                  //O: TV [9-][1]
	ServiceAreaList                                                             *ServiceAreaList                        //O: TLV [27][6-114]
	T3512Value                                                                  *GprsTimer3                             //O: TLV [5E][3]
	Non3gppDeRegistrationTimerValue                                             *GprsTimer2                             //O: TLV [5D][3]
	T3502Value                                                                  *GprsTimer2                             //O: TLV [16][3]
	EmergencyNumberList                                                         []byte                                  //O: TLV [34][5-50]
	ExtendedEmergencyNumberList                                                 []byte                                  //O: TLV-E [7A][7-65538]
	SorTransparentContainer                                                     []byte                                  //O: TLV-E [73][20-n]
	EapMessage                                                                  []byte                                  //O: TLV-E [78][7-1503]
	NssaiInclusionMode                                                          *uint8                                  //O: TV [A-][1]
	OperatorDefinedAccessCategoryDefinitions                                    []byte                                  //O: TLV-E [76][3-8323]
	NegotiatedDrxParameters                                                     *uint8                                  //O: TLV [51][3]
	Non3gppNwPolicies                                                           *uint8                                  //O: TV [D-][1]
	EpsBearerContextStatus                                                      *uint16                                 //O: TLV [60][4]
	NegotiatedExtendedDrxParameters                                             []byte                                  //O: TLV [6E][3-4]
	T3447Value                                                                  *GprsTimer3                             //O: TLV [6C][3]
	T3448Value                                                                  *GprsTimer2                             //O: TLV [6B][3]
	T3324Value                                                                  *GprsTimer3                             //O: TLV [6A][3]
	UeRadioCapabilityId                                                         []byte                                  //O: TLV [67][3-n]
	UeRadioCapabilityIdDeletionIndication                                       *uint8                                  //O: TV [E-][1]
	PendingNssai                                                                *Nssai                                  //O: TLV [39][4-146]
	CipheringKeyData                                                            []byte                                  //O: TLV-E [74][34-n]
	CagInformationList                                                          []byte                                  //O: TLV-E [75][3-n]
	TruncatedSTmsiConfiguration                                                 *uint8                                  //O: TLV [1B][3]
	NegotiatedWusAssistanceInformation                                          []byte                                  //O: TLV [1C][3-n]
	NegotiatedNbN1ModeDrxParameters                                             *uint8                                  //O: TLV [29][3]
	ExtendedRejectedNssai                                                       []byte                                  //O: TLV [68][5-90]
	ServiceLevelAaContainer                                                     []byte                                  //O: TLV-E [7B][6-n]
	NegotiatedPeipsAssistanceInformation                                        []byte                                  //O: TLV [33][3-n]
	AdditionalRequestResult                                                     *uint8                                  //O: TLV [35][3]
	NssrgInformation                                                            []byte                                  //O: TLV-E [70][7-4099]
	DisasterRoamingWaitRange                                                    *uint16                                 //O: TLV [14][4]
	DisasterReturnWaitRange                                                     *uint16                                 //O: TLV [2C][4]
	ListOfPlmnsToBeUsedInDisasterCondition                                      []byte                                  //O: TLV [13][2-n]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForRoaming                    *TrackingAreaIdentityList               //O: TLV [1D][9-114]
	ForbiddenTaiForTheListOfForbiddenTrackingAreasForregionalProvisionOfService *TrackingAreaIdentityList               //O: TLV [1E][9-114]
	ExtendedCagInformationList                                                  []byte                                  //O: TLV-E [71][3-n]
	NsagInformation                                                             []byte                                  //O: TLV-E [7C][9-3143]
}

func (msg *RegistrationAccept) encode() (wire []byte, err error) {
	defer func() {
		if err != nil {
			err = nasError("encoding RegistrationAccept", err)
		}
	}()
	var buf []byte
	// M: LV[2]
	if buf, err = encodeLV(false, uint16(1), uint16(1), &msg.RegistrationResult); err != nil {
		err = nasError("encoding RegistrationResult [M LV 2]", err)
		return
	}
	wire = append(wire, buf...)

	// O: TLV-E[14]
	if msg.Guti != nil {
		if buf, err = encodeLV(true, uint16(11), uint16(11), msg.Guti); err != nil {
			err = nasError("encoding Guti [O TLV-E 14]", err)
			return
		}
		wire = append(append(wire, 0x77), buf...)
	}

	// O: TLV[5-47]
	if msg.EquivalentPlmns != nil {
		if buf, err = encodeLV(false, uint16(3), uint16(45), msg.EquivalentPlmns); err != nil {
			err = nasError("encoding EquivalentPlmns [O TLV 5-47]", err)
			return
		}
		wire = append(append(wire, 0x4A), buf...)
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

	// O: TLV[4-42]
	if msg.RejectedNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(40), msg.RejectedNssai); err != nil {
			err = nasError("encoding RejectedNssai [O TLV 4-42]", err)
			return
		}
		wire = append(append(wire, 0x11), buf...)
	}

	// O: TLV[4-146]
	if msg.ConfiguredNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(144), msg.ConfiguredNssai); err != nil {
			err = nasError("encoding ConfiguredNssai [O TLV 4-146]", err)
			return
		}
		wire = append(append(wire, 0x31), buf...)
	}

	// O: TLV[3-5]
	if msg.NetworkFeatureSupport != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(3), msg.NetworkFeatureSupport); err != nil {
			err = nasError("encoding NetworkFeatureSupport [O TLV 3-5]", err)
			return
		}
		wire = append(append(wire, 0x21), buf...)
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
	if msg.PduSessionReactivationResult != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(32), msg.PduSessionReactivationResult); err != nil {
			err = nasError("encoding PduSessionReactivationResult [O TLV 4-34]", err)
			return
		}
		wire = append(append(wire, 0x26), buf...)
	}

	// O: TLV-E[5-515]
	if msg.PduSessionReactivationResultErrorCause != nil {
		if buf, err = encodeLV(true, uint16(2), uint16(512), msg.PduSessionReactivationResultErrorCause); err != nil {
			err = nasError("encoding PduSessionReactivationResultErrorCause [O TLV-E 5-515]", err)
			return
		}
		wire = append(append(wire, 0x72), buf...)
	}

	// O: TLV-E[12-1715]
	if msg.LadnInformation != nil {
		if buf, err = encodeLV(true, uint16(9), uint16(1712), msg.LadnInformation); err != nil {
			err = nasError("encoding LadnInformation [O TLV-E 12-1715]", err)
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

	// O: TLV[6-114]
	if msg.ServiceAreaList != nil {
		if buf, err = encodeLV(false, uint16(4), uint16(112), msg.ServiceAreaList); err != nil {
			err = nasError("encoding ServiceAreaList [O TLV 6-114]", err)
			return
		}
		wire = append(append(wire, 0x27), buf...)
	}

	// O: TLV[3]
	if msg.T3512Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3512Value); err != nil {
			err = nasError("encoding T3512Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x5E), buf...)
	}

	// O: TLV[3]
	if msg.Non3gppDeRegistrationTimerValue != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.Non3gppDeRegistrationTimerValue); err != nil {
			err = nasError("encoding Non3gppDeRegistrationTimerValue [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x5D), buf...)
	}

	// O: TLV[3]
	if msg.T3502Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3502Value); err != nil {
			err = nasError("encoding T3502Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x16), buf...)
	}

	// O: TLV[5-50]
	if len(msg.EmergencyNumberList) > 0 {
		tmp := newBytesEncoder(msg.EmergencyNumberList)
		if buf, err = encodeLV(false, uint16(3), uint16(48), tmp); err != nil {
			err = nasError("encoding EmergencyNumberList [O TLV 5-50]", err)
			return
		}
		wire = append(append(wire, 0x34), buf...)
	}

	// O: TLV-E[7-65538]
	if len(msg.ExtendedEmergencyNumberList) > 0 {
		tmp := newBytesEncoder(msg.ExtendedEmergencyNumberList)
		if buf, err = encodeLV(true, uint16(4), uint16(0), tmp); err != nil {
			err = nasError("encoding ExtendedEmergencyNumberList [O TLV-E 7-65538]", err)
			return
		}
		wire = append(append(wire, 0x7A), buf...)
	}

	// O: TLV-E[20-n]
	if len(msg.SorTransparentContainer) > 0 {
		tmp := newBytesEncoder(msg.SorTransparentContainer)
		if buf, err = encodeLV(true, uint16(17), uint16(0), tmp); err != nil {
			err = nasError("encoding SorTransparentContainer [O TLV-E 20-n]", err)
			return
		}
		wire = append(append(wire, 0x73), buf...)
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

	// O: TV[1]
	if msg.NssaiInclusionMode != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0A<<4)|(uint8(*msg.NssaiInclusionMode)&0x0f))
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

	// O: TLV[3]
	if msg.NegotiatedDrxParameters != nil {
		tmp := newUint8Encoder(*msg.NegotiatedDrxParameters)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding NegotiatedDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x51), buf...)
	}

	// O: TV[1]
	if msg.Non3gppNwPolicies != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0D<<4)|(uint8(*msg.Non3gppNwPolicies)&0x0f))
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
	if len(msg.NegotiatedExtendedDrxParameters) > 0 {
		tmp := newBytesEncoder(msg.NegotiatedExtendedDrxParameters)
		if buf, err = encodeLV(false, uint16(1), uint16(2), tmp); err != nil {
			err = nasError("encoding NegotiatedExtendedDrxParameters [O TLV 3-4]", err)
			return
		}
		wire = append(append(wire, 0x6E), buf...)
	}

	// O: TLV[3]
	if msg.T3447Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3447Value); err != nil {
			err = nasError("encoding T3447Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6C), buf...)
	}

	// O: TLV[3]
	if msg.T3448Value != nil {
		if buf, err = encodeLV(false, uint16(1), uint16(1), msg.T3448Value); err != nil {
			err = nasError("encoding T3448Value [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x6B), buf...)
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

	// O: TV[1]
	if msg.UeRadioCapabilityIdDeletionIndication != nil {
		// fill lefthalf with IEI and righthalf with value
		wire = append(wire, (0x0E<<4)|(uint8(*msg.UeRadioCapabilityIdDeletionIndication)&0x0f))
	}

	// O: TLV[4-146]
	if msg.PendingNssai != nil {
		if buf, err = encodeLV(false, uint16(2), uint16(144), msg.PendingNssai); err != nil {
			err = nasError("encoding PendingNssai [O TLV 4-146]", err)
			return
		}
		wire = append(append(wire, 0x39), buf...)
	}

	// O: TLV-E[34-n]
	if len(msg.CipheringKeyData) > 0 {
		tmp := newBytesEncoder(msg.CipheringKeyData)
		if buf, err = encodeLV(true, uint16(31), uint16(0), tmp); err != nil {
			err = nasError("encoding CipheringKeyData [O TLV-E 34-n]", err)
			return
		}
		wire = append(append(wire, 0x74), buf...)
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

	// O: TLV[3]
	if msg.TruncatedSTmsiConfiguration != nil {
		tmp := newUint8Encoder(*msg.TruncatedSTmsiConfiguration)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding TruncatedSTmsiConfiguration [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x1B), buf...)
	}

	// O: TLV[3-n]
	if len(msg.NegotiatedWusAssistanceInformation) > 0 {
		tmp := newBytesEncoder(msg.NegotiatedWusAssistanceInformation)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding NegotiatedWusAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x1C), buf...)
	}

	// O: TLV[3]
	if msg.NegotiatedNbN1ModeDrxParameters != nil {
		tmp := newUint8Encoder(*msg.NegotiatedNbN1ModeDrxParameters)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding NegotiatedNbN1ModeDrxParameters [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x29), buf...)
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
		wire = append(append(wire, 0x7B), buf...)
	}

	// O: TLV[3-n]
	if len(msg.NegotiatedPeipsAssistanceInformation) > 0 {
		tmp := newBytesEncoder(msg.NegotiatedPeipsAssistanceInformation)
		if buf, err = encodeLV(false, uint16(1), uint16(0), tmp); err != nil {
			err = nasError("encoding NegotiatedPeipsAssistanceInformation [O TLV 3-n]", err)
			return
		}
		wire = append(append(wire, 0x33), buf...)
	}

	// O: TLV[3]
	if msg.AdditionalRequestResult != nil {
		tmp := newUint8Encoder(*msg.AdditionalRequestResult)
		if buf, err = encodeLV(false, uint16(1), uint16(1), tmp); err != nil {
			err = nasError("encoding AdditionalRequestResult [O TLV 3]", err)
			return
		}
		wire = append(append(wire, 0x35), buf...)
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

	// O: TLV-E[3-n]
	if len(msg.ExtendedCagInformationList) > 0 {
		tmp := newBytesEncoder(msg.ExtendedCagInformationList)
		if buf, err = encodeLV(true, uint16(0), uint16(0), tmp); err != nil {
			err = nasError("encoding ExtendedCagInformationList [O TLV-E 3-n]", err)
			return
		}
		wire = append(append(wire, 0x71), buf...)
	}

	// O: TLV-E[9-3143]
	if len(msg.NsagInformation) > 0 {
		tmp := newBytesEncoder(msg.NsagInformation)
		if buf, err = encodeLV(true, uint16(6), uint16(3140), tmp); err != nil {
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
	// M LV[2]
	if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), &msg.RegistrationResult); err != nil {
		err = nasError("decoding RegistrationResult [M LV 2]", err)
		return
	}
	offset += consumed
	for offset < wireLen {
		iei := getIei(wire[offset])
		switch iei {
		case 0x77: //O: TLV-E[14]
			offset++ //consume IEI
			v := new(MobileIdentity)
			if consumed, err = decodeLV(wire[offset:], true, uint16(11), uint16(11), v); err != nil {
				err = nasError("decoding Guti [O TLV-E 14]", err)
				return
			}
			offset += consumed
			msg.Guti = v
		case 0x4A: //O: TLV[5-47]
			offset++ //consume IEI
			v := new(PlmnList)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(45), v); err != nil {
				err = nasError("decoding EquivalentPlmns [O TLV 5-47]", err)
				return
			}
			offset += consumed
			msg.EquivalentPlmns = v
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
		case 0x11: //O: TLV[4-42]
			offset++ //consume IEI
			v := new(RejectedNssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(40), v); err != nil {
				err = nasError("decoding RejectedNssai [O TLV 4-42]", err)
				return
			}
			offset += consumed
			msg.RejectedNssai = v
		case 0x31: //O: TLV[4-146]
			offset++ //consume IEI
			v := new(Nssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(144), v); err != nil {
				err = nasError("decoding ConfiguredNssai [O TLV 4-146]", err)
				return
			}
			offset += consumed
			msg.ConfiguredNssai = v
		case 0x21: //O: TLV[3-5]
			offset++ //consume IEI
			v := new(NetworkFeatureSupport)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(3), v); err != nil {
				err = nasError("decoding NetworkFeatureSupport [O TLV 3-5]", err)
				return
			}
			offset += consumed
			msg.NetworkFeatureSupport = v
		case 0x50: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(PduSessionStatus)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionStatus [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionStatus = v
		case 0x26: //O: TLV[4-34]
			offset++ //consume IEI
			v := new(PduSessionReactivationResult)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(32), v); err != nil {
				err = nasError("decoding PduSessionReactivationResult [O TLV 4-34]", err)
				return
			}
			offset += consumed
			msg.PduSessionReactivationResult = v
		case 0x72: //O: TLV-E[5-515]
			offset++ //consume IEI
			v := new(PduSessionReactivationResultErrorCause)
			if consumed, err = decodeLV(wire[offset:], true, uint16(2), uint16(512), v); err != nil {
				err = nasError("decoding PduSessionReactivationResultErrorCause [O TLV-E 5-515]", err)
				return
			}
			offset += consumed
			msg.PduSessionReactivationResultErrorCause = v
		case 0x79: //O: TLV-E[12-1715]
			offset++ //consume IEI
			v := new(LadnInformation)
			if consumed, err = decodeLV(wire[offset:], true, uint16(9), uint16(1712), v); err != nil {
				err = nasError("decoding LadnInformation [O TLV-E 12-1715]", err)
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
		case 0x27: //O: TLV[6-114]
			offset++ //consume IEI
			v := new(ServiceAreaList)
			if consumed, err = decodeLV(wire[offset:], false, uint16(4), uint16(112), v); err != nil {
				err = nasError("decoding ServiceAreaList [O TLV 6-114]", err)
				return
			}
			offset += consumed
			msg.ServiceAreaList = v
		case 0x5E: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3512Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3512Value = v
		case 0x5D: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer2)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding Non3gppDeRegistrationTimerValue [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.Non3gppDeRegistrationTimerValue = v
		case 0x16: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer2)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3502Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3502Value = v
		case 0x34: //O: TLV[5-50]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(48), v); err != nil {
				err = nasError("decoding EmergencyNumberList [O TLV 5-50]", err)
				return
			}
			offset += consumed
			msg.EmergencyNumberList = []byte(*v)
		case 0x7A: //O: TLV-E[7-65538]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedEmergencyNumberList [O TLV-E 7-65538]", err)
				return
			}
			offset += consumed
			msg.ExtendedEmergencyNumberList = []byte(*v)
		case 0x73: //O: TLV-E[20-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(17), uint16(0), v); err != nil {
				err = nasError("decoding SorTransparentContainer [O TLV-E 20-n]", err)
				return
			}
			offset += consumed
			msg.SorTransparentContainer = []byte(*v)
		case 0x78: //O: TLV-E[7-1503]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(4), uint16(1500), v); err != nil {
				err = nasError("decoding EapMessage [O TLV-E 7-1503]", err)
				return
			}
			offset += consumed
			msg.EapMessage = []byte(*v)
		case 0x0A: //O: TV[1]
			msg.NssaiInclusionMode = new(uint8)
			*msg.NssaiInclusionMode = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x76: //O: TLV-E[3-8323]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(8320), v); err != nil {
				err = nasError("decoding OperatorDefinedAccessCategoryDefinitions [O TLV-E 3-8323]", err)
				return
			}
			offset += consumed
			msg.OperatorDefinedAccessCategoryDefinitions = []byte(*v)
		case 0x51: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding NegotiatedDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.NegotiatedDrxParameters = (*uint8)(v)
		case 0x0D: //O: TV[1]
			msg.Non3gppNwPolicies = new(uint8)
			*msg.Non3gppNwPolicies = wire[offset] & 0x0f //take right 1/2
			offset++
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
				err = nasError("decoding NegotiatedExtendedDrxParameters [O TLV 3-4]", err)
				return
			}
			offset += consumed
			msg.NegotiatedExtendedDrxParameters = []byte(*v)
		case 0x6C: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer3)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3447Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3447Value = v
		case 0x6B: //O: TLV[3]
			offset++ //consume IEI
			v := new(GprsTimer2)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding T3448Value [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.T3448Value = v
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
		case 0x0E: //O: TV[1]
			msg.UeRadioCapabilityIdDeletionIndication = new(uint8)
			*msg.UeRadioCapabilityIdDeletionIndication = wire[offset] & 0x0f //take right 1/2
			offset++
		case 0x39: //O: TLV[4-146]
			offset++ //consume IEI
			v := new(Nssai)
			if consumed, err = decodeLV(wire[offset:], false, uint16(2), uint16(144), v); err != nil {
				err = nasError("decoding PendingNssai [O TLV 4-146]", err)
				return
			}
			offset += consumed
			msg.PendingNssai = v
		case 0x74: //O: TLV-E[34-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(31), uint16(0), v); err != nil {
				err = nasError("decoding CipheringKeyData [O TLV-E 34-n]", err)
				return
			}
			offset += consumed
			msg.CipheringKeyData = []byte(*v)
		case 0x75: //O: TLV-E[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding CagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.CagInformationList = []byte(*v)
		case 0x1B: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding TruncatedSTmsiConfiguration [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.TruncatedSTmsiConfiguration = (*uint8)(v)
		case 0x1C: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NegotiatedWusAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.NegotiatedWusAssistanceInformation = []byte(*v)
		case 0x29: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding NegotiatedNbN1ModeDrxParameters [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.NegotiatedNbN1ModeDrxParameters = (*uint8)(v)
		case 0x68: //O: TLV[5-90]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(3), uint16(88), v); err != nil {
				err = nasError("decoding ExtendedRejectedNssai [O TLV 5-90]", err)
				return
			}
			offset += consumed
			msg.ExtendedRejectedNssai = []byte(*v)
		case 0x7B: //O: TLV-E[6-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(3), uint16(0), v); err != nil {
				err = nasError("decoding ServiceLevelAaContainer [O TLV-E 6-n]", err)
				return
			}
			offset += consumed
			msg.ServiceLevelAaContainer = []byte(*v)
		case 0x33: //O: TLV[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(0), v); err != nil {
				err = nasError("decoding NegotiatedPeipsAssistanceInformation [O TLV 3-n]", err)
				return
			}
			offset += consumed
			msg.NegotiatedPeipsAssistanceInformation = []byte(*v)
		case 0x35: //O: TLV[3]
			offset++ //consume IEI
			v := new(uint8Decoder)
			if consumed, err = decodeLV(wire[offset:], false, uint16(1), uint16(1), v); err != nil {
				err = nasError("decoding AdditionalRequestResult [O TLV 3]", err)
				return
			}
			offset += consumed
			msg.AdditionalRequestResult = (*uint8)(v)
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
		case 0x71: //O: TLV-E[3-n]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(0), uint16(0), v); err != nil {
				err = nasError("decoding ExtendedCagInformationList [O TLV-E 3-n]", err)
				return
			}
			offset += consumed
			msg.ExtendedCagInformationList = []byte(*v)
		case 0x7C: //O: TLV-E[9-3143]
			offset++ //consume IEI
			v := new(bytesDecoder)
			if consumed, err = decodeLV(wire[offset:], true, uint16(6), uint16(3140), v); err != nil {
				err = nasError("decoding NsagInformation [O TLV-E 9-3143]", err)
				return
			}
			offset += consumed
			msg.NsagInformation = []byte(*v)
		default:
			err = ErrUnknownIei
			return
		}
	}
	return
}
