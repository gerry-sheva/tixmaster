package common

import "github.com/imagekit-developer/imagekit-go"

type ImageKit struct {
	Dir string
	*imagekit.ImageKit
}
