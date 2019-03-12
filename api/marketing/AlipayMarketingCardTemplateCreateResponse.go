package marketing

import (
  "github.com/solarhell/antsdk/api"
)

type AlipayMarketingCardTemplateCreateResponse struct {
  api.AlipayResponse
  TemplateId string `json:"template_id"`  // 支付宝卡模板ID
}
