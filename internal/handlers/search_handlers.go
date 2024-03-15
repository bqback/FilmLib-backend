package handlers

import (
	"filmlib/internal/service"
	"net/http"
)

type SearchHandler struct {
	as service.IActorService
	ms service.IMovieService
}

// @Summary Искать фильм
// @Description Поиск фильма по строке
// @Description Строка ищется в названии фильма и списке актёров
// @Tags movies
//
// @Accept  json
// @Produce  json
//
// @Param searchQuery body string true "Поисковый запрос"
//
// @Success 200  {object}  dto.SearchResult "Список результатов"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /movies/search/ [POST]
func (sh SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
}
