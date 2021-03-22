package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"raspberrypi-car/wheel"
)

type WebController struct {
	CarWheel  *wheel.CarWheel
	Templates *template.Template
}

func NewWebController(carWheel *wheel.CarWheel) *WebController {
	templates := template.Must(template.New("").Funcs(nil).ParseGlob("templates/*"))
	return &WebController{
		CarWheel:  carWheel,
		Templates: templates,
	}
}

func (c *WebController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "Welcome to RaspberryPi Car.")
	case "/control":
		c.CarWheel.Forward()
		w.Header().Set("Content-Type", "text/html")
		_ = c.Templates.ExecuteTemplate(w, "index.tmpl", map[string]interface{}{
			"name": "jtf",
		})
	}
}
