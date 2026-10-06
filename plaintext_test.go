package nas

import "testing"

// A UE holding a security context processes unprotected only the messages TS
// 24.501 4.4.4.2 lists. The list used to lack AUTHENTICATION RESULT and
// DEREGISTRATION ACCEPT, so a UE dropped the EAP-Success or EAP-Failure an
// AMF sent in plain text, and admitted a SECURITY MODE COMMAND, which is never
// sent unprotected, and a 5GMM STATUS, which the clause does not list.
func TestUeAcceptsPlainOnlyWhatTheClauseLists(t *testing.T) {
	accepted := map[uint8]bool{
		IdentityRequestMsgType:            true,
		AuthenticationRequestMsgType:      true,
		AuthenticationResultMsgType:       true,
		AuthenticationRejectMsgType:       true,
		RegistrationRejectMsgType:         true,
		DeregistrationAcceptFromUeMsgType: true,
		ServiceRejectMsgType:              true,
	}
	for _, msgType := range []uint8{
		IdentityRequestMsgType, AuthenticationRequestMsgType, AuthenticationResultMsgType,
		AuthenticationRejectMsgType, RegistrationRejectMsgType, DeregistrationAcceptFromUeMsgType,
		ServiceRejectMsgType, SecurityModeCommandMsgType, GmmStatusMsgType, RegistrationAcceptMsgType,
		ConfigurationUpdateCommandMsgType, DlNasTransportMsgType, DeregistrationRequestToUeMsgType,
		ServiceAcceptMsgType,
	} {
		if got := acceptPlaintextN1Mm(msgType, false); got != accepted[msgType] {
			t.Errorf("UE accepts message type 0x%02x plain = %v, want %v", msgType, got, accepted[msgType])
		}
	}
}

// End to end: a plain AUTHENTICATION RESULT reaches a UE that derived a key
// from the EAP challenge before it.
func TestUeDecodesAPlainAuthenticationResultWithAContext(t *testing.T) {
	msg := &AuthenticationResult{EapMessage: []byte{3, 1, 0, 4}} // EAP-Success
	msg.SetSecurityHeader(NasSecNone)
	wire, err := EncodeMm(nil, msg, true)
	if err != nil {
		t.Fatal(err)
	}
	ue := NewNasContext(false)
	if err := ue.DeriveKeys(AlgCiphering128NEA2, AlgIntegrity128NIA2, make([]byte, 32)); err != nil {
		t.Fatal(err)
	}
	if _, err := Decode(ue, wire, true); err != nil {
		t.Errorf("a UE with a security context drops a plain Authentication Result: %v", err)
	}
}
