package routes

import (
<<<<<<< HEAD
	"github.com/Levantate-Labs/sainterview-backend/auth-service/controllers"
	"github.com/Levantate-Labs/sainterview-backend/auth-service/middleware"
=======
	"github.com/AlfrinP/point_calculator/controllers"
	"github.com/AlfrinP/point_calculator/middleware"
>>>>>>> 52a2cfba8417f30f47f3a85feb3c92850e82f352
	"github.com/gofiber/fiber/v2"
)

func SetupProtectedRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/securitycheck", middleware.DeserializeUser, controllers.SecurityCheck)
}
