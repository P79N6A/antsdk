package data

import (
  "github.com/solarhell/antsdk/api"
)

type KoubeiMarketingDataCustomreportBatchqueryResponse struct {
  api.AlipayResponse
  ReportConditionList []CustomReportCondition `json:"report_condition_list"`  // 分页输出自定义开放数据规则列表
}
