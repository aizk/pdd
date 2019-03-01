package pdd

import "encoding/json"

type CMSPromUrl struct {
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

type ChannelType int

const (
	PackageMail19  ChannelType = 0 // 1.9 包邮
	TodayHotStyle  ChannelType = 1 // 今日爆款
	BrandClearance ChannelType = 2 // 品牌清仓
	PCTerminalMall ChannelType = 4 // PC端专属商城
	RaiseBabyCash  ChannelType = 5 // 养宝宝兑现金
)

func (d *DDK) CMSPromUrlGen(pids []string, channelType ChannelType, generateWeappWebview, weAppWebViewShortUrl, weAppWebViewUrl bool, notMustparams ...Params) (res *CMSPromUrl, err error) {
	res = new(CMSPromUrl)
	params := NewParamsWithType(DDK_CMSPromUrlGenerate, notMustparams...)
	params.Set("p_id_list", TransformPids(pids))
	params.Set("channel_type", int(channelType))
	params.Set("generate_weapp_webview", generateWeappWebview)
	params.Set("we_app_web_view_url", weAppWebViewShortUrl)
	params.Set("we_app_web_view_short_url", weAppWebViewUrl)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "cms_promotion_url_generate_response", "url_list")
	err = json.Unmarshal(bytes, &res)
	return
}
