package mysql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"server/global"
	"server/model/tables"
)

func GetTXHashList(userid string) ([]string, error) {
	// 定义正确的变量类型，确保与数据库查询结果匹配
	var tx []tables.MPSTransaction
	// 执行数据库查询，并捕获所有可能的错误
	err := global.MPS_DB.Where("user_id = ?", userid).Find(&tx).Error
	if err != nil {
		// 如果是记录未找到的错误，返回用户不存在的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, global.ErrorUserNotExist
		}
		// 对于其他错误，返回具体的数据库错误
		return nil, fmt.Errorf("database query failed: %w", err)
	}
	// 将查询结果转换为 *response.TXHashs 类型
	hashsList := make([]string, len(tx))
	for i, t := range tx {
		hashsList[i] = t.TxHash
	}
	return hashsList, nil
}

func CreateMPSRechargeOrder(order *tables.MPSRechargeOrder) error {

	return global.MPS_DB.Create(order).Error

}
