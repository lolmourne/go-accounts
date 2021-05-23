package profile

import "github.com/lolmourne/go-accounts/resource/s3"

type IUsecase interface {
	UploadFile([]byte) (string, error)
}

type Usecase struct {
	res s3.IS3
}

func NewUsecase(res s3.IS3) IUsecase {
	return &Usecase{
		res: res,
	}
}
