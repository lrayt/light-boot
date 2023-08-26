package file_utils

import (
	"os"
)

// IsFolder 判断是否是文件夹
func IsFolder(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err == nil && fileInfo.IsDir() {
		return true
	}
	return false
}
