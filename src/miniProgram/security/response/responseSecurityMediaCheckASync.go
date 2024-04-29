package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseSecurityMediaCheckASync struct {
	response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
}
