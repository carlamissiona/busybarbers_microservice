package router

import (
	"carlamissiona/golang-barbers/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type Router interface {
	InstallRouter(app *fiber.App)
}

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {

	controllers.Initcontroller()

	web := app.Group("", cors.New(), csrf.New())
	web.Get("/", controllers.RenderHome)
	web.Get("/services", controllers.RenderServices)
	web.Get("/about", controllers.RenderAbout)
	web.Get("/transaction", controllers.RenderPayment)
	web.Get("/transaction-paid", controllers.RenderPaid)
	web.Get("/contact", controllers.RenderContact)
	web.Get("/contact-submit", controllers.RenderContactSubmit)

}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
