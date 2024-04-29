package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentStrategyGet struct {
	response.ResponseWork

	Strategy *power.HashMap `json:"strategy"`
}
