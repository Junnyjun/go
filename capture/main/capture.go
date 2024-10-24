package main

import (
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	_ "github.com/asticode/go-astilectron-bootstrap"
	"log"
)

var AppName = "Drag to Capture"
var AppVersion = "1.0.0"

func init() {

}
func main() {
	// Astilectron 앱 생성
	var application, err = astilectron.New(nil, astilectron.Options{
		AcceptTCPTimeout:  3000,
		AppName:           "JUNNYLAND CAPTURE",
		BaseDirectoryPath: "./",
	})
	if err != nil {
		log.Fatal(err)
	}

	var window, _ = application.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})
	application.Start()

	println("Application started")
	window.Close()
	application.Close()
}
