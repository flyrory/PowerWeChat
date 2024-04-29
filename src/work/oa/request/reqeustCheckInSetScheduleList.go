package request

import "github.com/flyrory/PowerWeChat/v3/src/kernel/power"

type RequestCheckInSetScheduleList struct {
	GroupID   int              `json:"groupid"`
	Items     []*power.HashMap `json:"items"`
	YearMonth int              `json:"yearmonth"`
}
