package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileShare struct {
	response.ResponseWork

	ShareURL string `json:"share_url"`
}
