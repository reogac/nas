/**generated time: 2024-07-17 15:11:00.936588**/

package nas

type DecodedGmmMessage struct {
	MsgType   uint8 // message type
	MacFailed bool  // indicate that integrity check has failed
	SecHeader uint8 //outer security header type

	RegistrationRequest         *RegistrationRequest
	RegistrationAccept          *RegistrationAccept
	RegistrationComplete        *RegistrationComplete
	RegistrationReject          *RegistrationReject
	DeregistrationRequestFromUe *DeregistrationRequestFromUe
	DeregistrationAcceptFromUe  *DeregistrationAcceptFromUe
	DeregistrationRequestToUe   *DeregistrationRequestToUe
	DeregistrationAcceptToUe    *DeregistrationAcceptToUe
	ServiceRequest              *ServiceRequest
	ServiceReject               *ServiceReject
	ServiceAccept               *ServiceAccept
	ConfigurationUpdateCommand  *ConfigurationUpdateCommand
	ConfigurationUpdateComplete *ConfigurationUpdateComplete
	AuthenticationRequest       *AuthenticationRequest
	AuthenticationResponse      *AuthenticationResponse
	AuthenticationReject        *AuthenticationReject
	AuthenticationFailure       *AuthenticationFailure
	AuthenticationResult        *AuthenticationResult
	IdentityRequest             *IdentityRequest
	IdentityResponse            *IdentityResponse
	SecurityModeCommand         *SecurityModeCommand
	SecurityModeComplete        *SecurityModeComplete
	SecurityModeReject          *SecurityModeReject
	GmmStatus                   *GmmStatus
	Notification                *Notification
	NotificationResponse        *NotificationResponse
	UlNasTransport              *UlNasTransport
	DlNasTransport              *DlNasTransport
}
type DecodedGsmMessage struct {
	MsgType uint8 //message type

	PduSessionEstablishmentRequest      *PduSessionEstablishmentRequest
	PduSessionEstablishmentAccept       *PduSessionEstablishmentAccept
	PduSessionEstablishmentReject       *PduSessionEstablishmentReject
	PduSessionAuthenticationCommand     *PduSessionAuthenticationCommand
	PduSessionAuthenticationComplete    *PduSessionAuthenticationComplete
	PduSessionAuthenticationResult      *PduSessionAuthenticationResult
	PduSessionModificationRequest       *PduSessionModificationRequest
	PduSessionModificationReject        *PduSessionModificationReject
	PduSessionModificationCommand       *PduSessionModificationCommand
	PduSessionModificationComplete      *PduSessionModificationComplete
	PduSessionModificationCommandReject *PduSessionModificationCommandReject
	PduSessionReleaseRequest            *PduSessionReleaseRequest
	PduSessionReleaseReject             *PduSessionReleaseReject
	PduSessionReleaseCommand            *PduSessionReleaseCommand
	PduSessionReleaseComplete           *PduSessionReleaseComplete
	GsmStatus                           *GsmStatus
}

func decodePlainMm(wire []byte) (gmm DecodedGmmMessage, err error) {
	if len(wire) < 3 {
		err = ErrIncomplete
		return
	}
	header := mmHeaderFromBytes(wire[1:3]) //EPD is removed
	switch header.msgType {
	case RegistrationRequestMsgType:
		gmm.MsgType = RegistrationRequestMsgType
		msg := &RegistrationRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationRequest = msg
		}
	case RegistrationAcceptMsgType:
		gmm.MsgType = RegistrationAcceptMsgType
		msg := &RegistrationAccept{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationAccept = msg
		}
	case RegistrationCompleteMsgType:
		gmm.MsgType = RegistrationCompleteMsgType
		msg := &RegistrationComplete{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationComplete = msg
		}
	case RegistrationRejectMsgType:
		gmm.MsgType = RegistrationRejectMsgType
		msg := &RegistrationReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationReject = msg
		}
	case DeregistrationRequestFromUeMsgType:
		gmm.MsgType = DeregistrationRequestFromUeMsgType
		msg := &DeregistrationRequestFromUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationRequestFromUe = msg
		}
	case DeregistrationAcceptFromUeMsgType:
		gmm.MsgType = DeregistrationAcceptFromUeMsgType
		msg := &DeregistrationAcceptFromUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationAcceptFromUe = msg
		}
	case DeregistrationRequestToUeMsgType:
		gmm.MsgType = DeregistrationRequestToUeMsgType
		msg := &DeregistrationRequestToUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationRequestToUe = msg
		}
	case DeregistrationAcceptToUeMsgType:
		gmm.MsgType = DeregistrationAcceptToUeMsgType
		msg := &DeregistrationAcceptToUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationAcceptToUe = msg
		}
	case ServiceRequestMsgType:
		gmm.MsgType = ServiceRequestMsgType
		msg := &ServiceRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ServiceRequest = msg
		}
	case ServiceRejectMsgType:
		gmm.MsgType = ServiceRejectMsgType
		msg := &ServiceReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ServiceReject = msg
		}
	case ServiceAcceptMsgType:
		gmm.MsgType = ServiceAcceptMsgType
		msg := &ServiceAccept{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ServiceAccept = msg
		}
	case ConfigurationUpdateCommandMsgType:
		gmm.MsgType = ConfigurationUpdateCommandMsgType
		msg := &ConfigurationUpdateCommand{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ConfigurationUpdateCommand = msg
		}
	case ConfigurationUpdateCompleteMsgType:
		gmm.MsgType = ConfigurationUpdateCompleteMsgType
		msg := &ConfigurationUpdateComplete{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ConfigurationUpdateComplete = msg
		}
	case AuthenticationRequestMsgType:
		gmm.MsgType = AuthenticationRequestMsgType
		msg := &AuthenticationRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationRequest = msg
		}
	case AuthenticationResponseMsgType:
		gmm.MsgType = AuthenticationResponseMsgType
		msg := &AuthenticationResponse{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationResponse = msg
		}
	case AuthenticationRejectMsgType:
		gmm.MsgType = AuthenticationRejectMsgType
		msg := &AuthenticationReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationReject = msg
		}
	case AuthenticationFailureMsgType:
		gmm.MsgType = AuthenticationFailureMsgType
		msg := &AuthenticationFailure{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationFailure = msg
		}
	case AuthenticationResultMsgType:
		gmm.MsgType = AuthenticationResultMsgType
		msg := &AuthenticationResult{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationResult = msg
		}
	case IdentityRequestMsgType:
		gmm.MsgType = IdentityRequestMsgType
		msg := &IdentityRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.IdentityRequest = msg
		}
	case IdentityResponseMsgType:
		gmm.MsgType = IdentityResponseMsgType
		msg := &IdentityResponse{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.IdentityResponse = msg
		}
	case SecurityModeCommandMsgType:
		gmm.MsgType = SecurityModeCommandMsgType
		msg := &SecurityModeCommand{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.SecurityModeCommand = msg
		}
	case SecurityModeCompleteMsgType:
		gmm.MsgType = SecurityModeCompleteMsgType
		msg := &SecurityModeComplete{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.SecurityModeComplete = msg
		}
	case SecurityModeRejectMsgType:
		gmm.MsgType = SecurityModeRejectMsgType
		msg := &SecurityModeReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.SecurityModeReject = msg
		}
	case GmmStatusMsgType:
		gmm.MsgType = GmmStatusMsgType
		msg := &GmmStatus{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.GmmStatus = msg
		}
	case NotificationMsgType:
		gmm.MsgType = NotificationMsgType
		msg := &Notification{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.Notification = msg
		}
	case NotificationResponseMsgType:
		gmm.MsgType = NotificationResponseMsgType
		msg := &NotificationResponse{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.NotificationResponse = msg
		}
	case UlNasTransportMsgType:
		gmm.MsgType = UlNasTransportMsgType
		msg := &UlNasTransport{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.UlNasTransport = msg
		}
	case DlNasTransportMsgType:
		gmm.MsgType = DlNasTransportMsgType
		msg := &DlNasTransport{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DlNasTransport = msg
		}
	default:
		err = ErrUnknownMsgType
	}
	return
}
func decodeSm(wire []byte) (gsm DecodedGsmMessage, err error) {
	if len(wire) < 4 {
		err = ErrIncomplete
		return
	}
	header := smHeaderFromBytes(wire[1:4]) //EPD is removed
	switch header.msgType {
	case PduSessionEstablishmentRequestMsgType:
		gsm.MsgType = PduSessionEstablishmentRequestMsgType
		msg := &PduSessionEstablishmentRequest{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionEstablishmentRequest = msg
		}
	case PduSessionEstablishmentAcceptMsgType:
		gsm.MsgType = PduSessionEstablishmentAcceptMsgType
		msg := &PduSessionEstablishmentAccept{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionEstablishmentAccept = msg
		}
	case PduSessionEstablishmentRejectMsgType:
		gsm.MsgType = PduSessionEstablishmentRejectMsgType
		msg := &PduSessionEstablishmentReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionEstablishmentReject = msg
		}
	case PduSessionAuthenticationCommandMsgType:
		gsm.MsgType = PduSessionAuthenticationCommandMsgType
		msg := &PduSessionAuthenticationCommand{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionAuthenticationCommand = msg
		}
	case PduSessionAuthenticationCompleteMsgType:
		gsm.MsgType = PduSessionAuthenticationCompleteMsgType
		msg := &PduSessionAuthenticationComplete{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionAuthenticationComplete = msg
		}
	case PduSessionAuthenticationResultMsgType:
		gsm.MsgType = PduSessionAuthenticationResultMsgType
		msg := &PduSessionAuthenticationResult{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionAuthenticationResult = msg
		}
	case PduSessionModificationRequestMsgType:
		gsm.MsgType = PduSessionModificationRequestMsgType
		msg := &PduSessionModificationRequest{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationRequest = msg
		}
	case PduSessionModificationRejectMsgType:
		gsm.MsgType = PduSessionModificationRejectMsgType
		msg := &PduSessionModificationReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationReject = msg
		}
	case PduSessionModificationCommandMsgType:
		gsm.MsgType = PduSessionModificationCommandMsgType
		msg := &PduSessionModificationCommand{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationCommand = msg
		}
	case PduSessionModificationCompleteMsgType:
		gsm.MsgType = PduSessionModificationCompleteMsgType
		msg := &PduSessionModificationComplete{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationComplete = msg
		}
	case PduSessionModificationCommandRejectMsgType:
		gsm.MsgType = PduSessionModificationCommandRejectMsgType
		msg := &PduSessionModificationCommandReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationCommandReject = msg
		}
	case PduSessionReleaseRequestMsgType:
		gsm.MsgType = PduSessionReleaseRequestMsgType
		msg := &PduSessionReleaseRequest{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseRequest = msg
		}
	case PduSessionReleaseRejectMsgType:
		gsm.MsgType = PduSessionReleaseRejectMsgType
		msg := &PduSessionReleaseReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseReject = msg
		}
	case PduSessionReleaseCommandMsgType:
		gsm.MsgType = PduSessionReleaseCommandMsgType
		msg := &PduSessionReleaseCommand{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseCommand = msg
		}
	case PduSessionReleaseCompleteMsgType:
		gsm.MsgType = PduSessionReleaseCompleteMsgType
		msg := &PduSessionReleaseComplete{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseComplete = msg
		}
	case GsmStatusMsgType:
		gsm.MsgType = GsmStatusMsgType
		msg := &GsmStatus{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.GsmStatus = msg
		}
	default:
		err = ErrUnknownMsgType
	}
	return
}
