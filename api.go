package zkbtc_ffi

import "C"
import (
	"github.com/lightec-xyz/common/operations"
	"github.com/lightec-xyz/zkbtc-ffi/cgo"
	"github.com/lightec-xyz/zkbtc-ffi/types"
)

func BtcBaseProve(setupPath string, data *types.BaseLevelProofData) (*operations.Proof, error) {
	proof, err := cgo.BtcBaseProve(setupPath, data)
	if err != nil {
		return nil, err
	}
	return proof, nil
}
