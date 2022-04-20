package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {

	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("should bind json error", zap.Any("err", err.Error()))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic create post failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic GetPostById pid failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func GetPostListHandler2(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("get post list handler 2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
