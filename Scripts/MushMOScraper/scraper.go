package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
) 

type MushroomInstance struct {
	url string
}

var originalURLs = make(map[string]string)

func main() { 

	var mushroomInstances []MushroomInstance
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
	)

	var authenticityToken string

	// Handle the login page
	c.OnHTML("form[action='/account/login/new']", func(e *colly.HTMLElement) {
		if authenticityToken == "" {
			authenticityToken = e.ChildAttr("input[name='authenticity_token']", "value")
			fmt.Println("Authenticity Token:", authenticityToken)

			// Perform login
			err := c.Post("https://mushroomobserver.org/account/login/new", map[string]string{
				"login":             "**",
				"password":          "**",
				"authenticity_token": authenticityToken,
			})
			if err != nil {
				log.Fatal("Login failed:", err)
			}
		} else {
			e.Request.Visit("https://mushroomobserver.org/names?by=num_views&page=1&q=1pMfs")
		}
	})

	c.OnHTML("table", func(e *colly.HTMLElement) {
        mushroomInstance := MushroomInstance{
            url: e.ChildAttr("a", "href"),
        }
        mushroomInstances = append(mushroomInstances, mushroomInstance)
    })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Received response from", r.Request.URL)
		if r.Request.URL.String() == "https://mushroomobserver.org/account/login/new" && r.StatusCode == http.StatusOK {
			fmt.Println("On login page, looking for form...")
		} else {
			fmt.Println("Handling response for URL:", r.Request.URL)
		}
	})
	
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	c.Visit("https://mushroomobserver.org/names?by=num_views&page=1&q=1pMfs")

	c.Wait()
	  

	file, err := os.Create("products.csv") 
	if err != nil { 
		log.Fatalln("Failed to create output CSV file", err) 
	} 
	defer file.Close() 

	writer := csv.NewWriter(file)

	headers := []string{ 
		"url", 

	} 
	
	writer.Write(headers) 
	 

	for _, mushroomInstance  := range mushroomInstances { 
		record := []string{ 
			mushroomInstance.url,  
		} 
 
		writer.Write(record) 
	} 
	defer writer.Flush()

}