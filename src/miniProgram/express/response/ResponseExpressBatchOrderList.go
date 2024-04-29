package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressBatchOrderList struct {
	response.ResponseMiniProgram

	OrderList []*power.HashMap `json:"order_list"`
}
