package nas

import (
	"math"
	"testing"
)

func Test_Counter(t *testing.T) {
	c := newCounter(100, 10)
	if c.overflow() != 100 {
		t.Errorf("wrong overflow")
	}

	if c.sqn() != 10 {
		t.Errorf("wrong sqn")
	}
	c.inc()

	if c.sqn() != 11 {
		t.Errorf("wrong inc")
	}
	c.setSqn(math.MaxUint8)
	c.inc()
	if c.sqn() != 0 || c.overflow() != 101 {
		t.Errorf("fail increase sqn with overflow")
	}
}
