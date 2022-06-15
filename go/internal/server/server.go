package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"shima/internal/route"
)

const (
	server        = "127.0.0.1:7331"
	logging       = false
	loggingFormat = true
)

func Run() {
	ex, _   := os.Executable()
	devMode := os.Getenv("__GO_DEV__") == "1"

	if !loggingFormat {
		log.SetFlags(0)
	}

	log.Println(" ")
	log.Println("  [########################################################]")
	log.Println("  ♥                 (^.^)/   shima   \\(^.^)                ♥")
	log.Println(" ")
	log.Println(" ", "dirname:", filepath.Dir(ex))
	log.Println(" ", "server : http://"+server+"/img")
	log.Println(" ", "dev    :", devMode)
	log.Println(" ")
	log.Println(" ", `☕️ Server started!`)
	log.Println(" ", "🚀 Waits for incoming requests...")
	log.Println(" ")

	if !logging {
		log.SetOutput(ioutil.Discard)
		log.SetFlags(0)
	}

	mux := http.NewServeMux()
	route.HandleFunc(mux, "/img/", true, route.HandlerSvg)
	route.HandleFunc(mux, "/css/", devMode, route.HandlerCss)
	route.HandleFunc(mux, "/htm/", devMode, route.HandlerHtm)
	route.HandleFunc(mux, "/", true, route.HandlerDefault)
	err := http.ListenAndServe(server, mux)

	if err != nil {
		log.SetOutput(os.Stderr)
		log.Println("Unable to start the server!")
		log.Fatal(err)
	}
}
