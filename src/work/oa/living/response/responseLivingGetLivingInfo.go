package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetLivingInfo struct {
	response.ResponseWork

	LivingInfo *power.HashMap `json:"living_info"`
}
