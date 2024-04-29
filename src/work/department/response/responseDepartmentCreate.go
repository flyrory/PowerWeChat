package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	response.ResponseWork
	ID int `json:"id"`
}
