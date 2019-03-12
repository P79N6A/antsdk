package fund

import "github.com/solarhell/antsdk/api"

type AlipayFundTransToaccountTransferResponse struct {
	api.AlipayResponse
	OutBizNo string `json:"out_biz_no"` //商户转账唯一订单号
	OrderId  string `json:"order_id"`   //支付宝转账单据号，成功一定返回，失败可能不返回也可能返回 例:20160627110070001502260006780837
	PayDate  string `json:"pay_date"`   //支付时间：格式为yyyy-MM-dd HH:mm:ss，仅转账成功返回。
}
