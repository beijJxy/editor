package api

import (
	"github.com/beijJxy/editor/model"
	"github.com/beijJxy/editor/service"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EditorApi struct{}

// @Tags Editor
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /editor/routerName [post]
func (p *EditorApi) ApiName(c *gin.Context) {

	var plug model.Request
	_ = c.ShouldBindJSON(&plug)

	if res, err := service.ServiceGroupApp.PlugService(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {

		response.OkWithDetailed(res, "成功Y", c)
	}
}
