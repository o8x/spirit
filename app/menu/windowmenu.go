package menu

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/o8x/spirit/app/utils"
)

var (
	AlwaysOnTop = true
)

func NewWindowMenu(ctx context.Context) *menu.MenuItem {
	w := New("Window")

	text := w.AddText("")
	update := func() {
		text.Label = fmt.Sprintf("Always On Top: %v", AlwaysOnTop)
		runtime.MenuUpdateApplicationMenu(ctx)
	}

	w.AddSeparator()
	w.Add("Reload", func(data *menu.CallbackData) {
		runtime.WindowReload(ctx)
	})

	w.Add("Reload App", func(data *menu.CallbackData) {
		runtime.WindowReloadApp(ctx)
	})

	w.AddSeparator()
	w.Add("Default", func(data *menu.CallbackData) {
		utils.WindowSetDefaultSize(ctx)
	})

	w.Add("Maximise", func(data *menu.CallbackData) {
		runtime.WindowToggleMaximise(ctx)
	})

	w.Add("Minimise", func(data *menu.CallbackData) {
		runtime.WindowMinimise(ctx)
	})

	w.AddSeparator()
	w.Add("Always On Top", func(data *menu.CallbackData) {
		AlwaysOnTop = !AlwaysOnTop
		runtime.WindowSetAlwaysOnTop(ctx, AlwaysOnTop)
		update()
	})

	w.Add("Move To Center", func(data *menu.CallbackData) {
		runtime.WindowCenter(ctx)
	})

	update()
	return w.Build()
}
