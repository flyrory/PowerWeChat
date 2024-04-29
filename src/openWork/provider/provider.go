package provider

import "github.com/flyrory/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *AccessToken, error) {
	client, err := NewClient(app)
	if err != nil {
		return nil, nil, err
	}
	accessToken, err := NewAccessToken(&app)
	if err != nil {
		return nil, nil, err
	}
	return client, accessToken, nil
}
