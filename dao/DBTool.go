package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

//数据库对象和锁
var (
	db            *gorm.DB
	userlock      sync.Mutex
	videolock     sync.Mutex
	followerdlock sync.Mutex
	favoritedlock sync.Mutex
	commentlock   sync.Mutex
	mapLock       sync.Mutex
)

//初始化db
func InitDBTool() {
	dsn := "//先擦掉了"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
