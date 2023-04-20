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

type commentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) commentHandler {
	return commentHandler{
		commentService: commentService,
	}
}

// CreateNewComment godoc
// @Tags comment
// @Description Create New Comment Data
// @ID create-new-comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewCommentRequest true "request body json"
// @Success 201 {object} dto.NewCommentRequest
// @Router /comment [post]
func (p commentHandler) CreateComment(c *gin.Context) {
	var commentRequest dto.NewCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newComment, err := p.commentService.CreateComment(user.Id, commentRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newComment)
}

// GetComments godoc
// @Tags comment
// @Description Get All Comment Data
// @ID get-comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.GetCommentResponse true "request body json"
// @Success 200 {object} dto.GetCommentResponse
// @Router /comment [get]
func (p commentHandler) GetComments(c *gin.Context) {

	response, err := p.commentService.GetComments()

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetCommentById godoc
// @Tags comment
// @Description Get One Comment Data By Id
// @ID get-comment-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.CommentResponse true "request body json"
// @Success 200 {object} dto.CommentResponse
// @Router /comment/:commentId [get]
func (p commentHandler) GetCommentById(c *gin.Context) {

	commentId, err := helpers.GetParamId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.commentService.GetCommentById(commentId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteComment godoc
// @Tags comment
// @Description Delete Comment Data By Id
// @ID delete-comment-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewCommentResponse true "request body json"
// @Success 200 {object} dto.NewCommentResponse
// @Router /comment/:commentId [delete]
func (p commentHandler) DeleteCommentById(c *gin.Context) {

	commentId, err := helpers.GetParamId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.commentService.DeleteCommentById(commentId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateComment godoc
// @Tags comment
// @Description Update Comment Data By Id
// @ID update-comment-by-id
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewCommentResponse true "request body json"
// @Success 200 {object} dto.NewCommentResponse
// @Router /comment/:commentId [put]
func (p commentHandler) UpdateCommentById(c *gin.Context) {
	var commentRequest dto.NewCommentRequest

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	commentId, err := helpers.GetParamId(c, "commentId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.commentService.UpdateCommentById(commentId, commentRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
