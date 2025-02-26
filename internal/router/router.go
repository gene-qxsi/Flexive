package router

import (
	"fmt"
	"log"

	"github.com/gene-qxsi/Flexive/configs"
	controllersHTTP "github.com/gene-qxsi/Flexive/internal/controllers/http/v1"
	controllersWS "github.com/gene-qxsi/Flexive/internal/controllers/ws"
	"github.com/gene-qxsi/Flexive/internal/middleware"
	auth "github.com/gene-qxsi/Flexive/internal/repository"
	repository "github.com/gene-qxsi/Flexive/internal/repository/sqlrepo"
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
	subscriptionRepo := repository.NewSubscriptionRepo(postgres)
	authRepo := auth.NewAuthRepository(redis, conf)
	profileRepo := repository.NewProfileRepository(postgres)
	chatRepo := repository.NewChatRepo(postgres)

	userService := services.NewUserService(userRepo, hasher)
	channelService := services.NewChannelService(channelRepo)
	commentService := services.NewCommentService(commentRepo)
	postService := services.NewPostService(postRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	authService := services.NewAuthService(authRepo, conf)
	profileService := services.NewProfileService(profileRepo)
	chatService := services.NewChatService(chatRepo)

	authUsecase := usecase.NewAuthUseCase(userService, authService, profileService)
	profileUsecase := usecase.NewProfileUsecase(profileService)
	chatUsecase := usecase.NewChatUsecase(chatService)

	userHandler := controllersHTTP.NewUserController(userService)
	channelController := controllersHTTP.NewChannelController(channelService)
	commentController := controllersHTTP.NewCommentHandler(commentService)
	postController := controllersHTTP.NewPostHandler(postService)
	subscriptionController := controllersHTTP.NewSubscriptionController(subscriptionService)
	authController := controllersHTTP.NewAuthController(authUsecase)
	profileController := controllersHTTP.NewProfileController(profileUsecase)
	chatsController := controllersHTTP.NewChatController(chatUsecase)
	wsController := controllersWS.NewWSController(chatUsecase)

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
	profileGroup := router.Group("/profiles", authMiddleware.JWTAuth())
	{
		profileGroup.GET("/", profileController.GetProfiles)
		profileGroup.GET("/me", profileController.GetMyProfile)
		profileGroup.GET("/:userID", profileController.GetProfile)
		profileGroup.PUT("/me", profileController.UpdateProfile)
		// profileGroup.PUT("/:id", profileController.UpdateProfile)
	}
	// CHAT API
	chatGroup := router.Group("/chats")
	{
		chatGroup.GET("/", chatsController.GetChats)
		chatGroup.POST("/", chatsController.CreateChat)
	}
	// WEBSOCKET API
	wsGroup := router.Group("/websockets")
	{
		wsGroup.GET("/", wsController.ChatController)
	}

	return router
}
