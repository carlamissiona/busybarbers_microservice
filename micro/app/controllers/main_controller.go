package controllers
 
import (
	"carlamissiona/golang-barbers/pkg/database"
	"database/sql"
	_ "database/sql"
	"log"
	"os" 
	"github.com/gofiber/fiber/v2"
    
	m3o "go.m3o.com"
	"go.m3o.com/email"
)

 type Article struct {
 
	Title       string `json:"article_title"`
	Content     string  `json:"article_content"` 
	Link        string  `json:"article_link"`
	Changed_on interface{} `json:"article_change_date"`
     
} 
 type Maps struct {
    Description string `json:"map_description"`
	Title       string `json:"map_title"`
	Content     string  `json:"map_content"` 
	Link        string  `json:"map_link"`
	Changed_on interface{} `json:"map_change_date"`
}

var db *sql.DB 

func Initcontroller() {   
	log.Println("InitInitCon!")
	db = database.SetupDatabase()
}


func GetApi_Articles(c *fiber.Ctx) error {
	log.Println("API ARTICLES1") 
   
	rows, err := db.Query("Select title, content, link, changed_on from public.bbr_articles")
	
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
    
	log.Println("Struct")
	// database.OpeDatabase()
	var articles []Article
    for rows.Next() {
		log.Printf("Print Next Row")
		var ar Article
		err := rows.Scan( &ar.Title, &ar.Content, &ar.Link, &ar.Changed_on )  
        if err != nil {
			log.Printf("Err! %v", err)
        } 
        articles = append(articles, ar)
    } 
 

    if err := rows.Err(); err != nil {
		log.Printf("Err! %v", err)
         
        return c.JSON(&fiber.Map{
            "status": 500,
            "message": "Failed to fetch articles",
            "articles": nil,
            "total_fetched": "0",
        });
    }
	log.Println("Row! %v", articles)
	log.Println(articles);
	 
    return c.JSON(&fiber.Map{
        "status": 200,
        "message": "Successfully fetched articles",
        "articles": articles,
        "total_fetched": len(articles),
    });

}  

func GetApi_Users(c *fiber.Ctx) error {
 
	return c.SendString("GetApi_Users")
}

func GetApi_Maps(c *fiber.Ctx) error {

    log.Println("API maps1") 
   
	rows, err := db.Query("Select content, link, title,description,  changed_on from public.bbr_maps;")
	if err != nil {
		log.Fatalln(err)
		return c.JSON( &fiber.Map{"message":"An error occured on Get Api Maps"})
	} 
    
	log.Println("Struct") 
	var maps []Maps
    for rows.Next() {
		log.Printf("Print Next Row")
		var ar Maps
		err := rows.Scan( &ar.Title, &ar.Content, &ar.Link, &ar.Description, &ar.Changed_on )  
        if err != nil {
			log.Printf("Err! %v", err)
        } 
        maps = append(maps, ar)
    } 
    
    if err := rows.Err(); err != nil {
		log.Printf("Err! %v", err)
         
        return c.JSON(&fiber.Map{
            "status": 500,
            "message": "Failed to fetch maps",
            "maps": nil,
            "total_fetched": "0",
        });
    }
	 
    return c.JSON(&fiber.Map{
        "Etag" :  c.Get("If-None-Match"),     
        "status": 200,
        "message": "Successfully fetched maps",
        "maps": maps,
        "total_fetched": len(maps),
    });
}



func RenderHome(c *fiber.Ctx) error {
	 
	 
	rows, err := db.Query("Select title, content, link, changed_on from public.bbr_articles")
	 
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
    
	log.Println("Struct")
	// database.OpeDatabase()
	var articles []Article
    for rows.Next() {
		var ar Article
		err := rows.Scan( &ar.Title, &ar.Content, &ar.Link, &ar.Changed_on )
        if err != nil {
			log.Printf("Err! %v", err)
        } 
        articles = append(articles, ar) 
    }
    if err := rows.Err(); err != nil {
		log.Printf("Err! %v", err)
    }
	log.Println("Row! %v", articles[0].Link)
	log.Println("Row! %v", articles)
	log.Println(articles);
	return c.Render("index", fiber.Map{
	   "Articles" : articles,
	   "FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderPaid(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderPayment(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderAbout(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderContact(c *fiber.Ctx) error {
	log.Println("High")

	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}
func RenderContactSubmit(c *fiber.Ctx) error {
	log.Println("From", c.Params("from"))
	api_secret := os.Getenv("M3O_APIKEY")

	client := m3o.New(api_secret)
	rsp, err := client.Email.Send(&email.SendRequest{
		To:      "missiona.carla@gmail.com <missiona.carla@gmail.com>",
		From:    "Awesome Dot Com <codetuna@protonmail.com>",
		Subject: "Email verification",
		TextBody: `Hi there,

        Please verify your email by clicking this link: $micro_verification_link`,
	})
	log.Println(rsp, err)

	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}
