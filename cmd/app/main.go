package main

import (
	"inspiration-tech-case/configuration"
	"inspiration-tech-case/internal/api"
	"inspiration-tech-case/pkg/db"
	_ "inspiration-tech-case/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/swaggo/fiber-swagger"
)

// @title Inspiration Tech Case
// @version 1.0.0
// @description Web API for an issuer bank 
// @termsOfService http://swagger.io/terms/
// @host localhost:9999
// @BasePath /
// @schemes http https
func main() {
	config := configuration.New()
	database := db.NewDatabase(config)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	api.SetupRoutes(app, database, config)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	err := app.Listen(config.Get("SERVER.PORT"))
	if err != nil {
		panic(err)
	}

}
