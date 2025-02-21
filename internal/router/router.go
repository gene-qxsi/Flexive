package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gene-qxsi/Flexive/internal/delivery/http/controllers"
	"github.com/gene-qxsi/Flexive/internal/middleware"
	"github.com/gene-qxsi/Flexive/internal/repository"
	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gene-qxsi/Flexive/internal/storage"
	"github.com/gene-qxsi/Flexive/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func InitRouter() *gin.Engine {
	const op = "internal/router/router.go/InitRouter()"

	router := gin.Default()

	storage, err := storage.NewStorage()
	if err != nil {
		err = fmt.Errorf("❌ РОУТЕР-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Fatalln(err)
	}

	redisClient := redis.NewClient(&redis.Options{})
	hasher := services.NewSHA1Hasher(os.Getenv("SALT"))

	userRepo := repository.NewUserRepo(storage)
	channelRepo := repository.NewChannelRepo(storage)
	commentRepo := repository.NewCommentRepo(storage)
	postRepo := repository.NewPostRepo(storage)
	reactionRepo := repository.NewReactionRepo(storage)
	subscriptionRepo := repository.NewSubscriptionRepo(storage)
	authRepo := repository.NewAuthRepository(redisClient)

	userService := services.NewUserService(userRepo, hasher)
	channelService := services.NewChannelService(channelRepo)
	commentService := services.NewCommentService(commentRepo)
	postService := services.NewPostService(postRepo)
	reactionService := services.NewReactionService(reactionRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	authService := services.NewAuthService(authRepo)

	authUseCase := usecase.NewAuthUseCase(userService, authService)

	userHandler := controllers.NewUserController(userService)
	channelController := controllers.NewChannelController(channelService)
	commentController := controllers.NewCommentHandler(commentService)
	postController := controllers.NewPostHandler(postService)
	reactionController := controllers.NewReactionHandler(reactionService)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService)
	authController := controllers.NewAuthController(authUseCase)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	// USERS API
	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/", userHandler.GetUsers)
		usersGroup.GET("/:id", userHandler.GetUser)
		usersGroup.POST("/", userHandler.CreateUser)
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
	}

	return router
}
