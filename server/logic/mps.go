package logic

import (
	"server/dao/mysql"
)

func GetTXHashList(userId string) (out []string, err error) {
	// 查询数据库
	if out, err = mysql.GetTXHashList(userId); err != nil {
		return nil, err
	}
	// 返回信息
	return out, err
}
