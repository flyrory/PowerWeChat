package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCalendarAdd struct {
	response.ResponseWork

	CalID string `json:"cal_id"`
}
