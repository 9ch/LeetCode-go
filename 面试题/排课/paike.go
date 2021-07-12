package main

import (
	"fmt"
	"time"
)

//有一个数组存储当前老师已经确定上课的时间，例如【2018-05-04 09:00,】

func main() {
	timeZone, _ := time.LoadLocation("Asia/Shanghai")
	timeObj_1, _ := time.Parse("2006-01-02 15:04:05", "2020-12-01 00:00:00")
	fmt.Println("timeObj_1", timeObj_1.Unix())
	timeObj_2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-12-01 00:00:00", timeZone)
	fmt.Println("timeObj_2", timeObj_2.Unix())

}
