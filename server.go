package main

import (
	"log"

	c "github.com/baelorswift/home/controllers"
	m "github.com/baelorswift/home/models"
	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/jinzhu/configor"
	eztemplate "github.com/michelloworld/ez-gin-template"
	"gopkg.in/gin-gonic/gin.v1"
)

var config = struct {
	m.Context

	DSN string `json:"dsn"`
}{}

func main() {
	configor.Load(&config, "config/app.json")
	raven.SetDSN(config.DSN)

	// Create template renderer
	render := eztemplate.New()
	render.TemplatesDir = "views/"

	// Create controller context
	context := &m.Context{
		Address:    config.Address,
		APIAddress: config.APIAddress,
	}

	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	r.Static("/static", "./static/")
	r.HTMLRender = render.Init()

	// Innit Controllers
	c.NewHomeController(r, context)

	log.Fatal(r.Run(config.Address))
}
