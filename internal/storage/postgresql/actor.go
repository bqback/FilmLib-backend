package postgresql

import (
	"context"
	"database/sql"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/pkg/entities"
	"filmlib/internal/utils"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PgActorStorage struct {
	db *sqlx.DB
}

func NewActorStorage(db *sqlx.DB) *PgActorStorage {
	return &PgActorStorage{
		db: db,
	}
}

func (s *PgActorStorage) Create(ctx context.Context, info dto.NewActor) (*entities.Actor, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "CreateActor"

	newActor := &entities.Actor{
		Name:      info.Name,
		Gender:    info.Gender,
		BirthDate: info.BirthDate,
	}

	query, args, err := squirrel.
		Insert(actorTable).
		Columns(allActorInsertFields...).
		Values(info.Name, info.Gender, info.BirthDate).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var actorID int
	row := s.db.QueryRow(query, args...)
	if err := row.Scan(&actorID); err != nil {
		logger.DebugFmt("Actor insert failed with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrActorNotCreated
	}
	logger.DebugFmt("Actor created", requestID, funcName, nodeName)

	newActor.ID = uint64(actorID)

	return newActor, nil
}

func (s *PgActorStorage) Read(ctx context.Context, id dto.ActorID) (*entities.Actor, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "ReadActor"

	query, args, err := squirrel.
		Select(allActorSelectFields...).
		From(actorTable).
		Where(squirrel.Eq{actorIDField: id.Value}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var actor entities.Actor
	if err := s.db.Get(&actor, query, args...); err != nil {
		logger.DebugFmt("Actor select failed with error "+err.Error(), requestID, funcName, nodeName)
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrEmptyResult
		}
		return nil, apperrors.ErrActorNotSelected
	}
	logger.DebugFmt("Actor selected", requestID, funcName, nodeName)

	return &actor, nil
}

func (s *PgActorStorage) Delete(ctx context.Context, id dto.ActorID) error {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return err
	}

	funcName := "DeleteActor"

	query, args, err := squirrel.
		Delete(actorTable).
		Where(squirrel.Eq{actorIDField: id.Value}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	if _, err = s.db.Exec(query, args...); err != nil {
		logger.DebugFmt("Actor delete failed with error "+err.Error(), requestID, funcName, nodeName)
		if err == sql.ErrNoRows {
			return apperrors.ErrEmptyResult
		}
		return apperrors.ErrActorNotDeleted
	}
	logger.DebugFmt("Actor deleted", requestID, funcName, nodeName)

	return nil
}

func (s *PgActorStorage) GetActorMovies(ctx context.Context, id dto.ActorID) ([]dto.MovieInfo, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "GetActorMovies"

	query, args, err := squirrel.
		Select(movieInfoFields...).
		From(movieTable).
		LeftJoin(actorMovieTable + " ON " + actorMovieMovieIDField + "=" + movieIDField).
		Where(squirrel.Eq{actorMovieActorIDField: id.Value}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var movies []dto.MovieInfo
	err = s.db.Select(&movies, query, args...)
	if err != nil {
		logger.DebugFmt("Actor movies select failed with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrActorMoviesNotSelected
	}
	logger.DebugFmt("Actor movies selected", requestID, funcName, nodeName)

	return movies, nil
}
