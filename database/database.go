import "os"
package database

import (
	"fmt"
	"os"
  mySecret := os.Getenv("DB_USER")
  mySecret := os.Getenv("DB_NAME")
  mySecret := os.Getenv("DB_PORT")
  mySecret := os.Getenv("DB_HOST")
  mySecret := os.Getenv("DB_PASS")

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ error = godotenv.Load()

//config connecting database
func Database() *gorm.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}