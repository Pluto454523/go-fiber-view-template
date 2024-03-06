package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	// Pass the engine to Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Setup route
	app.Get("/", renderTemplate)

	port := ":" + os.Getenv("PORT")
	app.Listen(port)

}

func renderTemplate(c *fiber.Ctx) error {
	// Render the template with variable data
	return c.Render("template", fiber.Map{
		"Name": "World",
	})
}
