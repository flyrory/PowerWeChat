package request

import "github.com/flyrory/PowerWeChat/v3/src/kernel/power"

type RequestWebDriveSpaceCreate struct {
	UserID    string           `json:"userid"`
	SpaceName string           `json:"space_name"`
	AuthInfo  []*power.HashMap `json:"auth_info"`
}
