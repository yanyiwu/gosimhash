package gosimhash

import (
	"testing"
	"path"
	"runtime"
)

func getCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}

func TestSimhash(t *testing.T) {
	DICT_DIR := path.Join(path.Dir(getCurrentFilePath()), "deps/simhash/submodules/cppjieba/dict")
	DICT_PATH := path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH := path.Join(DICT_DIR, "hmm_model.utf8")
	IDF_PATH := path.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH := path.Join(DICT_DIR, "stop_words.utf8")
	hasher := New(
		DICT_PATH, 
		HMM_PATH, 
		IDF_PATH,
		STOP_WORDS_PATH,
	)
	defer hasher.Free()

	var s string
	var top_n int
	var expected uint64
	var actual uint64

	top_n = 5
	s = "我来到北京清华大学"
	expected = 0xfeb6372a8750eb1d
	actual = hasher.MakeSimhash(s, top_n)
	if expected != actual {
		t.Error(actual)
	}
}
