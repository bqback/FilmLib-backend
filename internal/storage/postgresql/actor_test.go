package postgresql

import (
	"context"
	"errors"
	"filmlib/internal/apperrors"
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
		{
			name: "Bad request (could not build query)",
			args: args{
				info: dto.NewActor{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Could not execute query",
			args: args{
				info: dto.NewActor{
					Name:      "Test name",
					Gender:    "male",
					BirthDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := squirrel.
						Insert(actorTable).
						Columns(allActorInsertFields...).
						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
						PlaceholderFormat(squirrel.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
						WillReturnError(errors.New("Mock insert query fail"))
				},
			},
			wantErr: true,
			err:     apperrors.ErrActorNotCreated,
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

func TestActorStorage_Read(t *testing.T) {
	t.Parallel()
	type args struct {
		id    dto.ActorID
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Successful read",
			args: args{
				id: dto.ActorID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := squirrel.
						Select(allActorSelectFields...).
						From(actorTable).
						Where(squirrel.Eq{actorIDField: args.id.Value}).
						PlaceholderFormat(squirrel.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows([]string{shortIDField}).
							AddRow(1, "Test actor", "female", time.Now()),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				id: dto.ActorID{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Could not execute query",
			args: args{
				id: dto.ActorID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := squirrel.
						Select(allActorSelectFields...).
						From(actorTable).
						Where(squirrel.Eq{actorIDField: args.id.Value}).
						PlaceholderFormat(squirrel.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnError(errors.New("Mock select query fail"))
				},
			},
			wantErr: true,
			err:     apperrors.ErrActorNotCreated,
		},
		{
			name: "No results",
			args: args{
				id: dto.ActorID{
					Value: 99,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := squirrel.
						Select(allActorSelectFields...).
						From(actorTable).
						Where(squirrel.Eq{actorIDField: args.id.Value}).
						PlaceholderFormat(squirrel.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows()
				},
			},
			wantErr: true,
			err:     apperrors.ErrEmptyResult,
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

			_, err = s.Read(ctx, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestActorStorage_Update(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.UpdatedActor
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Successful update",
			args: args{
				info: dto.UpdatedActor{
					ID: 1,
					Values: map[string]interface{}{
						"name":   "Test name",
						"gender": "male",
						"dob":    time.Now(),
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					id := 1
					sql, _, _ := squirrel.
						Insert(actorTable).
						Columns(allActorInsertFields...).
						Values(args.info.Values["name"], args.info.Values["gender"], args.info.Values["dob"]).
						PlaceholderFormat(squirrel.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Values["name"], args.info.Values["gender"], args.info.Values["dob"]).
						WillReturnRows(sqlmock.NewRows([]string{shortIDField}).
							AddRow(id),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				info: dto.UpdatedActor{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Could not execute query",
			args: args{
				info: dto.UpdatedActor{
					ID: 1,
					Values: map[string]interface{}{
						"name": "foo",
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := squirrel.
						Update(actorTable).
						Set(actorNameField, args.info.Values["name"]).
						Where(squirrel.Eq{actorIDField: args.info.ID}).
						PlaceholderFormat(squirrel.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.info.ID, args.info.Values["name"]).
						WillReturnError(errors.New("Mock update query fail"))
				},
			},
			wantErr: true,
			err:     apperrors.ErrActorNotCreated,
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

			_, err = s.Update(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

// func TestActorStorage_Delete(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		info  dto.NewActor
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{
// 		{
// 			name: "Successful creation",
// 			args: args{
// 				info: dto.NewActor{
// 					Name:      "Test name",
// 					Gender:    "male",
// 					BirthDate: time.Now(),
// 				},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 					id := 1
// 					sql, _, _ := squirrel.
// 						Insert(actorTable).
// 						Columns(allActorInsertFields...).
// 						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						PlaceholderFormat(squirrel.Dollar).
// 						Suffix("RETURNING id").
// 						ToSql()
// 					mock.ExpectQuery(regexp.QuoteMeta(sql)).
// 						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						WillReturnRows(sqlmock.NewRows([]string{shortIDField}).
// 							AddRow(id),
// 						)
// 				},
// 			},
// 			wantErr: false,
// 			err:     nil,
// 		},
// 		{
// 			name: "Bad request (could not build query)",
// 			args: args{
// 				info: dto.NewActor{},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 				},
// 			},
// 			wantErr: true,
// 			err:     apperrors.ErrCouldNotBuildQuery,
// 		},
// 		{
// 			name: "Could not execute query",
// 			args: args{
// 				info: dto.NewActor{
// 					Name:      "Test name",
// 					Gender:    "male",
// 					BirthDate: time.Now(),
// 				},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 					sql, _, _ := squirrel.
// 						Insert(actorTable).
// 						Columns(allActorInsertFields...).
// 						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						PlaceholderFormat(squirrel.Dollar).
// 						Suffix("RETURNING id").
// 						ToSql()
// 					mock.ExpectQuery(regexp.QuoteMeta(sql)).
// 						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						WillReturnError(errors.New("Mock insert query fail"))
// 				},
// 			},
// 			wantErr: true,
// 			err:     apperrors.ErrActorNotCreated,
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			db, mock, err := sqlmock.Newx()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			tt.args.query(mock, tt.args)

// 			ctx := context.WithValue(
// 				context.WithValue(context.Background(), dto.LoggerKey, logging.BasicTestLogger()),
// 				dto.RequestIDKey, uuid.New().String(),
// 			)

// 			s := NewActorStorage(db)

// 			_, err = s.Create(ctx, tt.args.info)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetUsers() error = %v, wantErr %v", err != nil, tt.wantErr)
// 			}

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

// func TestActorStorage_GetActorMovies(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		info  dto.NewActor
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{
// 		{
// 			name: "Successful creation",
// 			args: args{
// 				info: dto.NewActor{
// 					Name:      "Test name",
// 					Gender:    "male",
// 					BirthDate: time.Now(),
// 				},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 					id := 1
// 					sql, _, _ := squirrel.
// 						Insert(actorTable).
// 						Columns(allActorInsertFields...).
// 						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						PlaceholderFormat(squirrel.Dollar).
// 						Suffix("RETURNING id").
// 						ToSql()
// 					mock.ExpectQuery(regexp.QuoteMeta(sql)).
// 						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						WillReturnRows(sqlmock.NewRows([]string{shortIDField}).
// 							AddRow(id),
// 						)
// 				},
// 			},
// 			wantErr: false,
// 			err:     nil,
// 		},
// 		{
// 			name: "Bad request (could not build query)",
// 			args: args{
// 				info: dto.NewActor{},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 				},
// 			},
// 			wantErr: true,
// 			err:     apperrors.ErrCouldNotBuildQuery,
// 		},
// 		{
// 			name: "Could not execute query",
// 			args: args{
// 				info: dto.NewActor{
// 					Name:      "Test name",
// 					Gender:    "male",
// 					BirthDate: time.Now(),
// 				},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 					sql, _, _ := squirrel.
// 						Insert(actorTable).
// 						Columns(allActorInsertFields...).
// 						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						PlaceholderFormat(squirrel.Dollar).
// 						Suffix("RETURNING id").
// 						ToSql()
// 					mock.ExpectQuery(regexp.QuoteMeta(sql)).
// 						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						WillReturnError(errors.New("Mock insert query fail"))
// 				},
// 			},
// 			wantErr: true,
// 			err:     apperrors.ErrActorNotCreated,
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			db, mock, err := sqlmock.Newx()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			tt.args.query(mock, tt.args)

// 			ctx := context.WithValue(
// 				context.WithValue(context.Background(), dto.LoggerKey, logging.BasicTestLogger()),
// 				dto.RequestIDKey, uuid.New().String(),
// 			)

// 			s := NewActorStorage(db)

// 			_, err = s.Create(ctx, tt.args.info)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetUsers() error = %v, wantErr %v", err != nil, tt.wantErr)
// 			}

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

// func TestActorStorage_GetAll(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		info  dto.NewActor
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{
// 		{
// 			name: "Successful creation",
// 			args: args{
// 				info: dto.NewActor{
// 					Name:      "Test name",
// 					Gender:    "male",
// 					BirthDate: time.Now(),
// 				},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 					id := 1
// 					sql, _, _ := squirrel.
// 						Insert(actorTable).
// 						Columns(allActorInsertFields...).
// 						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						PlaceholderFormat(squirrel.Dollar).
// 						Suffix("RETURNING id").
// 						ToSql()
// 					mock.ExpectQuery(regexp.QuoteMeta(sql)).
// 						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						WillReturnRows(sqlmock.NewRows([]string{shortIDField}).
// 							AddRow(id),
// 						)
// 				},
// 			},
// 			wantErr: false,
// 			err:     nil,
// 		},
// 		{
// 			name: "Bad request (could not build query)",
// 			args: args{
// 				info: dto.NewActor{},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 				},
// 			},
// 			wantErr: true,
// 			err:     apperrors.ErrCouldNotBuildQuery,
// 		},
// 		{
// 			name: "Could not execute query",
// 			args: args{
// 				info: dto.NewActor{
// 					Name:      "Test name",
// 					Gender:    "male",
// 					BirthDate: time.Now(),
// 				},
// 				query: func(mock sqlmock.Sqlmock, args args) {
// 					sql, _, _ := squirrel.
// 						Insert(actorTable).
// 						Columns(allActorInsertFields...).
// 						Values(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						PlaceholderFormat(squirrel.Dollar).
// 						Suffix("RETURNING id").
// 						ToSql()
// 					mock.ExpectQuery(regexp.QuoteMeta(sql)).
// 						WithArgs(args.info.Name, args.info.Gender, args.info.BirthDate).
// 						WillReturnError(errors.New("Mock insert query fail"))
// 				},
// 			},
// 			wantErr: true,
// 			err:     apperrors.ErrActorNotCreated,
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			db, mock, err := sqlmock.Newx()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			tt.args.query(mock, tt.args)

// 			ctx := context.WithValue(
// 				context.WithValue(context.Background(), dto.LoggerKey, logging.BasicTestLogger()),
// 				dto.RequestIDKey, uuid.New().String(),
// 			)

// 			s := NewActorStorage(db)

// 			_, err = s.Create(ctx, tt.args.info)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetUsers() error = %v, wantErr %v", err != nil, tt.wantErr)
// 			}

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }
