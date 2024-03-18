package postgresql

import (
	"context"
	"filmlib/internal/logging"
	"filmlib/internal/pkg/dto"
	"regexp"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestActorStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewActor
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Successful creation",
			args: args{
				info: dto.NewActor{
					Name:      "Test name",
					Gender:    "male",
					BirthDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					id := 1
					sql, _, _ := squirrel.
						Insert(actorTable).
						Columns(allActorInsertFields...).
						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
						PlaceholderFormat(squirrel.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
						WillReturnRows(sqlmock.NewRows([]string{shortIDField}).
							AddRow(id),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.Newx()

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.args.query(mock, tt.args)

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, logging.BasicTestLogger()),
				dto.RequestIDKey, uuid.New().String(),
			)

			s := NewActorStorage(db)

			_, err = s.Create(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
