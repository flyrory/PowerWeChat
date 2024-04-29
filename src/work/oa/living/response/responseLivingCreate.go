package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingCreate struct {
	response.ResponseWork

	LivingID int `json:"livingid"`
}
