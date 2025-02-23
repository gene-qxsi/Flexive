package router

import (
	"fmt"
	"log"

	"github.com/gene-qxsi/Flexive/configs"
	"github.com/gene-qxsi/Flexive/internal/delivery/http/v1/controllers"
	"github.com/gene-qxsi/Flexive/internal/middleware"
	"github.com/gene-qxsi/Flexive/internal/repository"
	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gene-qxsi/Flexive/internal/storage"
	"github.com/gene-qxsi/Flexive/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitRouter(conf *configs.Config) *gin.Engine {
	const op = "internal/router/router.go/InitRouter()"

	router := gin.Default()

	postgres, err := storage.OpenDB(&gorm.Config{}, conf)
	if err != nil {
		err = fmt.Errorf("❌ РОУТЕР-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Fatalln(err)
	}

	redis, err := storage.OpenRedis(&redis.Options{})
	if err != nil {
		err = fmt.Errorf("❌ РОУТЕР-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Fatalln(err)
	}

	hasher := services.NewBcryptHasher(conf.Salt)

	userRepo := repository.NewUserRepo(postgres)
	channelRepo := repository.NewChannelRepo(postgres)
	commentRepo := repository.NewCommentRepo(postgres)
	postRepo := repository.NewPostRepo(postgres, redis, conf)
	reactionRepo := repository.NewReactionRepo(postgres)
	subscriptionRepo := repository.NewSubscriptionRepo(postgres)
	authRepo := repository.NewAuthRepository(redis, conf)
	profileRepo := repository.NewProfileRepository(postgres)

	userService := services.NewUserService(userRepo, hasher)
	channelService := services.NewChannelService(channelRepo)
	commentService := services.NewCommentService(commentRepo)
	postService := services.NewPostService(postRepo)
	reactionService := services.NewReactionService(reactionRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	authService := services.NewAuthService(authRepo, conf)
	profileService := services.NewProfileService(profileRepo)

	authUseCase := usecase.NewAuthUseCase(userService, authService)
	profileUseCase := usecase.NewProfileUsecase(profileService)

	userHandler := controllers.NewUserController(userService)
	channelController := controllers.NewChannelController(channelService)
	commentController := controllers.NewCommentHandler(commentService)
	postController := controllers.NewPostHandler(postService)
	reactionController := controllers.NewReactionHandler(reactionService)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService)
	authController := controllers.NewAuthController(authUseCase)
	profileController := controllers.NewProfileController(profileUseCase)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	// USERS API
	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/", userHandler.GetUsers)
		usersGroup.GET("/:id", userHandler.GetUser)
		// usersGroup.POST("/", userHandler.CreateUser)
		usersGroup.PATCH("/:id", userHandler.UpdateUser)
		usersGroup.DELETE("/:id", userHandler.DeleteUser)
	}
	// CHANNELS API
	channelsGroup := router.Group("/channels", authMiddleware.JWTAuth())
	{
		channelsGroup.GET("/", channelController.GetChannels)
		channelsGroup.GET("/:id", channelController.GetChannel)
		channelsGroup.POST("/", channelController.CreateChannel)
		channelsGroup.PATCH("/:id", channelController.UpdateChannel)
		channelsGroup.DELETE("/:id", channelController.DeleteChannel)
	}
	// COMMENTS API
	commentsGroup := router.Group("/comments", authMiddleware.JWTAuth())
	{
		commentsGroup.GET("/", commentController.GetComments)
		commentsGroup.GET("/:id", commentController.GetComment)
		commentsGroup.POST("/", commentController.CreateComment)
		commentsGroup.PATCH("/:id", commentController.UpdateComment)
		commentsGroup.DELETE("/:id", commentController.DeleteComment)
	}
	// POSTS API
	postsGroup := router.Group("/posts", authMiddleware.JWTAuth())
	{
		postsGroup.GET("/", postController.GetPosts)
		postsGroup.GET("/:id", postController.GetPost)
		postsGroup.POST("/", postController.CreatePost)
		postsGroup.PATCH("/:id", postController.UpdatePost)
		postsGroup.DELETE("/:id", postController.DeletePost)
	}
	// REACTIONS API
	reactionsGroup := router.Group("/reactions", authMiddleware.JWTAuth())
	{
		reactionsGroup.GET("/", reactionController.GetReactions)
		reactionsGroup.GET("/:userID/:postID", reactionController.GetReaction)
		reactionsGroup.POST("/", reactionController.CreateReaction)
		reactionsGroup.PATCH("/:userID/:postID", reactionController.UpdateReaction)
		reactionsGroup.DELETE("/:userID/:postID", reactionController.DeleteReaction)
	}
	// SUBSCRIPTIONS API
	subscriptionsGroup := router.Group("/subscriptions", authMiddleware.JWTAuth())
	subscriptionsGroup.Use()
	{
		subscriptionsGroup.GET("/", subscriptionController.GetSubscriptions)
		subscriptionsGroup.GET("/:userID/:channelID", subscriptionController.GetSubscription)
		subscriptionsGroup.POST("/", subscriptionController.CreateSubscription)
		subscriptionsGroup.PATCH("/:userID/:channelID", subscriptionController.UpdateSubscription)
		subscriptionsGroup.DELETE("/:userID/:channelID", subscriptionController.DeleteSubscription)
	}
	// AUTHORIZATION API
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/sign-in", authController.SignIn)
		authGroup.POST("/sign-up", authController.SignUp)
		authGroup.POST("/refresh", authController.RefreshToken)
		authGroup.POST("/sign-out", authController.SignOut)
	}

	// PROFILE API
	profileGroup := router.Group("/profile")
	{
		profileGroup.GET("/:id", profileController.GetProfile)
		profileGroup.PUT("/:id", profileController.UpdateProfile)
	}

	return router
}
