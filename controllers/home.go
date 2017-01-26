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

// NewHomeController ..
func NewHomeController(r *gin.Engine, c *m.Context) {
	ctrl := new(HomeController)
	ctrl.context = c

	r.GET("", ctrl.Index)
	r.GET("index", ctrl.Index)
}
