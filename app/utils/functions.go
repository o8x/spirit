package utils

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func WindowSetDefaultSize(ctx context.Context) {
	runtime.WindowSetSize(ctx, 350, 350)
}

func MessageBox(ctx context.Context, title, message string) {
	_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   title,
		Message: message,
	})
}

func SelectImage(ctx context.Context) string {
	file, err := runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title:                      "选择图片",
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: true,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "图片 (*.png;*.jpg;*.jpeg;*.bmp)",
				Pattern:     "*.png;*.jpg;*.jpeg;*.bmp",
			},
		},
	})

	if err != nil {
		return ""
	}

	return file
}
