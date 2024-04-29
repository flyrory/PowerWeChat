package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	response.ResponseMiniProgram

	UserName string `json:"username"`
}
