package handlers

import (
	"encoding/json"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/service"
	"filmlib/internal/utils"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
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
	funcName := "CreateActor"

	rCtx := r.Context()
	logger, err := utils.GetReqLogger(rCtx)
	if err != nil {
		log.Println(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	requestID, err := utils.GetReqID(rCtx)
	if err != nil {
		logger.Error(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
	}

	var newActor dto.NewActor
	err = json.NewDecoder(r.Body).Decode(&newActor)
	if err != nil {
		logger.DebugFmt("Failed to decode request: "+err.Error(), requestID, funcName, nodeName)
	}

	err = validate.Struct(newActor)
	if err != nil {
		logger.Error("Validation failed")
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		}

		for _, err := range err.(validator.ValidationErrors) {
			logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		}
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}

	actor, err := ah.as.Create(rCtx, newActor)
	if err != nil {
		logger.Error(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Actor created", requestID, funcName, nodeName)

	jsonResponse, err := json.Marshal(actor)
	if err != nil {
		logger.Error("Failed to marshal response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Failed to return response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()
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
