package pdd

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
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
	CreateAt             int      `json:"create_at"`               // 创建时间
	GoodsId              int      `json:"goods_id"`                // 商品id
	GoodsName            string   `json:"goods_name"`              // 商品名称
	GoodsDesc            string   `json:"goods_desc"`              // 商品描述
	GoodsThumbnailUrl    string   `json:"goods_thumbnail_url"`     // 商品缩略图
	GoodsImageUrl        string   `json:"goods_image_url"`         // 商品主图
	GoodsGalleryUrls     []string `json:"goods_gallery_urls"`      // 商品轮播图
	SoldQuantity         int      `json:"sold_quantity"`           // 已售卖件数
	MinGroupPrice        int      `json:"min_group_price"`         // 最小拼团价（单位为分） 最低价 sku 的拼团价，单位为分
	MinNormalPrice       int      `json:"min_normal_price"`        // 最小单买价格（单位为分）
	MallId               int      `json:"mall_id"`                 // 店铺 id
	MallName             string   `json:"mall_name"`               // 店铺名字
	MerchantType         int      `json:"merchant_type"`           // 店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
	MallCps              int      `json:"mall_cps"`                // 该商品所在店铺是否参与全店推广，0：否，1：是
	MallRate             int      `json:"mall_rate"`               // 全店推广佣金比率？ 千分比
	CategoryId           int      `json:"category_id"`             // 商品类目 id
	CategoryName         string   `json:"category_name"`           // 商品类目名称
	OptId                int      `json:"opt_id"`                  // 商品标签 id，使用 pdd.goods.opts.get 接口获取
	OptName              string   `json:"opt_name"`                // 商品标签名
	OptIds               []int    `json:"opt_ids"`                 // 商品标签id
	CatIds               []int    `json:"cat_ids"`                 // 商品类目 id 列表
	CatId                int      `json:"cat_id"`                  // 可能为 null
	HasCoupon            bool     `json:"has_coupon"`              // 是否有优惠券
	CouponMinOrderAmount int      `json:"coupon_min_order_amount"` // 优惠券门槛价格，单位为分
	CouponDiscount       int      `json:"coupon_discount"`         // 优惠券面额，单位为分
	CouponTotalQuantity  int      `json:"coupon_total_quantity"`   // 优惠券总数量
	CouponRemainQuantity int      `json:"coupon_remain_quantity"`  // 优惠券剩余数量
	CouponStartTime      int      `json:"coupon_start_time"`       // 优惠券生效时间，UNIX时间戳
	CouponEndTime        int      `json:"coupon_end_time"`         // 优惠券失效时间，UNIX时间戳
	PromotionRate        int      `json:"promotion_rate"`          // 佣金比例，千分比 300
	GoodsEvalScore       float32  `json:"goods_eval_score"`        // 商品评价分 4.81
	GoodsEvalCount       int      `json:"goods_eval_count"`        // 商品评价数量
	AvgDesc              int      `json:"avg_desc"`                // 描述评分
	AvgLgst              int      `json:"avg_lgst"`                // 物流评分
	AvgServ              int      `json:"avg_serv"`                // 服务评分
	DescPct              float32  `json:"desc_pct"`                // 描述分击败同类店铺百分比
	LgstPct              float32  `json:"lgst_pct"`                // 物流分击败同类店铺百分比
	ServPct              float32  `json:"serv_pct"`                // 服务分击败同类店铺百分比
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

func GoodsToGoodsModel(g *Goods) (r *GoodsModel) {
	r = new(GoodsModel)
	copier.Copy(r, g)
	r.GoodsGalleryUrls = g.MarshalGoodsGalleryUrls()
	r.OptIds = g.MarshalOptIds()
	r.CatIds = g.MarshalCatIds()
	return
}

func GoodsListToGoodsModelList(gs []*Goods) (gms []*GoodsModel) {
	for _, goods := range gs {
		gms = append(gms, GoodsToGoodsModel(goods))
	}
	return
}

// search interface with more params
func (d *DDK) GoodsSearch(notMustparams ...Params) (res *GoodsListResponse, err error) {
	params := NewParamsWithType(DDK_GoodsSearch, notMustparams...)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_search_response")
	res = new(GoodsListResponse)
	err = json.Unmarshal(bytes, res)
	return
}

func (d *DDK) GetExistGoods(ids []int) (res []*Goods, err error) {
	length := len(ids)
	start := 0
	pageSize := 100

	if length < pageSize {
		res, err = d.getGoodsByIds(ids[start:length])
		if err != nil {
			return
		}
		return
	}

	for length > pageSize {
		r, err1 := d.getGoodsByIds(ids[start:pageSize])
		if err1 != nil {
			err = err1
			return
		}
		res = append(res, r...)
		start += 100
		pageSize += 100
	}

	r, err := d.getGoodsByIds(ids[start:length])
	if err != nil {
		return
	}
	res = append(res, r...)

	return
}

func (d *DDK) getGoodsByIds(ids []int) (rsp []*Goods, err error) {
	idstr, _ := json.Marshal(ids)
	params := NewParams()
	params.Set("goods_id_list", string(idstr))
	r, err := d.GoodsSearch(params)
	rsp = r.OrderList
	return
}

// detail 接口查询返回的商品详情比搜索返回的多出一些信息：
// goods_gallery_urls
// goods_desc
// sold_quantity
func (d *DDK) GoodsDetail(goodsId int) (res *Goods, err error) {
	res = new(Goods)
	params := NewParamsWithType(DDK_GoodsDetail)
	params.Set("goods_id_list", fmt.Sprintf("[%d]", goodsId))

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "goods_detail_response", "goods_details")
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
func (d *DDK) GoodsPromotionUrlGenerate(pid string, goodsId int, notMustparams ...Params) (res *GoodsPromotionUrl, err error) {
	params := NewParamsWithType(DDK_GoodsPromotionUrlGenerate, notMustparams...)
	params.Set("p_id", pid)
	params.Set("goods_id_list", fmt.Sprintf("[%d]", goodsId))

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

//TopGoods 爆款排行商品
type TopGoods struct {
	GoodsID              int64    `json:"goods_id"`                //商品id
	GoodsName            string   `json:"goods_name"`              //商品名称
	GoodsDesc            string   `json:"goods_desc"`              //商品描述
	GoodsThumbnailURL    string   `json:"goods_thumbnail_url"`     //商品缩略图
	GoodsImageURL        string   `json:"goods_image_url"`         //商品主图
	GoodsGalleryUrls     []string `json:"goods_gallery_urls"`      //商品轮播图
	SoldQuantity         int64    `json:"sold_quantity"`           //已售卖件数
	MinGroupPrice        int64    `json:"min_group_price"`         //最小拼团价（单位为分）
	MinNormalPrice       int64    `json:"min_normal_price"`        //最小单买价格（单位为分）
	MallName             string   `json:"mall_name"`               //店铺名字
	MerchantType         int      `json:"merchant_type"`           //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
	CategoryID           int64    `json:"category_id"`             //商品类目ID，使用pdd.goods.cats.get接口获取
	CategoryName         string   `json:"category_name"`           //商品类目名
	OptID                int64    `json:"opt_id"`                  //商品标签ID，使用pdd.goods.opts.get接口获取
	OptName              string   `json:"opt_name"`                //商品标签名
	OptIds               []int64  `json:"opt_ids"`                 //商品标签id
	CatIds               []int64  `json:"cat_ids"`                 //商品类目id
	MallCps              int      `json:"mall_cps"`                //该商品所在店铺是否参与全店推广，0：否，1：是
	HasCoupon            bool     `json:"has_coupon"`              //商品是否有优惠券 true-有，false-没有
	CouponMinOrderAmount int64    `json:"coupon_min_order_amount"` //优惠券门槛价格，单位为分
	CouponDiscount       int64    `json:"coupon_discount"`         //优惠券面额，单位为分
	CouponTotalQuantity  int64    `json:"coupon_total_quantity"`   //优惠券总数量
	CouponRemainQuantity int64    `json:"coupon_remain_quantity"`  //优惠券剩余数量
	CouponStartTime      int64    `json:"coupon_start_time"`       //优惠券生效时间，UNIX时间戳
	CouponEndTime        int64    `json:"coupon_end_time"`         //优惠券失效时间，UNIX时间戳
	PromotionRate        int64    `json:"promotion_rate"`          //佣金比例，千分比
	GoodsEvalScore       float64  `json:"goods_eval_score"`        //商品评价分
	GoodsEvalCount       int64    `json:"goods_eval_count"`        //商品评价数量
	AvgDesc              int64    `json:"avg_desc"`                //描述评分
	AvgLgst              int64    `json:"avg_lgst"`                //物流评分
	AvgServ              int64    `json:"avg_serv"`                //服务评分
	DescPct              float64  `json:"desc_pct"`                //描述分击败同类店铺百分比
	LgstPct              float64  `json:"lgst_pct"`                //物流分击败同类店铺百分比
	ServPct              float64  `json:"serv_pct"`                //服务分击败同类店铺百分比
}

//TopGoodsList 爆款排行商品列表
type TopGoodsList struct {
	Goods []TopGoods `json:"list"`
	Total int64      `json:"total"`
}

//TopGoodsListQuery 多多客获取爆款排行商品接口
func (d *DDK) TopGoodsListQuery(notMustparams ...Params) (res TopGoodsList, err error) {
	params := NewParamsWithType(DDK_TopGoodsListQuery, notMustparams...)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "top_goods_list_get_response")
	err = json.Unmarshal(bytes, &res)
	return
}
