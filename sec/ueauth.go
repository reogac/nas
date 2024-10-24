package sec

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

const (
	FC_FOR_CK_PRIME_IK_PRIME_DERIVATION  = "20"
	FC_FOR_ALGORITHM_KEY_DERIVATION      = "69"
	FC_FOR_KAUSF_DERIVATION              = "6A"
	FC_FOR_RES_STAR_XRES_STAR_DERIVATION = "6B"
	FC_FOR_KSEAF_DERIVATION              = "6C"
	FC_FOR_KAMF_DERIVATION               = "6D"
	FC_FOR_KAMF_PRIME_DERIVATION         = "72"
	FC_FOR_KGNB_KN3IWF_DERIVATION        = "6E"
	FC_FOR_NH_DERIVATION                 = "6F"
)

func kdfLen(input []byte) []byte {
	r := make([]byte, 2)
	binary.BigEndian.PutUint16(r, uint16(len(input)))
	return r
}

// This function implements the KDF defined in TS.33220 cluase B.2.0.
//
// For P0-Pn, the ones that will be used directly as a string (e.g. "WLAN") should be type-casted by []byte(),
// and the ones originally in hex (e.g. "bb52e91c747a") should be converted by using hex.DecodeString().
func KDF(key []byte, FC string, param ...[]byte) (sum []byte, err error) {
	kdf := hmac.New(sha256.New, key)

	var s []byte
	if s, err = hex.DecodeString(FC); err != nil {
		return
	}

	for _, p := range param {
		s = append(append(s, p...), kdfLen(p)...)
	}

	if _, err = kdf.Write(s); err != nil {
		return
	}
	sum = kdf.Sum(nil)
	return
}

func SeafKey(key []byte, p ...[]byte) (sum []byte, err error) {
	sum, err = KDF(key, FC_FOR_KSEAF_DERIVATION, p...)
	return
}
func AlgKey(key []byte, p ...[]byte) (sum []byte, err error) {
	sum, err = KDF(key, FC_FOR_ALGORITHM_KEY_DERIVATION, p...)
	return
}

func RanKey(key []byte, p ...[]byte) (sum []byte, err error) {
	sum, err = KDF(key, FC_FOR_KGNB_KN3IWF_DERIVATION, p...)
	return
}
func NhKey(key []byte, p ...[]byte) (sum []byte, err error) {
	sum, err = KDF(key, FC_FOR_NH_DERIVATION, p...)
	return
}

func KAMF(kseaf, supi, abba []byte) (kamf []byte, err error) {
	kamf, err = KDF(kseaf, FC_FOR_KAMF_DERIVATION, supi, abba)
	return
}
func KamfPrime(kamf, direction, count []byte) (kamfprime []byte, err error) {
	kamfprime, err = KDF(kamf, FC_FOR_KAMF_PRIME_DERIVATION, direction, count)
	return
}

func KAUSF(ckik []byte, servingnet, sqnxorak []byte) (kausf []byte, err error) {
	kausf, err = KDF(ckik, FC_FOR_KAUSF_DERIVATION, servingnet, sqnxorak)
	return
}

func CkPrimeIkPrime(key []byte, servingnet []byte, sqnxorak []byte) (ck []byte, ik []byte, err error) {
	var buf []byte
	if buf, err = KDF(key, FC_FOR_CK_PRIME_IK_PRIME_DERIVATION, servingnet, sqnxorak); err == nil {
		l := len(buf) / 2
		ck = buf[:l]
		ik = buf[l:]
	}
	return
}

func ResstarXresstar(key, servingnet, rand, res []byte) (resstar []byte, xresstar []byte, err error) {
	var buf []byte
	if buf, err = KDF(key, FC_FOR_RES_STAR_XRES_STAR_DERIVATION, servingnet, rand, res); err == nil {
		l := len(buf) / 2
		resstar = buf[:l]
		xresstar = buf[l:]
	}
	return
}
