package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCalendarGet struct {
	response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`
}
