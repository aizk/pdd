package pdd

import (
	"encoding/json"
)

type OrderListResponse struct {
	OrderList  []*Order `json:"order_list"`
	TotalCount int      `json:"total_count"`
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
	OrderSettleTime       int    `json:"order_settle_time"`  // 结算时间
	OrderVerifyTime       int    `json:"order_verify_time"`  // 审核时间
	OrderReceiveTime      int    `json:"order_receive_time"` // 收货时间
	OrderPayTime          int    `json:"order_pay_time"`
	PromotionRate         int    `json:"promotion_rate"`
	PromotionAmount       int    `json:"promotion_amount"`
	BatchNo               string `json:"batch_no"`
	OrderStatus           int    `json:"order_status"`
	OrderStatusDesc       string `json:"order_status_desc"`
	VerifyTime            int    `json:"verify_time"`
	OrderGroupSuccessTime int    `json:"order_group_success_time"`
	OrderModifyAt         int    `json:"order_modify_at"`
	Type                  int    `json:"type"`
	GroupId               int    `json:"group_id"`
	AuthDuoId             int    `json:"auth_duo_id"`
	ZsDuoId               int    `json:"zs_duo_id"`
	CustomParameters      string `json:"custom_parameters"`
	Pid                   string `json:"pid"`
	PId                   string `json:"p_id"`
	MatchChannel          int    `json:"match_channel"`
	DuoCouponAmount       int    `json:"duo_coupon_amount"`
}

// page_size default 100
// page default  1
func (d *DDK) OrderListIncrementGet(startUpdateTime, endUpdateTime int64, notMustparams ...Params) (res *OrderListResponse, err error) {
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

	for _, ord := range res.OrderList {
		if ord.Pid == "" {
			ord.Pid = ord.PId
		}
	}
	return
}

func (d *DDK) ColorOrderIncrementGet(startUpdateTime, endUpdateTime int64, notMustparams ...Params) (res *OrderListResponse, err error) {
	params := NewParamsWithType(DDK_ColorOrderIncrementGet, notMustparams...)
	params.Set("start_update_time", startUpdateTime)
	params.Set("end_update_time", endUpdateTime)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "order_list_get_response")
	res = new(OrderListResponse)
	err = json.Unmarshal(bytes, res)

	for _, ord := range res.OrderList {
		if ord.Pid == "" {
			ord.Pid = ord.PId
		}
	}
	return
}
