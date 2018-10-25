package pdd

import (
	"strconv"
	"time"
	"sort"
	"net/url"
	"encoding/json"
)

// pdd params
type Params map[string]interface{}

func NewParams() Params {
	p := make(Params)
	return p
}

func NewParamsWithType(_type MethodType, params ...Params) Params {
	p := make(Params)
	p["type"] = _type
	p["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (p Params) Sign(context *Context) {
	p["client_id"] = context.ClientId
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := context.ClientSecret
	for _, key := range keys {
		signStr += key + getString(p[key])
	}
	signStr += context.ClientSecret
	p["sign"] = createSign(signStr)
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

func (p Params) GetQuery() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, getString(v))
	}
	return u.Encode()
}

func getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case MethodType:
		return string(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
	return ""
}