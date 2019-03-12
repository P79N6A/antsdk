package trade

import (
  "github.com/solarhell/antsdk/api"
)

type AlipayTradeCreateResponse struct {
  api.AlipayResponse
  OutTradeNo  string `json:"out_trade_no"`  // 商户订单号
  TradeNo     string `json:"trade_no"`      // 支付宝交易号
}
