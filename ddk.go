package pdd

import (
	"encoding/json"
)

type DDK struct {
	Context *Context
}

func NewDDK(cfg *Config) *DDK {
	return &DDK{Context: NewContext(cfg)}
}

func newDDKWithContext(ctx *Context) *DDK {
	return &DDK{Context: ctx}
}

type RedPackageUrl struct {
	Url                      string `json:"url"`
	ShortUrl                 string `json:"short_url"`
	MobileUrl                string `json:"mobile_url"`
	MobileShortUrl           string `json:"mobile_short_url"`
	WeAppWebViewUrl          string `json:"we_app_web_view_url"`
	WeAppWebViewShortUrl     string `json:"we_app_web_view_short_url"`
	MultiGroupUrl            string `json:"multi_group_url"`              // 红包推广多人团链接
	MultiGroupShortUrl       string `json:"multi_group_short_url"`        // 红包推广多人团短链接
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       // 红包推广多人团移动链接
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` // 红包推广多人团移动短链接
}

func (d *DDK) RPPromUrlGenerate(pids []string, generateWeappWebview bool, notMustparams ...Params) (res *RedPackageUrl, err error) {
	res = new(RedPackageUrl)
	params := NewParamsWithType(DDK_RPPromUrlGenerate, notMustparams...)
	params.Set("p_id_list", TransformPids(pids))
	params.Set("generate_weapp_webview", generateWeappWebview)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "rp_promotion_url_generate_response", "url_list")
	err = json.Unmarshal(bytes, &res)
	return
}

type LotteryUrl struct {
	Url                      string `json:"url"`
	ShortUrl                 string `json:"short_url"`
	MobileUrl                string `json:"mobile_url"`
	MobileShortUrl           string `json:"mobile_short_url"`
	WeAppWebViewUrl          string `json:"we_app_web_view_url"`
	WeAppWebViewShortUrl     string `json:"we_app_web_view_short_url"`
	MultiGroupUrl            string `json:"multi_group_url"`              // 红包推广多人团链接
	MultiGroupShortUrl       string `json:"multi_group_short_url"`        // 红包推广多人团短链接
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       // 红包推广多人团移动链接
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` // 红包推广多人团移动短链接
}

// lottery custom_parameters
func (d *DDK) LotteryUrlGen(pids []string, generateWeappWebview bool, notMustparams ...Params) (res *LotteryUrl, err error) {
	res = new(LotteryUrl)
	params := NewParamsWithType(DDK_LotteryUrlGen, notMustparams...)
	params.Set("pid_list", TransformPids(pids))
	params.Set("generate_weapp_webview", generateWeappWebview)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0,"lottery_url_response", "url_list")
	err = json.Unmarshal(bytes, &res)
	return
}
