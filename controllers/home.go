package controllers

import (
	"net/http"

	m "github.com/baelorswift/home/models"

	"gopkg.in/gin-gonic/gin.v1"
)

// HomeController ..
type HomeController struct {
	context *m.Context
}

// Index ..
func (ctrl HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", gin.H{
		"title":       "Home",
		"page":        "home",
		"api_address": ctrl.context.APIAddress,
	})
}

// About ..
func (ctrl HomeController) About(c *gin.Context) {
	c.HTML(http.StatusOK, "home/about", gin.H{
		"title": "About",
		"page":  "about",
	})
}

// NewHomeController ..
func NewHomeController(r *gin.Engine, c *m.Context) {
	ctrl := new(HomeController)
	ctrl.context = c

	r.GET("", ctrl.Index)
	r.GET("index", ctrl.Index)
	r.GET("about", ctrl.About)
}
