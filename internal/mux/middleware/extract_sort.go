package middleware

import (
	"context"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/utils"
	"net/http"
	"slices"
	"strconv"
)

func ExtractSortParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		funcName := "ExtractSortParams"

		ctx := r.Context()

		logger, requestID, err := utils.GetLoggerAndID(ctx)
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		sortTypeString := r.URL.Query().Get("sort")
		if sortTypeString == "" {
			logger.DebugFmt("'sort' param missing from query", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}
		sortType, err := strconv.Atoi(sortTypeString)
		if err != nil {
			logger.DebugFmt("Failed to parse query param as int", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}
		if !slices.Contains([]int{dto.TitleSort, dto.RatingSort, dto.ReleaseSort}, sortType) {
			logger.DebugFmt("Invalid sort type provided", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		sortOrderString := r.URL.Query().Get("order")
		var sortOrder int
		if sortOrderString == "" {
			sortOrder = dto.DefaultSort[sortType]
		} else {
			sortOrder, err = strconv.Atoi(sortOrderString)
			if err != nil {
				logger.DebugFmt("Failed to parse query param as int", requestID, funcName, nodeName)
				apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
				return
			}
			if !slices.Contains([]int{dto.AscSort, dto.DescSort}, sortOrder) {
				logger.DebugFmt("Invalid sort order provided", requestID, funcName, nodeName)
				apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
				return
			}
		}

		opts := dto.SortOptions{Type: sortType, Order: sortOrder}
		rCtx := context.WithValue(ctx, dto.SortOptionsKey, opts)
		logger.DebugFmt("Stored in context", requestID, funcName, nodeName)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
