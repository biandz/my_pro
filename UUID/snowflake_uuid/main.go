package main

import (
	"fmt"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

/**
雪花算法总共64位
     1bit	             41bit             5bit	     5bit	 12bit
符号位（保留字段）	时间戳(当前时间-纪元时间)	数据中心id	机器id	自增序列
**/
// 缺点:依赖时间戳，存在始终回拨导致出现重复的情况++‘
var (
	data_center_id int64 = 1
	worker_id      int64 = 1
)

func main() {
	newSnowflake, err := snowflake.NewSnowflake(data_center_id, worker_id)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < 10; i++ {
		fmt.Println(newSnowflake.NextVal())
	}
}
