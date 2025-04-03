package cmd

import (
	"fmt"
	"net/http"

	"github.com/flb-cc/submitter_search_engine/config"
	"github.com/flb-cc/submitter_search_engine/engine"
	"github.com/flb-cc/submitter_search_engine/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type PostJson struct {
	Urls    []string `json:"urls"`
	Engines []string `json:"engines"`
	Proxy   string   `json:"proxy"`
}

var (
	argHost string
	argPort string

	httpRun = &cobra.Command{
		Use:   "http",
		Short: `Run as an HTTP service. Method: POST, Data:Json, Example: {"urls":["http[s]://a.a"], "engines":["google", "bing"], "proxy":"socks5://[user:pass@]host:port"}`,
		RunE: func(cmd *cobra.Command, args []string) error {

			gin.SetMode(gin.ReleaseMode)

			router := gin.Default()

			router.GET("/robots.txt", func(c *gin.Context) {
				c.String(http.StatusOK, "User-agent: *\nDisallow:/")
			})
			router.GET("/favicon.ico", func(c *gin.Context) {
				c.String(http.StatusOK, "")
			})
			router.NoRoute(func(c *gin.Context) {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    404,
					"message": "404 Not Found",
				})
			})

			router.POST("/", func(c *gin.Context) {

				var postDataJson PostJson
				if err := c.ShouldBindJSON(&postDataJson); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if len(postDataJson.Urls) == 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "urls cannot be empty"})
					return
				}
				engineConfig, _ := config.NewConfig().GetSubmitConfig(postDataJson.Engines)
				for name, config := range engineConfig {
					result, _ := engine.Managers.Tasks[name].Submit(cmd.Context(), &model.EngineSubmitInput{
						Config: config.(map[string]interface{}),
						Urls:   postDataJson.Urls,
						Proxy:  postDataJson.Proxy,
					})
					resultMap[name] = result
				}

				c.JSON(http.StatusOK, resultMap)

			})

			addr := argHost + ":" + argPort
			fmt.Printf("Gin server is running on http://%s\n", addr)
			router.Run(addr)

			return nil
		},
	}
)

func init() {
	httpRun.PersistentFlags().StringVarP(&argHost, "host", "", "", "Listening Address. Example: 0.0.0.0 (required)")
	httpRun.PersistentFlags().StringVarP(&argPort, "port", "", "", "Listening Port. Example:8080 (required)")
}
