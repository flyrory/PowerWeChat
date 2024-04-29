package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInDatas struct {
	response.ResponseWork

	Datas []*power.HashMap `json:"datas"`
}
