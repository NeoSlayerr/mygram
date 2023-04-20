package handler

import (
	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/pkg/helpers"
	"mygram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoHandler {
	return photoHandler{
		photoService: photoService,
	}
}

// CreateNewPhoto godoc
// @Tags photo
// @Description Create New Photo Data
// @ID create-new-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewPhotoRequest true "request body json"
// @Success 201 {object} dto.NewPhotoRequest
// @Router /photo [post]
func (p photoHandler) CreatePhoto(c *gin.Context) {
	var photoRequest dto.NewPhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newPhoto, err := p.photoService.CreatePhoto(user.Id, photoRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newPhoto)
}

// GetPhotos godoc
// @Tags photo
// @Description Get All Photo Data
// @ID get-photos
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.GetPhotoResponse true "request body json"
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photo [get]
func (p photoHandler) GetPhotos(c *gin.Context) {

	response, err := p.photoService.GetPhotos()

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetPhotoById godoc
// @Tags photo
// @Description Get One Photo Data By Id
// @ID get-photo-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.PhotoResponse true "request body json"
// @Success 200 {object} dto.PhotoResponse
// @Router /photo/:photoId [get]
func (p photoHandler) GetPhotoById(c *gin.Context) {

	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.photoService.GetPhotoById(photoId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeletePhoto godoc
// @Tags photo
// @Description Delete Photo Data By Id
// @ID delete-photo-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewPhotoResponse true "request body json"
// @Success 200 {object} dto.NewPhotoResponse
// @Router /photo/:photoId [delete]
func (p photoHandler) DeletePhotoById(c *gin.Context) {

	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.photoService.DeletePhotoById(photoId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdatePhoto godoc
// @Tags photo
// @Description Update Photo Data By Id
// @ID update-photo-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewPhotoResponse true "request body json"
// @Success 200 {object} dto.NewPhotoResponse
// @Router /photo/:photoId [put]
func (p photoHandler) UpdatePhotoById(c *gin.Context) {
	var photoRequest dto.NewPhotoRequest

	if err := c.ShouldBindJSON(&photoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	photoId, err := helpers.GetParamId(c, "photoId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.photoService.UpdatePhotoById(photoId, photoRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
