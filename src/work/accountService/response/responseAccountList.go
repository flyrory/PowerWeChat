package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountList struct {
	response.ResponseWork

	AccountList []*power.HashMap `json:"account_list"`
}
