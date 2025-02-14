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
package secctx

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/reogac/nas"
	"github.com/reogac/utils/sec5g"
)

const (
	HDP_NONE uint8 = iota
	HDP_HANDOVER
	HDP_MOBILITY_UPDATE
)

type SecurityContext struct {
	isAmf bool

	//NAS contexts
	gppNas    *nas.NasContext
	nonGppNas *nas.NasContext

	ngKsi nas.KeySetIdentifier
	kamf  []byte

	//AS context
	kgnb   []uint8 //gnb key
	kn3iwf []uint8 //n3iwf key
	nh     []uint8 //next hop parameter (for AS security context)
	ncc    uint8   // next chain counter (for AS security context)

}

func NewSecurityContext(ksi *nas.KeySetIdentifier, kamf []byte, isAmf bool) *SecurityContext {
	ctx := &SecurityContext{
		ngKsi:     *ksi,
		isAmf:     isAmf,
		kamf:      make([]byte, len(kamf)),
		gppNas:    nas.NewNasContext(isAmf, nas.Bearer3GPP),
		nonGppNas: nas.NewNasContext(isAmf, nas.BearerNon3GPP),
	}
	copy(ctx.kamf, kamf)
	return ctx
}

func (ctx *SecurityContext) NgKsi() *nas.KeySetIdentifier {
	return &ctx.ngKsi
}

func (ctx *SecurityContext) NasContext(isGpp bool) *nas.NasContext {
	if isGpp {
		return ctx.gppNas
	}
	return ctx.nonGppNas
}

func (ctx *SecurityContext) MatchNgKsi(ngksi *nas.KeySetIdentifier) bool {
	return ctx.ngKsi.Tsc == ngksi.Tsc && ctx.ngKsi.Id == ngksi.Id
}

func (ctx *SecurityContext) Kamf() []byte {
	return ctx.kamf
}

func (ctx *SecurityContext) Kgnb() []byte {
	return ctx.kgnb
}

func (ctx *SecurityContext) Kn3iwf() []byte {
	return ctx.kn3iwf
}

// Access Network key Derivation function defined in TS 33.501 Annex A.9
func (ctx *SecurityContext) createAnKey() (err error) {

	P0 := make([]byte, 4)
	P1 := []byte{nas.AccessType3GPP}

	binary.BigEndian.PutUint32(P0, uint32(ctx.gppNas.UlCounter()))
	if ctx.kgnb, err = sec5g.RanKey(ctx.kamf, P0, P1); err != nil {
		return
	}
	P1[0] = nas.AccessTypeNon3GPP
	binary.BigEndian.PutUint32(P0, uint32(ctx.nonGppNas.UlCounter()))
	ctx.kn3iwf, err = sec5g.RanKey(ctx.kamf, P0, P1)
	return
}

// NH Derivation function defined in TS 33.501 Annex A.10
func (ctx *SecurityContext) createNh(syncinput []byte) (err error) {
	ctx.nh, err = sec5g.NhKey(ctx.kamf, syncinput)
	return
}

func (ctx *SecurityContext) DeriveAsKeys() (err error) {
	if err = ctx.createAnKey(); err != nil {
		return
	}

	err = ctx.createNh(ctx.kgnb)
	ctx.ncc = 1
	return
}

func (ctx *SecurityContext) NextNh() (ncc uint8, nh []uint8, err error) {
	ncc = (ctx.ncc + 1) & 0x07
	nh, err = sec5g.NhKey(ctx.kamf, ctx.nh)
	return
}

func (ctx *SecurityContext) UpdateNh() error {
	ctx.ncc++
	ctx.ncc &= 0x07
	return ctx.createNh(ctx.nh)
}

// Algorithm key Derivation function defined in TS 33.501 Annex A.9
func (ctx *SecurityContext) DeriveNasKeys(encAlg, intAlg, hdp uint8) (err error) {
	//first, derive KAMF if neccesary
	var p0 []byte
	var p1 [4]byte
	var kamf []byte
	switch hdp {
	case HDP_HANDOVER:
		//derive kamf prime
		p0, _ = hex.DecodeString("01")
		//NOTE: p1 = dlcount is only applied for 3gpp access (not sure about
		//non-3gpp (see TS 133.501 A.13)
		binary.BigEndian.PutUint32(p1[:], uint32(ctx.gppNas.DlCounter()))
		kamf, err = sec5g.KamfPrime(ctx.kamf, p0, p1[:])
	case HDP_MOBILITY_UPDATE:
		//derive kamf prime
		p0, _ = hex.DecodeString("00")
		binary.BigEndian.PutUint32(p1[:], uint32(ctx.gppNas.UlCounter()))
		kamf, err = sec5g.KamfPrime(ctx.kamf, p0, p1[:])
	default:
		kamf = ctx.kamf
	}
	if err != nil {
		return
	}
	//then derive key for nas contexts
	if err = ctx.gppNas.DeriveKeys(encAlg, intAlg, kamf); err != nil {
		return
	}

	if err = ctx.nonGppNas.DeriveKeys(encAlg, intAlg, kamf); err != nil {
		return
	}
	ctx.kamf = kamf
	return
}
