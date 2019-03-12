package marketing

import (
  "github.com/solarhell/antsdk/api"
)

type KoubeiMarketingCampaignActivityQueryResponse struct {
  api.AlipayResponse
  CampDetail CampDetail `json:"camp_detail"`  // 活动详情
}
