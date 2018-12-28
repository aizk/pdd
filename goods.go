package pdd

import (
	"encoding/json"
)

type GoodsAPI struct {
	Context *Context
}

func NewGoodsAPIWithContext(c *Context) *GoodsAPI {
	return &GoodsAPI{Context: c}
}

type Category struct {
    Level int `json:"level"`
	CatId int `json:"cat_id"`
    ParentCatId int `json:"parent_cat_id"`
    CatName string `json:"cat_name"`
}

func (g *GoodsAPI) GoodsCatGet(parentCatId int) (res []*Category, err error) {
	params := NewParamsWithType(GoodsCatsGet)
	params.Set("parent_cat_id", parentCatId)

	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_cats_get_response", "goods_cats_list")
	json.Unmarshal(bytes, &res)
	return
}