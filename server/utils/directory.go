package utils

import "os"

// PathExists 判断文件夹是否存在
func PathExists(fileName string) (bool, error) {
	fi, err := os.Stat(fileName) // 判断有没有这个文件名
	if err == nil {
		if fi.IsDir() { // 判断是不是一个目录
			return true, nil
		}
		return false, err
	}
	if os.IsNotExist(err) { // 文件名不存在
		return false, nil
	}
	return false, err
}
