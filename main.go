package main

import (
	db "go_jwt_samplev2/db"
	routes "go_jwt_samplev2/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()
	app := fiber.New()
	routes.Setup(app)
	


	app.Listen(":3000")
}
