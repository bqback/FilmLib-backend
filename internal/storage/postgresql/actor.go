package postgresql

import (
	"context"
	"database/sql"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/pkg/entities"
	"filmlib/internal/utils"

	"github.com/Masterminds/squirrel"
)

type PgActorStorage struct {
	db *sql.DB
}

func NewActorStorage(db *sql.DB) *PgActorStorage {
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

	query1, args, err := squirrel.
		Insert(actorTable).
		Columns(allActorInsertFields...).
		Values(info.Name, info.Gender, info.BirthDate).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		// logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	// logger.DebugFmt("Built query\n\t"+query1+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	// tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	// if err != nil {
	// 	// logger.DebugFmt("Failed to start transaction with error "+err.Error(), requestID.String(), funcName, nodeName)
	// 	return nil, apperrors.ErrCouldNotBeginTransaction
	// }
	// logger.DebugFmt("Transaction started", requestID.String(), funcName, nodeName)

	var actorID int
	row := s.db.QueryRow(query1, args...)
	if err := row.Scan(&actorID); err != nil {
		logger.DebugFmt("Actor insert failed with error "+err.Error(), requestID, funcName, nodeName)
		// err = tx.Rollback()
		// if err != nil {
		// 	logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		// 	return nil, apperrors.ErrCouldNotRollback
		// }
		return nil, apperrors.ErrActorNotCreated
	}
	// logger.DebugFmt("Board created", requestID.String(), funcName, nodeName)

	newActor.ID = uint64(actorID)

	return newActor, nil
}
