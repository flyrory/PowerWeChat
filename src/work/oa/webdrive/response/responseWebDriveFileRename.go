package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileRename struct {
	response.ResponseWork

	File *power.HashMap `json:"file"`
}
