package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrNilContext = errors.New("context is nil")
)

var (
	ErrInvalidLoggingLevel = errors.New("invalid logging level")
	ErrLoggerMissing       = errors.New("logger missing from context")
)

var (
	ErrRequestIDMissing = errors.New("request ID is missing from context")
)

var (
	ErrCouldNotParseURLParam = errors.New("failed to parse URL params")
	ErrURLParamMissing       = errors.New("url param is missing from context")
)

var (
	ErrEnvNotFound         = errors.New("unable to load .env file")
	ErrDatabaseUserMissing = errors.New("database user is missing from env")
	ErrDatabasePWMissing   = errors.New("database password is missing from env")
	ErrDatabaseNameMissing = errors.New("database name is missing from env")
)

var (
	ErrCouldNotBuildQuery       = errors.New("failed to build SQL query")
	ErrCouldNotPrepareStatement = errors.New("failed to prepare query statement")
	ErrCouldNotBeginTransaction = errors.New("failed to start DB transaction")
	ErrCouldNotRollback         = errors.New("failed to roll back after a failed query")
	ErrCouldNotCommit           = errors.New("failed to commit DB transaction changes")
)

var (
	ErrActorNotCreated  = errors.New("failed to insert actor into database")
	ErrActorNotSelected = errors.New("failed to select actor from database")
)

var (
	ErrMovieNotCreated  = errors.New("failed to insert movie into database")
	ErrMovieNotSelected = errors.New("failed to select movie from database")
)

var (
	ErrCouldNotLinkActor      = errors.New("failed to link actor to movie")
	ErrActorMoviesNotSelected = errors.New("failed to get actor's movies")
	ErrMovieActorsNotSelected = errors.New("failed to get movie's actors")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var BadRequestResponse = ErrorResponse{
	Code:    http.StatusBadRequest,
	Message: "Bad request",
}

var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "Internal error",
}

func ReturnError(err ErrorResponse, w http.ResponseWriter, r *http.Request) {
	// rCtx := r.Context()
	// logger := *utils.GetReqLogger(rCtx)
	// requestID := chimw.GetReqID(rCtx)
	// funcName := "ReturnError"

	w.WriteHeader(err.Code)
	// logger.DebugFmt(fmt.Sprintf("Wrote error code %v", err.Code), requestID, funcName, nodeName)
	_, _ = w.Write([]byte(err.Message))
	// logger.DebugFmt(fmt.Sprintf("Wrote message %v", err.Code), requestID, funcName, nodeName)
	r.Body.Close()
	// logger.DebugFmt("Request body closed", requestID, funcName, nodeName)
}
