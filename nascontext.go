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
package nas

import (
	"fmt"
	"github.com/reogac/nas/sec"
	"sync"
)

// TS 33.501 Annex A.8 Algorithm distinguisher For Knas_int Knas_enc
const (
	NNASEncAlg uint8 = 0x01
	NNASIntAlg uint8 = 0x02
	NRRCEncAlg uint8 = 0x03
	NRRCIntAlg uint8 = 0x04
	NUpEncAlg  uint8 = 0x05
	NUpIntAlg  uint8 = 0x06
)

// TS 33.501 5.11.1.1 Algorithm identifier values For integrity algorithm
const (
	AlgIntegrity128NIA0 uint8 = 0x00 // NULL
	AlgIntegrity128NIA1 uint8 = 0x01 // 128-Snow3G
	AlgIntegrity128NIA2 uint8 = 0x02 // 128-AES
	AlgIntegrity128NIA3 uint8 = 0x03 // 128-ZUC
)

// TS 33.501 5.11.1.1 Algorithm identifier values For ciphering algorithm
const (
	AlgCiphering128NEA0 uint8 = 0x00 // NULL
	AlgCiphering128NEA1 uint8 = 0x01 // 128-Snow3G
	AlgCiphering128NEA2 uint8 = 0x02 // 128-AES
	AlgCiphering128NEA3 uint8 = 0x03 // 128-ZUC
)

// 1bit
const (
	DirectionUplink   uint8 = 0x00
	DirectionDownlink uint8 = 0x01
)

// 5bits
const (
	OnlyOneBearer uint8 = 0x00
	Bearer3GPP    uint8 = 0x01
	BearerNon3GPP uint8 = 0x02
)

type NasContext struct {
	localCounter  Counter   //sending NAS counter
	remoteCounter Counter   //receiving NAS counter
	encAlg        uint8     //encryption algorithm
	intAlg        uint8     //integrity protection algorithm
	intKey        [16]uint8 //integrity protection key
	encKey        [16]uint8 //encryption key
	bearer        uint8
	emergency     bool
	isAmf         bool
	mutex         sync.Mutex
}

func NewEmergencyNasContext(isAmf bool, bearer uint8) *NasContext {
	ctx := &NasContext{
		emergency: true,
		bearer:    bearer,
	}
	ctx.setDirection(isAmf)
	return ctx
}

func (ctx *NasContext) setDirection(isAmf bool) {
	ctx.isAmf = isAmf
}

func (ctx *NasContext) selectAlgorithms(intOrder []byte, encOrder []byte, ueSecCap *UeSecurityCapability) {
	//for oai ue
	ctx.encAlg = AlgCiphering128NEA0
	ctx.intAlg = AlgIntegrity128NIA2
	supported := false
	for _, alg := range intOrder {
		switch alg {
		case AlgIntegrity128NIA0:
			supported = ueSecCap.GetIA(0)
		case AlgIntegrity128NIA1:
			supported = ueSecCap.GetIA(1)
		case AlgIntegrity128NIA2:
			supported = ueSecCap.GetIA(2)
		case AlgIntegrity128NIA3:
			supported = ueSecCap.GetIA(3)
		}
		if supported {
			ctx.intAlg = alg
			break
		}
	}

	supported = false
	for _, alg := range encOrder {
		switch alg {
		case AlgCiphering128NEA0:
			supported = ueSecCap.GetEA(0)
		case AlgCiphering128NEA1:
			supported = ueSecCap.GetEA(1)
		case AlgCiphering128NEA2:
			supported = ueSecCap.GetEA(2)
		case AlgCiphering128NEA3:
			supported = ueSecCap.GetEA(3)
		}
		if supported {
			ctx.encAlg = alg
			break
		}
	}

}

func (ctx *NasContext) setAlgorithms(encAlg uint8, intAlg uint8) {
	ctx.encAlg = encAlg
	ctx.intAlg = intAlg
}

func (ctx *NasContext) SelectedAlgorithms() (uint8, uint8) {
	return ctx.encAlg, ctx.intAlg
}

func (ctx *NasContext) deriveKeys(kAmf []byte) (err error) {
	// Encryption Key
	P0 := []byte{NNASEncAlg}
	P1 := []byte{ctx.encAlg}

	var kEnc, kInt []byte
	if kEnc, err = sec.AlgKey(kAmf, P0, P1); err != nil {
		return
	}

	// Integrity Key
	P0 = []byte{NNASIntAlg}
	P1 = []byte{ctx.intAlg}

	if kInt, err = sec.AlgKey(kAmf, P0, P1); err != nil {
		return
	}
	copy(ctx.encKey[:], kEnc[16:32])
	copy(ctx.intKey[:], kInt[16:32])
	ctx.localCounter.set(0, 0)
	ctx.remoteCounter.set(0, 0)
	return
}

func (ctx *NasContext) getDirection(isSending bool) (direction uint8, counter uint32) {
	if isSending { //for sending message
		if ctx.isAmf {
			direction = DirectionDownlink
		} else {
			direction = DirectionUplink
		}
		counter = uint32(ctx.localCounter)
	} else { //for receiving message
		if ctx.isAmf {
			direction = DirectionUplink
		} else {
			direction = DirectionDownlink
		}
		counter = uint32(ctx.remoteCounter)
	}
	return
}

func (ctx *NasContext) encrypt(payload []byte, isSending bool) (output []byte, err error) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	direction, counter := ctx.getDirection(isSending)
	switch ctx.encAlg {
	case AlgCiphering128NEA0:
		log.Debugf("Use NEA0")
		output = payload
	case AlgCiphering128NEA1:
		log.Debugln("Use NEA1")
		output, err = NEA1(ctx.encKey, counter, uint32(ctx.bearer), uint32(direction), payload, uint32(len(payload))*8)
	case AlgCiphering128NEA2:
		log.Debugln("Use NEA2")
		output, err = NEA2(ctx.encKey, counter, ctx.bearer, direction, payload)
	case AlgCiphering128NEA3:
		log.Debugln("Use NEA3")
		output, err = NEA3(ctx.encKey, counter, ctx.bearer, direction, payload, uint32(len(payload))*8)
	default:
		err = fmt.Errorf("Unknown Algorithm Identity[%d]", ctx.encAlg)
	}
	if err != nil {
		return
	}
	return
}

func (ctx *NasContext) calculateMac(payload []byte, isSending bool) (mac []byte, err error) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	direction, counter := ctx.getDirection(isSending)
	switch ctx.intAlg {
	case AlgIntegrity128NIA0:
		//log.Warningln("Integrity NIA0 is emergency.")
		mac = make([]byte, 4)
	case AlgIntegrity128NIA1:
		log.Debugf("Use NIA1")
		mac, err = NIA1(ctx.intKey, counter, ctx.bearer, uint32(direction), payload, uint64(len(payload))*8)
	case AlgIntegrity128NIA2:
		log.Debugf("Use NIA2")
		mac, err = NIA2(ctx.intKey, counter, ctx.bearer, direction, payload)
	case AlgIntegrity128NIA3:
		log.Debugf("Use NIA3")
		mac, err = NIA3(ctx.intKey, counter, ctx.bearer, direction, payload, uint32(len(payload))*8)
	default:
		err = fmt.Errorf("Unknown Algorithm Identity[%d]", ctx.intAlg)
	}

	return
}

func acceptPlaintextN1Mm(msgType uint8, isAmf bool) bool {
	if isAmf {
		// TS 24.501 4.4.4.3: Except the messages listed below, no NAS signalling messages shall be processed
		// by the receiving 5GMM entity in the AMF or forwarded to the 5GSM entity, unless the secure exchange
		// of NAS messages has been established for the NAS signalling connection
		switch msgType {
		case RegistrationRequestMsgType:
		case IdentityResponseMsgType:
		case AuthenticationResponseMsgType:
		case AuthenticationFailureMsgType:
		case SecurityModeRejectMsgType:
		case DeregistrationRequestFromUeMsgType:
		case DeregistrationAcceptToUeMsgType:
		default:
			return false
		}
		return true
	} else {
		switch msgType {
		case IdentityRequestMsgType:
		case AuthenticationRequestMsgType:
		case SecurityModeCommandMsgType:
		case ServiceRejectMsgType:
		case RegistrationRejectMsgType:
		case AuthenticationRejectMsgType:
		case GmmStatusMsgType:
		default:
			return false
		}
		return true

	}
}
