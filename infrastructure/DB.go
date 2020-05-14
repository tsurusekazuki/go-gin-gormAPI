package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Host string
	Username string
	Password string
	DBName string
	Connect *gorm.DB
}

func (d *DB) NewDB() *DB {
	c := NewConfig()
	d = &DB{
		Host: c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName: c.DB.Production.DBName,
	}
	return d.newDB()
}

func (d *DB) newDB() *DB {
	db, err := gorm.Open("mysql", d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	//db.AutoMigrate(&domain.Users{})
	//db.AutoMigrate(&domain.UsersForGet{})
	d.Connect = db
	fmt.Println("DB接続完了")

	return d
}

func (db *DB) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.Connect.First(out, where...)
}
