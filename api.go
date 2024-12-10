package nas

import (
	"bytes"
	//	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "nas"})
}

type Encoder interface {
	encode() ([]byte, error)
}

type Decoder interface {
	decode([]byte) error
}

type GmmMessage interface {
	Encoder
	GetSecurityHeader() uint8
}

type GsmMessage interface {
	Encoder
	SetPti(uint8)
	SetSessionId(uint8)
}

// hold decoded Nas message, either MM or SM
type NasMessage struct {
	Gmm *DecodedGmmMessage
	Gsm *DecodedGsmMessage
}

func EncodeSm(msg GsmMessage) (wire []byte, err error) {
	wire, err = msg.encode()
	return
}

func EncodeMm(ctx *NasContext, msg GmmMessage) (wire []byte, err error) {
	secType := msg.GetSecurityHeader()
	if ctx == nil && secType != NasSecNone {
		err = fmt.Errorf("No security context to encode protected message")
		return
	}

	//encode plain nas
	if wire, err = msg.encode(); err != nil {
		return
	}

	//log.Tracef("Encode message type %d with security header = %d\n", wire[2], secType)

	if ctx == nil {
		//return plain nas
		return
	}
	//set to plain nas
	wire[1] = NasSecNone

	// Security protected NAS Message
	// a security protected NAS message must be integrity protected, and ciphering is optional
	ciphering := false
	switch secType {
	case NasSecIntegrity:
	case NasSecBoth:
		ciphering = true
	case NasSecIntegrityNew:
		//log.Tracef("Message with new security context: %d", secType)
		ctx.localCounter.set(0, 0)
	case NasSecBothNew:
		//log.Tracef("Message with new security context: %d", secType)
		ciphering = true
		ctx.localCounter.set(0, 0)
	default:
		err = fmt.Errorf("Wrong security header type: 0x%0x", secType)
		return
	}

	if ciphering {
		//log.Tracef("Encode message type %d with counter = %d\n", wire[2], ctx.localCounter)
		if wire, err = ctx.encrypt(wire, true); err != nil {
			err = nasError("fail to encryp MM", err)
			return
		}
	}
	secHeader := []byte{EPD_5GMM, secType, 0, 0, 0, 0, ctx.localCounter.sqn()}
	wire = append(secHeader, wire...)

	var mac32 []byte
	if mac32, err = ctx.calculateMac(wire[6:], true); err != nil {
		err = nasError("fail to build MAC for MM", err)
		return
	}
	// Add mac value
	copy(wire[2:], mac32[0:4])

	// Increase local counter
	ctx.localCounter.inc()

	return
}

func (ctx *NasContext) EncryptMmContainer(plain []byte) (cipher []byte, err error) {
	cipher, err = ctx.encrypt(plain, true) //for sending
	return
}

func (ctx *NasContext) DecodeMmContainer(wire []byte) (gmm DecodedGmmMessage, err error) {
	if wire, err = ctx.encrypt(wire, false); err == nil { //for receiving
		err = nasError("fail to decrypt MM container", err)
		gmm, err = decodePlainMm(wire)
	}
	return
}

func Decode(ctx *NasContext, wire []byte) (nasMsg NasMessage, err error) {
	//get message type SM or MM
	wireLen := len(wire)
	if wireLen < 3 {
		err = ErrIncomplete
		return
	}
	switch wire[0] { //epd value; first byte
	case EPD_5GMM: //mobility management message
		var gmm DecodedGmmMessage
		if gmm, err = decodeMm(ctx, wire); err == nil {
			nasMsg.Gmm = &gmm
		} else {
			err = nasError("decode NasMm", err)
		}
	case EPD_5GSM: //session management message
		var gsm DecodedGsmMessage
		if gsm, err = decodeSm(wire); err == nil {
			nasMsg.Gsm = &gsm
		} else {
			err = nasError("decode NasMm", err)
		}
	default:
		err = nasError("decode Nas message", ErrUnknownEpd)
	}

	return
}

// wire length is not less than 3 (checked in previous step)
func decodeMm(ctx *NasContext, wire []byte) (gmm DecodedGmmMessage, err error) {
	//must have a header

	//get security header type
	secType := wire[1] & 0x0f

	switch secType {
	case NasSecNone: //decode plain text
		gmm, err = decodePlainMm(wire)
		gmm.SecHeader = secType
		gmm.MacFailed = false

		if err != nil {
			err = nasError("decode plain nas", err)
			return
		}
		// if having sec-context and not for emergency, check for message type

		//TODO: this logic is for downlink (AMF), we may need to add logic for
		//uplink (UE)
		if ctx != nil && !ctx.emergency {
			// TS 24.501 4.4.4.3: Except the messages listed below, no NAS signalling messages shall be processed
			// by the receiving 5GMM entity in the AMF or forwarded to the 5GSM entity, unless the secure exchange
			// of NAS messages has been established for the NAS signalling connection
			switch gmm.MsgType {
			case RegistrationRequestMsgType:
			case IdentityResponseMsgType:
			case AuthenticationResponseMsgType:
			case AuthenticationFailureMsgType:
			case SecurityModeRejectMsgType:
			case DeregistrationRequestFromUeMsgType:
			case DeregistrationAcceptToUeMsgType:
			default:
				err = fmt.Errorf("UE can not send plain nas for non-emergency service when there is a valid security context")
			}
		}
		return

	case NasSecIntegrity, NasSecIntegrityNew, NasSecBoth, NasSecBothNew:
		if gmm, err = decodeProtectedMm(ctx, wire); err != nil {
			err = nasError("decode protected nas", err)
		}
	default:
		err = ErrUnknownSec
	}
	return
}

func decodeProtectedMm(ctx *NasContext, wire []byte) (gmm DecodedGmmMessage, err error) {
	// message is security protected
	if len(wire) < 6 {
		err = ErrIncomplete
		return
	}

	secHeader := wire[0:6]
	secType := secHeader[1] & 0x0f
	seqNum := wire[6] //sequence number

	rmac32 := secHeader[2:] //receive mac32
	wire = wire[6:]

	newNasContext := false

	// determine if messag is encrypt and whether a new security context is
	// notified
	ciphered := false
	switch secType {
	case NasSecIntegrity:
	case NasSecBoth:
		ciphered = true
	case NasSecBothNew:
		ciphered = true
		newNasContext = true
	case NasSecIntegrityNew:
		newNasContext = true
	default:
		err = fmt.Errorf("Wrong security header type: 0x%0x", secType)
		return
	}

	if ciphered && ctx == nil {
		err = fmt.Errorf("No security context to decrypt")
		return
	}

	//no security context and message is not encrypted
	if ctx == nil {
		//decode plain text (event if it is integrity protected)
		if gmm, err = decodePlainMm(wire[1:]); err != nil { //remote sequence number
			err = nasError("fail to decode plain nas", err)
		}
		gmm.SecHeader = secType
		gmm.MacFailed = false
		return
	}
	//now, there must be a security context to decode
	if newNasContext { //note seqNum must be zero
		ctx.remoteCounter.set(0, 0)
	} else {
		//sync remote counter
		seqDiff := int(ctx.remoteCounter.sqn()) - int(seqNum)
		if seqDiff > int(NAS_COUNT_WINDOW) {
			ctx.remoteCounter.setOverflow(ctx.remoteCounter.overflow() + 1)
		}
		ctx.remoteCounter.setSqn(seqNum)
	}

	var mac32 []byte
	if mac32, err = ctx.calculateMac(wire, false); err != nil {
		err = nasError("fail to calculate message authentication code", err)
		return
	}

	if !bytes.Equal(mac32, rmac32) { //integrity check fails
		if ciphered {
			err = fmt.Errorf("Integrity check not passed:%x-%x", mac32, rmac32)
		} else {
			//no encryption, just decode plaintext
			if gmm, err = decodePlainMm(wire[1:]); err == nil { //remove sequence number
				gmm.SecHeader = secType
				gmm.MacFailed = true
			} else {
				err = nasError("fail to decode plain nas", err)
			}
		}
		return
	}

	//integrity-check passed
	if ciphered {
		// decrypt payload without sequence number (payload[1])
		if wire, err = ctx.encrypt(wire[1:], false); err != nil { //remove sequence number
			err = nasError("fail to decrypt message", err)
			return
		}
	}

	gmm, err = decodePlainMm(wire)
	gmm.SecHeader = secType
	gmm.MacFailed = false
	return
}
