package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentStrategyCreate struct {
	response.ResponseWork

	Strategy   []*power.HashMap `json:"strategy"`
	NextCursor string           `json:"next_cursor"`
}
