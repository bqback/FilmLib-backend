package mux

import (
	"net/http"

	"filmlib/internal/auth"
	"filmlib/internal/config"
	"filmlib/internal/handlers"
	"filmlib/internal/logging"
	"filmlib/internal/mux/middleware"

	_ "filmlib/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupMux(handlers *handlers.Handlers, config *config.Config, logger *logging.LogrusLogger) *http.Handler {
	mux := http.NewServeMux()

	baseUrl := config.API.BaseUrl

	manager := auth.NewManager(config.JWT)
	authMW := middleware.NewAuthMiddleware(manager)

	authUrl := baseUrl + "auth/"
	mux.Handle("POST "+authUrl, wrapHandleFunc(handlers.AuthHandler.Auth))

	actorsBaseUrl := baseUrl + "actors/"
	actorMux := actorMux(handlers.ActorHandler, actorsBaseUrl, authMW)
	mux.Handle(actorsBaseUrl, actorMux)

	moviesBaseUrl := baseUrl + "movies/"
	movieMux := movieMux(handlers.MovieHandler, moviesBaseUrl, authMW)
	mux.Handle(moviesBaseUrl, movieMux)

	searchBaseUrl := baseUrl + "search/"
	searchMux := searchMux(handlers.SearchHandler, searchBaseUrl, authMW)
	mux.Handle(searchBaseUrl, searchMux)

	swaggerUrl := "/swagger/*"
	mux.HandleFunc("GET "+swaggerUrl, httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),
	))

	mwStack := middleware.Stack(mux,
		middleware.JsonHeader, middleware.NewLogger(logger),
		middleware.RequestID, middleware.PanicRecovery,
	)

	return &mwStack
}

func actorMux(handlers handlers.ActorHandler, baseUrl string, authMiddleware middleware.AuthMiddleware) http.Handler {
	mux := http.NewServeMux()

	actorsSpecificUrl := baseUrl + "{id}/"

	mux.Handle("POST "+baseUrl, middleware.Stack(
		wrapHandleFunc(handlers.CreateActor),
		authMiddleware.CheckPerms,
	))
	mux.HandleFunc("GET "+baseUrl, handlers.GetActors)
	mux.Handle("GET "+actorsSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.ReadActor),
		middleware.ExtractID,
	))
	mux.Handle("PATCH "+actorsSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.UpdateActor),
		middleware.ExtractID,
		authMiddleware.CheckPerms,
	))
	mux.Handle("DELETE "+actorsSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.DeleteActor),
		middleware.ExtractID,
		authMiddleware.CheckPerms,
	))

	authMux := middleware.Stack(mux, authMiddleware.Auth)
	return authMux
}

func movieMux(handlers handlers.MovieHandler, baseUrl string, authMiddleware middleware.AuthMiddleware) http.Handler {
	mux := http.NewServeMux()

	moviesSpecificUrl := baseUrl + "{id}/"

	mux.Handle("GET "+baseUrl, middleware.Stack(
		wrapHandleFunc(handlers.GetMovies),
		middleware.ExtractSortParams,
	))
	mux.Handle("POST "+baseUrl, middleware.Stack(
		wrapHandleFunc(handlers.CreateMovie),
		authMiddleware.CheckPerms,
	))
	mux.Handle("GET "+moviesSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.ReadMovie),
		middleware.ExtractID,
	))
	mux.Handle("PATCH "+moviesSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.UpdateMovie),
		middleware.ExtractID,
		authMiddleware.CheckPerms,
	))
	mux.Handle("DELETE "+moviesSpecificUrl, middleware.Stack(
		wrapHandleFunc(handlers.DeleteMovie),
		middleware.ExtractID,
		authMiddleware.CheckPerms,
	))

	authMux := middleware.Stack(mux, authMiddleware.Auth)
	return authMux
}

func searchMux(handlers handlers.SearchHandler, baseUrl string, authMiddleware middleware.AuthMiddleware) http.Handler {
	mux := http.NewServeMux()

	searchMovieUrl := baseUrl + "movies/"

	mux.Handle("GET "+searchMovieUrl, middleware.Stack(
		wrapHandleFunc(handlers.SearchMovies),
		middleware.ExtractQuery,
	))

	authMux := middleware.Stack(mux, authMiddleware.Auth)
	return authMux
}

func wrapHandleFunc(hf http.HandlerFunc) http.Handler {
	return hf
}
