package main

import (
	"context"
	"os"
	"path/filepath"
	
	"github.com/rgehrsitz/config-vault/backend/models"
	"github.com/rgehrsitz/config-vault/backend/repository"
	"github.com/rgehrsitz/config-vault/backend/service"
)

// App struct
type App struct {
	ctx              context.Context
	attributeService *service.AttributeService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	
	// Initialize data directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get user home directory: " + err.Error())
	}
	
	dataDir := filepath.Join(homeDir, ".config-vault")
	attributeTypesPath := filepath.Join(dataDir, "attribute_types.json")
	
	// Initialize repositories
	attributeRepo, err := repository.NewFileAttributeTypeRepo(attributeTypesPath)
	if err != nil {
		panic("Failed to initialize attribute type repository: " + err.Error())
	}
	
	// Initialize services
	a.attributeService = service.NewAttributeService(attributeRepo)
}

// CreateAttributeType creates a new attribute type
func (a *App) CreateAttributeType(name, pattern string) (*models.AttributeType, error) {
	return a.attributeService.CreateAttributeType(a.ctx, name, pattern)
}

// GetAttributeType gets an attribute type by ID
func (a *App) GetAttributeType(id string) (*models.AttributeType, error) {
	return a.attributeService.GetAttributeType(a.ctx, id)
}

// UpdateAttributeType updates an existing attribute type
func (a *App) UpdateAttributeType(id, name, pattern string) (*models.AttributeType, error) {
	return a.attributeService.UpdateAttributeType(a.ctx, id, name, pattern)
}

// DeleteAttributeType deletes an attribute type
func (a *App) DeleteAttributeType(id string) error {
	return a.attributeService.DeleteAttributeType(a.ctx, id)
}

// ListAttributeTypes returns all attribute types
func (a *App) ListAttributeTypes() ([]*models.AttributeType, error) {
	return a.attributeService.ListAttributeTypes(a.ctx)
}

// ValidateAttributeValue validates a value against an attribute type
func (a *App) ValidateAttributeValue(attributeTypeName, value string) error {
	return a.attributeService.ValidateValue(a.ctx, attributeTypeName, value)
}
