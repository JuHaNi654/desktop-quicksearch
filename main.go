package main

import (
	"context"
	"embed"
	"quicksearch/internal"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	desktop := NewEntries()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Quicksearch",
		AlwaysOnTop:   true,
		MaxWidth:      500,
		MaxHeight:     300,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: internal.NewFileHandler(),
		},
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
		},
		Bind: []interface{}{
		  app,	
      desktop,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
