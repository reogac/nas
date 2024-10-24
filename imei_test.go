package nas

import (
	"testing"
)

func Test_Imei(t *testing.T) {
	imeiStr := "123203209090402"
	imei := new(Imei)
	var err error
	if err = imei.Parse(imeiStr); err != nil {
		t.Errorf("Parse IMEI fails: %+v", err)
	}
	var wire []byte
	if wire, err = imei.encode(); err != nil {
		t.Errorf("encode IMEI fails: %+v", err)
	}
	var newImei Imei
	if err = newImei.decode(wire); err != nil {
		t.Errorf("decode IMEI fails: %+v", err)
	}
	if imei.String() != newImei.String() {
		t.Errorf("not equal: %s != %s", imei.String(), newImei.String())
	}
}
