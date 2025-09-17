package cgo

/*
#cgo LDFLAGS: -L. -lzkbtc
#include "libzkbtc.h"
#include <stdlib.h>
*/
import "C"
import (
	"github.com/lightec-xyz/common/operations"
	"unsafe"
)

func BtcBaseProve(setupPath string, param string) (*operations.Proof, error) {
	cSetupPath := C.CString(setupPath)
	cParam := C.CString(param)
	defer C.free(unsafe.Pointer(cSetupPath))
	defer C.free(unsafe.Pointer(cParam))
	res := C.BtcBaseProve(cSetupPath, cParam)
	defer C.free(unsafe.Pointer(res))
	proof, err := parseRes(res)
	if err != nil {
		return nil, err
	}
	return proof, nil
}
