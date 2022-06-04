package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db            *gorm.DB
	userlock      sync.Mutex
	videolock     sync.Mutex
	followerdlock sync.Mutex
	favoritedlock sync.Mutex
	commentlock   sync.Mutex
	mapLock       sync.Mutex
)

func InitDBTool() {
	dsn := "ds:1234@tcp(139.224.105.6:3306)/dou_sheng?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
