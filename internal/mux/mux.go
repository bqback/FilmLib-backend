package mux

import (
	"net/http"

	"filmlib/internal/config"
)

func SetupMux(handlers *handlers.Handlers, config *config.Config) *http.ServeMux {
	mux := http.NewServeMux()

	baseUrl := config.API.BaseUrl

	actorsBaseUrl := baseUrl + "actors/"
	actorsSpecificUrl := actorsBaseUrl + "{id}/"

	moviesBaseUrl := baseUrl + "movies/"
	moviesSpecificUrl := moviesBaseUrl + "{id}/"
	moviesSearchUrl := moviesBaseUrl + "search/"

	mux.HandleFunc("POST "+actorsBaseUrl, handlers.CreateActor)
	mux.HandleFunc("GET "+actorsSpecificUrl, handlers.ReadActor)
	mux.HandleFunc("PATCH "+actorsSpecificUrl, handlers.UpdateActor)
	mux.HandleFunc("DELETE "+actorsSpecificUrl, handlers.DeleteActor)

	mux.HandleFunc("GET "+moviesBaseUrl, handlers.GetTopMovies)
	mux.HandleFunc("POST "+moviesBaseUrl, handlers.CreateMovie)
	mux.HandleFunc("GET "+moviesSpecificUrl, handlers.ReadMovie)
	mux.HandleFunc("PATCH "+moviesSpecificUrl, handlers.UpdateMovie)
	mux.HandleFunc("DELETE "+moviesSpecificUrl, handlers.DeleteMovie)
	mux.HandleFunc("GET "+moviesSearchUrl, handlers.SearchMovies)

	return mux
}

// func stackMiddleware(middlewares ...http.Handler) http.Handler {

// }
