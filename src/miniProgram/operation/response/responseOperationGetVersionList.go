package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
