package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"

	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/williamlim16/backend-trash/database"
	"github.com/williamlim16/backend-trash/models"
)

func CreateTrash(c *fiber.Ctx) error {
	var trash models.Trash

	if err := c.BodyParser(&trash); err != nil {
		fmt.Println("Unable to parse body")
		fmt.Println(err)
	}

	file, errFile := c.FormFile("image")

	if errFile != nil {
		fmt.Println("Unable to parse image")
	}

	now := time.Now()

	filename := strconv.Itoa(now.Day()) + file.Filename
	filepath := "/public/images/" + filename
	trash.Image = filepath

	errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))

	if errSaveFile != nil {
		log.Println("Save error")
	}

	if err := database.DB.Create(&trash).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Trash successfuly added",
	})
}

func GetTrash(c *fiber.Ctx) error {
	var getTrash []models.Trash
	database.DB.Preload("TrashCan").Find(&getTrash)

	return c.JSON(fiber.Map{
		"data": getTrash,
	})
}

func UpdateTrash(c *fiber.Ctx) error {
	trashID, _ := strconv.Atoi(c.Params("id"))
	trash := models.Trash{
		ID: uint(trashID),
	}

	if err := c.BodyParser(&trash); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&trash).Updates(trash)
	return c.JSON(fiber.Map{
		"message": "Trash updated successfuly",
	})
}

func DeleteTrash(c *fiber.Ctx) error {

	trashID, _ := strconv.Atoi(c.Params("id"))
	trash := models.Trash{
		ID: uint(trashID),
	}

	if err := c.BodyParser(&trash); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Delete(&trash)
	return c.JSON(fiber.Map{
		"message": "Trash " + strconv.Itoa(trashID) + "successfuly deleted.",
	})
}
