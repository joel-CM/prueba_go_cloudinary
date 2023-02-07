package pkgcld

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

func Get(ctx context.Context, cld *cloudinary.Cloudinary, id string) (urlImg string, err error) {
	asset, assetErr := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: id})

	if assetErr != nil {
		return "", nil
	}

	return asset.SecureURL, nil
}
