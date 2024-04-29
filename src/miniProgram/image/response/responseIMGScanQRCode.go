package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseIMGScanQRCode struct {
	response.ResponseMiniProgram
	CodeResults []*power.HashMap `json:"code_results"`
	ImgSize     power.HashMap    `json:"img_size"`
}
