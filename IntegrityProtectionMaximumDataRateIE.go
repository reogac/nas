package nas

// 9.11.4.7 Integrity protection maximum data rate [M V 2]
type IntegrityProtectionMaximumDataRate struct {
	uplink   uint8
	downlink uint8
}

func NewIntegrityProtectionMaximumDataRate(u, d uint8) IntegrityProtectionMaximumDataRate {
	return IntegrityProtectionMaximumDataRate{
		uplink:   u,
		downlink: d,
	}
}

func (ie *IntegrityProtectionMaximumDataRate) Uplink() uint8 {
	return ie.uplink
}

func (ie *IntegrityProtectionMaximumDataRate) Downlink() uint8 {
	return ie.downlink
}

func (ie *IntegrityProtectionMaximumDataRate) encode() (wire []byte, err error) {
	wire = []byte{ie.uplink, ie.downlink}
	return
}

func (ie *IntegrityProtectionMaximumDataRate) decode(wire []byte) (err error) {
	if len(wire) != 2 {
		err = ErrInvalidSize
		return
	}
	ie.uplink, ie.downlink = wire[0], wire[1]
	return
}
