package pdd

import (
	"encoding/json"
)

type DDK struct {
	Context *Context
}

func NewDDK(c *Config) *DDK {
	return &DDK{Context: NewContext(c.ClientId, c.ClientSecret)}
}

func NewDDKWithContext(c *Context) *DDK {
	return &DDK{Context: c}
}

type RedPackageUrl struct {
	Url                      string `json:"url"`
	ShortUrl                 string `json:"short_url"`
	MobileUrl                string `json:"mobile_url"`
	MobileShortUrl           string `json:"mobile_short_url"`
	MultiGroupUrl            string `json:"multi_group_url"`              // 红包推广多人团链接
	MultiGroupShortUrl       string `json:"multi_group_short_url"`        // 红包推广多人团短链接
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`       // 红包推广多人团移动链接
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"` // 红包推广多人团移动短链接
}

func (d *DDK) RPPromUrlGenerate(pids []string, generateWeappWebview bool, notMustparams ...Params) (res []*RedPackageUrl, err error) {
	params := NewParamsWithType(DDK_RPPromUrlGenerate, notMustparams...)
	params.Set("p_id_list", TransformPids(pids))
	params.Set("generate_weapp_webview", generateWeappWebview)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "rp_promotion_url_generate_response", "url_list")
	err = json.Unmarshal(bytes, &res)
	return
}
