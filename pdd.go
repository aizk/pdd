package pdd

import (
	"encoding/json"
)

type Pdd struct {
	Context *Context
}

type Config struct {
	ClientId     string
	ClientSecret string
}

func NewPdd(c *Config) *Pdd {
	return &Pdd{Context: NewContext(c.ClientId, c.ClientSecret)}
}

func (p *Pdd) GetDDK() *DDK {
	return NewDDKWithContext(p.Context)
}

func (p *Pdd) GetGoodsAPI() *GoodsAPI {
	return NewGoodsAPIWithContext(p.Context)
}

//GoodsOpt 商品类别
type GoodsOpt struct {
	Level       int    `json:"level"`         //层级，1-一级，2-二级，3-三级，4-四级
	ParentOptID int    `json:"parent_opt_id"` //id所属父ID，其中，parent_id=0时为顶级节点
	OptName     string `json:"opt_name"`      //商品标签名
	OptID       int    `json:"opt_id"`        //商品标签ID
}

//GoodsOptList 商品类别列表
type GoodsOptList struct {
	GoodsOptList []*GoodsOpt `json:"goods_opt_list"`
}

//PddGoodsOptGet 查询商品标签列表
func (d *DDK) PddGoodsOptGet(parentOptID int) (res *GoodsOptList, err error) {
	params := NewParamsWithType(GoodsOptGet)
	params.Set("parent_opt_id", parentOptID)
	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_opt_get_response")
	res = new(GoodsOptList)
	err = json.Unmarshal(bytes, res)
	return
}

//GoodsCat  商品标准类目
type GoodsCat struct {
	Level       int    `json:"level"`         //层级，1-一级，2-二级，3-三级，4-四级
	ParentCatID int    `json:"parent_cat_id"` //id所属父ID，其中，parent_id=0时为顶级节点
	CatName     string `json:"cat_name"`      //商品标准类目名
	CatID       int    `json:"cat_id"`        //商品标准类目ID
}

//GoodsCatsList 商品标准类目列表
type GoodsCatsList struct {
	CatsList []*GoodsCat `json:"goods_cats_list"`
}

//PddGoodsCatGet 查询商品标准类目列表
func (d *DDK) PddGoodsCatGet(ParentCatID int) (res *GoodsCatsList, err error) {
	params := NewParamsWithType(GoodsCatsGet)
	params.Set("parent_cat_id", ParentCatID)
	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_cats_get_response")
	res = new(GoodsCatsList)
	err = json.Unmarshal(bytes, res)
	return
}
