package zkbtc_ffi

import (
	"github.com/consensys/gnark/test"
	blockchainUtil "github.com/lightec-xyz/btc_provers/utils/blockchain"
	"github.com/lightec-xyz/btc_provers/utils/client"
	"github.com/lightec-xyz/zkbtc-ffi/helper"
	"testing"
)

func TestBtcBaseProve(t *testing.T) {
	url, user, pwd := helper.GetBtcUrl()
	assert := test.NewAssert(t)
	btcSetup, _ := helper.GetCircuitUrl()
	cl := client.NewJsonRpcClient(url, user, pwd)
	endBlockHeight := uint32(102311)
	proofData, err := blockchainUtil.GetBaseLevelProofData(cl, endBlockHeight)
	assert.NoError(err)
	proof, err := BtcBaseProve(btcSetup, proofData)
	assert.NoError(err)
	t.Log(proof)

}
