package security

import (
	"github.com/flyrory/PowerWeChat/v3/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	client, err := NewClient(&app)
	if err != nil {
		return nil, err
	}

	return client, nil

}
