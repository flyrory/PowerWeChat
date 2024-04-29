package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMenuCreate struct {
	response.ResponseWork
	Button []*power.HashMap `json:"button"`
}
