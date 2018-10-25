package pdd

import (
	"encoding/json"
)

type ThemeListResponse struct {
	ThemeList []*Theme `json:"theme_list"`
	Total     int      `json:"total"`
}

type Theme struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	GoodsNum int    `json:"goods_num"` // 商品数量
}

func (p *Pdd) GetThemeList(page, pageSize int) (res *ThemeListResponse, err error) {
	params := NewParamsWithType(ThemeListGet)
	params.Set("page", page)
	params.Set("page_size", pageSize)
	params.Sign(p)
	r, err := GetBytesPost(params.GetQuery())
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r,"theme_list_get_response")
	res = new(ThemeListResponse)
	err = json.Unmarshal(bytes, res)
	return
}
