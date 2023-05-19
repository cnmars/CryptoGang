package account

import (
	// "fmt"
	"golddigger/common/models/response"
	// "golddigger/global"

	// "context"
	// "time"

	"github.com/gin-gonic/gin"
)

type AccountApi struct{}

func (w *AccountApi) GetRandomSignMsg(c *gin.Context) {
	msg := accountService.GetRandomSignMsg()
	
	response.OkWithData(msg, "",c)
}
