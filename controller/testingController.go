package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	FirstName string `json:"fir" from:"wilyem"`
	LastName  string `json:"aaa" from:"ree"`
	Email     string `json:"wopp" form:"woop"`
	Password  []byte `json:"asa" form:"lol"`
}

func Register2(c *fiber.Ctx) error {
	body := struct {
		FirstName string `json:"wee"`
		LastName  string
		Email     string
		Password  string
		Phone     string
	}{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
	}
	cred := Person{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  []byte(body.Password),
	}
	log.Println(cred)

	return c.JSON(fiber.Map{
		"data": cred,
	})
}
