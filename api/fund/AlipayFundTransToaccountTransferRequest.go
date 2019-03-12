package fund

import (
	"github.com/solarhell/antsdk/api"
	"github.com/solarhell/antsdk/utils"
)

//单笔转账到支付宝账户接口
type AlipayFundTransToaccountTransferRequest struct {
	api.IAlipayRequest
	TerminalType string                                            `json:"terminal_type"`
	TerminalInfo string                                            `json:"terminal_info"`
	ProdCode     string                                            `json:"prod_code"`
	NotifyUrl    string                                            `json:"notify_url"`
	ReturnUrl    string                                            `json:"return_url"`
	BizContent   AlipayFundTransToaccountTransferRequestBizContent `json:"biz_content"`
}

type AlipayFundTransToaccountTransferRequestBizContent struct {
	OutBizNo      string `json:"out_biz_no"`      //商户转账唯一订单号
	PayeeType     string `json:"payee_type"`      //收款方账户类型.可取值：1、ALIPAY_USERID：支付宝账号对应的支付宝唯一用户号。以2088开头的16位纯数字组成。2、ALIPAY_LOGONID：支付宝登录号，支持邮箱和手机号格式。
	PayeeAccount  string `json:"payee_account"`   //收款方账户
	Amount        string `json:"amount"`          // 转账金额，单位：元。只支持2位小数，小数点前最大支持13位，金额必须大于等于0.1元。
	PayerShowName string `json:"payer_show_name"` //付款方姓名
	PayeeRealName string `json:"payee_real_name"` //收款方真实姓名
	Remark        string `json:"remark"`          //转账备注（支持200个英文/100个汉字）。 当付款方为企业账户，且转账金额达到（大于等于）50000元，remark不能为空。收款方可见，会展示在收款用户的收支详情中。
}

func (this *AlipayFundTransToaccountTransferRequest) GetApiMethodName() string {
	return "alipay.fund.trans.toaccount.transfer"
}

func (this *AlipayFundTransToaccountTransferRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayFundTransToaccountTransferRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayFundTransToaccountTransferRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayFundTransToaccountTransferRequest) GetNotifyUrl() string {
	return this.NotifyUrl
}

func (this *AlipayFundTransToaccountTransferRequest) GetReturnUrl() string {
	return this.ReturnUrl
}

func (this *AlipayFundTransToaccountTransferRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayFundTransToaccountTransferRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayFundTransToaccountTransferRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJson(this.BizContent))
	return txtParams
}