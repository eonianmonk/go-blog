package http

import (
	"fmt"

	"github.com/eonianmonk/go-blog/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/headers", handlers.GetHeaders)

	api.Get(fmt.Sprintf("/post/:%s", handlers.PostIdParameter), handlers.GetPost)
	api.Post("/post", handlers.CreatePost)
	api.Put("/post", handlers.UpdatePost)
	api.Delete("/post/:%s", handlers.DeletePost)
}
