package service

import (
	"github.com/beijJxy/editor/model"
)

type EditorService struct{}

func (e *EditorService) PlugService(req model.Request) (res model.Response, err error) {
	// 写你的业务逻辑
	return res, nil
}
