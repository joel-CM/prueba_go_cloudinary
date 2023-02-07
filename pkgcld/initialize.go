package pkgcld

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func Initialize() (*cloudinary.Cloudinary, context.Context, error) {
	CLD_KEY := os.Getenv("KEY")
	CLD_SECRET := os.Getenv("SECRET")
	CLD_CLOUD := os.Getenv("CLOUD")
	CLD_URL := fmt.Sprintf("cloudinary://%v:%v@%v", CLD_KEY, CLD_SECRET, CLD_CLOUD)

	cld, cldErr := cloudinary.NewFromURL(CLD_URL)
	if cldErr != nil {
		return nil, nil, cldErr
	}

	// context
	cotx := context.Background()

	return cld, cotx, nil
}
