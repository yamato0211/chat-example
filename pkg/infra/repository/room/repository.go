package room

import (
	"context"
	"sync"

	"github.com/sivchari/chat-example/pkg/domain/entity"
	"github.com/sivchari/chat-example/pkg/domain/repository/room"
)

type repository struct {
	mapByID map[string]*entity.Room
	// TODO: sync.RWMutexとの違いを考えて最適化しよう
	mu sync.RWMutex
}

func New() room.Repository {
	return &repository{
		mapByID: make(map[string]*entity.Room, 0),
	}
}

func (r *repository) Insert(_ context.Context, room *entity.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.mapByID[room.ID] = room
	return nil
}

func (r *repository) Select(_ context.Context, id string) (*entity.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	room, ok := r.mapByID[id]
	if !ok {
		return &entity.Room{}, nil
	}
	return room, nil
}

func (r *repository) SelectAll(_ context.Context) ([]*entity.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rooms := make([]*entity.Room, 0, len(r.mapByID))
	for _, e := range r.mapByID {
		rooms = append(rooms, e)
	}
	return rooms, nil
}
