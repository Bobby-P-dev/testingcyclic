package routers

import (
	"github.com/Bobby-P-dev/testingcyclic/controllers"
	"github.com/gin-gonic/gin"
)

func RouterA() *gin.Engine {
	r := gin.Default()

	r.POST("/testing", controllers.CreateData)
	r.GET("/geting", controllers.GetData)

	return r
}
