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

func parseRes(res *C.char) (*operations.Proof, error) {
	goRes := C.GoString(res)
	var output FFIRes
	err := json.Unmarshal([]byte(goRes), &output)
	if err != nil {
		return nil, err
	}
	if output.Code == 0 {
		return nil, errors.New(output.Err)
	}
	proofBytes, err := ParseProof([]byte(output.Proof))
	if err != nil {
		return nil, err
	}
	witnessBytes, err := ParseWitness([]byte(output.Witness))
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
