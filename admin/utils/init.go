/**
 * @Author: dexukong
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2022/08/08 10:36
 */

package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init error:", err)
	}

	return db
}
