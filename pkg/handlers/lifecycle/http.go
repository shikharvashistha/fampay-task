package lifecycle

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPHandlers(e *gin.RouterGroup, d Deployment) {
	ep := NewEndpoints(d)
	e.GET(GetData, ep.GetDataHandler())
	e.GET(SearchData, ep.SearchDataHandler())
	e.POST(SchduleCron, ep.CronHandler())
	e.DELETE(DeleteCron, ep.CronDeleteHandler())
}
