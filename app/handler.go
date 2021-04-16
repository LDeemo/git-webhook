package app

import (
	"encoding/json"
	"fmt"
	"git-webhook/common"
	"git-webhook/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Callback 回调方法
func Callback(ctx *gin.Context) {
	requestID := common.RandString(16)
	logger := common.GetLogger()
	logger.Info(requestID, "Get from github callback")
	//获取json数据
	data, err := ctx.GetRawData()
	if err != nil {
		logger.Error(requestID, "Get Json data failed, error: %s", err.Error())
		ctx.JSON(502, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
		})
		return
	}
	var jsonMap map[string]interface{}
	//反序列化
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		logger.Error(requestID, "Json unmarshal failed, error: %s", err.Error())
		ctx.JSON(502, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
		})
		return
	}
	//处理逻辑
	//fmt.Println(jsonMap)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"requestID": requestID,
	//	"msg": "可以",
	//})
	server := &service.WebServer{}
	msg, err := server.DoExec(requestID)
	if err != nil {
		logger.Error(requestID, "Do script failed, error: %+v", err)
		ctx.JSON(502, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
		})
		return
	}
	logger.Info(requestID, fmt.Sprintf("%s",msg))
	ctx.JSON(http.StatusOK, gin.H{
		"requestID": requestID,
		"msg":       "可以",
	})
	return
}
