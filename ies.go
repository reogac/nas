/** this file was generated at 2024-12-16 17:55:27.320479 by tqtung@etri.re.kr **/

package nas

/*
// 9.11.2.1 Additional information [O TLV 3-n]
type AdditionalInformation struct {
Bytes []byte
}
func (ie *AdditionalInformation) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *AdditionalInformation) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.2.10 Service-level-AA container [O TLV-E 6-n]
type ServiceLevelAaContainer struct {
Bytes []byte
}
func (ie *ServiceLevelAaContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ServiceLevelAaContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.2.1A Access type [M V 1/2]
type AccessType struct {
Value uint8
}
func (ie *AccessType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AccessType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.2.1B DNN [O TLV 3-102]
type Dnn struct {
Bytes []byte
}
func (ie *Dnn) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *Dnn) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.2.2 EAP message [O TLV-E 7-1503]
type EapMessage struct {
Bytes []byte
}
func (ie *EapMessage) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *EapMessage) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.2.3 GPRS timer [O TV 2]
type GprsTimer struct {
Value uint8
}
func (ie *GprsTimer) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *GprsTimer) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.2.4 GPRS timer 2 [O TLV 3]
type GprsTimer2 struct {
Value uint8
}
func (ie *GprsTimer2) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *GprsTimer2) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.2.5 GPRS timer 3 [O TLV 3]
type GprsTimer3 struct {
Value uint8
}
func (ie *GprsTimer3) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *GprsTimer3) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.2.8 S-NSSAI [O TLV 3-10]
type SNssai struct {
Bytes []byte
}
func (ie *SNssai) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *SNssai) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.1 5GMM capability [O TLV 3-15]
type GmmCapability struct {
Bytes []byte
}
func (ie *GmmCapability) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *GmmCapability) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.10 ABBA [M LV 3-n]
type Abba struct {
Bytes []byte
}
func (ie *Abba) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *Abba) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.12 Additional 5G security information [O TLV 3]
type AdditionalSecurityInformation struct {
Value uint8
}
func (ie *AdditionalSecurityInformation) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AdditionalSecurityInformation) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.12A Additional information requested [O TLV 3]
type AdditionalInformationRequested struct {
Value uint8
}
func (ie *AdditionalInformationRequested) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AdditionalInformationRequested) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.13 Allowed PDU session status [O TLV 4-34]
type AllowedPduSessionStatus struct {
Bytes []byte
}
func (ie *AllowedPduSessionStatus) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *AllowedPduSessionStatus) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.14 Authentication failure parameter [O TLV 16]
type AuthenticationFailureParameter struct {
Bytes []byte
}
func (ie *AuthenticationFailureParameter) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *AuthenticationFailureParameter) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.15 Authentication parameter AUTN [O TLV 18]
type AuthenticationParameterAutn struct {
Bytes []byte
}
func (ie *AuthenticationParameterAutn) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *AuthenticationParameterAutn) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.16 Authentication parameter RAND [O TV 17]
type AuthenticationParameterRand struct {
Bytes []byte
}
func (ie *AuthenticationParameterRand) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *AuthenticationParameterRand) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.17 Authentication response parameter [O TLV 18]
type AuthenticationResponseParameter struct {
Bytes []byte
}
func (ie *AuthenticationResponseParameter) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *AuthenticationResponseParameter) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.18 Configuration update indication [O TV 1]
type ConfigurationUpdateIndication struct {
Value uint8
}
func (ie *ConfigurationUpdateIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *ConfigurationUpdateIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.18A CAG information list [O TLV-E 3-n]
type CagInformationList struct {
Bytes []byte
}
func (ie *CagInformationList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *CagInformationList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.18C Ciphering key data [O TLV-E 34-n]
type CipheringKeyData struct {
Bytes []byte
}
func (ie *CipheringKeyData) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *CipheringKeyData) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.19 Daylight saving time [O TLV 3]
type DaylightSavingTime struct {
Value uint8
}
func (ie *DaylightSavingTime) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *DaylightSavingTime) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.2 5GMM cause [M V 1]
type GmmCause struct {
Value uint8
}
func (ie *GmmCause) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *GmmCause) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.20 De-registration type [M V 1/2]
type DeRegistrationType struct {
Value uint8
}
func (ie *DeRegistrationType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *DeRegistrationType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.23 Emergency number list [O TLV 5-50]
type EmergencyNumberList struct {
Bytes []byte
}
func (ie *EmergencyNumberList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *EmergencyNumberList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.23A EPS bearer context status [O TLV 4]
type EpsBearerContextStatus struct {
Value uint16
}
func (ie *EpsBearerContextStatus) encode() (wire []byte, err error) {
                wire = twoBytes(ie.Value)
return
}

func (ie *EpsBearerContextStatus) decode(wire []byte) (err error) {
                ie.Value, err = fromTwoBytes(wire)
return }
*/

/*
// 9.11.3.24 EPS NAS message container [O TLV-E 4-n]
type EpsNasMessageContainer struct {
Bytes []byte
}
func (ie *EpsNasMessageContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *EpsNasMessageContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.25 EPS NAS security algorithms [O TV 2]
type EpsNasSecurityAlgorithms struct {
Value uint8
}
func (ie *EpsNasSecurityAlgorithms) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *EpsNasSecurityAlgorithms) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.26 Extended emergency number list [O TLV-E 7-65538]
type ExtendedEmergencyNumberList struct {
Bytes []byte
}
func (ie *ExtendedEmergencyNumberList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ExtendedEmergencyNumberList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.26A Extended DRX parameters [O TLV 3-4]
type ExtendedDrxParameters struct {
Bytes []byte
}
func (ie *ExtendedDrxParameters) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ExtendedDrxParameters) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.28 IMEISV request [O TV 1]
type ImeisvRequest struct {
Value uint8
}
func (ie *ImeisvRequest) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *ImeisvRequest) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.29 LADN indication [O TLV-E 3-811]
type LadnIndication struct {
Bytes []byte
}
func (ie *LadnIndication) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *LadnIndication) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.2A 5GS DRX parameters [O TLV 3]
type DrxParameters struct {
Value uint8
}
func (ie *DrxParameters) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *DrxParameters) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.3 5GS identity type [M V 1/2]
type IdentityType struct {
Value uint8
}
func (ie *IdentityType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *IdentityType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.30 LADN information [O TLV-E 12-1715]
type LadnInformation struct {
Bytes []byte
}
func (ie *LadnInformation) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *LadnInformation) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.31 MICO indication [O TV 1]
type MicoIndication struct {
Value uint8
}
func (ie *MicoIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *MicoIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.31A MA PDU session information [O TV 1]
type MaPduSessionInformation struct {
Value uint8
}
func (ie *MaPduSessionInformation) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *MaPduSessionInformation) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.31B Mapped NSSAI [O TLV 3-42]
type MappedNssai struct {
Bytes []byte
}
func (ie *MappedNssai) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *MappedNssai) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.31C Mobile station classmark 2 [O TLV 5]
type MobileStationClassmark2 struct {
Bytes []byte
}
func (ie *MobileStationClassmark2) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *MobileStationClassmark2) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.32 key set identifier [M V 1/2]
type KeySetIdentifier struct {
Value uint8
}
func (ie *KeySetIdentifier) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *KeySetIdentifier) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.33 message container [O TLV-E 4-n]
type MessageContainer struct {
Bytes []byte
}
func (ie *MessageContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *MessageContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.34 security algorithms [M V 1]
type SecurityAlgorithms struct {
Value uint8
}
func (ie *SecurityAlgorithms) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *SecurityAlgorithms) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.35 Network name [O TLV 3-n]
type NetworkName struct {
Bytes []byte
}
func (ie *NetworkName) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *NetworkName) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.36 Network slicing indication [O TV 1]
type NetworkSlicingIndication struct {
Value uint8
}
func (ie *NetworkSlicingIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *NetworkSlicingIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.36A Non-3GPP NW provided policies [O TV 1]
type Non3gppNwProvidedPolicies struct {
Value uint8
}
func (ie *Non3gppNwProvidedPolicies) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *Non3gppNwProvidedPolicies) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.37 NSSAI [O TLV 4-74]
type Nssai struct {
Bytes []byte
}
func (ie *Nssai) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *Nssai) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.37A NSSAI inclusion mode [O TV 1]
type NssaiInclusionMode struct {
Value uint8
}
func (ie *NssaiInclusionMode) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *NssaiInclusionMode) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.38 Operator-defined access category definitions [O TLV-E 3-8323]
type OperatorDefinedAccessCategoryDefinitions struct {
Bytes []byte
}
func (ie *OperatorDefinedAccessCategoryDefinitions) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *OperatorDefinedAccessCategoryDefinitions) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.39 Payload container [O TLV-E 4-65538]
type PayloadContainer struct {
Bytes []byte
}
func (ie *PayloadContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PayloadContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.4 5GS mobile identity [M LV-E 6-n]
type MobileIdentity struct {
Bytes []byte
}
func (ie *MobileIdentity) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *MobileIdentity) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.40 Payload container type [O TV 1]
type PayloadContainerType struct {
Value uint8
}
func (ie *PayloadContainerType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *PayloadContainerType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.41 PDU session identity 2 [C TV 2]
type PduSessionIdentity2 struct {
Value uint8
}
func (ie *PduSessionIdentity2) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *PduSessionIdentity2) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.42 PDU session reactivation result [O TLV 4-34]
type PduSessionReactivationResult struct {
Bytes []byte
}
func (ie *PduSessionReactivationResult) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PduSessionReactivationResult) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.43 PDU session reactivation result error cause [O TLV-E 5-515]
type PduSessionReactivationResultErrorCause struct {
Bytes []byte
}
func (ie *PduSessionReactivationResultErrorCause) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PduSessionReactivationResultErrorCause) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.44 PDU session status [O TLV 4-34]
type PduSessionStatus struct {
Bytes []byte
}
func (ie *PduSessionStatus) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PduSessionStatus) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.45 PLMN list [O TLV 5-47]
type PlmnList struct {
Bytes []byte
}
func (ie *PlmnList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PlmnList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.46 Rejected NSSAI [O TLV 4-42]
type RejectedNssai struct {
Bytes []byte
}
func (ie *RejectedNssai) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *RejectedNssai) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.46A Release assistance indication [O TV 1]
type ReleaseAssistanceIndication struct {
Value uint8
}
func (ie *ReleaseAssistanceIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *ReleaseAssistanceIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.47 Request type [O TV 1]
type RequestType struct {
Value uint8
}
func (ie *RequestType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *RequestType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.48 S1 UE network capability [O TLV 4-15]
type S1UeNetworkCapability struct {
Bytes []byte
}
func (ie *S1UeNetworkCapability) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *S1UeNetworkCapability) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.48A S1 UE security capability [O TLV 4-7]
type S1UeSecurityCapability struct {
Bytes []byte
}
func (ie *S1UeSecurityCapability) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *S1UeSecurityCapability) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.49 Service area list [O TLV 6-114]
type ServiceAreaList struct {
Bytes []byte
}
func (ie *ServiceAreaList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ServiceAreaList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.5 5GS network feature support [O TLV 3-5]
type NetworkFeatureSupport struct {
Bytes []byte
}
func (ie *NetworkFeatureSupport) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *NetworkFeatureSupport) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.50 Service type [M V 1/2]
type ServiceType struct {
Value uint8
}
func (ie *ServiceType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *ServiceType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.50A SMS indication [O TV 1]
type SmsIndication struct {
Value uint8
}
func (ie *SmsIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *SmsIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.51 SOR transparent container [O TLV-E 20-n]
type SorTransparentContainer struct {
Bytes []byte
}
func (ie *SorTransparentContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *SorTransparentContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.51A Supported codec list [O TLV 5-n]
type SupportedCodecList struct {
Bytes []byte
}
func (ie *SupportedCodecList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *SupportedCodecList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.52 Time zone [O TV 2]
type TimeZone struct {
Value uint8
}
func (ie *TimeZone) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *TimeZone) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.53 Time zone and time [O TV 8]
type TimeZoneAndTime struct {
Bytes []byte
}
func (ie *TimeZoneAndTime) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *TimeZoneAndTime) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.54 UE security capability [O TLV 4-10]
type UeSecurityCapability struct {
Bytes []byte
}
func (ie *UeSecurityCapability) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *UeSecurityCapability) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.55 UE usage setting [O TLV 3]
type UeUsageSetting struct {
Value uint8
}
func (ie *UeUsageSetting) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *UeUsageSetting) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.56 UE status [O TLV 3]
type UeStatus struct {
Value uint8
}
func (ie *UeStatus) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *UeStatus) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.57 Uplink data status [O TLV 4-34]
type UplinkDataStatus struct {
Bytes []byte
}
func (ie *UplinkDataStatus) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *UplinkDataStatus) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.6 5GS registration result [M LV 2]
type RegistrationResult struct {
Value uint8
}
func (ie *RegistrationResult) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *RegistrationResult) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.68 UE radio capability ID [O TLV 3-n]
type UeRadioCapabilityId struct {
Bytes []byte
}
func (ie *UeRadioCapabilityId) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *UeRadioCapabilityId) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.69 UE radio capability ID deletion indication [O TV 1]
type UeRadioCapabilityIdDeletionIndication struct {
Value uint8
}
func (ie *UeRadioCapabilityIdDeletionIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *UeRadioCapabilityIdDeletionIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.7 5GS registration type [M V 1/2]
type RegistrationType struct {
Value uint8
}
func (ie *RegistrationType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *RegistrationType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.70 Truncated 5G-S-TMSI configuration [O TLV 3]
type TruncatedSTmsiConfiguration struct {
Value uint8
}
func (ie *TruncatedSTmsiConfiguration) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *TruncatedSTmsiConfiguration) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.71 WUS assistance information [O TLV 3-n]
type WusAssistanceInformation struct {
Bytes []byte
}
func (ie *WusAssistanceInformation) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *WusAssistanceInformation) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.72 N5GC indication [O TV 1]
type N5gcIndication struct {
Value uint8
}
func (ie *N5gcIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *N5gcIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.73 NB-N1 mode DRX parameters [O TLV 3]
type NbN1ModeDrxParameters struct {
Value uint8
}
func (ie *NbN1ModeDrxParameters) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *NbN1ModeDrxParameters) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.74 Additional configuration indication [O TV 1]
type AdditionalConfigurationIndication struct {
Value uint8
}
func (ie *AdditionalConfigurationIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AdditionalConfigurationIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.75 Extended rejected NSSAI [O TLV 5-90]
type ExtendedRejectedNssai struct {
Bytes []byte
}
func (ie *ExtendedRejectedNssai) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ExtendedRejectedNssai) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.76 UE request type [O TLV 3]
type UeRequestType struct {
Value uint8
}
func (ie *UeRequestType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *UeRequestType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.77 Paging restriction [O TLV 3-35]
type PagingRestriction struct {
Bytes []byte
}
func (ie *PagingRestriction) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PagingRestriction) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.79 NID [O TLV 8]
type Nid struct {
Bytes []byte
}
func (ie *Nid) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *Nid) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.8 5GS tracking area identity [O TV 7]
type TrackingAreaIdentity struct {
Bytes []byte
}
func (ie *TrackingAreaIdentity) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *TrackingAreaIdentity) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.80 PEIPS assistance information [O TLV 3-n]
type PeipsAssistanceInformation struct {
Bytes []byte
}
func (ie *PeipsAssistanceInformation) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PeipsAssistanceInformation) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.81 5GS additional request result [O TLV 3]
type AdditionalRequestResult struct {
Value uint8
}
func (ie *AdditionalRequestResult) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AdditionalRequestResult) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.82 NSSRG information [O TLV-E 7-4099]
type NssrgInformation struct {
Bytes []byte
}
func (ie *NssrgInformation) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *NssrgInformation) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.83 List of PLMNs to be used in disaster condition [O TLV 2-n]
type ListOfPlmnsToBeUsedInDisasterCondition struct {
Bytes []byte
}
func (ie *ListOfPlmnsToBeUsedInDisasterCondition) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ListOfPlmnsToBeUsedInDisasterCondition) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.84 Registration wait range [O TLV 4]
type RegistrationWaitRange struct {
Value uint16
}
func (ie *RegistrationWaitRange) encode() (wire []byte, err error) {
                wire = twoBytes(ie.Value)
return
}

func (ie *RegistrationWaitRange) decode(wire []byte) (err error) {
                ie.Value, err = fromTwoBytes(wire)
return }
*/

/*
// 9.11.3.85 PLMN identity [O TLV 5]
type PlmnIdentity struct {
Bytes []byte
}
func (ie *PlmnIdentity) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PlmnIdentity) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.86 Extended CAG information list [O TLV-E 3-n]
type ExtendedCagInformationList struct {
Bytes []byte
}
func (ie *ExtendedCagInformationList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ExtendedCagInformationList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.87 NSAG information [O TLV-E 9-3143]
type NsagInformation struct {
Bytes []byte
}
func (ie *NsagInformation) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *NsagInformation) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.9 5GS tracking area identity list [O TLV 9-114]
type TrackingAreaIdentityList struct {
Bytes []byte
}
func (ie *TrackingAreaIdentityList) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *TrackingAreaIdentityList) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.3.91 Priority indicator [O TV 1]
type PriorityIndicator struct {
Value uint8
}
func (ie *PriorityIndicator) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *PriorityIndicator) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.3.9A 5GS update type [O TLV 3]
type UpdateType struct {
Value uint8
}
func (ie *UpdateType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *UpdateType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.1 5GSM capability [O TLV 3-15]
type GsmCapability struct {
Bytes []byte
}
func (ie *GsmCapability) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *GsmCapability) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.10 PDU address [O TLV 11]
type PduAddress struct {
Bytes []byte
}
func (ie *PduAddress) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PduAddress) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.11 PDU session type [O TV 1]
type PduSessionType struct {
Value uint8
}
func (ie *PduSessionType) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *PduSessionType) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.12 QoS flow descriptions [O TLV-E 6-65538]
type QosFlowDescriptions struct {
Bytes []byte
}
func (ie *QosFlowDescriptions) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *QosFlowDescriptions) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.13 QoS rules [M LV-E 6-65538]
type QosRules struct {
Bytes []byte
}
func (ie *QosRules) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *QosRules) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.14 Session-AMBR [M LV 7]
type SessionAmbr struct {
Bytes []byte
}
func (ie *SessionAmbr) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *SessionAmbr) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.15 SM PDU DN request container [O TLV 3-255]
type SmPduDnRequestContainer struct {
Bytes []byte
}
func (ie *SmPduDnRequestContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *SmPduDnRequestContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.16 SSC mode [O TV 1]
type SscMode struct {
Value uint8
}
func (ie *SscMode) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *SscMode) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.17 Re-attempt indicator [O TLV 3]
type ReAttemptIndicator struct {
Value uint8
}
func (ie *ReAttemptIndicator) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *ReAttemptIndicator) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.18 5GSM network feature support [O TLV 3-15]
type GsmNetworkFeatureSupport struct {
Bytes []byte
}
func (ie *GsmNetworkFeatureSupport) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *GsmNetworkFeatureSupport) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.2 5GSM cause [O TV 2]
type GsmCause struct {
Value uint8
}
func (ie *GsmCause) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *GsmCause) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.20 Serving PLMN rate control [O TLV 4]
type ServingPlmnRateControl struct {
Value uint16
}
func (ie *ServingPlmnRateControl) encode() (wire []byte, err error) {
                wire = twoBytes(ie.Value)
return
}

func (ie *ServingPlmnRateControl) decode(wire []byte) (err error) {
                ie.Value, err = fromTwoBytes(wire)
return }
*/

/*
// 9.11.4.21 5GSM congestion re-attempt indicator [O TLV 3]
type GsmCongestionReAttemptIndicator struct {
Value uint8
}
func (ie *GsmCongestionReAttemptIndicator) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *GsmCongestionReAttemptIndicator) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.22 ATSSS container [O TLV-E 3-65538]
type AtsssContainer struct {
Bytes []byte
}
func (ie *AtsssContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *AtsssContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.23 Control plane only indication [O TV 1]
type ControlPlaneOnlyIndication struct {
Value uint8
}
func (ie *ControlPlaneOnlyIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *ControlPlaneOnlyIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.24 IP header compression configuration [O TLV 5-257]
type IpHeaderCompressionConfiguration struct {
Bytes []byte
}
func (ie *IpHeaderCompressionConfiguration) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *IpHeaderCompressionConfiguration) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.24 Header compression configuration [O TLV 5-257]
type HeaderCompressionConfiguration struct {
Bytes []byte
}
func (ie *HeaderCompressionConfiguration) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *HeaderCompressionConfiguration) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.25 DS-TT Ethernet port MAC address [O TLV 8]
type DsTtEthernetPortMacAddress struct {
Bytes []byte
}
func (ie *DsTtEthernetPortMacAddress) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *DsTtEthernetPortMacAddress) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.26 UE-DS-TT residence time [O TLV 10]
type UeDsTtResidenceTime struct {
Bytes []byte
}
func (ie *UeDsTtResidenceTime) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *UeDsTtResidenceTime) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.27 Port management information container [O TLV-E 8-65538]
type PortManagementInformationContainer struct {
Bytes []byte
}
func (ie *PortManagementInformationContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *PortManagementInformationContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.28 Ethernet header compression configuration [O TLV 3]
type EthernetHeaderCompressionConfiguration struct {
Value uint8
}
func (ie *EthernetHeaderCompressionConfiguration) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *EthernetHeaderCompressionConfiguration) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.3 Always-on PDU session indication [O TV 1]
type AlwaysOnPduSessionIndication struct {
Value uint8
}
func (ie *AlwaysOnPduSessionIndication) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AlwaysOnPduSessionIndication) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.30 Requested MBS container [O TLV-E 8-65538]
type RequestedMbsContainer struct {
Bytes []byte
}
func (ie *RequestedMbsContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *RequestedMbsContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.31 Received MBS container [O TLV-E 9-65538]
type ReceivedMbsContainer struct {
Bytes []byte
}
func (ie *ReceivedMbsContainer) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ReceivedMbsContainer) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.32 PDU session pair ID [O TLV 3]
type PduSessionPairId struct {
Value uint8
}
func (ie *PduSessionPairId) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *PduSessionPairId) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.33 RSN [O TLV 3]
type Rsn struct {
Value uint8
}
func (ie *Rsn) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *Rsn) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.4 Always-on PDU session requested [O TV 1]
type AlwaysOnPduSessionRequested struct {
Value uint8
}
func (ie *AlwaysOnPduSessionRequested) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AlwaysOnPduSessionRequested) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.5 Allowed SSC mode [O TV 1]
type AllowedSscMode struct {
Value uint8
}
func (ie *AllowedSscMode) encode() (wire []byte, err error) {
                wire = []byte{ie.Value}
return
}

func (ie *AllowedSscMode) decode(wire []byte) (err error) {
                if len(wire) != 1 {
err = ErrInvalidSize
return
}
                ie.Value = wire[0]
return }
*/

/*
// 9.11.4.6 Extended protocol configuration options [O TLV-E 4-65538]
type ExtendedProtocolConfigurationOptions struct {
Bytes []byte
}
func (ie *ExtendedProtocolConfigurationOptions) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *ExtendedProtocolConfigurationOptions) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.7 Integrity protection maximum data rate [M V 2]
type IntegrityProtectionMaximumDataRate struct {
Value uint16
}
func (ie *IntegrityProtectionMaximumDataRate) encode() (wire []byte, err error) {
                wire = twoBytes(ie.Value)
return
}

func (ie *IntegrityProtectionMaximumDataRate) decode(wire []byte) (err error) {
                ie.Value, err = fromTwoBytes(wire)
return }
*/

/*
// 9.11.4.8 Mapped EPS bearer contexts [O TLV-E 7-65538]
type MappedEpsBearerContexts struct {
Bytes []byte
}
func (ie *MappedEpsBearerContexts) encode() (wire []byte, err error) {
                wire = ie.Bytes
return
 }

func (ie *MappedEpsBearerContexts) decode(wire []byte) (err error) {
                ie.Bytes = make([]byte, len(wire))
copy(ie.Bytes, wire)
return }
*/

/*
// 9.11.4.9 Maximum number of supported packet filters [O TV 3]
type MaximumNumberOfSupportedPacketFilters struct {
Value uint16
}
func (ie *MaximumNumberOfSupportedPacketFilters) encode() (wire []byte, err error) {
                wire = twoBytes(ie.Value)
return
}

func (ie *MaximumNumberOfSupportedPacketFilters) decode(wire []byte) (err error) {
                ie.Value, err = fromTwoBytes(wire)
return }
*/
