package app

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	menu2 "github.com/o8x/spirit/app/menu"
)

type Application struct {
	ctx   context.Context
	menus *menu.Menu
}

func New() *Application {
	return &Application{}
}

func (a *Application) Startup(ctx context.Context, menus *menu.Menu) {
	a.ctx = ctx
	a.menus = menus

	menus.Append(menu2.NewImageMenu(ctx))
	menus.Append(menu2.NewWindowMenu(ctx))

	runtime.MenuSetApplicationMenu(ctx, menus)
	runtime.MenuUpdateApplicationMenu(ctx)
}

func (a *Application) ResizeMainWindow(width, height int) {
	if width > 1000 {
		width = 1000
	}
	runtime.WindowSetSize(a.ctx, width, height)
}

func (a *Application) DefaultSizeMainWindow() {
	runtime.WindowSetSize(a.ctx, 350, 350)
}
