package routes

import (
	controllers "go_jwt_samplev2/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/register", controllers.Register)
}
