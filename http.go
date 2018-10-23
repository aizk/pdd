package pdd

import (
	"github.com/parnurzeal/gorequest"
)

var Client *gorequest.SuperAgent

func init() {
	Client = gorequest.New()
}

// 简单封装一下 Post 函数
func Post(data interface{}, r interface{}) (err error) {
	_, r, errors := Client.Post("").
		Type("json").
		Send(data).
		EndStruct(r)
	for _, e := range errors {
		if e != nil {
			return e
		}
	}
	return
}