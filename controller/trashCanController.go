package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/williamlim16/backend-trash/database"
	"github.com/williamlim16/backend-trash/models"
)

func CreateTrashCan(c *fiber.Ctx) error {
	var trashCan models.TrashCan

	if err := c.BodyParser(&trashCan); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&trashCan).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Trash can successfuly added",
	})
}

func GetTrashCan(c *fiber.Ctx) error {
	var getTrashCan []models.TrashCan
	database.DB.Find(&getTrashCan)

	return c.JSON(fiber.Map{
		"data": getTrashCan,
	})
}

func UpdateTrashCan(c *fiber.Ctx) error {
	trashCanID, _ := strconv.Atoi(c.Params("id"))
	trashCan := models.TrashCan{
		ID: uint(trashCanID),
	}

	if err := c.BodyParser(&trashCan); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&trashCan).Updates(trashCan)
	return c.JSON(fiber.Map{
		"message": "Trash updated successfuly",
	})
}

func DeleteTrashCan(c *fiber.Ctx) error {
	trashCanID, _ := strconv.Atoi(c.Params("id"))
	trashCan := models.TrashCan{
		ID: uint(trashCanID),
	}
	if err := c.BodyParser(&trashCan); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Delete(&trashCan)
	return c.JSON(fiber.Map{
		"message": "Trash can " + strconv.Itoa(trashCanID) + " successfuly deleted.",
	})
}
