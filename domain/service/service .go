package service

import (
	"context"
)

type GetBaseURLRepo interface {
	GetBaseURL(ctx context.Context) string
}

type GenerateRandomStringService interface {
	GenerateRandomString(ctx context.Context) string
}

type HashPasswordService interface {
	HashPassword(ctx context.Context, plainPassword string) (string, error)
}

type VerifyPasswordService interface {
	VerifyPassword(ctx context.Context, req VerifyPasswordServiceRequest) error
}

type VerifyPasswordServiceRequest struct {
	PlainPassword  string
	HashedPassword string
}

// type ImageProcessingAndUploadRepo interface {
// 	ImageProcessingAndUpload(ctx context.Context, buffer []byte, quality int, dirname string, filename string, extension bimg.ImageType) error
// }
