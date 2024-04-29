package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomAdd struct {
	response.ResponseWork

	MeetingRoomID int `json:"meetingroom_id"`
}
