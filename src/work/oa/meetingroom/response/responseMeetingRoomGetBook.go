package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomGetBook struct {
	response.ResponseWork

	MeetingID  int `json:"meeting_id"`
	ScheduleID int `json:"schedule_id"`
}
