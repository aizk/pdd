package pdd

type MethodType string

const (
	// 多多客
	DDK_ThemeListGet              MethodType = "pdd.ddk.theme.list.get"
	DDK_RPPromUrlGenerate         MethodType = "pdd.ddk.rp.prom.url.generate" // 生成红包推广链接
	DDK_OrderListIncrementGet     MethodType = "pdd.ddk.order.list.increment.get"
	DDK_ColorOrderIncrementGet    MethodType = "pdd.ddk.color.order.increment.get"
	DDK_GoodsDetail               MethodType = "pdd.ddk.goods.detail"
	DDK_GoodsSearch               MethodType = "pdd.ddk.goods.search"
	DDK_GoodsPidQuery             MethodType = "pdd.ddk.goods.pid.query"
	DDK_GoodsPidGenerate          MethodType = "pdd.ddk.goods.pid.generate"
	DDK_GoodsPromotionUrlGenerate MethodType = "pdd.ddk.goods.promotion.url.generate"
	DDK_TopGoodsListQuery         MethodType = "pdd.ddk.top.goods.list.query" //多多客获取爆款排行商品接口

	//商品 API
	GoodsCatsGet MethodType = "pdd.goods.cats.get" //商品标准类目接口
	GoodsOptGet  MethodType = "pdd.goods.opt.get"  //查询商品标签列表
)
