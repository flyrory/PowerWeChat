package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingGetUserMeetingID struct {
	response.ResponseWork

	NextCursor    string   `json:"next_cursor"`
	MeetingidList []string `json:"meetingid_list"`
}
