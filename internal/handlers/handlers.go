package handlers

import (
	"encoding/json"
	"filmlib/internal/apperrors"
	"filmlib/internal/config"
	"filmlib/internal/logging"
	"filmlib/internal/service"
	"net/http"
)

type Handlers struct {
	ActorHandler
	MovieHandler
	SearchHandler
	AuthHandler
}

// var validate = validator.New(validator.WithRequiredStructEnabled())

const nodeName = "handler"

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services, config *config.Config) *Handlers {
	return &Handlers{
		ActorHandler:  *NewActorHandler(services.Actor),
		MovieHandler:  *NewMovieHandler(services.Movie),
		SearchHandler: *NewSearchHandler(services.Search),
		AuthHandler:   *NewAuthHandler(services.Auth),
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

// NewSearchHandler
// возвращает SearchHandler с необходимыми сервисами
func NewSearchHandler(ss service.ISearchService) *SearchHandler {
	return &SearchHandler{
		ss: ss,
	}
}

// NewAuthHandler
// возвращает AuthHandler с необходимыми сервисами
func NewAuthHandler(as service.IAuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
	}
}

// respondOnErr
// пишет в http.ResponseWriter ответ в зависимости от ошибки, отданной вызовом сервиса.
// Если в качестве obj передан nil, пишет код 204 в заголовок ответа.
// В остальных случаях пытается замаршалить obj и отдать его как JSON.
// Возвращаемое значение bool определяет, закрыто ли тело запроса.
func respondOnErr(
	err error, obj interface{},
	emptyResponse string,
	logger logging.ILogger, requestID string, funcName string,
	w http.ResponseWriter, r *http.Request,
) bool {
	closed := false
	switch err {
	case nil:
		switch obj {
		case nil:
			w.WriteHeader(http.StatusNoContent)
		default:
			jsonResponse, err := json.Marshal(obj)
			if err != nil {
				logger.Error("Failed to marshal response: " + err.Error())
				apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
				closed = true
			}

			_, err = w.Write(jsonResponse)
			if err != nil {
				logger.Error("Failed to return response: " + err.Error())
				apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
				closed = true
			}
		}
	case apperrors.ErrEmptyResult:
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(emptyResponse))
		if err != nil {
			logger.Error("Failed to return response: " + err.Error())
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			closed = true
		}
	default:
		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		closed = true
	}
	return closed
}
