package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/o8x/spirit/app"
	menu2 "github.com/o8x/spirit/app/menu"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	application := app.New()

	defaultMenu := menu.NewMenu()
	defaultMenu.Append(menu.AppMenu())
	defaultMenu.Append(menu.EditMenu())

	err := wails.Run(&options.App{
		Title:  "",
		Width:  350,
		Height: 350,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		AlwaysOnTop:      menu2.AlwaysOnTop,
		Menu:             defaultMenu,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup: func(ctx context.Context) {
			application.Startup(ctx, defaultMenu)
		},
		Bind: []interface{}{
			application,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
			},
			Appearance:           mac.NSAppearanceNameVibrantLight,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Spirit",
				Message: "© 2023 Alex",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
