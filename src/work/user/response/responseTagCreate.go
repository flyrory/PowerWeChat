package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagCreate struct {
	response.ResponseWork

	TagID string `json:"tagid"`
}
