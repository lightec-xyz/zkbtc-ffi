package cgo

/*
#cgo LDFLAGS: -L. -lzkbtc
#include "libzkbtc.h"
#include <stdlib.h>
*/
import "C"
import (
	"github.com/lightec-xyz/common/operations"
	"github.com/lightec-xyz/zkbtc-ffi/types"
	"unsafe"
)

func BtcBaseProve(setupPath string, data *types.BaseLevelProofData) (*operations.Proof, error) {
	cSetupPath := C.CString(setupPath)
	defer C.free(unsafe.Pointer(cSetupPath))
	req := BtcBaseReq{}
	defer req.Free()
	res := (Response)(C.BtcBaseProve(cSetupPath, (*C.BtcBaseReq)(&req)))
	defer res.Free()
	proof, err := parseResponse(res)
	if err != nil {
		return nil, err
	}
	return proof, nil
}

//func BtcMiddleProve(setupPath string, param string) (*operations.Proof, error) {
//	req := C.BtcMiddleReq{
//		id:   0,
//		name: C.CString("btcMiddle000001"),
//		age:  1,
//	}
//	ffiRes := C.BtcMiddleProve(req)
//	res := C.GoString(ffiRes)
//	return ParseFfiRes([]byte(res))
//}
