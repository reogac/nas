package nas

import (
	"testing"
)

func Test_UeSecurityCapability(t *testing.T) {
	secCap := &UeSecurityCapability{}

	secCap.SetEA(0, true)
	secCap.SetIA(5, true)
	secCap.SetExtra([]byte{1, 2, 3})

	if secCap.GetEA(1) || secCap.GetIA(7) {
		t.Errorf("Get bit return invalid values")
	}
	if !secCap.GetEA(0) || !secCap.GetIA(5) {
		t.Errorf("Set bit fails")
	}
	extra := secCap.GetExtra()
	if len(extra) != 3 || extra[2] != 3 {
		t.Errorf("set/get extra fails")
	}

	if buf, err := secCap.encode(); err != nil {
		t.Errorf("encode fails")
	} else {
		var newSecCap UeSecurityCapability
		if err = newSecCap.decode(buf); err != nil {
			t.Errorf("decode fails")
		} else {
			if !newSecCap.GetEA(0) || !newSecCap.GetIA(5) {
				t.Errorf("decode wrong")
			}
			extra = newSecCap.GetExtra()
			if len(extra) != 3 || extra[2] != 3 {
				t.Errorf("decode wrong (extra)")
			}
		}
	}

}
