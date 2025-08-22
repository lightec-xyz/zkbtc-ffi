package cgo

import "C"
import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/consensys/gnark-crypto/ecc"
	native_plonk "github.com/consensys/gnark/backend/plonk"
	plonk_bn254 "github.com/consensys/gnark/backend/plonk/bn254"
	"github.com/consensys/gnark/backend/witness"
	"github.com/lightec-xyz/common/operations"
)

func NewString(s string) CString {
	cs := C.CString(s)
	return CString(*cs)
}

func parseResponse(res Response) (*operations.Proof, error) {
	if res.Code == 0 {
		return nil, errors.New(C.GoString(res.Msg))
	}
	proofBytes, err := ParseProof([]byte(C.GoString(res.Proof)))
	if err != nil {
		return nil, err
	}
	witnessBytes, err := ParseWitness([]byte(C.GoString(res.Witness)))
	if err != nil {
		return nil, err
	}
	return &operations.Proof{
		Proof:   proofBytes,
		Witness: witnessBytes,
	}, nil

}

func ParseFfiRes(res []byte) (*operations.Proof, error) {
	var resp FFIRes
	err := json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code == 0 {
		return nil, errors.New(resp.Err)
	}
	var proof Proof
	err = json.Unmarshal([]byte(resp.Data), &proof)
	if err != nil {
		return nil, err
	}
	proofBytes, err := ParseProof([]byte(proof.Proof))
	if err != nil {
		return nil, err
	}
	witnessBytes, err := ParseWitness([]byte(proof.Witness))
	if err != nil {
		return nil, err
	}
	return &operations.Proof{
		Proof:   proofBytes,
		Witness: witnessBytes,
	}, nil

}

func ParseWitness(body []byte) (witness.Witness, error) {
	field := ecc.BN254.ScalarField()
	buffer := bytes.NewBuffer(body)
	wit, err := witness.New(field)
	if err != nil {
		return nil, err
	}
	_, err = wit.ReadFrom(buffer)
	if err != nil {
		return nil, err
	}
	return wit, nil
}
func ParseProof(proof []byte) (native_plonk.Proof, error) {
	reader := bytes.NewReader(proof)
	var bn254Proof plonk_bn254.Proof
	_, err := bn254Proof.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return &bn254Proof, nil
}
