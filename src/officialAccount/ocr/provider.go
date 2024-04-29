package ocr

import "github.com/flyrory/PowerWeChat/v3/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(app)

}
