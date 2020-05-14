package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/tsurusekazuki/clean-gin-gorm/app/interfaces/controllers"
)

type Routing struct {
	DB *DB
	Gin *gin.Engine
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB: db,
		Gin: gin.Default(),
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)
	// curl -i -H 'Content-Type:application/json' localhost:8080/users/1
	r.Gin.GET("/users/:id", func (c *gin.Context) { usersController.Get(c) })
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}
