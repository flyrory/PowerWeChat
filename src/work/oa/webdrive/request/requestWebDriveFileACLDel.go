package request

import "github.com/flyrory/PowerWeChat/v3/src/kernel/power"

type RequestWebDriveFileACLDel struct {
	UserID   string           `json:"userid"`
	FileID   string           `json:"fileid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}
