package menu

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/o8x/spirit/app/utils"
)

func NewImageMenu(ctx context.Context) *menu.MenuItem {
	w := New("Content")

	w.Add("Clean", func(data *menu.CallbackData) {
		runtime.EventsEmit(ctx, "clean_img")
		utils.WindowSetDefaultSize(ctx)
	})

	w.Add("Load Image", func(data *menu.CallbackData) {
		runtime.WindowReload(ctx)
		filename := utils.SelectImage(ctx)
		file, err := os.ReadFile(filename)
		if err != nil {
			utils.MessageBox(ctx, "提示", fmt.Sprintf("文件 '%s' 读取失败", filename))
			return
		}

		base64Img := base64.StdEncoding.EncodeToString(file)
		runtime.EventsEmit(ctx, "set_base64_img", fmt.Sprintf("data:image/png;base64,%s", base64Img))
	})

	return w.Build()
}
