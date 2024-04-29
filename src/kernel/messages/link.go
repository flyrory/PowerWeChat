package messages

import (
	"github.com/flyrory/PowerWeChat/v3/src/kernel/power"
)

type Link struct {
	*Message
}

func NewLink(items *power.HashMap) *Link {
	m := &Link{
		NewMessage(items),
	}
	m.Type = "link"

	m.Properties = []string{
		"title",
		"description",
		"url",
	}

	return m
}
