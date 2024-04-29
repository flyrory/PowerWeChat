package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetJsErrList struct {
	response.ResponseMiniProgram
	Success    bool             `json:"success"`
	OpenID     string           `json:"openid"`
	Data       []*power.HashMap `json:"data"`
	TotalCount int64            `json:"totalCount"`
}
