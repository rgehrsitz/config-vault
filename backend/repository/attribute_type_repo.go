package repository

import (
    "context"
    "sync"
    "github.com/rgehrsitz/config-vault/backend/models"
)

type AttributeTypeRepo interface {
    Create(ctx context.Context, at *models.AttributeType) (*models.AttributeType, error)
    List(ctx context.Context) ([]*models.AttributeType, error)
}

type InMemoryAttributeTypeRepo struct {
    mu sync.RWMutex
    data map[string]*models.AttributeType
}

func NewInMemoryAttributeTypeRepo() *InMemoryAttributeTypeRepo {
    return &InMemoryAttributeTypeRepo{
        data: make(map[string]*models.AttributeType),
    }
}

func (r *InMemoryAttributeTypeRepo) Create(ctx context.Context, at *models.AttributeType) (*models.AttributeType, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.data[at.ID] = at
    return at, nil
}

func (r *InMemoryAttributeTypeRepo) List(ctx context.Context) ([]*models.AttributeType, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    out := make([]*models.AttributeType, 0, len(r.data))
    for _, v := range r.data {
        out = append(out, v)
    }
    return out, nil
}
