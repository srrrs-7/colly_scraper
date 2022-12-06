package mecab

import (
	"fmt"

	"github.com/shogo82148/go-mecab"
)

func NewMecab() {
	tagger, err := mecab.New(map[string]string{"output-format-type": "wakati"})
	if err != nil {
		panic(err)
	}
	defer tagger.Destroy()

	parseMecab(tagger)
}

func parseMecab(tagger mecab.MeCab) mecab.Node {
	node, err := tagger.ParseToNode("こんにちは世界")
	if err != nil {
		panic(err)
	}
	fmt.Println(node)
	return node
}

func mecabDetail(node mecab.Node) {}

