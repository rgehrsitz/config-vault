package service

import (
    "context"
    "github.com/rgehrsitz/config-vault/backend/models"
    "github.com/rgehrsitz/config-vault/backend/repository"
    "github.com/google/uuid"
)

type AttributeService struct {
    repo repository.AttributeTypeRepo
}

func NewAttributeService(repo repository.AttributeTypeRepo) *AttributeService {
    return &AttributeService{repo: repo}
}

func (s *AttributeService) CreateAttributeType(ctx context.Context, name, pattern string) (*models.AttributeType, error) {
    at := &models.AttributeType{
        ID:      uuid.NewString(),
        Name:    name,
        Pattern: pattern,
    }
    return s.repo.Create(ctx, at)
}

func (s *AttributeService) ListAttributeTypes(ctx context.Context) ([]*models.AttributeType, error) {
    return s.repo.List(ctx)
}
