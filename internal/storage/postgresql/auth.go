package postgresql

import (
	"context"
	"database/sql"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/utils"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PgAuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *PgAuthStorage {
	return &PgAuthStorage{
		db: db,
	}
}

func (s *PgAuthStorage) Auth(ctx context.Context, info dto.LoginInfo) (bool, *dto.DBUser, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return false, nil, err
	}

	funcName := "Auth"

	query, args, err := squirrel.
		Select(allUserSelectFields...).
		From(userTable).
		Where(squirrel.Eq{userLoginField: info.Login}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return false, nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var user dto.DBUser
	if err := s.db.Get(&user, query, args...); err != nil {
		logger.DebugFmt("User select failed with error "+err.Error(), requestID, funcName, nodeName)
		if err == sql.ErrNoRows {
			return false, nil, apperrors.ErrEmptyResult
		}
		return false, nil, apperrors.ErrUserNotSelected
	}
	logger.DebugFmt("User selected", requestID, funcName, nodeName)

	err = utils.ComparePasswords(user.PasswordHash, info.Password)
	if err != nil {
		logger.DebugFmt("Passwords don't match", requestID, funcName, nodeName)
		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
		return false, nil, apperrors.ErrWrongPassword
	}
	logger.DebugFmt("Passwords match", requestID, funcName, nodeName)

	return true, &user, nil
}
