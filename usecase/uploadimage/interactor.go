package uploadimage

import (
	"context"
	"io"

	"github.com/h2non/bimg"
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
	path := "/public/images/"
	fullpath := path + filename

	file, err := req.Image.Open()
	if err != nil {
		return nil, apperror.ERR500.Var(err.Error())
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, apperror.ERR500.Var(err.Error())
	}

	err = util.CreateDirectoryIfNotExist("." + path)
	if err != nil {
		return nil, err
	}

	err = r.outport.ImageProcessingAndUpload(ctx, buffer, 50, path, filename, bimg.WEBP)
	if err != nil {
		return nil, err
	}

	res.Location = r.outport.GetBaseURL(ctx) + fullpath

	return res, nil
}
