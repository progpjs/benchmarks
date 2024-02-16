package main

import (
	"log"
)

// https://gofiber.io/

func main() {
	println("Fiber - Listening on port http://localhost:8000")

	app := fiber.New()
	app.Get("/", hello)
	log.Fatal(app.Listen(":8000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}
