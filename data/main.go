package main

import (
	"flag"
	"fmt"

	"github.com/yanjingang/gse"
)

var (
	text     = "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄"
	flagText = flag.String("text", text, "要分词的文本")
)

func main() {
	flag.Parse()

	var seg gse.Segmenter
	// seg.LoadDict("./data/dict/dictionary.txt")
	seg.LoadDict()

	// use DAG and HMM
	hmm := seg.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)
	//
	cut := seg.Cut(text)
	fmt.Println("cut: ", cut)

	hmm = seg.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)
	//
	cut = seg.CutSearch(text)
	fmt.Println("cut search: ", hmm)

	cut = seg.CutAll(text)
	fmt.Println("cut all: ", cut)

	segments := seg.Segment([]byte(*flagText))
	fmt.Println(gse.ToString(segments, true))
}
