package shared

import (
	"context"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"github.com/raismaulana/assetsP/application/apperror"
	"github.com/raismaulana/assetsP/domain/service"
	"github.com/raismaulana/assetsP/infrastructure/auth"
	"github.com/raismaulana/assetsP/infrastructure/envconfig"
	"github.com/raismaulana/assetsP/infrastructure/log"
	"golang.org/x/crypto/bcrypt"
)

type SharedGateway struct {
	Env      *envconfig.EnvConfig
	JWTToken *auth.JWTToken
}

func NewSharedGateway(env *envconfig.EnvConfig, jwtToken *auth.JWTToken) *SharedGateway {
	return &SharedGateway{
		Env:      env,
		JWTToken: jwtToken,
	}
}

func (r *SharedGateway) GetBaseURL(ctx context.Context) string {
	log.Info(ctx, "called")
	return r.Env.AppBaseURL
}

func (r *SharedGateway) HashPassword(ctx context.Context, plainPassword string) (string, error) {
	log.Info(ctx, "called")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return "", apperror.ERR500.Var(err)
	}

	return string(hashedPassword), nil
}

func (r *SharedGateway) GenerateRandomString(ctx context.Context) string {
	log.Info(ctx, "called")

	return uuid.NewString()
}

func (r *SharedGateway) VerifyPassword(ctx context.Context, req service.VerifyPasswordServiceRequest) error {
	log.Info(ctx, "called")

	err := bcrypt.CompareHashAndPassword([]byte(req.HashedPassword), []byte(req.PlainPassword))
	if err != nil {
		return apperror.ERR400.Var(err)
	}

	return nil
}

// The mime type of the image is changed, it is compressed and then saved in the specified folder.
func (r *SharedGateway) ImageProcessingAndUpload(ctx context.Context, buffer []byte, quality int, dirname string, filename string, extension bimg.ImageType) error {
	log.Info(ctx, "called")

	converted, err := bimg.NewImage(buffer).Convert(extension)
	if err != nil {
		return apperror.ERR500.Var(err)
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
	if err != nil {
		return apperror.ERR500.Var(err)
	}

	err = bimg.Write("."+dirname+filename, processed)
	if err != nil {
		return apperror.ERR500.Var(err)
	}

	return nil
}
