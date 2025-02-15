package router

import (
	"fmt"
	"log"

	"github.com/gene-qxsi/Flexive/internal/api/handlers"
	"github.com/gene-qxsi/Flexive/internal/api/repositories"
	"github.com/gene-qxsi/Flexive/internal/api/services"
	"github.com/gene-qxsi/Flexive/internal/storage"
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	const op = "internal/api/router/router.go/InitRouter()"

	router := chi.NewRouter()

	storage, err := storage.NewStorage()
	if err != nil {
		err = fmt.Errorf("❌ РОУТЕР-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Fatalln(err)
	}

	userRepo := repositories.NewUserRepo(storage)
	channelRepo := repositories.NewChannelRepo(storage)
	commentRepo := repositories.NewCommentRepo(storage)
	postRepo := repositories.NewPostRepo(storage)
	reactionRepo := repositories.NewReactionRepo(storage)
	subscriptionRepo := repositories.NewSubscriptionRepo(storage)

	userService := services.NewUserService(userRepo)
	channelService := services.NewChannelService(channelRepo)
	commentService := services.NewCommentService(commentRepo)
	postService := services.NewPostService(postRepo)
	reactionService := services.NewReactionService(reactionRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	userHandler := handlers.NewUserHandler(userService)
	channelHandler := handlers.NewChannelHandler(channelService)
	commentHandler := handlers.NewCommentHandler(commentService)
	postHandler := handlers.NewPostHandler(postService)
	reactionHandler := handlers.NewReactionHandler(reactionService)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)

	// USERS API
	router.Get("/users", userHandler.GetUsers)
	router.Get("/users/{id}", userHandler.GetUser)
	router.Post("/users", userHandler.CreateUser)
	router.Patch("/users/{id}", userHandler.UpdateUser)
	router.Delete("/users/{id}", userHandler.DeleteUser)
	// CHANNELS API
	router.Get("/channels", channelHandler.GetChannels)
	router.Get("/channels/{id}", channelHandler.GetChannel)
	router.Post("/channels", channelHandler.CreateChannel)
	router.Patch("/channels/{id}", channelHandler.UpdateChannel)
	router.Delete("/channels/{id}", channelHandler.DeleteChannel)
	// COMMENTS API
	router.Get("/comments", commentHandler.GetComments)
	router.Get("/comments/{id}", commentHandler.GetComment)
	router.Post("/comments", commentHandler.CreateComment)
	router.Patch("/comments/{id}", commentHandler.UpdateComment)
	router.Delete("/comments/{id}", commentHandler.DeleteComment)
	// POSTS API
	router.Get("/posts", postHandler.GetPosts)
	router.Get("/posts/{id}", postHandler.GetPost)
	router.Post("/posts", postHandler.CreatePost)
	router.Patch("/posts/{id}", postHandler.UpdatePost)
	router.Delete("/posts/{id}", postHandler.DeletePost)
	// REACTIONS API
	router.Get("/reactions", reactionHandler.GetReactions)
	router.Get("/reactions/{userID}/{postID}", reactionHandler.GetReaction)
	router.Post("/reactions", reactionHandler.CreateReaction)
	router.Patch("/reactions/{userID}/{postID}", reactionHandler.UpdateReaction)
	router.Delete("/reactions/{userID}/{postID}", reactionHandler.DeleteReaction)
	// SUBSCRIPTIONS API
	router.Get("/subscriptions", subscriptionHandler.GetSubscriptions)
	router.Get("/subscriptions/{userID}/{channelID}", subscriptionHandler.GetSubscription)
	router.Post("/subscriptions", subscriptionHandler.CreateSubscription)
	router.Patch("/subscriptions/{userID}/{channelID}", subscriptionHandler.UpdateSubscription)
	router.Delete("/subscriptions/{userID}/{channelID}", subscriptionHandler.DeleteSubscription)

	return router
}
