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
	ErrRequestIDMissing   = errors.New("request ID is missing from context")
	ErrSortOptionsMissing = errors.New("sort options are missing from context")
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
	ErrJWTSecretMissing    = errors.New("JWT secret is missing from env")
)

var (
	ErrCouldNotBuildQuery       = errors.New("failed to build SQL query")
	ErrCouldNotPrepareStatement = errors.New("failed to prepare query statement")
	ErrCouldNotBeginTransaction = errors.New("failed to start DB transaction")
	ErrCouldNotRollback         = errors.New("failed to roll back after a failed query")
	ErrCouldNotCommit           = errors.New("failed to commit DB transaction changes")
	ErrEmptyResult              = errors.New("no results for provided query")
)

var (
	ErrActorNotCreated  = errors.New("failed to insert actor into database")
	ErrActorNotSelected = errors.New("failed to select actor from database")
	ErrActorNotUpdated  = errors.New("failed to update actor in database")
	ErrActorNotDeleted  = errors.New("failed to delete actor from database")
)

var (
	ErrMovieNotCreated  = errors.New("failed to insert movie into database")
	ErrMovieNotSelected = errors.New("failed to select movie from database")
	ErrMovieNotUpdated  = errors.New("failed to update movie in database")
	ErrMovieNotDeleted  = errors.New("failed to delete movie from database")
)

var (
	ErrCouldNotLinkActor      = errors.New("failed to link actor to movie")
	ErrActorMoviesNotSelected = errors.New("failed to get actor's movies")
	ErrMovieActorsNotSelected = errors.New("failed to get movie's actors")
)

var (
	ErrUserNotSelected = errors.New("failed to select user")
	ErrWrongPassword   = errors.New("wrong password")
)

var (
	ErrCouldNotParseClaims = errors.New("failed to parse JWT claims")
	ErrTokenExpired        = errors.New("token expired")
	ErrInvalidIssuedTime   = errors.New("invalid issue timestamp")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var BadRequestResponse = ErrorResponse{
	Code:    http.StatusBadRequest,
	Message: http.StatusText(http.StatusBadRequest),
}

var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: http.StatusText(http.StatusInternalServerError),
}

var UnauthorizedResponse = ErrorResponse{
	Code:    http.StatusUnauthorized,
	Message: http.StatusText(http.StatusUnauthorized),
}

func ReturnError(err ErrorResponse, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(err.Code)
	_, _ = w.Write([]byte(err.Message))
	r.Body.Close()
}
