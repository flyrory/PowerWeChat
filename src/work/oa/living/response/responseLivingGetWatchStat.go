package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetWatchStat struct {
	response.ResponseWork

	NextKey  string         `json:"next_key"`
	StatInfo *power.HashMap `json:"stat_info"`
}
