package request

import "github.com/flyrory/PowerWeChat/v3/src/kernel/power"

type RequestMaterialUpdateNews struct {
	MediaID  int64            `json:"media_id"`
	Index    int64            `json:"index"`
	Articles []*power.HashMap `json:"articles"`
}
