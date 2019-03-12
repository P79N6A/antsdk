package trade

import (
  "github.com/solarhell/antsdk/api"
  "github.com/solarhell/antsdk/utils"
)

// 手机wap支付，add by solarhell
type AlipayTradeWapPayRequest struct {
  api.IAlipayRequest
  ProdCode      string                            `json:"prod_code"`
  NotifyUrl     string                            `json:"notify_url"`
  ReturnUrl     string                            `json:"return_url"`
  BizContent    AlipayTradeWapPayRequestBizContent   `json:"biz_content"`
}

type AlipayTradeWapPayRequestBizContent struct {
  OutTradeNo            string          `json:"out_trade_no"`           // 商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
  AuthCode              string          `json:"auth_code"`              // 支付授权码
  Subject               string          `json:"subject"`                // 订单标题
  SellerId              string          `json:"seller_id"`              // 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
  TotalAmount           float64         `json:"total_amount"`           // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入【可打折金额】和【不可打折金额】，该参数可以不用传入； 如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
  Body                  string          `json:"body"`                   // 订单描述
}

func (this *AlipayTradeWapPayRequest) GetTerminalType() string {
  return ""
}

func (this *AlipayTradeWapPayRequest) GetTerminalInfo() string {
  return ""
}


func (this *AlipayTradeWapPayRequest) GetApiMethodName() string {
  return "alipay.trade.wap.pay"
}

func (this *AlipayTradeWapPayRequest) GetApiVersion() string {
  return "1.0"
}


func (this *AlipayTradeWapPayRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayTradeWapPayRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayTradeWapPayRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayTradeWapPayRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayTradeWapPayRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
