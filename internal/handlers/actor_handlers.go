package handlers

import (
	"encoding/json"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/service"
	"filmlib/internal/storage/postgresql"
	"filmlib/internal/utils"
	"net/http"
)

type ActorHandler struct {
	as service.IActorService
}

// @Summary Создать актёра
// @Description
// @Tags Актёры
//
// @Accept  json
// @Produce  json
//
// @Param actorData body dto.NewActor true "Данные о новом актёре"
//
// @Success 200  {object}  entities.Actor "Объект нового актёра"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/ [post]
func (ah ActorHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	funcName := "CreateActor"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var newActor dto.NewActor
	err = json.NewDecoder(r.Body).Decode(&newActor)
	if err != nil {
		logger.DebugFmt("Failed to decode request: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}

	// err = validate.Struct(newActor)
	// if err != nil {
	// 	logger.Error("Validation failed")
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}
	// 	apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
	// 	return
	// }

	actor, err := ah.as.Create(rCtx, newActor)
	if closed := respondOnErr(err, actor, "Actor not created", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}

// @Summary Получить данные об актёре
// @Description Получить данные об актёре по его ID
// @Tags Актёры
//
// @Produce  json
//
// @Param id path uint true "ID актёра"
//
// @Success 200  {object}  entities.Actor "Объект актёра"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/{id}/ [get]
func (ah ActorHandler) ReadActor(w http.ResponseWriter, r *http.Request) {
	funcName := "ReadActor"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var actorID dto.ActorID
	id, err := utils.GetIDParam(rCtx)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		logger.Error(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	actorID.Value = id
	logger.DebugFmt("Extracted actor ID", requestID, funcName, nodeName)

	actor, err := ah.as.Read(rCtx, actorID)
	if closed := respondOnErr(err, actor, "No actor found with that ID", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}

// @Summary Изменить данные об актёре
// @Description Изменить данные об актёре по его ID.
// @Description Принимается любой набор параметров, но обязателен хотя бы один из них.
// @Tags Актёры
//
// @Accept  json
// @Produce  json
//
// @Param id path uint true "ID актёра"
// @Param actorData body dto.ExpectedActorUpdate true "Обновлённые данные актёра"
//
// @Success 200  {object}  entities.Actor
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/{id}/ [patch]
func (ah ActorHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	funcName := "UpdateActor"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	id, err := utils.GetIDParam(rCtx)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		logger.Error(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Extracted actor ID", requestID, funcName, nodeName)

	var updatedActor dto.UpdatedActor
	var updateValues map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updateValues)
	if err != nil {
		logger.DebugFmt("Failed to decode request: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	ok := postgresql.ValidateActorUpdate(updateValues)
	if !ok {
		logger.DebugFmt("Invalid update data", requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	updatedActor.ID = id
	updatedActor.Values = updateValues

	actor, err := ah.as.Update(rCtx, updatedActor)
	if closed := respondOnErr(err, actor, "No actor found with that ID", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}

// @Summary Удалить данные об актёре
// @Description Удалить данные об актёре по его ID
// @Tags Актёры
//
// @Produce  json
//
// @Param id path uint true "ID актёра"
//
// @Success 204  {string}  "no response"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/{id}/ [delete]
func (ah ActorHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	funcName := "DeleteActor"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var actorID dto.ActorID
	id, err := utils.GetIDParam(rCtx)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		logger.Error(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	actorID.Value = id
	logger.DebugFmt("Extracted actor ID", requestID, funcName, nodeName)

	err = ah.as.Delete(rCtx, actorID)
	if closed := respondOnErr(err, nil, "No actor found with that ID", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}

// @Summary Получить список актёров
// @Description Получить список всех актёров
// @Tags Актёры
//
// @Produce  json
//
// @Success 200  {object}  []entities.Actor "Список актёров"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /actors/ [get]
func (ah ActorHandler) GetActors(w http.ResponseWriter, r *http.Request) {
	funcName := "GetActors"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	actors, err := ah.as.GetActors(rCtx)
	if closed := respondOnErr(err, actors, "No actors found", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}
