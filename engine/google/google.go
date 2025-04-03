package google

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/flb-cc/submitter_search_engine/model"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/indexing/v3"
	"google.golang.org/api/option"
)

// Bing
var Google = &cGoogle{name: "google"}

type cGoogle struct {
	name string
}

func (s *cGoogle) GetName() string {
	return s.name
}

func (s *cGoogle) Submit(ctx context.Context, input *model.EngineSubmitInput) (resp interface{}, err error) {

	jsonStr, err := json.Marshal(input.Config)
	if err != nil {
		return
	}
	googleConfig, err := google.JWTConfigFromJSON(jsonStr, indexing.IndexingScope)
	if err != nil {
		return
	}

	googleClient := googleConfig.Client(ctx)

	if input.Proxy != "" {
		if proxyURL, err := url.Parse(input.Proxy); err == nil {
			googleClient.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
		}
	}

	svc, err := indexing.NewService(ctx, option.WithHTTPClient(googleClient))
	if err != nil {
		return
	}

	var respMap map[string]interface{} = make(map[string]interface{})

	for _, url := range input.Urls {
		googleResp, err := svc.UrlNotifications.Publish(&indexing.UrlNotification{
			Url:  url,
			Type: "URL_UPDATED",
		}).Do()
		if err != nil {
			respMap[url] = err.Error()
			continue
		}
		respMap[url] = googleResp
	}

	resp = respMap

	return
}
