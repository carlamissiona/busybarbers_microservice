package controllers

import (
	"carlamissiona/golang-barbers/pkg/database"
	"database/sql"
	_ "database/sql"
	_ "encoding/json"
	"github.com/gofiber/fiber/v2"
	m3o "go.m3o.com"
	"go.m3o.com/email"
	"log"
	"net/http"
	"os"
	// "os/exec"
	"github.com/Jeffail/gabs"
	"io/ioutil"
	_ "reflect"
)

var db *sql.DB

func Initcontroller() {
	log.Println("InitInitCon!")
	db = database.OpenDatabase()
}

type Article struct {
	Title      string      `json:"article_title"`
	Content    string      `json:"article_content"`
	Link       string      `json:"article_link"`
	Changed_on interface{} `json:"article_change_date"`
}
type Maps struct {
	Description string      `json:"map_description"`
	Title       string      `json:"map_title"`
	Content     string      `json:"map_content"`
	Link        string      `json:"map_link"`
	Changed_on  interface{} `json:"map_change_date"`
}

func RenderPaid(c *fiber.Ctx) error {

	return c.JSON("An error occured")

}

func RenderHome(c *fiber.Ctx) error {

	log.Println("API ARTICLES1")
	db = database.OpenDatabase()
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
		err := rows.Scan(&ar.Title, &ar.Content, &ar.Link, &ar.Changed_on)
		if err != nil {
			log.Printf("Err! %v", err)
		}
		articles = append(articles, ar)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Err! %v", err)

		return c.JSON(&fiber.Map{
			"status":        500,
			"message":       "Failed to fetch articles",
			"articles":      nil,
			"total_fetched": "0",
		})
	}
	log.Println("Row! %v", articles)
	log.Println(articles)

	log.Println(articles)
	return c.Render("index", fiber.Map{
		"Articles":   articles,
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderAbout(c *fiber.Ctx) error {
	urlmcsv := os.Getenv("SVC_MAPS_URL")
	log.Println(urlmcsv)
	response, err := http.Get(urlmcsv)
	if err != nil {
		log.Printf("No response from request %v", err)

	}
	defer response.Body.Close()
	 
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("No response from request")
		fmt.Println(err)

		c.Redirect(http.StatusFound, location.RequestURI())

	}
	var objParsed interface{}
	err = json.Unmarshal(body, &objParsed)
	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Unmarshal Error from request")
		fmt.Println(err)
	}

	return c.Render("about-page", fiber.Map{
		"Title":    "About",
		"Articles": "articles",
	}, "layouts/htm")

}

func RenderServices(c *fiber.Ctx) error {

	// urlmcsv := os.Getenv("SERVICES_URL")
	// log.Println(urlmcsv)

	// curl := exec.Command("curl", "-k", "-s", "-v", urlmcsv) // this line is modified
	// out, err := curl.Output()
	// if err != nil {
	// 	log.Println("erorr", err)

	// }
	// log.Println(string(out))

	return c.Render("services", fiber.Map{
		"Title":    "Services ",
		"Articles": "articles",
	}, "layouts/htm")
}

func RenderPayment(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

// func RenderAbout(c *fiber.Ctx) error {
// 	log.Println("Higshsss!")
// 	return c.Render("index", fiber.Map{
// 		"FiberTitle": "Hello From Fiber Html Engine",
// 	}, "layouts/htm")
// }

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
