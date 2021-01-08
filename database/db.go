package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect - Conexão com banco de dados
func Connect() (db *gorm.DB) {
	dbServer := "10.0.0.3:3306"
	dbUser := "root"
	dbPass := "123456"
	dbName := "ac_persona"

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbServer + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}
