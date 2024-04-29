package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetUserBehaviorData struct {
	response.ResponseWork

	MomentList []*power.HashMap `json:"behavior_data"`
}
