package imageapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.FailUnmarshalResponseBodyError))
				return
			}
			var errorMessage string
			for _, e := range errs {
				errorMessage = errorMessage + e.Translate(util.Trans) + "\n"
			}

			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.ERR400.Var(errorMessage)))
			return
		}

		log.Info(ctx, util.MustJSON(req))

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
