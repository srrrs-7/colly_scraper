package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr. Smith"

func NewNltk() {
	str := strings.Replace(st, "'", " ' ", -1)
	str = strings.Replace(str, ".", " . ", -1)
	s := strings.Split(str, " ")


	fmt.Println(str)
	fmt.Println(s)
}