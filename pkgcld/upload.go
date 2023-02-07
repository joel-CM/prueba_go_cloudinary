package pkgcld

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Upload(ctx context.Context, cld *cloudinary.Cloudinary, urlImg string, id string) (string, error) {
	uploadResult, uploadErr := cld.Upload.Upload(ctx, urlImg, uploader.UploadParams{PublicID: id})

	if uploadErr != nil {
		return "", uploadErr
	}

	return uploadResult.SecureURL, nil
}
