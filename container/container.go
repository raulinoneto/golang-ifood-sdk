package container

import (
	"github.com/raulinoneto/golang-ifood-sdk/authentication"
	"github.com/raulinoneto/golang-ifood-sdk/httpadapter"
	"github.com/raulinoneto/golang-ifood-sdk/mocks"
	"net/http"
	"time"
)

type container struct {
	env int
	timeout time.Duration
	httpadapter HttpAdapter
	authService authentication.Service
}

func New(env int, timeout time.Duration) *container {
	return &container{env: env,timeout: timeout}
}

func (c *container)GetHttpAdapter() HttpAdapter{
	if c.httpadapter != nil {
		return c.httpadapter
	}
	client := &http.Client{
		Timeout:       c.timeout,
	}
	switch c.env {
	case EnvDevelopment:
		c.httpadapter = httpadapter.New(new(mocks.HttpClientMock), "")
	case EnvProduction:
		c.httpadapter = httpadapter.New(client, urlProduction)
	case EnvSandBox:
		c.httpadapter = httpadapter.New(client, urlSandbox)
	}

	return c.httpadapter
}

func (c container) GetAuthenticationService(clientId, clientSecret string) authentication.Service{
	if c.authService == nil {
		c.authService = authentication.New(c.GetHttpAdapter(), clientId, clientSecret)
	}
	return c.authService
}