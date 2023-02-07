package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	cld, cldErr := cloudinary.NewFromParams(os.Getenv("CLOUD"), os.Getenv("KEY"), os.Getenv("SECRET"))
	if cldErr != nil {
		fmt.Println("Error initialize cloudinary, " + cldErr.Error())
	}

	// context
	ctx := context.Background()

	uploadResult, uploadErr := cld.Upload.Upload(ctx, "https://avatars.githubusercontent.com/u/84867719", uploader.UploadParams{PublicID: "joel"})
	if uploadErr != nil {
		fmt.Println("error upload, " + uploadErr.Error())
	}

	fmt.Println("secure ulr: " + uploadResult.SecureURL)

	// asset, assetErr := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "joel"})
	// if assetErr != nil {
	// 	fmt.Println("error asset, " + assetErr.Error())
	// }

	// fmt.Printf("asset name: %v, asset url: %v", asset.PublicID, asset.SecureURL)
}
