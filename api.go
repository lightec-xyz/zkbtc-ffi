package zkbtc_ffi

import "C"
import (
	blockchainUtil "github.com/lightec-xyz/btc_provers/utils/blockchain"
	"github.com/lightec-xyz/common/operations"
	"github.com/lightec-xyz/zkbtc-ffi/cgo"
	"github.com/lightec-xyz/zkbtc-ffi/helper"
)

func BtcBaseProve(setupPath string, data *blockchainUtil.BaseLevelProofData) (*operations.Proof, error) {
	param, err := helper.ToJson(data)
	if err != nil {
		return nil, err
	}
	proof, err := cgo.BtcBaseProve(setupPath, param)
	if err != nil {
		return nil, err
	}
	return proof, nil
}
