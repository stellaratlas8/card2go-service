package routes

import (
	"card2go_service/controller/admin"
	"card2go_service/controller/auth"
	"card2go_service/controller/destinations"
	"card2go_service/controller/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	RegisterAdmin(app)
	RegisterAuth(app)
	RegisterDestinations(app)
}

func RegisterAdmin(app *fiber.App) {
	path := app.Group("/admin", middleware.RequireDatabase)
	path.Get("/destinations/clear", admin.HandleClearDestinations)
	path.Get("/destinations/dummy", admin.HandleCreateDummyDestinations)
}

func RegisterAuth(app *fiber.App) {
	path := app.Group("/auth", middleware.RequireDatabase)
	path.Post("/", auth.HandleAuthentication)
	path.Post("/register", auth.HandleRegister)
}

func RegisterDestinations(app *fiber.App) {
	path := app.Group("/destinations", middleware.RequireDatabase)
	path.Get("/", destinations.HandleFeed)
	path.Get("/:id", destinations.HandleDestination)

	path.Post("/:id/book", middleware.RequireAuth, destinations.HandleBook)
	path.Post("/:id/book/:pid", middleware.RequireAuth, destinations.HandleBook)
}
