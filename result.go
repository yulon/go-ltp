package ltp

import (
	"C"
	"encoding/json"
)

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

func JsonToResult(j []byte) (r *Result) {
	r = &Result{}
	json.Unmarshal(j, r)
	return
}

func Analyze(line string, flag C.int) *Result {
	return JsonToResult(AnalyzeJson(line, flag))
}
