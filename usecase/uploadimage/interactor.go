package uploadimage

import (
	"bytes"
	"context"

	"github.com/disintegration/imaging"
	"github.com/raismaulana/assetsP/application/apperror"
	"github.com/raismaulana/assetsP/infrastructure/util"
)

//go:generate mockery --name Outport -output mocks/

type uploadImageInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UploadImage
func NewUsecase(outputPort Outport) Inport {
	return &uploadImageInteractor{
		outport: outputPort,
	}
}

// Execute the usecase UploadImage
func (r *uploadImageInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	filename := r.outport.GenerateRandomString(ctx)
	path := "./public/images/"
	fullpath := path + filename + ".jpg"

	file, err := req.Image.Open()
	if err != nil {
		return nil, apperror.ERR500.Var(err.Error())
	}
	defer file.Close()

	srcImage, err := imaging.Decode(file, imaging.AutoOrientation(true))
	if err != nil {
		return nil, apperror.ERR500.Var(err.Error())
	}

	err = util.CreateDirectoryIfNotExist(path)
	if err != nil {
		return nil, err
	}

	buffer := &bytes.Buffer{}
	err = imaging.Encode(buffer, srcImage, imaging.JPEG, imaging.JPEGQuality(50))
	if err != nil {
		return nil, apperror.ERR500.Var(err.Error())
	}

	img, err := imaging.Decode(buffer)
	if err != nil {
		return nil, apperror.ERR500.Var(err.Error())
	}

	err = imaging.Save(img, fullpath, imaging.JPEGQuality(50))
	if err != nil {
		return nil, apperror.ERR500.Var(err.Error())
	}

	res.Location = r.outport.GetBaseURL(ctx) + fullpath

	return res, nil
}
