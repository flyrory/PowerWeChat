package response

import "github.com/flyrory/PowerWeChat/v3/src/kernel/response"

type ResponseExpressGetQuota struct {
	response.ResponseMiniProgram

	QuotaNum string `json:"quota_num"`
}
