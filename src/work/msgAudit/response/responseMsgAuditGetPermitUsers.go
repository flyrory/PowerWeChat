package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	response.ResponseWork
	IDs []string `json:"ids"`
}
