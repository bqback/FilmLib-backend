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

type PgMovieStorage struct {
	db *sqlx.DB
}

func NewMovieStorage(db *sqlx.DB) *PgMovieStorage {
	return &PgMovieStorage{
		db: db,
	}
}

func (s *PgMovieStorage) GetMovies(ctx context.Context, opts dto.SortOptions) ([]*entities.Movie, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "GetAllMovie"

	query, args, err := squirrel.
		Select(movieGetAllFields...).
		From(movieTable).
		InnerJoin(actorMovieOnMovieID).
		InnerJoin(actorOnActorID).
		GroupBy(movieIDField).
		OrderBy(SortOptionsMap[opts.Type] + " " + SortOrderMap[opts.Order]).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var tempMovies []dto.GetAllMovie
	if err := s.db.Select(&tempMovies, query, args...); err != nil {
		logger.DebugFmt("Movie select failed with error "+err.Error(), requestID, funcName, nodeName)
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrEmptyResult
		}
		return nil, apperrors.ErrMovieNotSelected
	}
	logger.DebugFmt("Movie selected", requestID, funcName, nodeName)

	movies := make([]*entities.Movie, len(tempMovies))
	for i, tm := range tempMovies {
		movie := &entities.Movie{
			ID:          tm.ID,
			Title:       tm.Title,
			Description: tm.Description,
			ReleaseDate: tm.ReleaseDate,
			Rating:      tm.Rating,
			Actors:      tm.Actors,
		}
		movies[i] = movie
	}

	return movies, nil
}

func (s *PgMovieStorage) Create(ctx context.Context, info dto.NewMovie) (*entities.Movie, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "CreateMovie"

	newMovie := &entities.Movie{
		Title:       info.Title,
		Description: info.Description,
		ReleaseDate: info.ReleaseDate,
		Rating:      info.Rating,
	}

	query1, args, err := squirrel.
		Insert(movieTable).
		Columns(allMovieInsertFields...).
		Values(info.Title, info.Description, info.ReleaseDate, info.Rating).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Movie query built", requestID, funcName, nodeName)

	tx, err := s.db.Begin()
	if err != nil {
		logger.DebugFmt("Failed to start transaction with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBeginTransaction
	}
	logger.DebugFmt("Transaction started", requestID, funcName, nodeName)

	var movieID int
	row := tx.QueryRow(query1, args...)
	if err := row.Scan(&movieID); err != nil {
		logger.DebugFmt("Movie insert failed with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID, funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrMovieNotCreated
	}
	logger.DebugFmt("Movie created", requestID, funcName, nodeName)

	newMovie.ID = uint64(movieID)

	insertBase := squirrel.
		Insert(actorMovieTable).
		Columns(actorMovieFields...)
	for _, actorID := range info.Actors {
		insertBase = insertBase.Values(actorID, movieID)
	}
	query2, args, err := insertBase.
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Movie query built", requestID, funcName, nodeName)

	_, err = tx.Exec(query2, args...)
	if err != nil {
		logger.DebugFmt("Failed to execute query with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID, funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotLinkActor
	}
	logger.DebugFmt("Actors linked to movie", requestID, funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.DebugFmt("Failed to commit changes", requestID, funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID, funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
	}
	logger.DebugFmt("Changes commited", requestID, funcName, nodeName)

	return newMovie, nil
}

func (s *PgMovieStorage) Read(ctx context.Context, id dto.MovieID) (*entities.Movie, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "ReadMovie"

	query, args, err := squirrel.
		Select(allMovieSelectFields...).
		From(movieTable).
		Where(squirrel.Eq{movieIDField: id.Value}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var movie entities.Movie
	if err := s.db.Get(&movie, query, args...); err != nil {
		logger.DebugFmt("Movie select failed with error "+err.Error(), requestID, funcName, nodeName)
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrEmptyResult
		}
		return nil, apperrors.ErrMovieNotSelected
	}
	logger.DebugFmt("Movie selected", requestID, funcName, nodeName)

	return &movie, nil
}

func (s *PgMovieStorage) Update(ctx context.Context, info dto.UpdatedMovie) (*entities.Movie, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "UpdateMovie"

	query, args, err := squirrel.
		Update(movieTable).
		SetMap(info.Values).
		Where(squirrel.Eq{movieIDField: info.ID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(movieUpdateReturnSuffix).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)
	logger.DebugFmt(query, requestID, funcName, nodeName)

	var movie entities.Movie
	if err := s.db.Get(&movie, query, args...); err != nil {
		logger.DebugFmt("Movie update failed with error "+err.Error(), requestID, funcName, nodeName)
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrEmptyResult
		}
		return nil, apperrors.ErrMovieNotUpdated
	}
	logger.DebugFmt("Movie updated", requestID, funcName, nodeName)

	return &movie, nil
}

func (s *PgMovieStorage) Delete(ctx context.Context, id dto.MovieID) error {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return err
	}

	funcName := "DeleteMovie"

	query, args, err := squirrel.
		Delete(movieTable).
		Where(squirrel.Eq{movieIDField: id.Value}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	if _, err = s.db.Exec(query, args...); err != nil {
		logger.DebugFmt("Movie delete failed with error "+err.Error(), requestID, funcName, nodeName)
		if err == sql.ErrNoRows {
			return apperrors.ErrEmptyResult
		}
		return apperrors.ErrMovieNotDeleted
	}
	logger.DebugFmt("Movie deleted", requestID, funcName, nodeName)

	return nil
}

func (s *PgMovieStorage) GetMovieActors(ctx context.Context, id dto.MovieID) ([]dto.ActorInfo, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "GetMovieActors"

	query, args, err := squirrel.
		Select(actorInfoFields...).
		From(actorTable).
		LeftJoin(actorMovieTable + " ON " + actorMovieActorIDField + "=" + actorIDField).
		Where(squirrel.Eq{actorMovieMovieIDField: id.Value}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var actors []dto.ActorInfo
	err = s.db.Select(&actors, query, args...)
	if err != nil {
		logger.DebugFmt("Movie select failed with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrMovieActorsNotSelected
	}
	logger.DebugFmt("Movie actors selected", requestID, funcName, nodeName)

	return actors, nil
}

func (s *PgMovieStorage) FindByString(ctx context.Context, search_term string) ([]*entities.Movie, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "FindByString"

	matchQuery, args, err := squirrel.
		Select(movieIDField).
		From(movieTable).
		InnerJoin(actorMovieOnMovieID).
		InnerJoin(actorOnActorID).
		Where(squirrel.Or{
			squirrel.Like{tlMovieTitleField: "%" + search_term + "%"},
			squirrel.Like{tlActorNameField: "%" + search_term + "%"},
		}).
		GroupBy(movieIDField).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var matchIDs []int
	if err := s.db.Select(&matchIDs, matchQuery, args...); err != nil {
		logger.DebugFmt("Movie select failed with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrMovieNotSelected
	}
	logger.DebugFmt("Select query executed", requestID, funcName, nodeName)
	if len(matchIDs) == 0 {
		return nil, apperrors.ErrEmptyResult
	}

	movieQuery, args, err := squirrel.
		Select(movieGetAllFields...).
		From(movieTable).
		InnerJoin(actorMovieOnMovieID).
		InnerJoin(actorOnActorID).
		Where(squirrel.Eq{movieIDField: matchIDs}).
		GroupBy(movieIDField).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var tempMovies []dto.GetAllMovie
	if err := s.db.Select(&tempMovies, movieQuery, args...); err != nil {
		logger.DebugFmt("Movie select failed with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrMovieNotSelected
	}
	logger.DebugFmt("Select query executed", requestID, funcName, nodeName)
	if len(tempMovies) == 0 {
		return nil, apperrors.ErrEmptyResult
	}

	movies := make([]*entities.Movie, len(tempMovies))
	for i, tm := range tempMovies {
		movie := &entities.Movie{
			ID:          tm.ID,
			Title:       tm.Title,
			Description: tm.Description,
			ReleaseDate: tm.ReleaseDate,
			Rating:      tm.Rating,
			Actors:      tm.Actors,
		}
		movies[i] = movie
	}

	return movies, nil
}
