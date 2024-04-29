package response

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/response"
)

// ------------------------------------------------------------------------------------
type ContactWayID struct {
	ConfigID string `json:"config_id"`
}

type ResponseListContactWay struct {
	response.ResponseWork

	ContactWayIDs []*ContactWayID `json:"contact_way"`
	NextCursor    string          `json:"next_cursor"`
}
