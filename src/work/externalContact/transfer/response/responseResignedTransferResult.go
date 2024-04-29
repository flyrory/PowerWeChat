package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseResignedTransferResult struct {
	response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}
