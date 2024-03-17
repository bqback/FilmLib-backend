package handlers

import (
	"filmlib/internal/apperrors"
	"filmlib/internal/service"
	"filmlib/internal/utils"
	"net/http"
)

type SearchHandler struct {
	ss service.ISearchService
}

// @Summary Искать фильм
// @Description Поиск фильма по строке
// @Description Строка ищется в названии фильма и списке актёров
// @Tags Поиск
//
// @Accept  json
// @Produce  json
//
// @Param query query string true "Поисковый запрос"
//
// @Success 200  {object}  []entities.Movie "Список результатов"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /search/movie/ [get]
func (sh SearchHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	funcName := "SearchMovies"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	query, err := utils.GetSearchQuery(rCtx)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		logger.Error(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Extracted search query", requestID, funcName, nodeName)

	movies, err := sh.ss.FindMovies(rCtx, query)
	if closed := respondOnErr(err, movies, "No movies found", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}
