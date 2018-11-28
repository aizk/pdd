package ddk

import (
	"encoding/json"
	. "github.com/liunian1004/pdd/util"
	"fmt"
)

type Range struct {
	RangeId   int `json:"range_id"`
	RangeFrom int `json:"range_from"`
	RangeTo   int `json:"range_to"`
}

type GoodsListResponse struct {
	OrderList  []*Goods `json:"goods_list"`
	TotalCount int      `json:"total_count"`
}

type Goods struct {
	CreateAt             int      `json:"create_at"`
	GoodsId              int      `json:"goods_id"`
	GoodsName            string   `json:"goods_name"`
	GoodsDesc            string   `json:"goods_desc"`
	GoodsThumbnailUrl    string   `json:"goods_thumbnail_url"`
	GoodsImageUrl        string   `json:"goods_image_url"`
	GoodsGalleryUrls     []string `json:"goods_gallery_urls"`
	SoldQuantity         int      `json:"sold_quantity"`
	MinGroupPrice        int      `json:"min_group_price"` // 最低价sku的拼团价，单位为分
	MinNormalPrice       int      `json:"min_normal_price"`
	MallId               int      `json:"mall_id"`
	MallName             string   `json:"mall_name"`
	MerchantType         int      `json:"merchant_type"`
	MallCps              int      `json:"mall_cps"`
	CategoryId           int      `json:"category_id"`
	CategoryName         string   `json:"category_name"`
	OptId                int      `json:"opt_id"`
	OptName              string   `json:"opt_name"`
	OptIds               []int    `json:"opt_ids"`
	CatIds               []int    `json:"cat_ids"`
	CatId                int      `json:"cat_id"`
	HasCoupon            bool     `json:"has_coupon"`
	CouponMinOrderAmount int      `json:"coupon_min_order_amount"`
	CouponDiscount       int      `json:"coupon_discount"` // 优惠券面额，单位为分
	CouponTotalQuantity  int      `json:"coupon_total_quantity"`
	CouponRemainQuantity int      `json:"coupon_remain_quantity"`
	CouponStartTime      int      `json:"coupon_start_time"`
	CouponEndTime        int      `json:"coupon_end_time"`
	PromotionRate        int      `json:"promotion_rate"` // 佣金比例，千分比
	GoodsEvalScore       float32  `json:"goods_eval_score"`
	GoodsEvalCount       int      `json:"goods_eval_count"`
	AvgDesc              int      `json:"avg_desc"`
	AvgLgst              int      `json:"avg_lgst"`
	AvgServ              int      `json:"avg_serv"`
	DescPct              float32  `json:"desc_pct"`
	LgstPct              float32  `json:"lgst_pct"`
	ServPct              float32  `json:"serv_pct"`
}

func (g *Goods) MarshalGoodsGalleryUrls() string {
	r, _ := json.Marshal(g.GoodsGalleryUrls)
	return string(r)
}

func (g *Goods) MarshalOptIds() string {
	r, _ := json.Marshal(g.OptIds)
	return string(r)
}

func (g *Goods) MarshalCatIds() string {
	r, _ := json.Marshal(g.CatIds)
	return string(r)
}

// GoodsModel for gorm
type GoodsModel struct {
	CreateAt             int     `json:"create_at"`
	GoodsId              int     `json:"goods_id" gorm:"type:bigint;not null;unique"`
	GoodsName            string  `json:"goods_name"`
	GoodsDesc            string  `json:"goods_desc" gorm:"type:text"`
	GoodsThumbnailUrl    string  `json:"goods_thumbnail_url"`
	GoodsImageUrl        string  `json:"goods_image_url"`
	GoodsGalleryUrls     string  `json:"goods_gallery_urls" gorm:"type:text"`
	SoldQuantity         int     `json:"sold_quantity"`
	MinGroupPrice        int     `json:"min_group_price"`
	MinNormalPrice       int     `json:"min_normal_price"`
	MallId               int     `json:"mall_id"`
	MallName             string  `json:"mall_name"`
	MerchantType         int     `json:"merchant_type"`
	MallCps              int     `json:"mall_cps"`
	CategoryId           int     `json:"category_id"`
	CategoryName         string  `json:"category_name"`
	OptId                int     `json:"opt_id"`
	OptName              string  `json:"opt_name"`
	OptIds               string  `json:"opt_ids"`
	CatIds               string  `json:"cat_ids"`
	CatId                int     `json:"cat_id"`
	HasCoupon            bool    `json:"has_coupon"`
	CouponMinOrderAmount int     `json:"coupon_min_order_amount"`
	CouponDiscount       int     `json:"coupon_discount"`
	CouponTotalQuantity  int     `json:"coupon_total_quantity"`
	CouponRemainQuantity int     `json:"coupon_remain_quantity"`
	CouponStartTime      int     `json:"coupon_start_time"`
	CouponEndTime        int     `json:"coupon_end_time"`
	PromotionRate        int     `json:"promotion_rate"`
	GoodsEvalScore       float32 `json:"goods_eval_score"`
	GoodsEvalCount       int     `json:"goods_eval_count"`
	AvgDesc              int     `json:"avg_desc"`
	AvgLgst              int     `json:"avg_lgst"`
	AvgServ              int     `json:"avg_serv"`
	DescPct              float32 `json:"desc_pct"`
	LgstPct              float32 `json:"lgst_pct"`
	ServPct              float32 `json:"serv_pct"`
}

// search interface with more params
func (d *DDK) GoodsSearch(sortType int, withCoupon bool, notMustparams ...Params) (res *GoodsListResponse, err error) {
	params := NewParamsWithType(DDK_GoodsSearch, notMustparams...)
	params.Set("sort_type", sortType)
	params.Set("with_coupon", withCoupon)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_search_response")
	res = new(GoodsListResponse)
	err = json.Unmarshal(bytes, res)
	return
}

// one goods search
func (d *DDK) GoodsDetail(goodsId int) (res *Goods, err error) {
	params := NewParamsWithType(DDK_GoodsDetail)
	params.Set("goods_id_list", fmt.Sprintf("[%d]", goodsId))

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "goods_detail_response", "goods_details")
	res = new(Goods)
	err = json.Unmarshal(bytes, res)
	return
}

type GoodsPromotionUrl struct {
	Url                  string `json:"url"`
	ShortUrl             string `json:"short_url"`
	MobileUrl            string `json:"mobile_url"`
	MobileShortUrl       string `json:"mobile_short_url"`
	WeAppWebViewUrl      string `json:"we_app_web_view_url"`
	WeAppWebViewShortUrl string `json:"we_app_web_view_short_url"`
}

// create promotion url
func (d *DDK) GoodsPromotionUrlGenerate(pId, goodsId string, notMustparams ...Params) (res *GoodsPromotionUrl, err error) {
	params := NewParamsWithType(DDK_GoodsPromotionUrlGenerate, notMustparams...)
	params.Set("p_id", pId)
	params.Set("goods_id_list", fmt.Sprintf("[%s]", goodsId))

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_promotion_url_generate_response", "goods_promotion_url_list")
	var urls []*GoodsPromotionUrl
	err = json.Unmarshal(bytes, &urls)
	res = urls[0]
	return
}
