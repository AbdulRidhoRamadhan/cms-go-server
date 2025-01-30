package cloudinary

import (
	"context"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

func InitCloudinary() error {
	var err error
	cld, err = cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	return err
}

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buffer := make([]byte, file.Size)
	if _, err := src.Read(buffer); err != nil {
		return "", err
	}

	b64 := base64.StdEncoding.EncodeToString(buffer)
	dataURL := fmt.Sprintf("data:%s;base64,%s", file.Header.Get("Content-Type"), b64)

	uploadResult, err := cld.Upload.Upload(
		context.Background(),
		dataURL,
		uploader.UploadParams{
			ResourceType: "auto",
			PublicID:    file.Filename,
		},
	)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
} 