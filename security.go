package nas

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"

	"github.com/aead/cmac"

	"nas/snow3g"
	"nas/zuc"
)

func NEA1(ck [16]byte, countC, bearer, direction uint32, ibs []byte, length uint32) (obs []byte, err error) {
	var k [4]uint32
	for i := uint32(0); i < 4; i++ {
		k[i] = binary.BigEndian.Uint32(ck[4*(3-i) : 4*(3-i+1)])
	}
	iv := [4]uint32{(bearer << 27) | (direction << 26), countC, (bearer << 27) | (direction << 26), countC}

	l := (length + 31) / 32
	r := length % 32
	ks := snow3g.GetKeyStream(k, iv, int(l))
	// Clear keystream bits which exceed length
	if r != 0 {
		ks[l-1] &= ^((1 << (32 - r)) - 1)
	}

	obs = make([]byte, len(ibs))
	var i uint32
	for i = 0; i < length/32; i++ {
		for j := uint32(0); j < 4; j++ {
			obs[4*i+j] = ibs[4*i+j] ^ byte((ks[i]>>(8*(3-j)))&0xff)
		}
	}
	if r != 0 {
		ll := (r + 7) / 8
		for j := uint32(0); j < ll; j++ {
			obs[4*i+j] = ibs[4*i+j] ^ byte((ks[i]>>(8*(3-j)))&0xff)
		}
	}
	return obs, nil
}

// ibs: input bit stream, obs: output bit stream
func NEA2(key [16]byte, count uint32, bearer uint8, direction uint8,
	ibs []byte,
) (obs []byte, err error) {
	// Couter[0..32] | BEARER[0..4] | DIRECTION[0] | 0^26 | 0^64
	couterBlk := make([]byte, 16)
	// First 32 bits are count
	binary.BigEndian.PutUint32(couterBlk, count)
	// Put Bearer and direction together
	couterBlk[4] = (bearer << 3) | (direction << 2)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	obs = make([]byte, len(ibs))

	stream := cipher.NewCTR(block, couterBlk)
	stream.XORKeyStream(obs, ibs)
	return obs, nil
}

// NEA3 ibs: input bit stream, obs: output bit stream
// ref: https://www.gsma.com/security/wp-content/uploads/2019/05/EEA3_EIA3_specification_v1_8.pdf
func NEA3(ck [16]byte, count uint32, bearer uint8, direction uint8,
	ibs []byte, length uint32,
) (obs []byte, err error) {
	iv := make([]byte, 16)
	binary.BigEndian.PutUint32(iv, count)
	iv[4] = (bearer << 3) | (direction << 2)

	for i := 0; i < 8; i++ {
		iv[i+8] = iv[i]
	}

	l := (length + 31) / 32
	stream := zuc.Zuc(ck[:], iv, l)

	obs = make([]byte, len(ibs))

	for i := 0; i < int(l); i++ {
		for j := 0; j < 4 && (i*4+j) < int((length+7)/8); j++ {
			obs[i*4+j] = ibs[i*4+j] ^ byte((stream[i]>>(8*(3-j)))&0xff)
		}
	}

	if length%8 != 0 {
		obs[length/8] &= (uint8(0xff) << (8 - length%8))
	}

	for j := int(length/8 + 1); j < len(obs); j++ {
		obs[j] = 0
	}
	return obs, nil
}

// mulx() is for NIA1()
func mulx(V, c uint64) uint64 {
	if V&0x8000000000000000 != 0 {
		return (V << 1) ^ c
	} else {
		return V << 1
	}
}

// mulxPow() is for NIA1()
func mulxPow(V, i, c uint64) uint64 {
	if i == 0 {
		return V
	} else {
		return mulx(mulxPow(V, i-1, c), c)
	}
}

// mul() is for NIA1()
func mul(V, P, c uint64) uint64 {
	rst := uint64(0)
	for i := uint64(0); i < 64; i++ {
		if (P>>i)&1 == 1 {
			rst ^= mulxPow(V, i, c)
		}
	}
	return rst
}

func NIA1(ik [16]byte, countI uint32, bearer byte, direction uint32, msg []byte, length uint64) (
	mac []byte, err error,
) {
	fresh := uint32(bearer) << 27
	var k [4]uint32
	for i := uint32(0); i < 4; i++ {
		k[i] = binary.BigEndian.Uint32(ik[4*(3-i) : 4*(3-i+1)])
	}
	iv := [4]uint32{fresh ^ (direction << 15), countI ^ (direction << 31), fresh, countI}
	D := ((length + 63) / 64) + 1
	z := snow3g.GetKeyStream(k, iv, 5)

	P := (uint64(z[0]) << 32) | uint64(z[1])
	Q := (uint64(z[2]) << 32) | uint64(z[3])

	var Eval uint64 = 0
	for i := uint64(0); i < D-2; i++ {
		M := binary.BigEndian.Uint64(msg[8*i:])
		Eval = mul(Eval^M, P, 0x000000000000001b)
	}

	tmp := make([]byte, 8)
	copy(tmp, msg[8*(D-2):])
	M := binary.BigEndian.Uint64(tmp)
	Eval = mul(Eval^M, P, 0x000000000000001b)

	Eval = Eval ^ length
	Eval = mul(Eval, Q, 0x000000000000001b)
	MacI := uint32(Eval>>32) ^ z[4]
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, MacI)
	return b, nil
}

func NIA2(key [16]byte, count uint32, bearer uint8, direction uint8, msg []byte) (mac []byte, err error) {
	// Couter[0..32] | BEARER[0..4] | DIRECTION[0] | 0^26
	m := make([]byte, len(msg)+8)
	// First 32 bits are count
	binary.BigEndian.PutUint32(m, count)
	// Put Bearer and direction together
	m[4] = (bearer << 3) | (direction << 2)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	copy(m[8:], msg)

	mac, err = cmac.Sum(m, block, 16)
	if err != nil {
		return nil, err
	}
	// only get the most significant 32 bits to be mac value
	mac = mac[:4]
	return mac, nil
}

// NIA3 ibs: input bit stream, obs: output bit stream (mac)
// ref: https://www.gsma.com/security/wp-content/uploads/2019/05/EEA3_EIA3_specification_v1_8.pdf
func NIA3(ik [16]byte, count uint32, bearer uint8, direction uint8,
	msg []byte, length uint32,
) (mac []byte, err error) {
	iv := make([]byte, 16)
	binary.BigEndian.PutUint32(iv, count)
	iv[4] = (bearer << 3) & 0xF8
	iv[5], iv[6], iv[7] = 0, 0, 0
	iv[8] = (direction << 7) ^ iv[0]

	for i := 0; i < 7; i++ {
		iv[i+9] = iv[i+1]
	}

	iv[14] = (direction << 7) ^ iv[6]
	l := (length+31)/32 + 2
	stream := zuc.Zuc(ik[:], iv, l)
	mac = genMac(msg, stream, int(length))
	return mac, nil
}

func genMac(m []byte, stream []uint32, blength int) []byte {
	var t uint32 = 0

	l := len(stream)

	for i := 0; i < blength; i++ {
		if m[i/8]&(1<<(7-(i%8))) != 0 { // GET_BIT
			t ^= getWord(stream, i)
		}
	}

	t ^= getWord(stream, blength)
	t ^= getWord(stream, 32*(l-1))

	mac := make([]byte, 4)
	binary.BigEndian.PutUint32(mac, t)
	return mac
}

func getWord(stream []uint32, i int) (zi uint32) {
	cntBackBit := i % 32
	cntFrontBit := 32 - cntBackBit
	loc := i / 32

	if cntBackBit == 0 {
		zi = stream[loc]
	} else {
		zi = stream[loc]<<(cntBackBit) | (stream[loc+1] >> cntFrontBit)
	}
	return zi
}
