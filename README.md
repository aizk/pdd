# README

本项目为 `go` 语言实现的拼多多开放平台 SDK，调用方式简单粗暴。

对于未实现的接口，欢迎 pr 交流。

`go get github.com/liunian1004/pdd`

```go
import github.com/liunian1004/pdd

p := pdd.NewPdd(&pdd.Config{
    ClientId: "your client id",
    ClientSecret: "your client secret",
})

// 初始化多多客相关 API 调用
d := NewDDK(p)

// or
d := ddk.NewDDK(&pdd.Config{
    ClientId: "your client id",
    ClientSecret: "your client secret",
})

// 获取主题列表
r, err := d.ThemeListGet(1, 20)

// 设置接口调用失败重试次数
pdd.SetRetryTimes(3)

// 开启 Debug 模式
pdd.Debug(true)
```
## 非必须参数

通过自定义 Params 定制非必须参数，在函数的最后一个参数传入 Params 对象。

```go
d := NewDDK(p)

params := pdd.NewParams()
// 设置非必传参数
params.Set("custom_parameters", "test")
params.Set("generate_short_url", true)
s, err := d.RPPromUrlGenerate([]string{"test"}, true, params)
```

## Todo

- [] 实现 ddk 多多客 API 相关接口
    - [x] `OrderListIncrementGet()` pdd.ddk.order.list.increment.get 最后更新时间段增量同步推广订单信息
    - [x] `GoodsDetail()` pdd.ddk.goods.detail 多多进宝商品详情查询
    - [x] `GoodsSearch()` pdd.ddk.goods.search 多多进宝商品查询
    - [x] `GoodsPidQuery()` pdd.ddk.goods.pid.query 查询已经生成的推广位信息
    - [x] `GoodsPidGenerate()` pdd.ddk.goods.pid.generate 创建多多进宝推广位
    - [x] `GoodsPromotionUrlGenerate()` pdd.ddk.goods.promotion.url.generate 多多进宝推广链接生成
    - [x] `RPPromUrlGenerate()` pdd.ddk.rp.prom.url.generate 生成红包推广链接 （**需要对应权限**）
    - [] pdd.ddk.cms.prom.url.generate 生成商城-频道推广链接
    - [x] `ThemeListGet()` pdd.ddk.theme.list.get 多多进宝主题列表查询
    - [] pdd.ddk.theme.goods.search 多多进宝主题商品查询
    - [] pdd.ddk.theme.prom.url.generate 多多进宝主题推广链接生成
    - [] pdd.ddk.app.new.bill.list.get 多多客拉新账单
    - [] pdd.ddk.direct.goods.query 定向推广商品查询接口
    - [] pdd.ddk.goods.zs.unit.url.gen 多多进宝转链接口
    - [] pdd.ddk.zs.unit.goods.query 查询招商推广计划商品
    - [] pdd.ddk.weapp.qrcode.url.gen 多多客生成单品推广小程序二维码 url
    - [] pdd.ddk.goods.basic.info.get 获取商品基本信息接口
    - [] pdd.ddk.goods.recommend.get 运营频道商品查询 API
    - [] pdd.ddk.order.detail.get 查询订单详情
    - [] pdd.ddk.mall.goods.list.get 查询店铺商品
    - [] pdd.ddk.mall.url.gen 多多客生成店铺推广链接 API
    - [] pdd.ddk.lottery.url.gen 多多客生成转盘抽免单 url （**需要对应权限**）
    - [] pdd.ddk.lottery.new.list.get 多多客查询转盘拉新订单列表
    - [] pdd.ddk.resource.url.gen 生成多多进宝频道推广
    - [] pdd.ddk.merchant.list.get 多多客查店铺列表接口

## 多多客工具 API

提供给第三方开发者为多多进宝推广者提供第三方工具的 API。

- pdd.ddk.oauth.goods.pid.generate（多多进宝推广位创建接口）
- pdd.ddk.oauth.goods.pid.query（多多客已生成推广位信息查询）
- pdd.ddk.oauth.goods.prom.url.generate（生成多多进宝推广链接）
- pdd.ddk.oauth.order.list.increment.get（按照更新时间段增量同步推广订单信息）
- pdd.ddk.oauth.order.list.range.get（按照时间段获取多多进宝推广订单信息）

## 注意事项

1. 多多客工具API需要多多客授权才能使用，授权服务URL为 http://jinbao.pinduoduo.com/open.html ，根据接入指南提示组装合法的 URL，示例：http://jinbao.pinduoduo.com/open.html?client_id=89daeb71b2e546318fbb53cd04d0329c&redirect_uri=www.pinduoduo.com&response_type=code，之后引导多多客授权获取code即可。
2. code、access_token、refresh_token 失效时间与商家授权一致，分别为10分钟、24小时、30天，刷新token后分别延长access_token和refresh_token的失效时间，延长时间等同于授权时长，获取token教程详见：http://open.pinduoduo.com/#/document
3. 调用多多客工具 API 时，必须入参 access_token 方可正常调用
4. 多多客工具 API 可以通过创建多多客联盟应用获取权限