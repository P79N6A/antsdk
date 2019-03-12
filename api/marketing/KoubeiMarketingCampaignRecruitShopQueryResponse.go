package marketing

import (
  "github.com/solarhell/antsdk/api"
)

type KoubeiMarketingCampaignRecruitShopQueryResponse struct {
  api.AlipayResponse
  ShopCount string        `json:"shop_count"`
  PlanId    string        `json:"plan_id"`
  ShopInfo  []RecruitInfo `json:"shop_info"`
}
