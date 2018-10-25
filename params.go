package pdd

import (
	"strconv"
	"time"
	"sort"
	"net/url"
)

type MethodType string

const (
	ThemeListGet MethodType = "pdd.ddk.theme.list.get"
)

// pdd params
type Params map[string]interface{}

func NewParams() Params {
	p := make(Params)
	p["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	return p
}

func NewParamsWithType(_type MethodType) Params {
	p := make(Params)
	p["type"] = _type
	p["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	return p
}

func (p Params) Sign(pdd *Pdd) {
	p["client_id"] = pdd.Context.ClientId
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := pdd.Context.ClientSecret
	for _, key := range keys {
		signStr += key + getString(p[key])
	}
	signStr += pdd.Context.ClientSecret
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
	}
	return ""
}