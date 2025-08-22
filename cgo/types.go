package cgo

/*
#include "libzkbtc.h"
#include <stdlib.h>
*/
import "C"

type ICgo interface {
	Free()
}

type BtcBaseReq C.BtcBaseReq

func (b *BtcBaseReq) Free() {
	//todo
}

type CString C.char

type Response C.Response

func (r *Response) Free() {
	//todo
}

type Proof struct {
	Proof   string
	Witness string
}

type FFIRes struct {
	Data string `json:"data"`
	Err  string `json:"err"`
	Code int    `json:"code"`
}
