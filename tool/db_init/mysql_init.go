package db_init

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnMysql(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 启用日志记录
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("mysql连接成功...")
	}
	return db
}
