package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-vue-sandbox-api/models"
)

var ctx = context.Background()

func GetTasks(c *gin.Context) {
	tasks, err := models.Tasks().All(ctx, models.DB)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}
