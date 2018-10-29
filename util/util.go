package util

import (
	"github.com/bitly/go-simplejson"
	"fmt"
	. "github.com/liunian1004/pdd/context"
)

// the basic call method
func Call(context *Context, params Params) (r []byte, err error) {
	params.Sign(context)
	r, err = Post(context, params.GetQuery())
	return
}

// 兼容多层 Key 读取
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

func GetResponseArrayIndexBytes(data []byte, index int, keys ...string) (b []byte, err error) {
	js, err := simplejson.NewJson(data)
	if err != nil {
		return
	}
	for _, key := range keys {
		js = js.Get(key)
	}

	js = js.GetIndex(index)

	b, err = js.Encode()
	return
}

func TransformPids(pids []string) (r string) {
	if len(pids) == 0 {
		return ""
	}
	if len(pids) == 1 {
		return fmt.Sprintf(`["%s"]`, pids[0])
	}
	r += `["`
	for k, v := range pids {
		if k == 0 {
			r += v
		} else {
			r += "," + v
		}
	}
	r += `"]`
	return
}
