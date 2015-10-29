package ltp

/*
#cgo CFLAGS: -O2
#cgo LDFLAGS: -ljsonltp

#include "jsonltp.h"
*/
import "C"

import (
	"encoding/json"
	"strings"
)

func Init(dataDir string) {
	C.jsonltp_init(C.CString(dataDir))
}

func Close() {
	C.jsonltp_close()
}

const (
	FlagTag = C.JSONLTP_FLAG_TAG
	FlagNer = C.JSONLTP_FLAG_NER
	FlagDp = C.JSONLTP_FLAG_DP
	FlagSrl = C.JSONLTP_FLAG_SRL
	FlagAll = C.JSONLTP_FLAG_ALL
)

func DoJson(line string, buf *C.char, flag C.int) int {
	return int(C.jsonltp(C.CString(line), buf, flag))
}

var spstr = strings.Repeat(" ", 32768)

func Do(line string, flag C.int) (r *Result) {
	buf := C.CString(spstr)
	leng := int(C.jsonltp(C.CString(line), buf, flag))
	if leng > 32768 {
		buf = C.CString(strings.Repeat(" ", leng))
		C.jsonltp(C.CString(line), buf, flag)
	}
	r = &Result{}
	json.Unmarshal([]byte(C.GoString(buf)), r)
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
