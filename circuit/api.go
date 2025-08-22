package main

/*
#include <stdlib.h>
typedef struct {
    int Code;
	char* Msg;
    char* Proof;
    char* Witness;
} Response;
typedef struct {
    int id;
    char* name;
    int age;
} BtcBaseReq;
typedef struct {
    int id;
    char* name;
    int age;
} BtcMiddleReq;
*/
import "C"
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	native_plonk "github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	"github.com/lightec-xyz/btc_provers/circuits/blockchain/baselevel"
	blockchainUtil "github.com/lightec-xyz/btc_provers/utils/blockchain"
	"github.com/lightec-xyz/common/operations"
	"reflect"
)

//export BtcBaseProve
func BtcBaseProve(path *C.char, req *C.BtcBaseReq) C.Response {
	setupDir := C.GoString(path)
	proofData := &blockchainUtil.BaseLevelProofData{}
	fmt.Printf("btcBaseProve setupDir: %v\n", setupDir)
	return ErrResp(errors.New("not implemented"))
	resp, err := baselevel.Prove(setupDir, proofData)
	if err != nil {
		return ErrResp(err)
	}
	proof, witness, err := ProofToStr(resp)
	if err != nil {
		return ErrResp(err)
	}
	return OKResp(proof, witness)
}

//export BtcMiddleProve
func BtcMiddleProve(req C.BtcMiddleReq) C.Response {
	name := C.GoString(req.name)
	fmt.Printf("btcMiddleProve name: %v\n", name)
	return OKResp("OK", "")
}

func main() {}

func ParseObj(data []byte, dst interface{}) error {
	if reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}
	err := json.Unmarshal(data, dst)
	if err != nil {
		return err
	}
	return nil
}

func ErrResp(err error) C.Response {
	return C.Response{
		Code: 0,
		Msg:  C.CString(fmt.Sprintf("cgo error: %s", err.Error())),
	}

}

func OKResp(proof, witness string) C.Response {
	return C.Response{
		Code:    1,
		Msg:     C.CString("OK"),
		Proof:   C.CString(proof),
		Witness: C.CString(witness),
	}
}

type FFIRes struct {
	Data string `json:"data"`
	Err  string `json:"err"`
	Code int    `json:"code"`
}

type HexProof struct {
	Proof   string // hex
	Witness string
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
