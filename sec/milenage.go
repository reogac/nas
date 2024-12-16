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
package sec

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "milenage"})
}

type Milenage struct {
	block   cipher.Block
	opc     [16]uint8
	randxor [16]uint8 //E_k(rand XOR opc)
	rand    [16]uint8
	reader  io.Reader
}

func NewMilenage(k []uint8, opopc []uint8, isopc bool) (m *Milenage, err error) {
	//default rand reader
	m, err = NewMilenageEx(k, rand.Reader, opopc, isopc)
	return
}

// with customized rand reader
func NewMilenageEx(k []uint8, r io.Reader, opopc []uint8, isopc bool) (m *Milenage, err error) {
	if len(k) != 16 || len(opopc) != 16 {
		err = fmt.Errorf("Wrong input size")
		return
	}
	m = &Milenage{
		reader: r,
	}
	if r == nil {
		m.reader = rand.Reader
	}

	log.Tracef("milenage key=%x", k)
	if m.block, err = aes.NewCipher(k); err != nil {
		return
	}

	//generate opc if it is needed
	if !isopc {
		m.block.Encrypt(m.opc[:], opopc)
		for i := 0; i < 16; i++ {
			m.opc[i] ^= opopc[i]
		}
	} else {
		copy(m.opc[:], opopc[:])

	}
	log.Tracef("milenage opc=%x", m.opc[:])
	m.Refresh()
	return
}

// prepare a new random vector
func (m *Milenage) Refresh() {
	m.reader.Read(m.rand[:])
	//	RAND, _ := hex.DecodeString("FEE14720A39CC164BCC0E3628452847C")
	//copy(m.rand[:], RAND)
	var tmp [16]uint8
	for i := 0; i < 16; i++ {
		tmp[i] = m.rand[i] ^ m.opc[i]
	}
	m.block.Encrypt(m.randxor[:], tmp[:])
}

// set a new random vector
func (m *Milenage) SetRand(r []uint8) error {
	if len(r) != 16 {
		return fmt.Errorf("Wrong rand size")
	}
	copy(m.rand[:], r)
	var tmp [16]uint8
	for i := 0; i < 16; i++ {
		tmp[i] = m.rand[i] ^ m.opc[i]
	}
	m.block.Encrypt(m.randxor[:], tmp[:])
	return nil
}

func (m *Milenage) GetRand() []uint8 {
	return m.rand[:]
}

// f1 and f1star
func (m *Milenage) F1(sqn, amf []uint8) (maca []uint8, macs []uint8, err error) {
	if len(sqn) != 6 || len(amf) != 2 {
		err = fmt.Errorf("Wrong size input")
		return
	}
	var a, b, c [16]uint8
	//b = sqn || amf || sqn || amf
	copy(b[0:], sqn[:])
	copy(b[6:], amf[:])
	copy(b[8:], b[0:8])

	// c = (rot(b XOR opc), 8) XOR randxor
	var j int
	for i := 0; i < 16; i++ {
		j = (i + 8) % 16
		c[j] = b[i] ^ m.opc[i] ^ m.randxor[j]
	}

	// a = E_k(c) XOR opc
	m.block.Encrypt(a[:], c[:])
	for i := 0; i < 16; i++ {
		a[i] ^= m.opc[i]
	}
	maca = a[0:8]
	macs = a[8:]
	return
}

func (m *Milenage) F2F5() ([]uint8, []uint8) {
	tmp := m.operation(0, 1)
	return tmp[8:16], tmp[:6] // res, ak
}

func (m *Milenage) F3() []uint8 {
	return m.operation(12, 2) //ck
}

func (m *Milenage) F4() []uint8 {
	return m.operation(8, 4) //ik
}

func (m *Milenage) F5star() []uint8 {
	tmp := m.operation(4, 8)
	return tmp[:6] //akstar
}

func (m *Milenage) operation(rot int, v uint8) []uint8 {
	var a, b, c [16]uint8
	c[15] = v
	var j int
	//a= rotate(randxor XOR opc, rot) XOR c
	for i := 0; i < 16; i++ {
		j = (i + rot) % 16
		a[j] = m.randxor[i] ^ m.opc[i] ^ c[j]
	}

	//b = E_k(a) XOR opc
	m.block.Encrypt(b[:], a[:])
	for i := 0; i < 16; i++ {
		b[i] ^= m.opc[i]
	}
	return b[:]
}

func (m *Milenage) ValidateAuts(auts, randv []byte) (sqn [6]uint8, err error) {

	if len(auts) != 14 || len(randv) != 16 {
		err = fmt.Errorf("Wrong input size:auts[%d],rand[%d]", len(auts), len(randv))
		return
	}

	m.SetRand(randv) //never fails

	var amf [2]uint8 //resync: dummy amf='0000'

	ak_r := m.F5star()
	for i := 0; i < 6; i++ {
		sqn[i] = ak_r[i] ^ auts[i]
	}
	_, macs, _ := m.F1(sqn[:], amf[:]) //never fails

	if bytes.Compare(macs, auts[6:]) != 0 {
		err = fmt.Errorf("MAC failed: calculated MAC=%x, received MAC=%x", macs, auts[6:])
	}
	return
}

/*
func OPC(k []uint8, op []uint8) (opc [16]uint8, err error) {
	if len(k) != 16 || len(op) != 16 {
		err = fmt.Errorf("Wrong parameter size")
		return
	}
	var block cipher.Block
	if block, err = aes.NewCipher(k); err != nil {
		return
	}
	block.Encrypt(opc[:], op)
	for i := 0; i < 16; i++ {
		opc[i] ^= op[i]
	}
	return
}
*/
