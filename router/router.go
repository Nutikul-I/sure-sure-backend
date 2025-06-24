package router

import (
	// "os"
	"context"
	"strings"
	"time"

	"github.com/robfig/cron"
	"github.com/textures1245/payso-check-slip-backend/controller"
	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/repository"
	"github.com/textures1245/payso-check-slip-backend/service"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func RouterInit(app *fiber.App) *fiber.App {

	api := app.Group("/", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	user_service := controller.NewUserController(service.NewUserService(handler.NewUserHandler()))
	package_service := controller.NewPackageController(service.NewPackageService(handler.NewPackageHandler()))
	order_package_service := controller.NewOrderPackageController(service.NewOrderPackageService(handler.NewOrderPackageHandler()))
	rooms_service := controller.NewRoomsController(service.NewRoomsService(handler.NewRoomsHandler()))
	banks_service := controller.NewBanksController(service.NewBanksService(handler.NewBanksHandler()))
	transaction_service := controller.NewTransactionController(service.NewTransactionService(handler.NewTransactionHandler()))
	log_service := controller.NewLogController(service.NewLogService(handler.NewLogHandler()))

	// ------------------------------------------------------------------------------------------------------------------------------

	api.Post("/api/v1/login", user_service.GetOrCreateUser)

	user := api.Group("/api/v1/user")
	user.Get("/get", user_service.GetUserAll)
	user.Get("/get/:id", user_service.GetUserByID)
	user.Post("/create", user_service.CreateUser)
	user.Put("/update", user_service.UpdateUser)
	user.Delete("/delete/:id", user_service.DeleteUser)
	user.Get("/category/get", user_service.GetCategoryAll)

	packages := api.Group("/api/v1/package")
	packages.Get("/get", package_service.GetPackageAll)
	packages.Get("/get/:id", package_service.GetPackageByID)
	packages.Post("/create", package_service.CreatePackage)
	packages.Put("/update", package_service.UpdatePackage)
	packages.Delete("/delete/:id", package_service.DeletePackage)

	order_packages := api.Group("/api/v1/order-package")
	order_packages.Get("/get", order_package_service.GetOrderPackageAll)
	order_packages.Get("/get/:id", order_package_service.GetOrderPackageByID)
	order_packages.Get("/get/refno/:RefNo", order_package_service.GetOrderPackageByRefNo)
	order_packages.Post("/create", order_package_service.CreateOrderPackage)
	order_packages.Put("/update", order_package_service.UpdateOrderPackage)
	order_packages.Delete("/delete/:id", order_package_service.DeleteOrderPackage)

	rooms := api.Group("/api/v1/room2")
	rooms.Get("/get", rooms_service.GetAllRooms)
	rooms.Get("/get/:id", rooms_service.GetRoomByID)
	rooms.Post("/create", rooms_service.CreateRoom)
	rooms.Put("/update", rooms_service.UpdateRoom)
	rooms.Delete("/delete/:id", rooms_service.DeleteRoom)
	rooms.Get("/howto/:id/:user_id", rooms_service.HowTo)

	banks := api.Group("/api/v1/bank2")
	banks.Get("/get", banks_service.GetAllBanks)
	banks.Get("/get/:id", banks_service.GetBankByID)
	banks.Post("/create", banks_service.CreateBank)
	banks.Put("/update", banks_service.UpdateBank)
	banks.Delete("/delete/:id", banks_service.DeleteBank)

	transaction := api.Group("/api/v1/transaction")
	transaction.Get("/get", transaction_service.GetTransactionAll)
	transaction.Get("/get/:id", transaction_service.GetTransactionByID)
	transaction.Post("/create", transaction_service.CreateTransaction)
	transaction.Put("/update", transaction_service.UpdateTransaction)
	transaction.Delete("/delete/:id", transaction_service.DeleteTransaction)

	logs := api.Group("/api/v1/log")
	logs.Get("/get", log_service.GetLogAll)
	logs.Get("/get/:id", log_service.GetLogByID)
	logs.Post("/create", log_service.CreateLog)
	logs.Put("/update", log_service.UpdateLog)
	logs.Delete("/delete/:id", log_service.DeleteLog)

	api.Get("/api/v1/healthcheck", func(c *fiber.Ctx) error {
		conn := repository.ConnectDB()
		ctx := context.Background()
		err := conn.PingContext(ctx)
		if err != nil {
			log.Error("HEALTH_CHECK_FAILED:" + err.Error())
			log.Fatal("HEALTH_CHECK_FATAL: Closing the application")
		}
		return c.Status(fiber.StatusOK).JSON("version: 1.0.0")
	})

	// Start the cron job
	go StartCronJob()

	return app
}

func StartCronJob() {
	c := cron.New()

	// Run every minute (adjust as needed)
	err := c.AddFunc("*/60 * * * *", func() {
		log.Println("Running scheduled task at:", time.Now().Format(time.RFC3339))
		order_package_service := controller.NewOrderPackageController(service.NewOrderPackageService(handler.NewOrderPackageHandler()))

		order_package_service.CheckPaymentTransaction()
		// Your scheduled task logic here (e.g., clean up logs, fetch data, etc.)
	})

	if err != nil {
		log.Fatalf("Failed to start cron job: %v", err)
	}

	c.Start()
}
