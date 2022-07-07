package lifecycle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shikharvashistha/fampay/pkg/types"
)

func NewEndpoints(dep Deployment) *endpoints {
	return &endpoints{
		dep: dep,
	}
}

type endpoints struct {
	dep Deployment
}

func (e *endpoints) GetDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.Query("page")
		limit := c.Query("limit")
		genere := c.Query("genere")
		response, err := e.dep.GetData(genere, page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
func (e *endpoints) SearchDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		description := c.Query("description")
		response, err := e.dep.SearchData(title, description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
func (e *endpoints) CronHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.CronRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := e.dep.CronSchedule(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
func (e *endpoints) CronDeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.CronDeleteRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := e.dep.CronDelete(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Cron deleted successfully"})
	}
}
