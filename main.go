package main

import (
	"fmt"

	"github.com/joel-CM/prueba_go_cloudinary/pkgcld"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	cld, ctx, initCldErr := pkgcld.Initialize()
	if initCldErr != nil {
		fmt.Println("Error initialize cloudinary, " + initCldErr.Error())
	}

	urlImg, urlErr := pkgcld.Upload(ctx, cld, "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSvrxxc_aRuyD2mEd5PJf29VGwGJ50fzh1NJ1bwhpYCrw&s", "messi")
	if urlErr != nil {
		fmt.Println("Error: " + urlErr.Error())
	}

	fmt.Print(urlImg)
}
