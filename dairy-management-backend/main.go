package main

import (
	"dairy-management-backend/config"
	"dairy-management-backend/controllers"
	"dairy-management-backend/middleware"
	repositories "dairy-management-backend/repositories"
	"dairy-management-backend/routes"
	"dairy-management-backend/usecases"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database connection
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Run database migrations
	if err := config.MigrateDB(); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	// Create a new Fiber app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Allow frontend
		AllowMethods:     "GET,POST,PUT,DELETE",   // Allowed methods
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true, // Allow cookies and authentication headers
	}))
	// Initialize repositories
	userRepo := repositories.NewUserRepository(config.DB)
	cowRepo := repositories.NewCowRepository(config.DB) // ✅ Add Cow Repository

	// Initialize use cases
	userUC := usecases.NewUserUseCase(userRepo)
	cowUC := usecases.NewCowUseCase(cowRepo) // ✅ Add Cow Use Case
	authUC := usecases.NewAuthUseCase(userRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userUC)
	cowController := controllers.NewCowController(cowUC) // ✅ Add Cow Controller
	authController := controllers.NewAuthController(authUC, userUC)
	// Middleware
	app.Use(middleware.LoggerMiddleware)

	// Register API routes
	routes.RegisterUserRoutes(app, userController)
	routes.RegisterCowRoutes(app, cowController) // ✅ Add Cow Routes
	routes.RegisterAuthRoutes(app, authController)
	// Start server
	port := ":8081"
	fmt.Printf("✅ Server running on port %s\n", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
