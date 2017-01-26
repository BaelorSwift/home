package controllers

import (
	"net/http"

	"github.com/baelorswift/home/models"

	"gopkg.in/gin-gonic/gin.v1"
)

// DocsController ..
type DocsController struct {
	context *models.Context
}

// Index ..
func (ctrl DocsController) Index(c *gin.Context) {
	manifest := models.NewManifest()

	c.HTML(http.StatusOK, "docs/index", gin.H{
		"title":       "Docs",
		"page":        "docs",
		"api_address": ctrl.context.APIAddress,
		"manifest":    &manifest,
	})
}

// NewDocsController ..
func NewDocsController(r *gin.Engine, c *models.Context) {
	ctrl := new(DocsController)
	ctrl.context = c

	r.GET("docs", ctrl.Index)
}
