package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	
)

type MainController struct {
	beego.Controller
}

type CarController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Student struct {
  StudentId   string
  StudentName string
  ClassName   string
}

func (c *CarController) Get() {
  carId := c.Ctx.Input.Param(":carId")
  res, err := http.Get("https://www.weblio.jp/content/"+carId)
	if err != nil {
		log.Fatal(err)
	}
  defer res.Body.Close()
 /* sitemap, err := ioutil.ReadAll(res.Body)
  if err != nil {
		log.Fatal(err)
  }
  sitemap = nil
  fmt.Printf("%s", sitemap)*/
  
  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n\n", doc)
  ss:= Student{}
  
  
  // Find the review items
	doc.Find(".midashigo").Each(func(i int, s *goquery.Selection) {
	// For each item found, get the band and title
	band := s.Text()
	if i == 0{
		ss.StudentName = band
	
	}
	//title := s.Find("i").Text()
	fmt.Printf("Review %d: %s \n", i, band)
	})
	
	
  
  
  c.Data["json"] = ss
  
  c.ServeJSON()
}
