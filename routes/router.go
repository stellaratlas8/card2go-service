package routes

import (
	"card2go_service/controller/auth"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	RegisterAdmins(app)
	RegisterAuth(app)
	RegisterTours(app)
	RegisterHotels(app)
	RegisterPOI(app)
}

func RegisterAdmins(app *fiber.App) {

}

func RegisterAuth(app *fiber.App) {
	path := app.Group("/auth")
	path.Post("/", auth.HandleAuthentication)
	path.Post("/register", auth.HandleRegister)
}

func RegisterTours(app *fiber.App) {

}

func RegisterHotels(app *fiber.App) {

}

func RegisterPOI(app *fiber.App) {

}
