package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr. Smith"

func NewNltk() {
	s := []string{}
	s = strings.Split(st, " ")
	str = strings.Replace(st, "'", " ' ", -1)

	fmt.Println(s)
}