package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var Test *mgo.Session

func Setup(url string) {
	Test, databaseErr := mgo.Dial(url)
	fmt.Println(url)
	//连接失败时终止
	if databaseErr != nil {
		panic(databaseErr)
	}
	fmt.Println("连接成功")
	//延迟关闭，释放资源
	defer Test.Close()
	//设置模式
	Test.SetMode(mgo.Monotonic, true)
}