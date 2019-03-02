package pdd

import (
	"encoding/json"
)

type GoodsAPI struct {
	Context *Context
}

func NewGoodsAPI(cfg *Config) *GoodsAPI {
	return &GoodsAPI{Context: NewContext(cfg)}
}

func newGoodsAPIWithContext(c *Context) *GoodsAPI {
	return &GoodsAPI{Context: c}
}

type Category struct {
    Level int `json:"level"` // 层级，1-一级，2-二级，3-三级，4-四级
	CatId int `json:"cat_id"` //
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

type Opt struct {
	Level       int    `json:"level"`         // 层级，1-一级，2-二级，3-三级，4-四级
	ParentOptID int    `json:"parent_opt_id"` // id 所属父 Id，其中，parent_id = 0 时为顶级节点
	OptName     string `json:"opt_name"`      // 商品标签名
	OptID       int    `json:"opt_id"`        // 商品标签ID
}

func (g *GoodsAPI) GoodsOptGet(parentOptId int) (res []*Opt, err error) {
	params := NewParamsWithType(GoodsOptGet)
	params.Set("parent_opt_id", parentOptId)

	r, err := Call(g.Context, params)
	if err != nil {
		return
	}

	bytes, err := GetResponseBytes(r, "goods_opt_get_response", "goods_opt_list")
	json.Unmarshal(bytes, &res)
	return
}