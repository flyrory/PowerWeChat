package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveSpaceInfo struct {
	response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}
