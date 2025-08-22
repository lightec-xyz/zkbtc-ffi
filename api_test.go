package zkbtc_ffi

import "testing"

func TestBtcBaseProve(t *testing.T) {
	proof, err := BtcBaseProve("btcSetupPath", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(proof)
}
