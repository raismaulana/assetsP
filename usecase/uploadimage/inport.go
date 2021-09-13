package uploadimage

import (
	"context"
	"mime/multipart"
)

// Inport of UploadImage
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UploadImage
type InportRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

// InportResponse is response payload after running the usecase UploadImage
type InportResponse struct {
	Location string `json:"location"`
}
