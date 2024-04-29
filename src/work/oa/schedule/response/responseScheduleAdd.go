package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseScheduleAdd struct {
	response.ResponseWork

	ScheduleID string `json:"schedule_id"`
}
