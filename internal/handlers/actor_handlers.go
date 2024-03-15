package handlers

import (
	"filmlib/internal/service"
	"net/http"
)

type ActorHandler struct {
	as service.IActorService
}

// @Summary Создать актёра
// @Description
// @Tags actors
//
// @Accept  json
// @Produce  json
//
// @Param actorData body dto.NewActor true "Данные о новом актёре"
//
// @Success 200  {object}  doc_structs.ActorResponse "Объект нового актёра"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/ [post]
func (ah ActorHandler) CreateActor(w http.ResponseWriter, r *http.Request) {

}

// @Summary Получить данные об актёре
// @Description Получить данные об актёре по его ID
// @Tags actors
//
// @Accept  json
// @Produce  json
//
// @Param actorID path dto.ActorID true "ID актёра"
//
// @Success 200  {object}  doc_structs.ActorResponse "Объект актёра"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/{id}/ [get]
func (ah ActorHandler) ReadActor(w http.ResponseWriter, r *http.Request) {
}

// @Summary Изменить данные об актёре
// @Description Изменить данные об актёре по его ID
// @Tags actors
//
// @Accept  json
// @Produce  json
//
// @Param actorID path dto.ActorID true "ID актёра"
// @Param actorData body dto.UpdatedActor true "Обновлённые данные актёра"
//
// @Success 204  {string}  "no response"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/{id}/ [patch]
func (ah ActorHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
}

// @Summary Удалить данные об актёре
// @Description Удалить данные об актёре по его ID
// @Tags actors
//
// @Accept  json
// @Produce  json
//
// @Param actorID path dto.ActorID true "ID актёра"
//
// @Success 204  {string}  "no response"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/{id}/ [delete]
func (ah ActorHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
}

// @Summary Получить список актёров
// @Description Получить список всех актёров
// @Tags actors
//
// @Accept  json
// @Produce  json
//
// @Success 200  {object}  doc_structs.ActorListResponse "Список актёров"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/ [get]
func (ah ActorHandler) GetActors(w http.ResponseWriter, r *http.Request) {
}
