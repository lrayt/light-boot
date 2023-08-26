package file_utils

import "os"

// PathExists 目录是否存在
func PathExists(target string) bool {
	_, err := os.Stat(target)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
