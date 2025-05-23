# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

ConfigVault is a cross-platform desktop application for managing hierarchical equipment inventories with Git-backed configuration snapshots and attribute-level diffing. Built with:

- **Backend**: Go with Wails v2.10.1 framework
- **Frontend**: Idiomatic Svelte 5 + TypeScript + Tailwind CSS 4 + Vite
- **Architecture**: Desktop app with embedded web frontend

## Development Commands

### Running the Application

```bash
wails dev                    # Start development server with hot reload
```

### Building

```bash
wails build                  # Build production executable for current platform
```

### Frontend Development

```bash
cd frontend
npm install                  # Install frontend dependencies
npm run dev                  # Run Vite dev server (standalone)
npm run build               # Build frontend for production
npm run check               # Run Svelte type checking
```

### Backend Development

```bash
go mod tidy                  # Clean up Go dependencies
go run .                     # Run Go backend directly (for testing)
```

## Architecture

### Directory Structure

- `/` - Root contains main Wails application (`main.go`, `app.go`)
- `/backend/` - Go backend services and business logic
  - `/models/` - Data structures (Node, AttributeType, etc.)
  - `/repository/` - Data access layer interfaces and implementations
  - `/service/` - Business logic layer
  - `/util/` - Utilities (JSON schema validation, Git operations)
  - `/schema/` - JSON schema definitions
- `/frontend/` - Svelte frontend application
  - `/src/components/` - Reusable Svelte components
  - `/src/api/` - Generated Wails API bindings
  - `/src/stores/` - Svelte state management
  - `/src/plugin-sdk.js` - Plugin system API

### Key Patterns

- **Repository Pattern**: Data access abstracted through interfaces in `/backend/repository/`
- **Service Layer**: Business logic in `/backend/service/` coordinates between repositories
- **Component Contracts**: Frontend components follow contracts defined in `SVELTE_COMPONENT_CONTRACTS.md`
- **Plugin Architecture**: Extensible via WASM/JS plugins with defined hook points

### Backend Services

The application uses a layered architecture:

1. **Models** - Domain entities mirroring JSON schema
2. **Repositories** - Data access (in-memory, file-based, Git-backed)
3. **Services** - Business logic and coordination
4. **Wails Bindings** - Bridge between Go backend and Svelte frontend

### Frontend Architecture

- **Three-column layout**: Hierarchy tree, details panel, snapshots timeline
- **State management**: Svelte stores for application state
- **Component-based**: Reusable components with defined prop/event contracts
- **API integration**: Wails-generated bindings for Go backend calls

## Key Features Being Implemented

- Hierarchical node management with drag-and-drop
- Custom attribute types with regex validation
- Git-backed configuration snapshots
- Side-by-side visual diffing
- Offline capability with sync
- Plugin system for extensibility
- Role-based access control
- Export capabilities (PDF, Markdown, JSON)

## Testing

No specific test commands are configured yet. When implementing tests:

- Go: Use standard `go test` for backend
- Frontend: Consider adding test scripts to `package.json`

## Notes

- Application title currently shows "Wails/Vite/Svelte/Tailwind/Typescript" - update in `main.go`
- This is an early-stage project transitioning from template to ConfigVault implementation
- Plugin system architecture is designed but not yet implemented
- Git integration utilities are planned in `/backend/util/git.go`
