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

func (s *PgActorStorage) Update(ctx context.Context, info dto.UpdatedActor) (*entities.Actor, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "UpdateActor"

	// Черновой вариант на трудный день
	// updateBase := squirrel.Update(actorTable)
	// map[string]interface{} проще, но как будто опаснее
	// values := reflect.ValueOf(info)
	// types := values.Type()
	// for i := 1; i < values.NumField(); i++ {
	// 	if !values.Field(i).IsNil() {
	// 		updateBase = updateBase.Set(actorUpdateStructToField[types.Field(i).Name], values.Field(i).)
	// 	}
	// }

	// query, args, err := updateBase.
	// 	Where(squirrel.Eq{actorIDField: info.ID}).
	// 	PlaceholderFormat(squirrel.Dollar).
	// 	ToSql()
	query, args, err := squirrel.Update(actorTable).
		SetMap(info.Values).
		Where(squirrel.Eq{actorIDField: info.ID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(actorUpdateReturnSuffix).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)
	logger.DebugFmt(query, requestID, funcName, nodeName)

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
		LeftJoin(actorMovieOnMovieID).
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

func (s *PgActorStorage) GetAll(ctx context.Context) ([]*entities.Actor, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "GetAllActors"

	query, args, err := squirrel.
		Select(actorGetAllFields...).
		From(actorTable).
		InnerJoin(actorMovieOnActorID).
		InnerJoin(movieOnMovieID).
		GroupBy(actorIDField).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var tempActors []dto.GetAllActor
	if err := s.db.Select(&tempActors, query, args...); err != nil {
		logger.DebugFmt("Actor select failed with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrActorNotSelected
	}
	logger.DebugFmt("Select query executed", requestID, funcName, nodeName)
	if len(tempActors) == 0 {
		return nil, apperrors.ErrEmptyResult
	}

	actors := make([]*entities.Actor, len(tempActors))
	for i, ta := range tempActors {
		actor := &entities.Actor{
			ID:        ta.ID,
			Name:      ta.Name,
			Gender:    ta.Gender,
			BirthDate: ta.BirthDate,
			Movies:    ta.Movies,
		}
		actors[i] = actor
	}

	return actors, nil
}

// func (s *PgActorStorage) FindByName(ctx context.Context, search_term string) ([]*entities.Actor, error) {
// 	logger, requestID, err := utils.GetLoggerAndID(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	funcName := "GetAllActors"

// 	query, args, err := squirrel.
// 		Select(actorGetAllFields...).
// 		From(actorTable).
// 		InnerJoin(actorMovieOnActorID).
// 		InnerJoin(movieOnMovieID).
// 		Where(squirrel.Like{actorNameField: "%" + search_term + "%"}).
// 		GroupBy(actorIDField).
// 		PlaceholderFormat(squirrel.Dollar).
// 		ToSql()
// 	if err != nil {
// 		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
// 		return nil, apperrors.ErrCouldNotBuildQuery
// 	}
// 	logger.DebugFmt("Query built", requestID, funcName, nodeName)

// 	var tempActors []dto.GetAllActor
// 	if err := s.db.Select(&tempActors, query, args...); err != nil {
// 		logger.DebugFmt("Actor select failed with error "+err.Error(), requestID, funcName, nodeName)
// 		if err == sql.ErrNoRows {
// 			return nil, apperrors.ErrEmptyResult
// 		}
// 		return nil, apperrors.ErrActorNotSelected
// 	}
// 	logger.DebugFmt("Actor selected", requestID, funcName, nodeName)

// 	actors := make([]*entities.Actor, len(tempActors))
// 	for i, ta := range tempActors {
// 		actor := &entities.Actor{
// 			ID:        ta.ID,
// 			Name:      ta.Name,
// 			Gender:    ta.Gender,
// 			BirthDate: ta.BirthDate,
// 			Movies:    ta.Movies,
// 		}
// 		actors[i] = actor
// 	}

// 	return actors, nil
// }
