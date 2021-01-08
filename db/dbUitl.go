package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DbObj 连接MySQL
func DbObj() {
	db, err := gorm.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
