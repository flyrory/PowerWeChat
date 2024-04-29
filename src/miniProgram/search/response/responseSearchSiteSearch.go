package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

type ResponseSearchSiteSearch struct {
	response.ResponseMiniProgram

	Items        []*power.HashMap `json:"items"`
	HasNextPage  int              `json:"has_next_page"`
	HitCount     int              `json:"hit_count"`
	NextPageInfo string           `json:"next_page_info"`
}
