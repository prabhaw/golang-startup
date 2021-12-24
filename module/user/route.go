package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)



func UserRouter(route fiber.Router){
	route.Use("/get-user", func (c *fiber.Ctx) error {
		fmt.Println("Hello this is route")
		return c.Next()
	})
	route.Get("/get-user", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":true,
			"data": fiber.Map{
				"hello": "This is Routes",
			},
		})
	})
}