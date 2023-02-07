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

	urlImg, urlErr := pkgcld.Get(ctx, cld, "messi")
	if urlErr != nil {
		fmt.Println("Error: " + urlErr.Error())
	}

	fmt.Println(urlImg)
}