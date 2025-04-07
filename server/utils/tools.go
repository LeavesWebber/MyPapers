package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseTime(tString string) (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, tString)
	if err != nil {
		fmt.Println("error parsing tString:", err)
		return
	}
	formattedDate := t.Format("2006-01-02")
	t, err = time.Parse("2006-01-02", formattedDate)
	return
}

// Int64SliceToString 将int64切片转换成字符串
func Int64SliceToString(s []int64) (str string) {
	for _, v := range s {
		str += strconv.FormatInt(v, 10) + ","
	}
	// 去掉最后一个逗号
	str = str[:len(str)-1]
	return
}

func StringToSlice(input string) []string {
	// 去掉字符串中的 "[" 和 "]"
	input = strings.TrimSuffix(strings.TrimPrefix(input, "["), "]")
	// 去掉字符串中的引号和空格
	input = strings.ReplaceAll(input, `"`, "")
	input = strings.TrimSpace(input)

	// 使用逗号分隔符将字符串分割为字符串切片
	strSlice := strings.Split(input, ",")

	return strSlice
}
