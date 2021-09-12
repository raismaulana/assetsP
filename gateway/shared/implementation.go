package shared

import (
	"context"

	"github.com/raismaulana/assetsP/infrastructure/auth"
	"github.com/raismaulana/assetsP/infrastructure/envconfig"
	"github.com/raismaulana/assetsP/infrastructure/log"
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
