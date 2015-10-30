package ltp

/*
#cgo CFLAGS: -O2
#cgo LDFLAGS: -ljsonltp

#include "jsonltp.h"
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"unsafe"
)

func Init(dataDir string) {
	C.jsonltp_init(C.CString(dataDir))
}

func Release() {
	C.jsonltp_release()
}

const (
	FlagTag = C.JSONLTP_FLAG_TAG
	FlagNer = C.JSONLTP_FLAG_NER
	FlagDp = C.JSONLTP_FLAG_DP
	FlagSrl = C.JSONLTP_FLAG_SRL
	FlagAll = C.JSONLTP_FLAG_ALL
)

func Analyze(line string, flag C.int) (r *Result) {
	cLine := C.CString(line)
	cJson := C.jsonltp_analyze(cLine, flag)
	C.free(unsafe.Pointer(cLine))

	j := []byte(C.GoString(cJson))
	C.free(unsafe.Pointer(cJson))

	r = &Result{}
	json.Unmarshal(j, r)
	return
}

type Result struct {
	Words []string `json:"words"`
	Tags []string `json:"tags"`
	Nes []struct {
		I int `json:"i"`
		Ne string `json:"ne"`
	} `json:"nes"`
	Parse []struct {
		Parent int `json:"parent"`
		Deprel string `json:"deprel"`
	} `json:"parse"`
	Srl []struct {
		I int `json:"i"`
		Args []struct {
			Type string `json:"type"`
			Beg int `json:"beg"`
			End int `json:"end"`
		} `json:"args"`
	} `json:"srl"`
}
