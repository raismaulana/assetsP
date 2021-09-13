package imageapi

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/assetsP/infrastructure/auth"
	"github.com/raismaulana/assetsP/infrastructure/envconfig"
	"github.com/raismaulana/assetsP/usecase/uploadimage"
)

type Controller struct {
	JWTToken          auth.JWTToken
	Env               envconfig.EnvConfig
	Enforcer          casbin.Enforcer
	Router            gin.IRouter
	UploadImageInport uploadimage.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/image/upload", r.authorized(), r.uploadImageHandler(r.UploadImageInport))
}
