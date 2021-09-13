package imageapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/assetsP/application/apperror"
	"github.com/raismaulana/assetsP/infrastructure/log"
	"github.com/raismaulana/assetsP/infrastructure/util"
	"github.com/raismaulana/assetsP/usecase/uploadimage"
)

// uploadImageHandler ...
func (r *Controller) uploadImageHandler(inputPort uploadimage.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var req uploadimage.InportRequest
		if err := c.ShouldBind(&req); err != nil {
			log.Error(ctx, "bind", err.Error())
			errorMessage := util.GetValidationErrorMessage(err)
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.ERR400.Var(errorMessage)))
			return
		}

		log.Info(ctx, util.MustJSON(req))
		if !strings.HasPrefix(req.Image.Header.Get("Content-Type"), "image/") {
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.UnsupportedImageFormat))
			return
		}
		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(apperror.GetErrorCode(err), NewErrorResponse(err))
			return
		}

		log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
