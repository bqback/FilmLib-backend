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
