package ltp

/*
#cgo CFLAGS: -O2
#cgo LDFLAGS: -ljsonltp

#include "jsonltp.h"
#include <stdlib.h>
*/
import "C"

import (
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

func AnalyzeJson(line string, flag C.int) (j []byte) {
	cLine := C.CString(line)
	cJson := C.jsonltp_analyze(cLine, flag)
	C.free(unsafe.Pointer(cLine))

	j = []byte(C.GoString(cJson))
	C.free(unsafe.Pointer(cJson))
	return
}
