package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamlim16/backend-trash/controller"
	"github.com/williamlim16/backend-trash/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/login", controller.Login)
	app.Post("/api/register", controller.Register)
	app.Use(middleware.IsAuthenticate)
	app.Get("/api/trashcans", controller.GetTrashCan)
	app.Post("/api/trashcan", controller.CreateTrashCan)
	app.Put("/api/trashcan/:id", controller.UpdateTrashCan)
	app.Delete("/api/trashcan/:id", controller.DeleteTrashCan)
	app.Get("/api/trash", controller.GetTrash)
	app.Post("/api/trash", controller.CreateTrash)
	app.Put("/api/trash/:id", controller.UpdateTrash)
	app.Delete("/api/trash/:id", controller.DeleteTrash)
}
