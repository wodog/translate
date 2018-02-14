package main

import (
	"os"
	"strings"

	"github.com/wodog/translate/baidu"
	"github.com/wodog/translate/youdao"
)

type translator interface {
	Translate(q string)
}

func main() {
	q := strings.Join(os.Args[1:], " ")

	var y translator
	y = youdao.New()
	y.Translate(q)

	var b translator
	b = baidu.New()
	b.Translate(q)
}
