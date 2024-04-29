package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseIMGSuperResolution struct {
	response.ResponseMiniProgram
	MediaID string `json:"media_id"`
}
