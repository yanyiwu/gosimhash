package main

import (
	"flag"
	"fmt"

	"github.com/yanyiwu/gosimhash"
)

var sentence = flag.String("sentence", "我来到北京清华大学", "")
var top_n = flag.Int("top_n", 5, "")

func main() {
	flag.Parse()

	hasher := gosimhash.New()
	defer hasher.Free()
	fingerprint := hasher.MakeSimhash(*sentence, *top_n)
	fmt.Printf("%s simhash: %x\n", *sentence, fingerprint)
}
