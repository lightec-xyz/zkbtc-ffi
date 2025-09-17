package cgo

type FFIRes struct {
	Proof   string `json:"proof"`
	Witness string `json:"witness"`
	Err     string `json:"err"`
	Code    int    `json:"code"`
}
