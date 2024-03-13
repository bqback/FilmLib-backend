package handlers

import (
	"net/http"
	"proxy_server/internal/config"
	"proxy_server/internal/service"
)

type Handlers struct {
	ActorHandler
	MovieHandler
}

const nodeName = "handler"

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services, config *config.Config) *Handlers {
	return &Handlers{
		ActorHandler: *NewActorHandler(services.Request),
		MovieHandler: *NewMovieHandler(services.Repeat, config),
	}
}

// NewMovieHandler
// возвращает MovieHandler с необходимыми сервисами
func NewMovieHandler(reqs service.IRequestService) *MovieHandler {
	return &MovieHandler{
		rs: reqs,
	}
}

// NewActorHandler
// возвращает ActorHandler с необходимыми сервисами
func NewActorHandler(reps service.IRepeatService, config *config.Config) *ActorHandler {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(config.Proxy.URL),
		},
	}
	return &ActorHandler{
		rs:     reps,
		client: client,
	}
}

// NewScanHandler
// возвращает ScanHandler с необходимыми сервисами
func NewScanHandler(scans service.IScanService, config *config.Config) *ScanHandler {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(config.Proxy.URL),
		},
	}
	dictLocation := config.FileAttack.DictFile
	return &ScanHandler{
		ss:           scans,
		client:       client,
		dictLocation: dictLocation,
	}
}
