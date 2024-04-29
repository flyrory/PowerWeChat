package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
