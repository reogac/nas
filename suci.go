package nas

import (
	"encoding/hex"
	//	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// TS 24.501 9.11.3.4
const (
	SupiFormatImsi uint8 = 0x00
	SupiFormatNai  uint8 = 0x01
	SupiFormatGci  uint8 = 0x02
	SupiFormatGli  uint8 = 0x03
)

// TS 24.501 9.11.3.4
const (
	ProtectionSchemeNullScheme    uint8 = 0
	ProtectionSchemeECIESProfileA uint8 = 1
	ProtectionSchemeECIESProfileB uint8 = 2
)

type SuciContent interface {
	Decoder
	Encoder
	getSupiFormat() uint8
	String() string
}

type Suci struct {
	Content SuciContent //SupiImsi or SupiOther
}

func (id *Suci) GetSupiFormat() uint8 {
	return id.Content.getSupiFormat()
}

func (id *Suci) getIdentityType() uint8 {
	return MobileIdentity5GSTypeSuci //Suci type
}

func (id *Suci) String() string {
	return id.Content.String()
}

func (id *Suci) Parse(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("Empty identity")
	}
	parts := strings.Split(s, "-")
	switch parts[0] {
	case "nai", "gci", "gli":
		c := new(SupiOther)
		if parts[0] == "nai" {
			c.format = SupiFormatNai
		} else if parts[0] == "gci" {
			c.format = SupiFormatGci
		} else {
			c.format = SupiFormatGli
		}
		if len(parts) != 2 {
			return fmt.Errorf("invalid supi format")
		}

		if err := c.Parse(parts[1]); err != nil {
			return nasError("parse suci (not imsi)", err)
		}
		id.Content = c

	case "suci", "imsi":
		c := new(SupiImsi)
		if err := c.Parse(parts[1:]); err != nil {
			return nasError("parse suci Imsi", err)
		}
		id.Content = c
	default:
		return fmt.Errorf("Invalid prefix for suci: %s", parts[0])
	}
	return nil
}

func (id *Suci) encode() (wire []byte, err error) {
	if wire, err = id.Content.encode(); err != nil {
		return
	}
	//get id type
	idType := id.getIdentityType()
	supiFormat := id.Content.getSupiFormat()
	//add supi format
	idType = idType&0x0f + supiFormat<<4
	wire = append([]byte{idType}, wire...)
	return
}

// wire not empty
func (id *Suci) decode(wire []byte) (err error) {
	//get supi format
	supiFormat := (wire[0] & 0x70) >> 4
	switch supiFormat {
	case SupiFormatNai, SupiFormatGci, SupiFormatGli:
		content := &SupiOther{
			format: supiFormat,
		}
		if err = content.decode(wire[1:]); err == nil {
			id.Content = content
		} else {
			err = nasError("decode suci content (not IMSI)", err)
		}
	default: //SupiFormatImsi uint8 = 0x00
		content := &SupiImsi{}
		if err = content.decode(wire[1:]); err == nil {
			id.Content = content
		} else {
			err = nasError("decode suci content (IMSI)", err)
		}
	}

	return
}

// Supi NAI/GCI/GLI
type SupiOther struct {
	format uint8
	Bytes  []byte
}

func NewSupiOther(format uint8, content []byte) *SupiOther {
	return &SupiOther{
		format: format,
		Bytes:  content,
	}
}

// format  = [nai, gci, gli]-content
func (c *SupiOther) Parse(s string) (err error) {
	c.Bytes, err = hex.DecodeString(s)
	return
}

func (c *SupiOther) String() string {
	return fmt.Sprintf("%s-%x", SupiFormatString(c.format), c.Bytes)
}

func (c *SupiOther) getSupiFormat() uint8 {
	return c.format
}
func (c *SupiOther) encode() (wire []byte, err error) {
	wire = c.Bytes
	return
}
func (c *SupiOther) decode(wire []byte) (err error) {
	c.Bytes = make([]byte, len(wire))
	copy(c.Bytes, wire)
	return
}

type SupiImsi struct {
	PlmnId                 PlmnId
	RoutingInd             RoutingIndicator
	ProtectionScheme       uint8
	HomeNetworkPublicKeyId uint8
	SchemeOutput           []byte
}

func (c *SupiImsi) getSupiFormat() uint8 {
	return SupiFormatImsi
}

func (c *SupiImsi) encode() (wire []byte, err error) {
	wire = make([]byte, 7)
	copy(wire, c.PlmnId.bytes[:])         //3 bytes
	copy(wire[3:], c.RoutingInd.bytes[:]) //2 bytes
	wire[5] = c.ProtectionScheme & 0x0f
	wire[6] = c.HomeNetworkPublicKeyId
	wire = append(wire, c.SchemeOutput...)
	return
}
func (c *SupiImsi) decode(wire []byte) (err error) {
	if len(wire) < 7 {
		err = nasError("decode supi IMSI", err)
		return
	}
	c.PlmnId.decode(wire[0:3])     //no error
	c.RoutingInd.decode(wire[3:5]) //no error
	c.ProtectionScheme = wire[5]
	c.HomeNetworkPublicKeyId = wire[6]
	c.SchemeOutput = make([]byte, len(wire)-7)
	copy(c.SchemeOutput, wire[7:])
	return
}

// suci-mcc-mnc-routeId-scheme-keyId-output (suci)
// imsi-plmnid-msin (supi-imsi)
func (c *SupiImsi) Parse(parts []string) (err error) {
	if len(parts) == 2 { //supi
		if err = c.PlmnId.Parse(parts[0]); err != nil {
			return nasError("Parse PlmnId", err)
		}
		if c.SchemeOutput, err = ParseMsin(parts[1]); err != nil {
			return nasError("Parse MSIN", err)
		}
		c.ProtectionScheme = ProtectionSchemeNullScheme

	} else if len(parts) == 6 { //concealed supi
		//plmnid
		if err = c.PlmnId.Parse(parts[0] + parts[1]); err != nil {
			return nasError("Parse PlmnId", err)
		}
		//routing indicator
		if err = c.RoutingInd.Parse(parts[2]); err != nil {
			return nasError("Parse routing indicator", err)
		}
		//protection scheme
		var tmp int
		if tmp, err = strconv.Atoi(parts[3]); err != nil {
			return fmt.Errorf("Invalid protection scheme")
		}
		switch uint8(tmp) {
		case ProtectionSchemeNullScheme:
		case ProtectionSchemeECIESProfileA:
		case ProtectionSchemeECIESProfileB:
		default:
			return fmt.Errorf("Unknown protection scheme")
		}
		c.ProtectionScheme = uint8(tmp)

		//home network public key id
		if tmp, err = strconv.Atoi(parts[4]); err != nil {
			return fmt.Errorf("Invalid home network public key id")
		}
		c.HomeNetworkPublicKeyId = uint8(tmp)
		if c.ProtectionScheme == ProtectionSchemeNullScheme {
			//parse MSIN
			c.SchemeOutput, err = ParseMsin(parts[5])
		} else {
			c.SchemeOutput, err = hex.DecodeString(parts[5])
		}

	} else {
		err = fmt.Errorf("Invalid suci format")
	}
	return
}

// suci-mcc-mnc-routeId-scheme-keyId-output
// or imsi-plmnId-msin
func (c *SupiImsi) String() string {
	mcc, mnc := c.PlmnId.Get()
	if c.ProtectionScheme == ProtectionSchemeNullScheme { //plain supi-imsi
		msin := MsinFromBytes(c.SchemeOutput)
		return fmt.Sprintf("imsi-%s%s-%s", mcc, mnc, msin)
	}
	//concealed supi (aka suci)
	rId := c.RoutingInd.String()
	return fmt.Sprintf("suci-%s-%s-%s-%d-%d-%x", mcc, mnc, rId, c.ProtectionScheme, c.HomeNetworkPublicKeyId, c.SchemeOutput)
}

type RoutingIndicator struct {
	bytes [2]byte
}

func (ri *RoutingIndicator) String() string {
	return fmt.Sprintf("%d%d%d%d", ri.bytes[0]&0x0f, (ri.bytes[0]&0xf0)>>4, ri.bytes[1]&0x0f, (ri.bytes[1]&0xf0)>>4)
}

func (ri *RoutingIndicator) Parse(s string) (err error) {
	var digits []byte
	if digits, err = decimalBytes(s); err != nil {
		err = nasError("parse routing indicator", err)
		return
	}

	if len(digits) != 4 {
		err = fmt.Errorf("Routing indicator must be 4 digits")
		return
	}
	ri.bytes[0] = digits[0] + digits[1]<<4
	ri.bytes[1] = digits[2] + digits[3]<<4
	return
}

func (ri *RoutingIndicator) encode() (wire []byte, err error) {
	wire = ri.bytes[:]
	return
}

func (ri *RoutingIndicator) decode(wire []byte) (err error) {
	if len(wire) != 2 {
		err = ErrInvalidSize
		return
	}
	copy(ri.bytes[:], wire)
	return
}

func SupiFormatString(format uint8) string {
	switch format {
	case SupiFormatNai:
		return "nai"
	case SupiFormatGci:
		return "gci"
	case SupiFormatGli:
		return "gli"
	default: //SupiFormatImsi uint8 = 0x00
	}
	return "imsi"
}

func MsinFromBytes(buf []byte) string {
	strBytes := make([]byte, 2*len(buf))
	i := 0
	for _, b := range buf {
		strBytes[i] = b&0x0f + '0'
		strBytes[i+1] = b>>4 + '0'
		i += 2
	}
	l := len(strBytes)
	if strBytes[l-1] > '9' {
		return string(strBytes[0 : l-1])
	}
	return string(strBytes)
}

func ParseMsin(msin string) (wire []byte, err error) {
	var digits []byte
	if digits, err = decimalBytes(msin); err != nil {
		return
	}
	if len(digits)%2 == 1 {
		digits = append(digits, 0x0f)
	}
	wire = make([]byte, len(digits)/2)
	for i := 0; i < len(wire); i++ {
		wire[i] = digits[2*i]&0x0f + digits[2*i+1]<<4
	}
	return
}

// get supi in original form [mcc+mnc+msin]
func GetSupiString(supi *SupiImsi) string {
	if supi.ProtectionScheme != ProtectionSchemeNullScheme {
		return ""
	}
	return supi.PlmnId.String() + MsinFromBytes(supi.SchemeOutput)
}
