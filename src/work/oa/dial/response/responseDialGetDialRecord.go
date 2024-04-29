package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseDialGetDialRecord struct {
	response.ResponseWork

	MeetingidList []*power.HashMap `json:"record"`
}
