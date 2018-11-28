# 多多客工具 API

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