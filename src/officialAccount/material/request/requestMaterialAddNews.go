package request

import "github.com/flyrory/PowerWeChat/v3/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}
