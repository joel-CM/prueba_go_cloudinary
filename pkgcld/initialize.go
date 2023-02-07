package pkgcld

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func Initialize() (*cloudinary.Cloudinary, context.Context, error) {
	cld, cldErr := cloudinary.NewFromParams(os.Getenv("CLOUD"), os.Getenv("KEY"), os.Getenv("SECRET"))
	if cldErr != nil {
		return nil, nil, cldErr
	}

	// context
	cotx := context.Background()

	return cld, cotx, nil
}
