package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/long2ice/fibers"
	"github.com/long2ice/fibers/security"
)

func main() {
	app := fibers.New(NewSwagger(), fiber.Config{})
	app.Use(
		logger.New(),
		recover.New(),
		cors.New(),
	)
	subApp := fibers.New(NewSwagger(), fiber.Config{})
	subApp.Get("/noModel", noModel)
	app.Mount("/sub", subApp)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))
	queryGroup := app.Group("/query", fibers.Tags("Query"))
	queryGroup.Get("/list", queryList)
	queryGroup.Get("/:id", queryPath)
	queryGroup.Delete("", query)

	app.Get("/noModel", noModel)

	formGroup := app.Group("/form", fibers.Tags("Form"), fibers.Security(&security.Bearer{}))
	formGroup.Post("/encoded", formEncode)
	formGroup.Put("", body)
	formGroup.Post("/file", file)

	log.Fatal(app.Listen(":8080"))
}
