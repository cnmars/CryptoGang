package apis

import (
	"golddigger/common/models/response"
	request "golddigger/app/system/models/request/common"
	"golddigger/global"
	system"golddigger/app/system/models"
	systemReq "golddigger/app/system/models/request"
	// systemRes "golddigger/app/system/models/response"
	"golddigger/utils"
	// "golddigger/utils/jwtutil"
	// "fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FeedbackApi struct{}
func (s *FeedbackApi) CreateFeedback(c *gin.Context) {
	var fb system.FeedBack
	err := c.ShouldBindJSON(&fb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(fb, utils.FeedbackVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = feedbackService.CreateFeedback(fb)
	if err != nil {
		global.GD_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}


func (s *FeedbackApi) SetReaded(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = feedbackService.SetReaded(idInfo.ID)
	if err != nil {
		global.GD_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}


func (s *FeedbackApi) DeleteById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = feedbackService.DeleteById(idInfo.ID)
	if err != nil {
		global.GD_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}



func (s *FeedbackApi) GetApiList(c *gin.Context) {
	var pageInfo systemReq.SearchFeedbackParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := feedbackService.GetFeedbackList(pageInfo.FeedBack , pageInfo.PageInfo , pageInfo.OrderKey , pageInfo.Desc )
	if err != nil {
		global.GD_LOG.Error("反馈信息获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
