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
	cld, _, initCldErr := pkgcld.Initialize()
	if initCldErr != nil {
		fmt.Println("Error initialize cloudinary, " + initCldErr.Error())
	}

	transforms := ""
	urlImg, urlErr := pkgcld.Transform(cld, transforms, "messi")
	if urlErr != nil {
		fmt.Print("Error: " + urlErr.Error())
	}

	fmt.Println(urlErr)
}
