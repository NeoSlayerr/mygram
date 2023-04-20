package handler

import (
	"mygram/database"
	"mygram/docs"
	"mygram/repository/comment_repository/comment_pg"
	"mygram/repository/photo_repository/photo_pg"
	"mygram/repository/social_media_repository/social_media_pg"
	"mygram/repository/user_repository/user_pg"
	"mygram/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	const PORT = ":3000"
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	socialMediaRepo := social_media_pg.NewSocialMediaPG(db)

	socialMediaService := service.NewSocialMediaService(socialMediaRepo)

	socialMediaHandler := NewSocialMediaHandler(socialMediaService)

	commentRepo := comment_pg.NewCommentPG(db)

	commentService := service.NewCommentService(commentRepo)

	commentHandler := NewCommentHandler(commentService)

	photoRepo := photo_pg.NewPhotoPG(db)

	photoService := service.NewPhotoService(photoRepo)

	photoHandler := NewPhotoHandler(photoService)

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialMediaRepo)

	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "MyGram"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoutes := route.Group("/user")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
	}

	photoRoute := route.Group("/photo")
	{
		photoRoute.GET("/", authService.Authentication(), photoHandler.GetPhotos)

		photoRoute.GET("/:photoId", authService.Authentication(), photoHandler.GetPhotoById)

		photoRoute.POST("/", authService.Authentication(), photoHandler.CreatePhoto)

		photoRoute.PUT("/:photoId", authService.Authentication(), authService.PhotoAuthorization(), photoHandler.UpdatePhotoById)

		photoRoute.DELETE("/:photoId", authService.Authentication(), authService.PhotoAuthorization(), photoHandler.DeletePhotoById)
	}

	commentRoute := route.Group("/comment")
	{
		commentRoute.GET("/", authService.Authentication(), commentHandler.GetComments)

		commentRoute.GET("/:commentId", authService.Authentication(), commentHandler.GetCommentById)

		commentRoute.POST("/", authService.Authentication(), commentHandler.CreateComment)

		commentRoute.PUT("/:commentId", authService.Authentication(), authService.CommentAuthorization(), commentHandler.UpdateCommentById)

		commentRoute.DELETE("/:commentId", authService.Authentication(), authService.CommentAuthorization(), commentHandler.DeleteCommentById)
	}

	socialMediaRoute := route.Group("/social_media")
	{
		socialMediaRoute.GET("/", authService.Authentication(), socialMediaHandler.GetSocialMedias)

		socialMediaRoute.GET("/:socialMediaId", authService.Authentication(), socialMediaHandler.GetSocialMediaById)

		socialMediaRoute.POST("/", authService.Authentication(), socialMediaHandler.CreateSocialMedia)

		socialMediaRoute.PUT("/:socialMediaId", authService.Authentication(), authService.SocialMediaAuthorization(), socialMediaHandler.UpdateSocialMediaById)

		socialMediaRoute.DELETE("/:socialMediaId", authService.Authentication(), authService.SocialMediaAuthorization(), socialMediaHandler.DeleteSocialMediaById)
	}

	route.Run(PORT)
}
