package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

// GenerateRandomNumericCode 生成指定长度的随机数字代码。
// 参数:
//
//	length - 指定生成代码的长度。必须大于0
//
// 返回值:
//
//	生成的随机数字代码字符串。
//	如果长度参数无效，则返回错误。
func GenerateRandomNumericCode(length int) (string, error) {
	// 检查长度是否合法
	if length <= 0 {
		return "", fmt.Errorf("invalid length: %d. Length must be greater than 0", length)
	}
	// 定义字符集为常量，这里仅使用数字字符
	const letters = "0123456789"
	// 使用字节切片提高性能，预分配所需的空间
	b := make([]byte, length)

	// 使用 math/big.Int 的 Rand 方法生成范围内的随机数
	for i := range b {
		// 生成一个介于0和字符集长度之间的随机数
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			// 如果生成随机数失败，返回错误
			return "", fmt.Errorf("failed to generate random number: %v", err)
		}
		// 根据生成的随机数选择对应的字符
		b[i] = letters[num.Int64()]
	}

	// 将字节切片转换为字符串返回
	return string(b), nil
}
