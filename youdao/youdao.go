package youdao

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)

type youdaoTranslator struct {
	apiAddress string
	keyfrom    string
	key        string
	doctype    string
}

func New() *youdaoTranslator {
	return &youdaoTranslator{
		apiAddress: "http://fanyi.youdao.com/openapi.do",
		keyfrom:    "afdsfsdfdsfdsjkh",
		key:        "600847748",
		doctype:    "json",
	}
}

func (t *youdaoTranslator) Translate(q string) {
	q = url.QueryEscape(q)
	url := fmt.Sprintf("%s?keyfrom=%s&key=%s&type=data&doctype=%s&version=1.1&q=%s", t.apiAddress, t.keyfrom, t.key, t.doctype, q)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result := gjson.GetBytes(data, "translation.0")
	fmt.Println(result)
}
