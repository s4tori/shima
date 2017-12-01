// Package route manages HTTP routes.
package route

import (
	"fmt"
	"log"
	"net/http"
	"ga"
	"grid"
)

type handle func(http.ResponseWriter, *http.Request, string)

// HandleFunc calls "handle" when the prefix matches the URL.
func HandleFunc(mux *http.ServeMux, prefix string, include bool, h handle) {
	if !include {
		return
	}

	fHandle := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len(prefix):]
		log.Println("----------------------")
		log.Printf("Path: %s", path)
		ga.TrackPageEvent(r)

		h(w, r, path)
	}

	mux.HandleFunc(prefix, fHandle)
}

// HandlerCss returns the grid (base64 encoding).
func HandlerCss(w http.ResponseWriter, r *http.Request, path string) {
	w.Header().Set("Content-Type", "image/svg+xml")
	grid.RenderBase64(w, path)
}

// HandlerSvg returns the grid (text encoding).
func HandlerSvg(w http.ResponseWriter, r *http.Request, path string) {
	w.Header().Set("Content-Type", "image/svg+xml")
	grid.Render(w, path)
}

// HandlerHtm returns an example (HTML page).
func HandlerHtm(w http.ResponseWriter, r *http.Request, path string) {
	fmt.Fprintln(w, `<html><body style="margin:0; pading: 0"> `)
	fmt.Fprint(w  , `<img style="width: 100%; height: 100%" src="`)
	grid.RenderBase64(w, path)
	fmt.Fprintln(w, `" />`)
	fmt.Fprintln(w, "</body></html>")
}

// HandlerDefault returns a 404 error.
func HandlerDefault(w http.ResponseWriter, r *http.Request, path string) {
	//path := r.URL.Path
	//if path == "/favicon.ico" || path == "/robots.txt" {
	//	http.NotFound(w, r)
	//	return
	//}

	http.NotFound(w, r)
}
