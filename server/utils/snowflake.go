package utils

import (
	"server/global"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// InitSnowFlake 初始化
func InitSnowFlake() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", global.MPS_CONFIG.System.StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(global.MPS_CONFIG.System.MachineID)
	return
}

// GenID 生成id
func GenID() int64 {
	return node.Generate().Int64()
}
