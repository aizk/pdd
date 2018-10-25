package pdd

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

// the basic call method
func Call(context *Context, params Params) (r []byte, err error) {
	params.Sign(context)
	r, err = Post(params.GetQuery())
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

func SetDebug(d bool) {
	Debug = d
}

// 设置请求失败重试次数
func SetRetryTimes(i int) {
	Retry = i
}