// @title SureSure Public API
// @version 1.0
// @description Public API for SureSure services
// @BasePath /api/v1
// @schemes http https
package main

import (
	_ "image/jpeg"

	"github.com/joho/godotenv"
	"github.com/textures1245/payso-check-slip-backend/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/textures1245/payso-check-slip-backend/docs"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file ")
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Swagger UI Route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// ReDoc UI Route
	app.Get("/docs", func(c *fiber.Ctx) error {
		html := `<!DOCTYPE html>
		<html>
		  <head>
		    <meta charset="utf-8" />
		    <meta name="viewport" content="width=device-width, initial-scale=1">
		    <title>SureSure Public API Docs</title>
		    <style>body { margin: 0; padding: 0; } #redoc { height: 100vh; }</style>
		  </head>
		  <body>
		    <div id="redoc"></div>
		    <script src="https://cdn.jsdelivr.net/npm/redoc@latest/bundles/redoc.standalone.js"></script>
		    <script>
		      Redoc.init('/swagger/doc.json', { expandResponses: '200,201' }, document.getElementById('redoc'))
		    </script>
		  </body>
		</html>`
		return c.Type("html").SendString(html)
	})

	log.Info("-= Start External Service =-")
	router.RouterInit(app)
	app.Listen(":4567")

}
