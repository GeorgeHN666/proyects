package main

import (
	"github.com/gofiber/fiber/v2"
	tp "github.com/gofiber/template/html"
)

func main() {
	// Template Html
	engine := tp.New("./Views", ".html")
	// Reload Html
	engine.Reload(true)
	// Init App
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// GET Route
	app.Get("/", func(c *fiber.Ctx) error {
		testMessage := "sample Text"
		return c.Render("index", fiber.Map{
			"sample": testMessage,
		})
	})
	// POST Route
	app.Post("/", func(c *fiber.Ctx) error {
		postmessage := "From Post"
		return c.Render("index", fiber.Map{
			"postsample": postmessage,
		})
	})
	// Listen
	app.Listen(":8000")
}

/*
	***Were Gonna Built A Facebook Clone***
	This Will Be A Simple Clone Of The Profile Of The Website Facebook
	We Will Use Go With the Package fiber as a framework, We also are going to use html/templates by fiber, Html5 With Bootstrap

	+++Packages We Will Use:
	-fiber
	-html templates
*/
