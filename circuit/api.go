package main

import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	native_plonk "github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	"github.com/lightec-xyz/btc_provers/circuits/blockchain/baselevel"
	blockchainUtil "github.com/lightec-xyz/btc_provers/utils/blockchain"
	"github.com/lightec-xyz/common/operations"
	"reflect"
)

//export BtcBaseProve
func BtcBaseProve(path *C.char, req *C.char) *C.char {
	setupDir := C.GoString(path)
	param := C.GoString(req)
	var data blockchainUtil.BaseLevelProofData
	err := ToObj(param, &data)
	if err != nil {
		return ErrResp(err)
	}
	fmt.Printf("btcBaseProve setupDir: %v %v \n", setupDir, param)
	resp, err := baselevel.Prove(setupDir, &data)
	if err != nil {
		return ErrResp(err)
	}
	return OKResp(resp)
}

func main() {}

func ToObj(s string, obj interface{}) error {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}
	return json.Unmarshal([]byte(s), obj)
}

func ToJson(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type FFIRes struct {
	Proof   string `json:"proof"`
	Witness string `json:"witness"`
	Err     string `json:"err"`
	Code    int    `json:"code"`
}

func ErrResp(err error) *C.char {
	res, _ := ToJson(FFIRes{
		Code: 0,
		Err:  err.Error(),
	})
	return C.CString(res)

}

func OKResp(proof *operations.Proof) *C.char {
	p, w, err := ProofToStr(proof)
	if err != nil {
		return ErrResp(err)
	}
	res, _ := ToJson(FFIRes{
		Code:    1,
		Proof:   p,
		Witness: w,
	})
	return C.CString(res)
}

func ProofToStr(proof *operations.Proof) (string, string, error) {
	proofBytes, err := ProofToBytes(proof.Proof)
	if err != nil {
		return "", "", err
	}
	witnessBytes, err := WitnessToBytes(proof.Witness)
	if err != nil {
		return "", "", err
	}
	return string(proofBytes), string(witnessBytes), nil
}

func ProofToBytes(proof native_plonk.Proof) ([]byte, error) {
	var buf bytes.Buffer
	_, err := proof.WriteTo(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func WitnessToBytes(witness witness.Witness) ([]byte, error) {
	var buf bytes.Buffer
	pubWit, err := witness.Public()
	if err != nil {
		return nil, err
	}
	_, err = pubWit.WriteTo(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
