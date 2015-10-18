package gosimhash

/*
#cgo CXXFLAGS: -DLOGGER_LEVEL=LL_WARN -O3 -Wall
#include <stdlib.h>
#include "simhash.h"
*/
import "C"

type Simhasher struct {
	hasher C.Simhasher
}

func New(dict_path, hmm_path, idf_path, stop_words_path string) Simhasher {
	var x Simhasher
	x.hasher = C.NewSimhasher(C.CString(dict_path), C.CString(hmm_path), C.CString(idf_path), C.CString(stop_words_path))
	return x
}

func (x Simhasher) Free() {
	C.FreeSimhasher(x.hasher)
}

func (x Simhasher) MakeSimhash(s string, top_n int) uint64 {
	return uint64(C.MakeSimhash(x.hasher, C.CString(s), C.int(top_n)))
}
