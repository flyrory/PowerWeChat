package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageGetPubTemplateTitleList struct {
	response.ResponseMiniProgram
	Count int              `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
