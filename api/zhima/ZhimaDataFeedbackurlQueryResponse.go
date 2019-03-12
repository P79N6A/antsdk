package zhima

import (
  "github.com/solarhell/antsdk/api"
)

type ZhimaDataFeedbackurlQueryResponse struct {
  api.AlipayResponse
  FeedbackUrl string `json:"feedback_url"`  // 反馈模板地址
}
