package main

import (
	"github.com/TestGo/controller"
	"github.com/gin-gonic/gin"
)

// type Employee struct {
// 	Position_id int
// 	Name        string
// 	Age         int
// 	Create_at   string
// 	Update_at   string
// }

func main() {
	controller.Connection()
	r := gin.Default()
	r.GET("/pepole", controller.ShowData)
	r.GET("/pepole/:id", controller.ShowFindId)
	r.POST("/pepole", controller.InsertData)
	r.PUT("/updatePepole/:id", controller.EditData)
	r.DELETE("/deletePepole/:id", controller.Delete)
	r.Run() //

}
