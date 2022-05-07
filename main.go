package main

import (
	"fmt"
	"html/template"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/kbinani/screenshot"
)

func main() {

	fmt.Println("Start :80")

	http.HandleFunc("/", handler)
	http.HandleFunc("/screen.jpg", handler_screen)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
		panic(err)
	}

	display1, err := strconv.Atoi(r.FormValue("display"))
	if err != nil {
		display1 = 0
	}

	param := struct {
		Display int
	}{Display: display1}

	t.ExecuteTemplate(w, "index", param) //{{.}}

}

func handler_screen(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Hello")

	display, err := strconv.Atoi(r.FormValue("display"))
	if err != nil {
		display = 0
	}

	if display >= screenshot.NumActiveDisplays() {
		display = 0
	}

	bounds := screenshot.GetDisplayBounds(display)

	img, err := screenshot.CaptureRect(bounds) //*image.RGBA
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)

}
