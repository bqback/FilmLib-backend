package actor

import (
	"context"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/pkg/entities"
	"filmlib/internal/storage"
)

type ActorService struct {
	as storage.IActorStorage
}

func NewActorService(actorStorage storage.IActorStorage) *ActorService {
	return &ActorService{
		as: actorStorage,
	}
}

func (s *ActorService) Create(ctx context.Context, info dto.NewActor) (*entities.Actor, error) {
	return s.as.Create(ctx, info)
}

func (s *ActorService) Read(ctx context.Context, id dto.ActorID) (*entities.Actor, error) {
	actor, err := s.as.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	movies, err := s.as.GetActorMovies(ctx, dto.ActorID{Value: actor.ID})
	if err != nil {
		return nil, err
	}
	actor.Movies = movies
	return actor, nil
}

func (s *ActorService) Update(ctx context.Context, info dto.UpdatedActor) (*entities.Actor, error) {
	actor, err := s.as.Update(ctx, info)
	if err != nil {
		return nil, err
	}
	movies, err := s.as.GetActorMovies(ctx, dto.ActorID{Value: actor.ID})
	if err != nil {
		return nil, err
	}
	actor.Movies = movies
	return actor, nil
}

func (s *ActorService) Delete(ctx context.Context, id dto.ActorID) error {
	return s.as.Delete(ctx, id)
}

func (s *ActorService) GetActors(ctx context.Context) ([]*entities.Actor, error) {
	actors, err := s.as.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return actors, nil
}
