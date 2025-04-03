package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/flb-cc/submitter_search_engine/config"
	"github.com/flb-cc/submitter_search_engine/engine"
	"github.com/flb-cc/submitter_search_engine/model"
	"github.com/spf13/cobra"
)

var (
	argEngines string
	argUrls    string

	proxy string
	urls  []string

	resultMap = make(map[string]interface{})

	cmdRun = &cobra.Command{
		Use:   "cmd",
		Short: "Run in command-line mode",
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			if argUrls == "" {
				return errors.New("urls cannot be empty")
			}
			urls = strings.Split(argUrls, ",")

			var engineConfig map[string]interface{}

			engineConfig, err = config.NewConfig().GetSubmitConfig(strings.Split(argEngines, ","))
			if err != nil {
				return
			}
			for name, config := range engineConfig {
				result, _ := engine.Managers.Tasks[name].Submit(cmd.Context(), &model.EngineSubmitInput{
					Config: config.(map[string]interface{}),
					Urls:   urls,
					Proxy:  proxy,
				})
				resultMap[name] = result
			}

			jsonData, _ := json.Marshal(resultMap)
			fmt.Println(string(jsonData))

			return nil
		},
	}
)

func init() {
	cmdRun.PersistentFlags().StringVarP(&argUrls, "urls", "", "", "Submit multiple URLs separated by commas. Example: url1,url2,url3 (required)")
	cmdRun.PersistentFlags().StringVarP(&argEngines, "engines", "", "", "Query multiple search engines specified in a comma-separated list. Example: google,bing,baidu")
	cmdRun.PersistentFlags().StringVarP(&proxy, "proxy", "", "", "SOCKS5 proxy address (format: socks5://[user:pass@]host:port)")
}
