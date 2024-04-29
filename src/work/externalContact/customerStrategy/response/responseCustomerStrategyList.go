package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyList struct {
	response.ResponseWork

	Strategy []*power.HashMap `json:"momentStrategy"`
}
