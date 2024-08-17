package gosimhash

/*
#cgo CXXFLAGS: -I./deps/simhash/include -I./deps/simhash/submodules/cppjieba/include -I./deps/simhash/submodules/cppjieba/deps/limonp/include -DLOGGING_LEVEL=LL_WARNING -O3 
#include <stdlib.h>
#include "simhash.h"
*/
import "C"

type Simhasher struct {
	hasher C.Simhasher
}

func New(paths ...string) Simhasher {
	dictpaths := getDictPaths(paths...)

	dpath, hpath, _, ipath, spath := C.CString(dictpaths[0]), C.CString(dictpaths[1]), C.CString(dictpaths[2]), C.CString(dictpaths[3]), C.CString(dictpaths[4])
	var x Simhasher
	x.hasher = C.NewSimhasher(dpath, hpath, ipath, spath)
	return x
}

func (x Simhasher) Free() {
	C.FreeSimhasher(x.hasher)
}

func (x Simhasher) MakeSimhash(s string, top_n int) uint64 {
	return uint64(C.MakeSimhash(x.hasher, C.CString(s), C.int(top_n)))
}
