package main

import (
<<<<<<< HEAD
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	config "github.com/mochammadshenna/aplikasi-po/configs"
	"github.com/mochammadshenna/aplikasi-po/internal/app/database"
	"github.com/mochammadshenna/aplikasi-po/internal/app/middleware"
	"github.com/mochammadshenna/aplikasi-po/internal/controller"
	"github.com/mochammadshenna/aplikasi-po/internal/repository"
	"github.com/mochammadshenna/aplikasi-po/internal/routes"
	"github.com/mochammadshenna/aplikasi-po/internal/service"
	"github.com/mochammadshenna/aplikasi-po/internal/state"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
	"github.com/mochammadshenna/aplikasi-po/internal/util/logger"
)

func main() {
	config.Init(state.App.Environment)
	logger.Init()

	validate := validator.New()
	db := database.NewDb()

	purchaseRepository := repository.NewPurchaseRepository()
	purchaseService := service.NewPurchaseOrderService(purchaseRepository, db, validate)
	purchaseController := controller.NewPurchaseOrderController(purchaseService)
	router := routes.NewRouter(purchaseController)

	host := fmt.Sprintf("%s:%d", config.Get().Server.Host, config.Get().Server.Port)
	fmt.Printf("Server running on host:%d \n", config.Get().Server.Port)

	server := http.Server{
		Addr:    host,
		Handler: middleware.MultipleMiddleware(middleware.NewHttpMiddleware(router)),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
=======
	"log"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/mochammadshenna/aplikasi-po/internal/app"
	"github.com/mochammadshenna/aplikasi-po/internal/controller"
	"github.com/mochammadshenna/aplikasi-po/internal/repository"
	"github.com/mochammadshenna/aplikasi-po/internal/service"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize DB connection
	db, err := app.NewDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Get absolute path for templates
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	templatePath := filepath.Join(currentDir, "templates")

	// Initialize template engine
	engine := html.New("./templates", ".html")
	engine.Layout("layout")
	engine.Reload(true)
	engine.Debug(true)

	// Initialize your service and controller
	purchaseOrderRepository := repository.NewPurchaseRepository()
	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepository, db, validator.New())
	purchaseOrderController := controller.NewPurchaseOrderController(purchaseOrderService)

	// Create fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Root route (Login page)
	app.Get("/", purchaseOrderController.HandleLoginPage)

	// Login routes
	app.Post("/login", purchaseOrderController.Login)
	app.Get("/auth/google", purchaseOrderController.GoogleLogin)

	// Protected routes (you should add authentication middleware)
	app.Get("/dashboard", purchaseOrderController.HandleDashboardPage)

	// Purchase Orders routes
	app.Get("/purchase-orders", purchaseOrderController.HandlePurchaseOrdersPage)
	app.Get("/purchase-orders/form", purchaseOrderController.HandlePurchaseOrderForm)
	app.Get("/purchase-orders/form/:id", purchaseOrderController.HandlePurchaseOrderForm)

	// API routes
	api := app.Group("/api")
	{
		api.Get("/purchase", purchaseOrderController.HandleGetPurchaseOrders)
		api.Get("/purchase/:id", purchaseOrderController.HandleGetPurchaseOrderById)
		api.Post("/purchase", purchaseOrderController.HandleSavePurchaseOrder)
		api.Put("/purchase/:id", purchaseOrderController.HandleUpdatePurchaseOrder)
		api.Delete("/purchase/:id", purchaseOrderController.HandleDeletePurchaseOrder)

		// Factory routes
		api.Get("/factories/production", purchaseOrderController.HandleGetProductionFactories)
		api.Get("/factories/finishing", purchaseOrderController.HandleGetFinishingFactories)
	}

	// Add a test route
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString(`
			<html>
				<head><title>Test</title></head>
				<body>
					<h1>Test Page</h1>
					<p>If you can see this, the server is working.</p>
				</body>
			</html>
		`)
	})

	// Serve static files
	app.Static("/static", "./static")

	// Debug route
	app.Get("/debug-template-content", func(c *fiber.Ctx) error {
		files := make(map[string]string)

		// Read all template files
		err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".html" {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				relPath, _ := filepath.Rel(templatePath, path)
				files[relPath] = string(content)
			}
			return nil
		})

		return c.JSON(fiber.Map{
			"templateDir": templatePath,
			"files":       files,
			"error":       err,
		})
	})

	// Add the logout route
	app.Get("/logout", purchaseOrderController.HandleLogout)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
}
