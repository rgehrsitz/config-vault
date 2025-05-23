package service

import (
    "context"
    "fmt"
    "regexp"
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
    if err := s.validateAttributeType(name, pattern); err != nil {
        return nil, err
    }
    
    if _, err := s.repo.GetByName(ctx, name); err == nil {
        return nil, fmt.Errorf("attribute type with name '%s' already exists", name)
    }
    
    at := &models.AttributeType{
        ID:      uuid.NewString(),
        Name:    name,
        Pattern: pattern,
    }
    return s.repo.Create(ctx, at)
}

func (s *AttributeService) GetAttributeType(ctx context.Context, id string) (*models.AttributeType, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *AttributeService) GetAttributeTypeByName(ctx context.Context, name string) (*models.AttributeType, error) {
    return s.repo.GetByName(ctx, name)
}

func (s *AttributeService) UpdateAttributeType(ctx context.Context, id, name, pattern string) (*models.AttributeType, error) {
    if err := s.validateAttributeType(name, pattern); err != nil {
        return nil, err
    }
    
    existing, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    if existing.Name != name {
        if _, err := s.repo.GetByName(ctx, name); err == nil {
            return nil, fmt.Errorf("attribute type with name '%s' already exists", name)
        }
    }
    
    updated := &models.AttributeType{
        ID:      id,
        Name:    name,
        Pattern: pattern,
    }
    
    return s.repo.Update(ctx, updated)
}

func (s *AttributeService) DeleteAttributeType(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}

func (s *AttributeService) ListAttributeTypes(ctx context.Context) ([]*models.AttributeType, error) {
    return s.repo.List(ctx)
}

func (s *AttributeService) ValidateValue(ctx context.Context, attributeTypeName, value string) error {
    at, err := s.repo.GetByName(ctx, attributeTypeName)
    if err != nil {
        return fmt.Errorf("unknown attribute type: %s", attributeTypeName)
    }
    
    matched, err := regexp.MatchString(at.Pattern, value)
    if err != nil {
        return fmt.Errorf("invalid regex pattern in attribute type '%s': %w", attributeTypeName, err)
    }
    
    if !matched {
        return fmt.Errorf("value '%s' does not match pattern for attribute type '%s'", value, attributeTypeName)
    }
    
    return nil
}

func (s *AttributeService) validateAttributeType(name, pattern string) error {
    if name == "" {
        return fmt.Errorf("attribute type name cannot be empty")
    }
    
    if pattern == "" {
        return fmt.Errorf("attribute type pattern cannot be empty")
    }
    
    if _, err := regexp.Compile(pattern); err != nil {
        return fmt.Errorf("invalid regex pattern: %w", err)
    }
    
    return nil
}
