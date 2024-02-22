package controller

import (
	"Glue-API/httputil"
	"Glue-API/model"
	"Glue-API/utils"
	"Glue-API/utils/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FsStatus godoc
//
//	@Summary		Show Status and List of Glue FS
//	@Description	GlueFS의 상태값과 리스트를 보여줍니다.
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	FsStatus
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs [get]
func (c *Controller) FsStatus(ctx *gin.Context) {
	dat, err := fs.FsStatus()
	dat2, err := fs.FsList()
	value := model.FsSum{
		FsStatus: dat,
		FsList:   dat2,
	}
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.IndentedJSON(http.StatusOK, value)
}

// FsCreate godoc
//
//	@Summary		Create of Glue FS
//	@Description	GlueFS를 생성합니다.
//	@param			fs_name 	path	string	true	"Glue FS Name"
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{string}	string	"Success"
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs/{fs_name} [post]
func (c *Controller) FsCreate(ctx *gin.Context) {
	fs_name := ctx.Param("fs_name")
	dat, err := fs.FsCreate(fs_name)
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.IndentedJSON(http.StatusOK, dat)
}

// FsDelete godoc
//
//	@Summary		Delete of Glue FS
//	@Description	GlueFS를 삭제합니다.
//	@param			fs_name 	path	string	true	"Glue FS Name"
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{string}	string	"Success"
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs/{fs_name} [delete]
func (c *Controller) FsDelete(ctx *gin.Context) {
	fs_name := ctx.Param("fs_name")
	dat, err := fs.FsDelete(fs_name)
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.IndentedJSON(http.StatusOK, dat)
}
func (c *Controller) FsDeleteOptions(ctx *gin.Context) {
	SetOptionHeader(ctx)
	ctx.IndentedJSON(http.StatusOK, nil)
}

// FsGetInfo godoc
//
//	@Summary		Detail Info of Glue FS
//	@Description	GlueFS의 상세 정보를 보여줍니다.
//	@param			fs_name 	path	string	true	"Glue FS Name"
//	@Tags			GlueFS
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	FsGetInfo
//	@Failure		400	{object}	httputil.HTTP400BadRequest
//	@Failure		404	{object}	httputil.HTTP404NotFound
//	@Failure		500	{object}	httputil.HTTP500InternalServerError
//	@Router			/api/v1/gluefs/info/{fs_name} [get]
func (c *Controller) FsGetInfo(ctx *gin.Context) {
	fs_name := ctx.Param("fs_name")
	dat, err := fs.FsGetInfo(fs_name)
	if err != nil {
		utils.FancyHandleError(err)
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	// Print the output
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.IndentedJSON(http.StatusOK, dat)
}
