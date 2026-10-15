package gosimhash

/*
#cgo CXXFLAGS: -I./deps/simhash/include -I./deps/simhash/submodules/cppjieba/include -I./deps/simhash/submodules/cppjieba/deps/limonp/include -DLOGGING_LEVEL=LL_WARNING -O3 -std=c++11
#include <stdlib.h>
#include "simhash.h"
*/
import "C"

import "unsafe"

type Simhasher struct {
	hasher C.Simhasher
}

func New(paths ...string) Simhasher {
	dictpaths := getDictPaths(paths...)

	dpath, hpath, ipath, spath := C.CString(dictpaths[0]), C.CString(dictpaths[1]), C.CString(dictpaths[3]), C.CString(dictpaths[4])
	defer C.free(unsafe.Pointer(dpath))
	defer C.free(unsafe.Pointer(hpath))
	defer C.free(unsafe.Pointer(ipath))
	defer C.free(unsafe.Pointer(spath))

	var x Simhasher
	x.hasher = C.NewSimhasher(dpath, hpath, ipath, spath)
	return x
}

func (x Simhasher) Free() {
	C.FreeSimhasher(x.hasher)
}

func (x Simhasher) MakeSimhash(s string, top_n int) uint64 {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	return uint64(C.MakeSimhash(x.hasher, cs, C.int(top_n)))
}
