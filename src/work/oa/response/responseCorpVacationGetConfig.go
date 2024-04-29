package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpVacationGetConfig struct {
	response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
