package main

import (
	"embed"
	"log"

	"github.com/gsvd/goeland/internal/store"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	store, err := store.New()
	if err != nil {
		log.Fatal("main: store.New:", err)
	}
	defer func() {
		if err := store.Close(); err != nil {
			log.Println("main: store.Close:", err)
		}
	}()

	app := NewApp(store)

	err = wails.Run(&options.App{
		Title:  "Goeland",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("main: wails.Run:", err.Error())
	}
}
