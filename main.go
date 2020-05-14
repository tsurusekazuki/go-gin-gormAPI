package main

import "github.com/tsurusekazuki/clean-gin-gorm/app/infrastructure"

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run(":8080")
}