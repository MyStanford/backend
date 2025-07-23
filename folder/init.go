package folder

import (
	"mystanford/utils"
	"os"
	"path/filepath"
)

func InitFolder() {
	rootPath, _ := os.Getwd()
	createDir(filepath.Join(rootPath, "data"))
}

func createDir(_path string) {
	if !utils.FileExist(_path) {
		os.Mkdir(_path, os.ModePerm)
	}
}
