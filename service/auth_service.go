package service

import (
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/pkg/helpers"
	"mygram/repository/social_media_repository"
	"mygram/repository/comment_repository"
	"mygram/repository/photo_repository"
	"mygram/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	PhotoAuthorization() gin.HandlerFunc
	CommentAuthorization() gin.HandlerFunc
	SocialMediaAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo  user_repository.UserRepository
	photoRepo photo_repository.PhotoRepository
	commentRepo comment_repository.CommentRepository
	socialMediaRepo social_media_repository.SocialMediaRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, commentRepo comment_repository.CommentRepository, socialMediaRepo social_media_repository.SocialMediaRepository) AuthService {
	return &authService{
		userRepo:  userRepo,
		photoRepo: photoRepo,
		commentRepo: commentRepo,
		socialMediaRepo: socialMediaRepo,
	}
}

func (a *authService) SocialMediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		socialMediaId, err := helpers.GetParamId(ctx, "socialMediaId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		socialMedia, err := a.socialMediaRepo.GetSocialMediaById(socialMediaId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if socialMedia.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the social media data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()
	}
}

func (a *authService) CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		commentId, err := helpers.GetParamId(ctx, "commentId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		comment, err := a.commentRepo.GetCommentById(commentId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if comment.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the comment data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()
	}
}

func (a *authService) PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		photoId, err := helpers.GetParamId(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		photo, err := a.photoRepo.GetPhotoById(photoId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the photo data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User // User{Id:0, Email: ""}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}
