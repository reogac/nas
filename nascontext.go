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
	"github.com/reogac/utils/sec5g"
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

// counters are the NAS COUNTs of one access: a security context used over both
// 3GPP and non-3GPP access has a pair for each (TS 24.501 4.4.3.1). One pair
// for both had messages over one access advance the COUNTs the other expected,
// and the peer drop them as replays or fail their check.
type counters struct {
	local  Counter      //sending NAS counter
	remote replayWindow //receiving NAS counter: the COUNTs received messages verified with
	//remoteLast is the COUNT the most recently verified message was protected
	//with, which a NAS container it carried is ciphered with. It is the highest
	//except when that message arrived late
	remoteLast Counter
}

type NasContext struct {
	pairs     [2]counters //by access: 3GPP, then non-3GPP
	encAlg    uint8       //encryption algorithm
	intAlg    uint8       //integrity protection algorithm
	intKey    [16]uint8   //integrity protection key
	encKey    [16]uint8   //encryption key
	emergency bool
	isAmf     bool
	mutex     sync.Mutex
}

func NewNasContext(isAmf bool) *NasContext {
	ctx := &NasContext{
		isAmf: isAmf,
	}
	return ctx
}

// pair is the COUNTs of the access bearer names.
func (ctx *NasContext) pair(bearer uint8) *counters {
	if bearer == BearerNon3GPP {
		return &ctx.pairs[1]
	}
	return &ctx.pairs[0]
}

// UlCounter is the uplink NAS COUNT over 3GPP access; UlCounterFor names the
// access.
func (ctx *NasContext) UlCounter() uint32 { return ctx.UlCounterFor(true) }

// DlCounter is the downlink NAS COUNT over 3GPP access; DlCounterFor names the
// access.
func (ctx *NasContext) DlCounter() uint32 { return ctx.DlCounterFor(true) }

// UlCounterFor is the uplink NAS COUNT over 3GPP access, or non-3GPP access.
func (ctx *NasContext) UlCounterFor(isGpp bool) uint32 {
	p := ctx.pair(getBearer(isGpp))
	if !ctx.isAmf {
		return uint32(p.local)
	}
	return uint32(p.remote.highest)
}

// DlCounterFor is the downlink NAS COUNT over 3GPP access, or non-3GPP access.
func (ctx *NasContext) DlCounterFor(isGpp bool) uint32 {
	p := ctx.pair(getBearer(isGpp))
	if ctx.isAmf {
		return uint32(p.local)
	}
	return uint32(p.remote.highest)
}

// CountNearLimit reports a NAS COUNT of either access, sent or received, within
// NAS_COUNT_MARGIN of the last one its keys may protect. The AMF then takes new
// keys into use, or the messages run out (TS 24.501 4.4.3.5).
func (ctx *NasContext) CountNearLimit() bool {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	for i := range ctx.pairs {
		p := &ctx.pairs[i]
		if p.local > NAS_COUNT_MAX-NAS_COUNT_MARGIN || p.remote.highest > NAS_COUNT_MAX-NAS_COUNT_MARGIN {
			return true
		}
	}
	return false
}

func (ctx *NasContext) SelectedAlgorithms() (uint8, uint8) {
	return ctx.encAlg, ctx.intAlg
}

// DeriveKeys takes a new K_AMF into use: the NAS keys of the algorithms named
// are derived from it, and the NAS COUNTs of both accesses start again from
// zero (TS 33.501 6.4.3.1).
func (ctx *NasContext) DeriveKeys(encAlg, intAlg uint8, kAmf []byte) (err error) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if err = ctx.deriveAlgorithmKeys(encAlg, intAlg, kAmf); err != nil {
		return
	}
	for i := range ctx.pairs {
		ctx.pairs[i].local.set(0, 0)
		ctx.pairs[i].remote.reset()
		ctx.pairs[i].remoteLast = 0
	}
	return
}

// ChangeAlgorithms changes the NAS algorithms of a context already in use, its
// K_AMF the same: the keys of the algorithms named are derived from kAmf and
// the NAS COUNTs of both accesses go on where they are, which is what a NAS
// Security Mode Command changing only the algorithms of the current 5G NAS
// security context asks of both ends (TS 24.501 5.4.2.1 a)). Only a new K_AMF
// starts them again, through DeriveKeys.
func (ctx *NasContext) ChangeAlgorithms(encAlg, intAlg uint8, kAmf []byte) error {
	return ctx.deriveAlgorithmKeys(encAlg, intAlg, kAmf)
}

// deriveAlgorithmKeys derives the NAS encryption and integrity keys of the
// algorithms named from kAmf (TS 33.501 A.8) and takes them into use; the
// caller holds the mutex.
func (ctx *NasContext) deriveAlgorithmKeys(encAlg, intAlg uint8, kAmf []byte) (err error) {
	ctx.encAlg = encAlg
	ctx.intAlg = intAlg
	// Encryption Key
	P0 := []byte{NNASEncAlg}
	P1 := []byte{encAlg}

	var kEnc, kInt []byte
	if kEnc, err = sec5g.AlgKey(kAmf, P0, P1); err != nil {
		return
	}

	// Integrity Key
	P0 = []byte{NNASIntAlg}
	P1 = []byte{intAlg}

	if kInt, err = sec5g.AlgKey(kAmf, P0, P1); err != nil {
		return
	}
	copy(ctx.encKey[:], kEnc[16:32])
	copy(ctx.intKey[:], kInt[16:32])
	return
}

func (ctx *NasContext) getDirection(isSending bool, bearer uint8) (direction uint8, counter uint32) {
	p := ctx.pair(bearer)
	if isSending { //for sending message
		if ctx.isAmf {
			direction = DirectionDownlink
		} else {
			direction = DirectionUplink
		}
		counter = uint32(p.local)
	} else { //for receiving message
		if ctx.isAmf {
			direction = DirectionUplink
		} else {
			direction = DirectionDownlink
		}
		counter = uint32(p.remoteLast)
	}
	return
}

func (ctx *NasContext) encrypt(payload []byte, isSending bool, bearer uint8) (output []byte, err error) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	direction, counter := ctx.getDirection(isSending, bearer)
	return ctx.cipher(payload, direction, counter, bearer)
}

// cipher runs the ciphering algorithm with an explicit COUNT; the caller holds
// the mutex
func (ctx *NasContext) cipher(payload []byte, direction uint8, counter uint32, bearer uint8) (output []byte, err error) {
	switch ctx.encAlg {
	case AlgCiphering128NEA0:
		//log.Debugf("Use NEA0")
		output = payload
	case AlgCiphering128NEA1:
		//log.Debugln("Use NEA1")
		output, err = NEA1(ctx.encKey, counter, uint32(bearer), uint32(direction), payload, uint32(len(payload))*8)
	case AlgCiphering128NEA2:
		//log.Debugln("Use NEA2")
		output, err = NEA2(ctx.encKey, counter, bearer, direction, payload)
	case AlgCiphering128NEA3:
		//log.Debugln("Use NEA3")
		output, err = NEA3(ctx.encKey, counter, bearer, direction, payload, uint32(len(payload))*8)
	default:
		err = fmt.Errorf("Unknown Algorithm Identity[%d]", ctx.encAlg)
	}
	if err != nil {
		return
	}
	return
}

func (ctx *NasContext) calculateMac(payload []byte, isSending bool, bearer uint8) (mac []byte, err error) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	direction, counter := ctx.getDirection(isSending, bearer)
	return ctx.mac(payload, direction, counter, bearer)
}

// mac computes the message authentication code with an explicit COUNT; the
// caller holds the mutex
func (ctx *NasContext) mac(payload []byte, direction uint8, counter uint32, bearer uint8) (mac []byte, err error) {
	switch ctx.intAlg {
	case AlgIntegrity128NIA0:
		//log.Warningln("Integrity NIA0 is emergency.")
		mac = make([]byte, 4)
	case AlgIntegrity128NIA1:
		//log.Debugf("Use NIA1")
		mac, err = NIA1(ctx.intKey, counter, bearer, uint32(direction), payload, uint64(len(payload))*8)
	case AlgIntegrity128NIA2:
		//log.Debugf("Use NIA2")
		mac, err = NIA2(ctx.intKey, counter, bearer, direction, payload)
	case AlgIntegrity128NIA3:
		//log.Debugf("Use NIA3")
		mac, err = NIA3(ctx.intKey, counter, bearer, direction, payload, uint32(len(payload))*8)
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
		// TS 24.501 4.4.4.2: the messages the receiving 5GMM entity in the UE
		// processes without integrity protection. The clause's conditions -
		// an IDENTITY REQUEST asking for the SUCI, a DEREGISTRATION ACCEPT for
		// a de-registration that was not a switch-off, a reject whose cause is
		// not #76 or #78 - are in the message, not its type, and are left to
		// the caller.
		switch msgType {
		case IdentityRequestMsgType:
		case AuthenticationRequestMsgType:
		case AuthenticationResultMsgType:
		case AuthenticationRejectMsgType:
		case RegistrationRejectMsgType:
		case DeregistrationAcceptFromUeMsgType:
		case ServiceRejectMsgType:
		default:
			return false
		}
		return true
	}
}
