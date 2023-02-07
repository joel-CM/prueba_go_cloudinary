package pkgcld

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func Transform(cld *cloudinary.Cloudinary, transforms string, id string) (urlImg string, err error) {
	asset, assetErr := cld.Image(id)
	if assetErr != nil {
		return "", assetErr
	}

	asset.Transformation = transforms

	url, urlErr := asset.String()
	if urlErr != nil {
		return "", urlErr
	}

	return url, nil
}
