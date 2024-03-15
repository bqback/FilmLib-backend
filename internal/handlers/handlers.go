package handlers

import (
	"filmlib/internal/config"
	"filmlib/internal/service"
)

type Handlers struct {
	ActorHandler
	MovieHandler
	SearchHandler
}

// var validate = validator.New(validator.WithRequiredStructEnabled())

const nodeName = "handler"

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services, config *config.Config) *Handlers {
	return &Handlers{
		ActorHandler:  *NewActorHandler(services.Actor),
		MovieHandler:  *NewMovieHandler(services.Movie),
		SearchHandler: *NewSearchHandler(services.Actor, services.Movie),
	}
}

// NewMovieHandler
// возвращает MovieHandler с необходимыми сервисами
func NewMovieHandler(ms service.IMovieService) *MovieHandler {
	return &MovieHandler{
		ms: ms,
	}
}

// NewActorHandler
// возвращает ActorHandler с необходимыми сервисами
func NewActorHandler(as service.IActorService) *ActorHandler {
	return &ActorHandler{
		as: as,
	}
}

// NewScanHandler
// возвращает ScanHandler с необходимыми сервисами
func NewSearchHandler(as service.IActorService, ms service.IMovieService) *SearchHandler {
	return &SearchHandler{
		as: as,
		ms: ms,
	}
}
