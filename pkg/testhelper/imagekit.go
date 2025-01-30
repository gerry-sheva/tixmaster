package testhelper

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gerry-sheva/tixmaster/pkg/common"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/joho/godotenv"
)

func CreateImageKit() (common.ImageKit, error) {
	err := godotenv.Load(filepath.Join("..", "..", ".env"))
	if err != nil {
		log.Fatal(err.Error())
	}

	imagekit, err := imagekit.New()
	if err != nil {
		panic(err)
	}

	println(os.Getenv("IMAGEKIT_PRIVATE_KEY"))

	ik := common.ImageKit{
		Dir:      "/test",
		ImageKit: imagekit,
	}

	return ik, nil
}
