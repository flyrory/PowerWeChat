package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInRecords struct {
	response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
