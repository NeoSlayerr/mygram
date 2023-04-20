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

type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) socialMediaHandler {
	return socialMediaHandler{
		socialMediaService: socialMediaService,
	}
}

// CreateNewSocialMedia godoc
// @Tags social_media
// @Description Create New Social Media Data
// @ID create-new-social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewSocialMediaRequest true "request body json"
// @Success 201 {object} dto.NewSocialMediaRequest
// @Router /social_media [post]
func (p socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var socialMediaRequest dto.NewSocialMediaRequest

	if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newSocialMedia, err := p.socialMediaService.CreateSocialMedia(user.Id, socialMediaRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newSocialMedia)
}

// GetSocialMedias godoc
// @Tags social_media
// @Description Get All Social Media Data
// @ID get-social-medias
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.GetSocialMediaResponse true "request body json"
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /social_media [get]
func (p socialMediaHandler) GetSocialMedias(c *gin.Context) {

	response, err := p.socialMediaService.GetSocialMedias()

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetSocialMediaById godoc
// @Tags social_media
// @Description Get One Social Media Data By Id
// @ID get-social-media-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.SocialMediaResponse true "request body json"
// @Success 200 {object} dto.SocialMediaResponse
// @Router /social_media/:socialMediaId [get]
func (p socialMediaHandler) GetSocialMediaById(c *gin.Context) {

	socialMediaId, err := helpers.GetParamId(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.socialMediaService.GetSocialMediaById(socialMediaId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteSocialMedia godoc
// @Tags social_media
// @Description Delete Social Media Data By Id
// @ID delete-social-media-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewSocialMediaResponse true "request body json"
// @Success 200 {object} dto.NewSocialMediaResponse
// @Router /social_media/:socialMediaId [delete]
func (p socialMediaHandler) DeleteSocialMediaById(c *gin.Context) {

	socialMediaId, err := helpers.GetParamId(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.socialMediaService.DeleteSocialMediaById(socialMediaId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateSocialMedia godoc
// @Tags social_media
// @Description Update Social Media Data By Id
// @ID update-social-media-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewSocialMediaResponse true "request body json"
// @Success 200 {object} dto.NewSocialMediaResponse
// @Router /social_media/:socialMediaId [put]
func (p socialMediaHandler) UpdateSocialMediaById(c *gin.Context) {
	var socialMediaRequest dto.NewSocialMediaRequest

	if err := c.ShouldBindJSON(&socialMediaRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	socialMediaId, err := helpers.GetParamId(c, "socialMediaId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.socialMediaService.UpdateSocialMediaById(socialMediaId, socialMediaRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
