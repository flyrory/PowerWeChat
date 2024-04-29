package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseSearchImageSearch struct {
	response.ResponseMiniProgram

	Items []*power.HashMap `json:"items"`
}
