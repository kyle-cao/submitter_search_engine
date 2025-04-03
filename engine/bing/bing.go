package bing

import (
	"context"
	"encoding/json"

	"github.com/flb-cc/submitter_search_engine/model"
	"github.com/flb-cc/submitter_search_engine/utils/curl"
)

// Bing
var Bing = &cBing{name: "bing", postUrl: "https://ssl.bing.com/webmaster/api.svc/json/SubmitUrlbatch"}

type cBing struct {
	name    string
	postUrl string
}

func (s *cBing) GetName() string {
	return s.name
}

func (s *cBing) Submit(ctx context.Context, input *model.EngineSubmitInput) (resp interface{}, err error) {

	fullUrl := s.postUrl + "?apikey=" + input.Config["apiKey"].(string)
	headers := map[string]string{
		"Content-Type": "application/json",
		"Host":         "ssl.bing.com",
	}
	options := make(map[string]string)
	if input.Proxy != "" {
		options["proxy"] = input.Proxy
	}

	postParams := make(map[string]interface{})
	postParams["siteUrl"] = input.Config["siteUrl"]
	postParams["urlList"] = input.Urls
	postData, err := json.Marshal(postParams)
	if err != nil {
		return
	}
	respBody, _, err := curl.Post(fullUrl, string(postData), headers, options)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(respBody), &resp)
	if err != nil {
		return
	}

	return
}
