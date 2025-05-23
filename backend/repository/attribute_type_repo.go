package repository

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "sync"
    "github.com/rgehrsitz/config-vault/backend/models"
)

type AttributeTypeRepo interface {
    Create(ctx context.Context, at *models.AttributeType) (*models.AttributeType, error)
    GetByID(ctx context.Context, id string) (*models.AttributeType, error)
    GetByName(ctx context.Context, name string) (*models.AttributeType, error)
    List(ctx context.Context) ([]*models.AttributeType, error)
    Update(ctx context.Context, at *models.AttributeType) (*models.AttributeType, error)
    Delete(ctx context.Context, id string) error
}

type FileAttributeTypeRepo struct {
    mu       sync.RWMutex
    dataPath string
    data     map[string]*models.AttributeType
}

func NewFileAttributeTypeRepo(dataPath string) (*FileAttributeTypeRepo, error) {
    repo := &FileAttributeTypeRepo{
        dataPath: dataPath,
        data:     make(map[string]*models.AttributeType),
    }
    
    if err := repo.ensureDataDir(); err != nil {
        return nil, err
    }
    
    if err := repo.load(); err != nil {
        return nil, err
    }
    
    return repo, nil
}

func (r *FileAttributeTypeRepo) ensureDataDir() error {
    dir := filepath.Dir(r.dataPath)
    return os.MkdirAll(dir, 0755)
}

func (r *FileAttributeTypeRepo) load() error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, err := os.Stat(r.dataPath); os.IsNotExist(err) {
        return nil
    }
    
    data, err := os.ReadFile(r.dataPath)
    if err != nil {
        return fmt.Errorf("failed to read attribute types file: %w", err)
    }
    
    var types []*models.AttributeType
    if err := json.Unmarshal(data, &types); err != nil {
        return fmt.Errorf("failed to unmarshal attribute types: %w", err)
    }
    
    r.data = make(map[string]*models.AttributeType)
    for _, at := range types {
        r.data[at.ID] = at
    }
    
    return nil
}

func (r *FileAttributeTypeRepo) save() error {
    types := make([]*models.AttributeType, 0, len(r.data))
    for _, at := range r.data {
        types = append(types, at)
    }
    
    data, err := json.MarshalIndent(types, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal attribute types: %w", err)
    }
    
    if err := os.WriteFile(r.dataPath, data, 0644); err != nil {
        return fmt.Errorf("failed to write attribute types file: %w", err)
    }
    
    return nil
}

func (r *FileAttributeTypeRepo) Create(ctx context.Context, at *models.AttributeType) (*models.AttributeType, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, exists := r.data[at.ID]; exists {
        return nil, fmt.Errorf("attribute type with ID %s already exists", at.ID)
    }
    
    r.data[at.ID] = at
    
    if err := r.save(); err != nil {
        delete(r.data, at.ID)
        return nil, err
    }
    
    return at, nil
}

func (r *FileAttributeTypeRepo) GetByID(ctx context.Context, id string) (*models.AttributeType, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    at, exists := r.data[id]
    if !exists {
        return nil, fmt.Errorf("attribute type with ID %s not found", id)
    }
    
    return at, nil
}

func (r *FileAttributeTypeRepo) GetByName(ctx context.Context, name string) (*models.AttributeType, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    for _, at := range r.data {
        if at.Name == name {
            return at, nil
        }
    }
    
    return nil, fmt.Errorf("attribute type with name %s not found", name)
}

func (r *FileAttributeTypeRepo) List(ctx context.Context) ([]*models.AttributeType, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    types := make([]*models.AttributeType, 0, len(r.data))
    for _, at := range r.data {
        types = append(types, at)
    }
    
    return types, nil
}

func (r *FileAttributeTypeRepo) Update(ctx context.Context, at *models.AttributeType) (*models.AttributeType, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, exists := r.data[at.ID]; !exists {
        return nil, fmt.Errorf("attribute type with ID %s not found", at.ID)
    }
    
    r.data[at.ID] = at
    
    if err := r.save(); err != nil {
        return nil, err
    }
    
    return at, nil
}

func (r *FileAttributeTypeRepo) Delete(ctx context.Context, id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    if _, exists := r.data[id]; !exists {
        return fmt.Errorf("attribute type with ID %s not found", id)
    }
    
    delete(r.data, id)
    
    return r.save()
}
