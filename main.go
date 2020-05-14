package main

import "github.com/tsurusekazuki/clean-gin-gorm/app/infrastructure"

func main() {
	d := infrastructure.DB{}
	db := d.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run(":8080")
}