package main

import (
	_ "image/jpeg"

	"github.com/joho/godotenv"
	_ "github.com/textures1245/payso-check-slip-backend/docs"
	"github.com/textures1245/payso-check-slip-backend/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	log.Info("-= Start External Service =-")
	router.RouterInit(app)
	app.Listen(":4567")

}
