package engine

import (
	"github.com/flb-cc/submitter_search_engine/model"

	"github.com/flb-cc/submitter_search_engine/engine/baidu"
	"github.com/flb-cc/submitter_search_engine/engine/bing"
	"github.com/flb-cc/submitter_search_engine/engine/google"
)

func init() {
	var engines = []model.Engine{
		baidu.BaiDu,
		bing.Bing,
		google.Google,
	}

	for _, engine := range engines {
		Register(engine)
	}
}

var Managers = &EngineManager{
	Tasks: make(map[string]model.Engine),
}

// EngineManager 搜索引擎管理者
type EngineManager struct {
	Tasks map[string]model.Engine
}

func Register(c model.Engine) {
	Managers.Tasks[c.GetName()] = c
}
