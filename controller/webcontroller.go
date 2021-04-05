package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"raspberrypi-car/camera"
	"raspberrypi-car/wheel"
)

type WebController struct {
	Cam       *camera.Camera
	CarWheel  *wheel.CarWheel
	Templates *template.Template
}

func NewWebController(carWheel *wheel.CarWheel, cam *camera.Camera) *WebController {
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
	case "/forward":
		c.CarWheel.Forward()
		//w.Header().Set("Content-Type", "text/html")
		//_ = c.Templates.ExecuteTemplate(w, "index.tmpl", map[string]interface{}{
		//	"name": "jtf",
		//})
		fmt.Fprintf(w, "forward")
	case "/backward":
		c.CarWheel.Backward()
		fmt.Fprintf(w, "backward")
	case "/stopped":
		c.CarWheel.Stopped()
		fmt.Fprintf(w, "stopped")
	case "/parallel_left":
		c.CarWheel.ParallelLeft()
		fmt.Fprintf(w, "parallel left")
	case "/parallel_right":
		c.CarWheel.ParallelRight()
		fmt.Fprintf(w, "parallel right")
	}
}
