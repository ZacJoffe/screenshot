package cmd

import (
	"os"

	"github.com/ZacJoffe/uploader-libs/imgur"
)

const CLIENT_ID = "0d297558de98a48"

func uploadImage(image *os.File) (string, error) {
	return imgur.UploadImage(image, CLIENT_ID)
}
