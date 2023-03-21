package editor

import (
	"editor/global"
	"editor/router"
	"github.com/gin-gonic/gin"
)

type EditorPlugin struct {
}

func CreateEditorPlug(Name string, Creator string) *EditorPlugin {
	global.GlobalConfig.Name = Name
	global.GlobalConfig.Creator = Creator

	return &EditorPlugin{}
}

func (*EditorPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEditorRouter(group)
}

func (*EditorPlugin) RouterPath() string {
	return "editor"
}
