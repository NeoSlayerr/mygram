package handler

import (
	"mygram/dto"
	"mygram/pkg/errs"
	"mygram/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// Register godoc
// @Tags user
// @Description Create New User
// @ID register
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewUserRequest true "request body json"
// @Success 201 {object} dto.NewUserRequest
// @Router /user/register [post]
func (uh *userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)

}

// Login godoc
// @Tags user
// @Description Login using created account
// @ID login
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewLoginRequest true "request body json"
// @Success 201 {object} dto.NewLoginRequest
// @Router /user/login [post]
func (uh *userHandler) Login(ctx *gin.Context) {
	var newLoginRequest dto.NewLoginRequest

	if err := ctx.ShouldBindJSON(&newLoginRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Login(newLoginRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
