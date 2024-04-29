package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomList struct {
	response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}
