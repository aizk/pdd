package pdd

import "github.com/bitly/go-simplejson"

// 兼容多层 Key 调用
func GetResponseBytes(data []byte, keys ...string) (b []byte, err error) {
	js, err := simplejson.NewJson(data)
	if err != nil {
		return
	}
	for _, key := range keys {
		js = js.Get(key)
	}
	b, err = js.Encode()
	return
}
