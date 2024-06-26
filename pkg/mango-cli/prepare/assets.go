package prepare

import (
	"embed"
	"io/fs"
	"os"
	"path"

	"github.com/kefniark/mango/pkg/mango-cli/config"
)

//go:embed static/**/*
var static embed.FS

type AssetsFilePrepare struct{}

func (prepare AssetsFilePrepare) Name() string {
	return "Static Preparer"
}

func (prepare AssetsFilePrepare) Execute(app string) error {
	copySubFolder(app, "")
	return nil
}

func copySubFolder(app string, folder string) {
	files, err := static.ReadDir(path.Join("static", folder))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			copySubFolder(app, path.Join(folder, file.Name()))
			continue
		}

		os.MkdirAll(path.Join(app, ".mango", folder), os.ModePerm)

		config.Logger.Debug().Str("app", app).Str("path", path.Join(app, ".mango", folder, file.Name())).Msg("Generated File")
		f, err := fs.ReadFile(static, path.Join("static", folder, file.Name()))
		if err != nil {
			continue
		}

		os.WriteFile(path.Join(app, ".mango", folder, file.Name()), f, os.ModePerm)
	}
}
