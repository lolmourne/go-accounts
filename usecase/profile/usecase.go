package profile

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	uuid "github.com/satori/go.uuid"
)

const AllowedExtensions = ".jpeg,.jpg"

func (u *Usecase) UploadFile(file []byte) (string, error) {
	mime := mimetype.Detect(file)
	if strings.Index(AllowedExtensions, mime.Extension()) == -1 {
		return "", errors.New("File Type is not allowed, file type: " + mime.Extension())
	}
	log.Println(mime)
	uid := uuid.NewV4()

	fileName := fmt.Sprintf("image/profile/%s.%s", uid.String(), mime.Extension())
	err := u.res.Put(file, fileName, mime.Extension())
	if err != nil {
		log.Print(err)
		return "", err
	}

	return "https://skilvul-course.s3-ap-southeast-1.amazonaws.com/" + fileName, nil
}
