package baidu

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/flb-cc/submitter_search_engine/model"
	"github.com/flb-cc/submitter_search_engine/utils/curl"
)

// BaiDu
var BaiDu = &cBaiDu{
	name:    "baidu",
	postUrl: "http://data.zz.baidu.com/urls",
}

type cBaiDu struct {
	name    string
	postUrl string
}

func (s *cBaiDu) GetName() string {
	return s.name
}

func (s *cBaiDu) Submit(ctx context.Context, input *model.EngineSubmitInput) (resp interface{}, err error) {

	fullUrl := s.postUrl + "?site=" + input.Config["site"].(string) + "&token=" + input.Config["token"].(string)
	headers := map[string]string{
		"Content-Type": "text/plain",
	}
	options := make(map[string]string)
	respBody, _, err := curl.Post(fullUrl, strings.Join(input.Urls, "\n"), headers, options)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(respBody), &resp)
	if err != nil {
		return
	}

	return
}
