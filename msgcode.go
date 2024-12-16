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
/** this file was generated at 2024-12-16 17:55:27.321873 by tqtung@etri.re.kr **/

package nas

const (
	RegistrationRequestMsgType                 uint8 = 65
	RegistrationAcceptMsgType                  uint8 = 66
	RegistrationCompleteMsgType                uint8 = 67
	RegistrationRejectMsgType                  uint8 = 68
	DeregistrationRequestFromUeMsgType         uint8 = 69
	DeregistrationAcceptFromUeMsgType          uint8 = 70
	DeregistrationRequestToUeMsgType           uint8 = 71
	DeregistrationAcceptToUeMsgType            uint8 = 72
	ServiceRequestMsgType                      uint8 = 76
	ServiceRejectMsgType                       uint8 = 77
	ServiceAcceptMsgType                       uint8 = 78
	ConfigurationUpdateCommandMsgType          uint8 = 84
	ConfigurationUpdateCompleteMsgType         uint8 = 85
	AuthenticationRequestMsgType               uint8 = 86
	AuthenticationResponseMsgType              uint8 = 87
	AuthenticationRejectMsgType                uint8 = 88
	AuthenticationFailureMsgType               uint8 = 89
	AuthenticationResultMsgType                uint8 = 90
	IdentityRequestMsgType                     uint8 = 91
	IdentityResponseMsgType                    uint8 = 92
	SecurityModeCommandMsgType                 uint8 = 93
	SecurityModeCompleteMsgType                uint8 = 94
	SecurityModeRejectMsgType                  uint8 = 95
	GmmStatusMsgType                           uint8 = 100
	NotificationMsgType                        uint8 = 101
	NotificationResponseMsgType                uint8 = 102
	UlNasTransportMsgType                      uint8 = 103
	DlNasTransportMsgType                      uint8 = 104
	PduSessionEstablishmentRequestMsgType      uint8 = 193
	PduSessionEstablishmentAcceptMsgType       uint8 = 194
	PduSessionEstablishmentRejectMsgType       uint8 = 195
	PduSessionAuthenticationCommandMsgType     uint8 = 197
	PduSessionAuthenticationCompleteMsgType    uint8 = 198
	PduSessionAuthenticationResultMsgType      uint8 = 199
	PduSessionModificationRequestMsgType       uint8 = 201
	PduSessionModificationRejectMsgType        uint8 = 202
	PduSessionModificationCommandMsgType       uint8 = 203
	PduSessionModificationCompleteMsgType      uint8 = 204
	PduSessionModificationCommandRejectMsgType uint8 = 205
	PduSessionReleaseRequestMsgType            uint8 = 209
	PduSessionReleaseRejectMsgType             uint8 = 210
	PduSessionReleaseCommandMsgType            uint8 = 211
	PduSessionReleaseCompleteMsgType           uint8 = 212
	GsmStatusMsgType                           uint8 = 214
)
