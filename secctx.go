package nas

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/reogac/nas/sec"
)

const (
	HDP_NONE uint8 = iota
	HDP_HANDOVER
	HDP_MOBILITY_UPDATE
)

type SecurityContext struct {
	useSecCap UeSecurityCapability
	nasCtx    NasContext
	ngKsi     KeySetIdentifier
	kamf      []byte
	kgnb      []uint8 //gnb key
	kn3iwf    []uint8 //n3iwf key
	nh        []uint8 //next hop parameter (for AS security context)
	ncc       uint8   // next chain counter (for AS security context)
	active    bool
}

func NewSecurityContext(ksi *KeySetIdentifier, kamf []byte, isDownlink bool, bearer uint8) *SecurityContext {
	ctx := &SecurityContext{
		ngKsi:  *ksi,
		active: false,
	}
	ctx.kamf = make([]byte, len(kamf))
	copy(ctx.kamf, kamf)
	ctx.nasCtx.isDownlink = isDownlink
	ctx.nasCtx.bearer = bearer
	return ctx
}

func (ctx *SecurityContext) NgKsi() *KeySetIdentifier {
	return &ctx.ngKsi
}

func (ctx *SecurityContext) NasContext() *NasContext {
	if ctx.active {
		return &ctx.nasCtx
	}
	return nil
}

func (ctx *SecurityContext) SelectAlgorithms(intOrder []byte, encOrder []byte, ueSecCap *UeSecurityCapability) {
	ctx.nasCtx.selectAlgorithms(intOrder, encOrder, ueSecCap)
}

func (ctx *SecurityContext) MatchNgKsi(ngksi *KeySetIdentifier) bool {
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
func (ctx *SecurityContext) createAnKey(access uint8) (err error) {

	P0 := make([]byte, 4)
	binary.BigEndian.PutUint32(P0, uint32(ctx.nasCtx.localCounter))
	P1 := []byte{access}

	switch access {
	case AccessType3GPP:
		ctx.kgnb, err = sec.RanKey(ctx.kamf, P0, P1)
	case AccessTypeNon3GPP:
		ctx.kn3iwf, err = sec.RanKey(ctx.kamf, P0, P1)
	default:
		err = fmt.Errorf("Unsupported access type %d", access)
	}
	return
}

// NH Derivation function defined in TS 33.501 Annex A.10
func (ctx *SecurityContext) createNh(syncinput []byte) (err error) {
	ctx.nh, err = sec.NhKey(ctx.kamf, syncinput)
	return
}

func (ctx *SecurityContext) Update(access uint8) (err error) {
	if err = ctx.createAnKey(access); err != nil {
		return
	}

	switch access {
	case AccessType3GPP:
		ctx.createNh(ctx.kgnb)
	case AccessTypeNon3GPP:
		ctx.createNh(ctx.kn3iwf)
	default:
		err = fmt.Errorf("Unsupported access type %d", access)
		return
	}
	ctx.ncc = 1
	return
}

func (ctx *SecurityContext) UpdateNh() error {
	ctx.ncc++
	return ctx.createNh(ctx.nh)
}

// Algorithm key Derivation function defined in TS 33.501 Annex A.9
func (ctx *SecurityContext) DeriveAlgKeys(hdp uint8) (err error) {
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
		binary.BigEndian.PutUint32(p1[:], uint32(ctx.nasCtx.localCounter))
		kamf, err = sec.KamfPrime(ctx.kamf, p0, p1[:])
	case HDP_MOBILITY_UPDATE:
		//derive kamf prime
		p0, _ = hex.DecodeString("00")
		binary.BigEndian.PutUint32(p1[:], uint32(ctx.nasCtx.remoteCounter))
		kamf, err = sec.KamfPrime(ctx.kamf, p0, p1[:])
	default:
		kamf = ctx.kamf
	}
	if err != nil {
		return
	}

	//then derive key for nas context
	log.Tracef("Security context's KAMF = %x", ctx.kamf)
	if err = ctx.nasCtx.DeriveKeys(kamf); err == nil {
		ctx.kamf = kamf
		ctx.active = true //activate the security context
	}
	return
}
