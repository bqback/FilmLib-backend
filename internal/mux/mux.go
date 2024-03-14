package mux

import (
	"net/http"

	"filmlib/internal/config"
	"filmlib/internal/handlers"
	"filmlib/internal/logging"
)

func SetupMux(handlers *handlers.Handlers, config *config.Config, logger *logging.LogrusLogger) *http.ServeMux {
	mux := http.NewServeMux()

	baseUrl := config.API.BaseUrl

	actorsBaseUrl := baseUrl + "actors/"
	actorsSpecificUrl := actorsBaseUrl + "{id}/"

	moviesBaseUrl := baseUrl + "movies/"
	movieSortParams := moviesBaseUrl + "{type}/{order}/"
	moviesSpecificUrl := moviesBaseUrl + "{id}/"
	moviesSearchUrl := moviesBaseUrl + "search/"

	mux.HandleFunc("POST "+actorsBaseUrl, handlers.ActorHandler.CreateActor)
	mux.HandleFunc("GET "+actorsBaseUrl, handlers.ActorHandler.GetActors)
	mux.HandleFunc("GET "+actorsSpecificUrl, handlers.ActorHandler.ReadActor)
	mux.HandleFunc("PATCH "+actorsSpecificUrl, handlers.ActorHandler.UpdateActor)
	mux.HandleFunc("DELETE "+actorsSpecificUrl, handlers.ActorHandler.DeleteActor)

	mux.HandleFunc("GET "+movieSortParams, handlers.MovieHandler.GetMovies)
	mux.HandleFunc("POST "+moviesBaseUrl, handlers.MovieHandler.CreateMovie)
	mux.HandleFunc("GET "+moviesSpecificUrl, handlers.MovieHandler.ReadMovie)
	mux.HandleFunc("PATCH "+moviesSpecificUrl, handlers.MovieHandler.UpdateMovie)
	mux.HandleFunc("DELETE "+moviesSpecificUrl, handlers.MovieHandler.DeleteMovie)

	mux.HandleFunc("POST "+moviesSearchUrl, handlers.SearchHandler.Search)

	return mux
}

// func stackMiddleware(middlewares ...http.Handler) http.Handler {

// }
