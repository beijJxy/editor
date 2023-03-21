package router

import (
	"github.com/beijJxy/editor/api"
	"github.com/gin-gonic/gin"
)

type EditorRouter struct {
}

func (s *EditorRouter) InitEditorRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.EditorApi
	{
		plugRouter.POST("routerName", plugApi.ApiName)
	}
}
