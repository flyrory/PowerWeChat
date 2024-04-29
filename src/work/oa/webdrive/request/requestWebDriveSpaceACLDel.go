package request

import "github.com/flyrory/PowerWeChat/v3/src/kernel/power"

type RequestWebDriveSpaceACLDel struct {
	UserID   string           `json:"userid"`
	SpaceID  string           `json:"spaceid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}
