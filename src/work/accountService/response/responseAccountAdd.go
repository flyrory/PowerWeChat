package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountAdd struct {
	response.ResponseWork

	OpenKFID string `json:"open_kfid"`
}
