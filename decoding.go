/*
* Copyright [2024] [Quang Tung Thai <tqtung@etri.re.kr>]
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
/** this file was generated at 2024-12-16 17:55:27.322054 by tqtung@etri.re.kr **/

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
	gmm.MsgType = header.msgType
	switch header.msgType {
	case RegistrationRequestMsgType:
		msg := &RegistrationRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationRequest = msg
		} else {
			err = nasError("decode RegistrationRequest message body", err)
		}
	case RegistrationAcceptMsgType:
		msg := &RegistrationAccept{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationAccept = msg
		} else {
			err = nasError("decode RegistrationAccept message body", err)
		}
	case RegistrationCompleteMsgType:
		msg := &RegistrationComplete{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationComplete = msg
		} else {
			err = nasError("decode RegistrationComplete message body", err)
		}
	case RegistrationRejectMsgType:
		msg := &RegistrationReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.RegistrationReject = msg
		} else {
			err = nasError("decode RegistrationReject message body", err)
		}
	case DeregistrationRequestFromUeMsgType:
		msg := &DeregistrationRequestFromUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationRequestFromUe = msg
		} else {
			err = nasError("decode DeregistrationRequestFromUe message body", err)
		}
	case DeregistrationAcceptFromUeMsgType:
		msg := &DeregistrationAcceptFromUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationAcceptFromUe = msg
		} else {
			err = nasError("decode DeregistrationAcceptFromUe message body", err)
		}
	case DeregistrationRequestToUeMsgType:
		msg := &DeregistrationRequestToUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationRequestToUe = msg
		} else {
			err = nasError("decode DeregistrationRequestToUe message body", err)
		}
	case DeregistrationAcceptToUeMsgType:
		msg := &DeregistrationAcceptToUe{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DeregistrationAcceptToUe = msg
		} else {
			err = nasError("decode DeregistrationAcceptToUe message body", err)
		}
	case ServiceRequestMsgType:
		msg := &ServiceRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ServiceRequest = msg
		} else {
			err = nasError("decode ServiceRequest message body", err)
		}
	case ServiceRejectMsgType:
		msg := &ServiceReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ServiceReject = msg
		} else {
			err = nasError("decode ServiceReject message body", err)
		}
	case ServiceAcceptMsgType:
		msg := &ServiceAccept{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ServiceAccept = msg
		} else {
			err = nasError("decode ServiceAccept message body", err)
		}
	case ConfigurationUpdateCommandMsgType:
		msg := &ConfigurationUpdateCommand{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ConfigurationUpdateCommand = msg
		} else {
			err = nasError("decode ConfigurationUpdateCommand message body", err)
		}
	case ConfigurationUpdateCompleteMsgType:
		msg := &ConfigurationUpdateComplete{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.ConfigurationUpdateComplete = msg
		} else {
			err = nasError("decode ConfigurationUpdateComplete message body", err)
		}
	case AuthenticationRequestMsgType:
		msg := &AuthenticationRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationRequest = msg
		} else {
			err = nasError("decode AuthenticationRequest message body", err)
		}
	case AuthenticationResponseMsgType:
		msg := &AuthenticationResponse{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationResponse = msg
		} else {
			err = nasError("decode AuthenticationResponse message body", err)
		}
	case AuthenticationRejectMsgType:
		msg := &AuthenticationReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationReject = msg
		} else {
			err = nasError("decode AuthenticationReject message body", err)
		}
	case AuthenticationFailureMsgType:
		msg := &AuthenticationFailure{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationFailure = msg
		} else {
			err = nasError("decode AuthenticationFailure message body", err)
		}
	case AuthenticationResultMsgType:
		msg := &AuthenticationResult{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.AuthenticationResult = msg
		} else {
			err = nasError("decode AuthenticationResult message body", err)
		}
	case IdentityRequestMsgType:
		msg := &IdentityRequest{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.IdentityRequest = msg
		} else {
			err = nasError("decode IdentityRequest message body", err)
		}
	case IdentityResponseMsgType:
		msg := &IdentityResponse{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.IdentityResponse = msg
		} else {
			err = nasError("decode IdentityResponse message body", err)
		}
	case SecurityModeCommandMsgType:
		msg := &SecurityModeCommand{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.SecurityModeCommand = msg
		} else {
			err = nasError("decode SecurityModeCommand message body", err)
		}
	case SecurityModeCompleteMsgType:
		msg := &SecurityModeComplete{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.SecurityModeComplete = msg
		} else {
			err = nasError("decode SecurityModeComplete message body", err)
		}
	case SecurityModeRejectMsgType:
		msg := &SecurityModeReject{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.SecurityModeReject = msg
		} else {
			err = nasError("decode SecurityModeReject message body", err)
		}
	case GmmStatusMsgType:
		msg := &GmmStatus{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.GmmStatus = msg
		} else {
			err = nasError("decode GmmStatus message body", err)
		}
	case NotificationMsgType:
		msg := &Notification{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.Notification = msg
		} else {
			err = nasError("decode Notification message body", err)
		}
	case NotificationResponseMsgType:
		msg := &NotificationResponse{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.NotificationResponse = msg
		} else {
			err = nasError("decode NotificationResponse message body", err)
		}
	case UlNasTransportMsgType:
		msg := &UlNasTransport{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.UlNasTransport = msg
		} else {
			err = nasError("decode UlNasTransport message body", err)
		}
	case DlNasTransportMsgType:
		msg := &DlNasTransport{
			MmHeader: *header,
		}
		if err = msg.decodeBody(wire[3:]); err == nil {
			gmm.DlNasTransport = msg
		} else {
			err = nasError("decode DlNasTransport message body", err)
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
	gsm.MsgType = header.msgType
	switch header.msgType {
	case PduSessionEstablishmentRequestMsgType:
		msg := &PduSessionEstablishmentRequest{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionEstablishmentRequest = msg
		} else {
			err = nasError("decode PduSessionEstablishmentRequest message body", err)
		}

	case PduSessionEstablishmentAcceptMsgType:
		msg := &PduSessionEstablishmentAccept{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionEstablishmentAccept = msg
		} else {
			err = nasError("decode PduSessionEstablishmentAccept message body", err)
		}

	case PduSessionEstablishmentRejectMsgType:
		msg := &PduSessionEstablishmentReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionEstablishmentReject = msg
		} else {
			err = nasError("decode PduSessionEstablishmentReject message body", err)
		}

	case PduSessionAuthenticationCommandMsgType:
		msg := &PduSessionAuthenticationCommand{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionAuthenticationCommand = msg
		} else {
			err = nasError("decode PduSessionAuthenticationCommand message body", err)
		}

	case PduSessionAuthenticationCompleteMsgType:
		msg := &PduSessionAuthenticationComplete{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionAuthenticationComplete = msg
		} else {
			err = nasError("decode PduSessionAuthenticationComplete message body", err)
		}

	case PduSessionAuthenticationResultMsgType:
		msg := &PduSessionAuthenticationResult{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionAuthenticationResult = msg
		} else {
			err = nasError("decode PduSessionAuthenticationResult message body", err)
		}

	case PduSessionModificationRequestMsgType:
		msg := &PduSessionModificationRequest{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationRequest = msg
		} else {
			err = nasError("decode PduSessionModificationRequest message body", err)
		}

	case PduSessionModificationRejectMsgType:
		msg := &PduSessionModificationReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationReject = msg
		} else {
			err = nasError("decode PduSessionModificationReject message body", err)
		}

	case PduSessionModificationCommandMsgType:
		msg := &PduSessionModificationCommand{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationCommand = msg
		} else {
			err = nasError("decode PduSessionModificationCommand message body", err)
		}

	case PduSessionModificationCompleteMsgType:
		msg := &PduSessionModificationComplete{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationComplete = msg
		} else {
			err = nasError("decode PduSessionModificationComplete message body", err)
		}

	case PduSessionModificationCommandRejectMsgType:
		msg := &PduSessionModificationCommandReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionModificationCommandReject = msg
		} else {
			err = nasError("decode PduSessionModificationCommandReject message body", err)
		}

	case PduSessionReleaseRequestMsgType:
		msg := &PduSessionReleaseRequest{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseRequest = msg
		} else {
			err = nasError("decode PduSessionReleaseRequest message body", err)
		}

	case PduSessionReleaseRejectMsgType:
		msg := &PduSessionReleaseReject{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseReject = msg
		} else {
			err = nasError("decode PduSessionReleaseReject message body", err)
		}

	case PduSessionReleaseCommandMsgType:
		msg := &PduSessionReleaseCommand{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseCommand = msg
		} else {
			err = nasError("decode PduSessionReleaseCommand message body", err)
		}

	case PduSessionReleaseCompleteMsgType:
		msg := &PduSessionReleaseComplete{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.PduSessionReleaseComplete = msg
		} else {
			err = nasError("decode PduSessionReleaseComplete message body", err)
		}

	case GsmStatusMsgType:
		msg := &GsmStatus{
			SmHeader: *header,
		}
		if err = msg.decodeBody(wire[4:]); err == nil {
			gsm.GsmStatus = msg
		} else {
			err = nasError("decode GsmStatus message body", err)
		}

	default:
		err = ErrUnknownMsgType
	}
	return
}
