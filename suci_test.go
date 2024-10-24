package nas

import (
	"testing"
)

func Test_SuciNai(t *testing.T) {
	testCases := make(map[string]string)
	testCases["gci-"] = "gci-"
	testCases["nai-324024243243"] = "nai-324024243243"
	testCases["suci-12929-6783198712"] = "imsi-12929-6783198712"
	testCases["imsi-129290-0000000001"] = "imsi-129290-0000000001"
	testCases["suci-129-29-0001-1-1-6783198712"] = "suci-129-29-0001-1-1-6783198712"
	testCases["suci-129-29-0001-0-1-6783198712"] = "imsi-12929-6783198712"
	var err error
	for suciStr, expectedSuciStr := range testCases {
		suci := new(Suci)
		if err = suci.Parse(suciStr); err != nil {
			t.Errorf("NAI parse fail: %+v", err)
			continue
		}
		var buf []byte
		if buf, err = suci.encode(); err != nil {
			t.Errorf("encode suci fail: %+v", err)
			return
		}
		var newSuci = new(Suci)

		if err = newSuci.decode(buf); err != nil {
			t.Errorf("decode suci fail: %+v", err)
			return
		}

		if expectedSuciStr != newSuci.String() {
			t.Errorf("expected suci=%s, parsed: %s", expectedSuciStr, newSuci.String())
			return
		}
	}
}
