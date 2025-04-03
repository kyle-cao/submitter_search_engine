package model

import (
	"context"
)

type Engine interface {
	GetName() string
	Submit(ctx context.Context, input *EngineSubmitInput) (resBody interface{}, err error)
}

type EngineSubmitInput struct {
	Config map[string]interface{}
	Urls   []string
	Proxy  string
}

type EngineSubmitOutput struct {
	Engine  string `json:"engine"`
	Message string `json:"message"`
}
