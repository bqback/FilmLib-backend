package handlers

import (
	"encoding/json"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/service"
	"filmlib/internal/utils"
	"net/http"
)

type AuthHandler struct {
	as service.IAuthService
}

// @Summary Авторизоваться
// @Description
// @Tags Авторизация
//
// @Accept  json
// @Produce  json
//
// @Param actorData body dto.LoginInfo true "Данные для авторизации"
//
// @Success 200  {object}  dto.JWT "JWT-токен для "
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/ [post]
func (ah AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	funcName := "Auth"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var info dto.LoginInfo
	err = json.NewDecoder(r.Body).Decode(&info)
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

	token, err := ah.as.Authenticate(rCtx, info)
	if closed := respondOnErr(err, token, "failed to authorize", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}
