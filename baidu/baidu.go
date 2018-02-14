package baidu

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)

type baiduTranslator struct {
	from   string
	to     string
	appid  string
	appkey string
	salt   string
}

func New() *baiduTranslator {
	return &baiduTranslator{
		from:   "auto",
		to:     "zh",
		appid:  "20180210000122212",
		appkey: "YJU4LcX3XzaQR1fFXQtW",
		salt:   "wodog",
	}
}

func (t *baiduTranslator) Translate(q string) {
	sign := fmt.Sprintf("%x", md5.Sum([]byte(t.appid+q+t.salt+t.appkey)))
	q = url.QueryEscape(q)
	url := fmt.Sprintf("http://api.fanyi.baidu.com/api/trans/vip/translate?from=%s&to=%s&appid=%s&salt=%s&sign=%s&q=%s", t.from, t.to, t.appid, t.salt, sign, q)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result := gjson.GetBytes(data, "trans_result.0.dst").String()
	fmt.Println(result)
}
