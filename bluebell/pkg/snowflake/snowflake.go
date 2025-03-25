package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(machineId int64, startTiem string) {
	// 计算时间偏移量
	var st time.Time
	st, err := time.Parse("2006-01-02", startTiem)
	if err != nil {
		return
	}
	// 设置起始时间
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineId)
	if err != nil {
		return
	}
}

func GenID() int64 {
	return node.Generate().Int64()
}
