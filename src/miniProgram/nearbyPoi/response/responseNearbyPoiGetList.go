package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseNearbyPoiGetList struct {
	response.ResponseMiniProgram
	Data *power.HashMap `json:"data"`
}
