package controllers

import (  
	"github.com/gofiber/fiber/v2" 
)


func GetApi_Articles(c *fiber.Ctx) error {
 
	return c.SendString("GetApi_Articles")
}

func GetApi_Users(c *fiber.Ctx) error {
 
	return c.SendString("GetApi_Users")
}

func GetApi_Maps(c *fiber.Ctx) error {
 
	return c.SendString("GetApi_Maps")
}
