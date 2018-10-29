package ddk

import (
	"encoding/json"
	"fmt"
	//. "github.com/liunian1004/pdd"
	. "github.com/liunian1004/pdd/context"
	. "github.com/liunian1004/pdd/util"
)

type DDK struct {
	Context *Context
}

//func NewDDK(p *Pdd) *DDK {
//	return &DDK{Context: p.Context}
//}

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

type OrderListResponse struct {
	OrderList []*Order `json:"order_list"`
	Total     int      `json:"total"`
}

type Order struct {
	OrderSn               string `json:"order_sn"`
	GoodsId               int    `json:"goods_id"`
	GoodsName             string `json:"goods_name"`
	GoodsThumbnailUrl     string `json:"goods_thumbnail_url"`
	GoodsQuantity         int    `json:"goods_quantity"`
	GoodsPrice            int    `json:"goods_price"`
	OrderAmount           int    `json:"order_amount"`
	OrderCreateTime       int    `json:"order_create_time"`
	OrderSettleTime       int    `json:"order_settle_time"`
	OrderVerifyTime       int    `json:"order_verify_time"`
	OrderReceiveTime      int    `json:"order_receive_time"`
	OrderPayTime          int    `json:"order_pay_time"`
	PromotionRate         int    `json:"promotion_rate"`
	PromotionAmount       int    `json:"promotion_amount"`
	BatchNo               string `json:"batch_no"`
	OrderStatus           int    `json:"order_status"`
	OrderStatusDesc       string `json:"order_status_desc"`
	VerifyTime            int    `json:"verify_time"`
	OrderGroupSuccessTime int    `json:"order_group_success_time"`
	OrderModifyAt         int    `json:"order_modify_at"`
	Status                int    `json:"status"`
	Type                  int    `json:"type"`
	GroupId               int    `json:"group_id"`
	AuthDuoId             int    `json:"auth_duo_id"`
	ZsDuoId               int    `json:"zs_duo_id"`
	CustomParameters      string `json:"custom_parameters"`
	Pid                   string `json:"pid"`
	ShowStatus            int    `json:"show_status"`
	MatchChannel          int    `json:"match_channel"`
	DuoCouponAmount       int    `json:"duo_coupon_amount"`
}

func (d *DDK) OrderListIncrementGet(startUpdateTime, endUpdateTime int, notMustparams ...Params) (res *OrderListResponse, err error) {
	params := NewParamsWithType(DDK_OrderListIncrementGet, notMustparams...)
	params.Set("start_update_time", startUpdateTime)
	params.Set("end_update_time", endUpdateTime)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "order_list_get_response")
	res = new(OrderListResponse)
	err = json.Unmarshal(bytes, res)
	return
}

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
	MinGroupPrice        int      `json:"min_group_price"`
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
	CouponDiscount       int      `json:"coupon_discount"`
	CouponTotalQuantity  int      `json:"coupon_total_quantity"`
	CouponRemainQuantity int      `json:"coupon_remain_quantity"`
	CouponStartTime      int      `json:"coupon_start_time"`
	CouponEndTime        int      `json:"coupon_end_time"`
	PromotionRate        int      `json:"promotion_rate"`
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

type PIdQueryListResponse struct {
	PIDList    []*PIdQuery `json:"p_id_list"`
	TotalCount int         `json:"total_count"`
}

type PIdQuery struct {
	PId        string `json:"p_id"`
	CreateTime string `json:"create_time"`
}

func (d *DDK) GoodsPidQuery(notMustparams ...Params) (res *PIdQueryListResponse, err error) {
	params := NewParamsWithType(DDK_GoodsPidQuery, notMustparams...)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "p_id_query_response")
	res = new(PIdQueryListResponse)
	err = json.Unmarshal(bytes, res)
	return
}

type PIdGenerateListResponse struct {
	PIDList    []*PIdGenerate `json:"p_id_list"`
	TotalCount int            `json:"total_count"`
}

type PIdGenerate struct {
	PId     string `json:"p_id"`
	PIdName string `json:"p_id_name"`
}

func (d *DDK) GoodsPidGenerate(number int, notMustparams ...Params) (res *PIdGenerateListResponse, err error) {
	params := NewParamsWithType(DDK_GoodsPidGenerate, notMustparams...)
	params.Set("number", number)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "p_id_generate_response")
	res = new(PIdGenerateListResponse)
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

func (d *DDK) ThemeListGet(page, pageSize int) (res *ThemeListResponse, err error) {
	params := NewParamsWithType(DDK_ThemeListGet)
	params.Set("page", page)
	params.Set("page_size", pageSize)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "theme_list_get_response")
	res = new(ThemeListResponse)
	err = json.Unmarshal(bytes, res)
	return
}
