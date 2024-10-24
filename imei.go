package nas

import (
	// "encoding/hex"
	//	"encoding/binary"
	"fmt"
)

type Imei struct {
	IsSv   bool
	digits []byte //valid decimal digits
}

func (id *Imei) encode() (wire []byte, err error) {
	oneByte := id.getIdentityType() & 0x07 //last 3 bits
	if len(id.digits)%2 == 1 {             //Odd number of digits
		oneByte += 0x08 //set bit 4
	}
	wire = []byte{oneByte} //add first byte
	halfByte := false      //has a digit occupied the right half?
	for i, d := range id.digits {
		if i == 0 { //first digit
			wire[0] += d << 4 //first digit occupies left half of byte 0
			continue
		}
		if !halfByte { //right half not yet occupied
			oneByte = d & 0x0f
			halfByte = true
		} else {
			oneByte += d << 4 //fill left half
			//add the full byte
			wire = append(wire, oneByte)
			halfByte = false
		}
	}
	if halfByte { //last byte is half-occupied
		wire = append(wire, oneByte)
	}
	return
}
func (id *Imei) decode(wire []byte) (err error) {
	numDigits := (len(wire)-1)*2 + int(getBit(wire[0], 4)) //twice remaining octets and oddity
	id.digits = make([]byte, numDigits)
	octetId := 0
	for i := 0; i < numDigits; i++ {
		if i == 0 { //first digits
			id.digits[i] = wire[octetId] >> 4 //left half
			octetId++
			continue
		}
		if i%2 == 0 { //event-index digit = left half
			id.digits[i] = wire[octetId] >> 4
			octetId++
		} else {
			id.digits[i] = wire[octetId] & 0x0f //right half
		}
	}
	//check for digit validity
	for i := 0; i < numDigits; i++ {
		if id.digits[i] > 9 {
			err = fmt.Errorf("digit %d[%d] is invalid", i+1, id.digits[i])
			return
		}
	}
	return
}

func (id *Imei) getIdentityType() uint8 {
	if id.IsSv {
		return MobileIdentity5GSTypeImeisv //ImeiSv type
	}
	return MobileIdentity5GSTypeImei //Imei type
}

func (id *Imei) Parse(s string) (err error) {
	id.digits, err = decimalBytes(s)
	return
}
func (id *Imei) String() string {
	digitStr := make([]byte, len(id.digits))
	for i := 0; i < len(digitStr); i++ {
		digitStr[i] = id.digits[i] + '0'
	}
	return string(digitStr)
}
