package gosimhash

import (
	"testing"
)

func TestSimhash(t *testing.T) {
	hasher := New()
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
