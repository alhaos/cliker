package webService

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

const (
	getDataEndpoint      = "/api/data"
	ajaxTestPageEndpoint = "/"
)

//go:embed templates/* resource/css/*
var templateFS embed.FS

// getData Handler from AJAX request
func (s *service) getData(context *gin.Context) {
	data := map[string]interface{}{
		"message": "Hello from server",
		"status":  "success",
	}
	context.JSON(http.StatusOK, data)
}

// registerRoutes register routs
func (s *service) registerRoutes(engine *gin.Engine) {
	engine.GET(getDataEndpoint, s.getData)
	engine.GET(ajaxTestPageEndpoint, s.ajaxController)
}

// ajaxController controllers for ajax test page
func (s *service) ajaxController(context *gin.Context) {
	context.HTML(http.StatusOK, "ajax.html", nil)
}

// loadTemplates load embed templates
func (s *service) loadTemplates(engine *gin.Engine) error {

	tmpl, err := template.ParseFS(templateFS, "templates/*")
	if err != nil {
		return err
	}

	engine.SetHTMLTemplate(tmpl)

	return nil
}
