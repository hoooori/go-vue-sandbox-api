package routers

import (
	"github.com/gin-gonic/gin"
	"go-vue-sandbox-api/controllers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("tasks", v1.GetTasks)
	}

	return r
}
