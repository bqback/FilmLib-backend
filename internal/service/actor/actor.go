package actor

import "filmlib/internal/storage"

type ActorService struct {
	as storage.IActorStorage
}

func NewActorService(actorStorage storage.IActorStorage) *ActorService {
	return &ActorService{
		as: actorStorage,
	}
}
