package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"   
  "carlamissiona/golang-barbers/app/controllers"
    "carlamissiona/golang-barbers/pkg/router"
    "log"  
	_ "github.com/lib/pq"  
  "github.com/gofiber/fiber/v2/middleware/etag"
  _ "github.com/jozsefsallai/fiber-bearer-token"
  _ "carlamissiona/golang-barbers/pkg/database"
)   

 
func NewApplication(modeparam string) *fiber.App {
  if modeparam == "api" {
    
        log.Println("Init Db!") 
        controllers.Initcontroller()  
        engine := html.New("./templates", ".html")
        
        app := fiber.New(fiber.Config{Views: engine})
        // app := fiber.New()
        app.Use(recover.New())  
        app.Use(logger.New())
        log.Println("API ARTICLES")  
        log.Println( app.Use(etag.New()) )
 

        web := app.Group("v1")
        web.Get("/articles", controllers.GetApi_Articles)
        web.Get("/users", controllers.GetApi_Users)
        web.Get("/maps", controllers.GetApi_Maps)  

        
        return app

    }else{

        engine := html.New("./templates", ".html")
        
        app := fiber.New(fiber.Config{Views: engine})
        app.Static("/", "./assets")
        log.Println("main monolith") 
        
        log.Println("main monolith") 
        app.Use(recover.New()) 
        app.Use(logger.New())
        
        app.Get("/dashboard", monitor.New())
        var r router.Router = nil
        r = router.NewHttpRouter()
        r.InstallRouter(app)
        return app

  }
  
}

 