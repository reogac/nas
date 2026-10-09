package nas

import "testing"

// An algorithm change keeps the NAS COUNTs of both ends where they are and
// keys both with the new algorithms: the UE's next message, protected with
// them under the next COUNT, verifies, and a message the old keys protected
// does not.
func TestAnAlgorithmChangeKeepsTheCounts(t *testing.T) {
	ue, amf := peers(t)
	for i := 0; i < 3; i++ {
		if !verified(amf, uplink(t, ue, NasSecBoth)) {
			t.Fatalf("message %d did not verify before the change", i)
		}
	}
	ul, dl := ue.UlCounter(), amf.DlCounter()
	kAmf := make([]byte, 32)
	for i := range kAmf {
		kAmf[i] = byte(i)
	}
	stale := uplink(t, ue, NasSecBoth)
	for _, ctx := range []*NasContext{ue, amf} {
		if err := ctx.ChangeAlgorithms(AlgCiphering128NEA1, AlgIntegrity128NIA1, kAmf); err != nil {
			t.Fatal(err)
		}
	}
	if enc, integ := amf.SelectedAlgorithms(); enc != AlgCiphering128NEA1 || integ != AlgIntegrity128NIA1 {
		t.Errorf("algorithms %d/%d after the change", enc, integ)
	}
	if ue.UlCounter() != ul+1 || amf.DlCounter() != dl {
		t.Errorf("COUNTs ul %d dl %d after the change, want %d and %d kept", ue.UlCounter(), amf.DlCounter(), ul+1, dl)
	}
	if verified(amf, stale) {
		t.Error("a message the old keys protected verified under the new ones")
	}
	if !verified(amf, uplink(t, ue, NasSecBoth)) {
		t.Error("the UE's next message under the new algorithms did not verify")
	}
}

// DeriveKeys, which takes a new K_AMF into use, still starts the COUNTs again.
func TestDerivingKeysStartsTheCountsAgain(t *testing.T) {
	ue, amf := peers(t)
	verified(amf, uplink(t, ue, NasSecBoth))
	if err := ue.DeriveKeys(AlgCiphering128NEA2, AlgIntegrity128NIA2, make([]byte, 32)); err != nil {
		t.Fatal(err)
	}
	if ue.UlCounter() != 0 {
		t.Errorf("uplink COUNT %d after a new K_AMF, want 0", ue.UlCounter())
	}
}
