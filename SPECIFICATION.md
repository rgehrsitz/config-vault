# ConfigVault: Comprehensive Design Specification

## Executive Summary

**ConfigVault** is a cross‑platform desktop application (Go/Wails backend, Svelte/Tailwind/TypeScript frontend) for managing hierarchical equipment inventories, capturing Git‑backed configuration snapshots, and performing interactive, attribute‑level diffs. It targets diverse personas—Lab Managers, Field Technicians, Compliance Auditors, Junior Viewers, and Administrators—addressing offline workflows, bulk updates, custom validations, and audit‑grade histories. This document serves as the single source of truth for product scope, architecture, UI, APIs, plugins, and roadmap.

---

## 1. Personas & User Context

### 1.1 Lab Manager: *Sophia Chen*

* **Context & Environment**: Central R\&D office on Windows desktop; oversees 5 labs, \~200 devices each; uses both UI and occasional CLI scripting.
* **Day‑in‑Life Scenario**: Sophia arrives, reviews last night’s automated test‑bench firmware updates, audits any failures, and adds new racks to next sprint’s test plan.
* **Goals**:

  * Rapidly onboard new rack hardware into CMDB.
  * Enforce component version policies.
  * Audit historical configurations for incident analysis.
* **Frustrations**:

  * Divergent spreadsheets; lost email attachments.
  * Manual merge conflicts when multiple engineers update same rack.

### 1.2 Field Technician: *Alex Martinez*

* **Context & Environment**: Laptop-bound (Windows/macOS), sometimes air‑gapped customer/test sites; occasionally on Linux VM; limited or no network.
* **Day‑in‑Life**: Drive to a remote test facility, snapshot current rack config, apply firmware patches, diff to confirm updates, later sync changes back to central repo.
* **Goals**:

  * Offline snapshotting with seamless sync on reconnect.
  * Quick inline fixes of typos or serial mismatches.
* **Frustrations**:

  * JSON hard to read without good UI.
  * Lack of built‑in diffing tools for nested attributes.

### 1.3 Compliance Auditor: *Jordan Patel*

* **Context**: Corporate office with tight security; reviews system configuration records for compliance (ISO, FDA, etc.).
* **Use Case**: Pull two snapshots six months apart, diff firmware changes on network cards, generate report for auditors.
* **Goals**:

  * Attribute‑level diff with clear visual highlights.
  * Exportable, human‑readable compliance reports.
* **Frustrations**:

  * High‑level CMDB logs that miss nested changes.
  * Snapshots scattered across systems.

### 1.4 Junior Technician / Viewer: *Chris Green*

* **Context**: IT support on Mac; read‑only role for ad‑hoc lookups.
* **Goals**:

  * Navigate hierarchy quickly.
  * View current or past states without confusion.
* **Frustrations**:

  * Cluttered UI with editing options they can’t use.
  * Difficulty locating the “latest” snapshot.

### 1.5 Administrator: *Pat Lee*

* **Context**: IT sysadmin; sets up ConfigVault instances; manages users and global settings; edits config files.
* **Goals**:

  * RBAC: assign roles (Viewer, Editor, Auditor, Admin).
  * Configure repo location and schema version.
* **Frustrations**:

  * Lack of granular permission controls.
  * Inconsistent repo paths across OS installs.

---

## 2. User Stories & Acceptance Criteria

| ID | Persona | Story                                        | Acceptance Criteria      | Priority |
| -- | ------- | -------------------------------------------- | ------------------------ | -------- |
| 1  | Sophia  | Define custom attribute types (name + regex) | - UI for defining types. |          |

* Regex tested live with sample input.
* New types appear in import/edit dropdowns.                  | Must     |
  \| 2   | Sophia               | Create/rename/move/re‑parent hierarchy nodes                                              | - Context menu & drag‑drop.
* Undo capability.
* Tree updates in real time.                                       | Must     |
  \| 3   | Alex                 | Import hierarchy from JSON                                                                | - Schema validation errors shown.
* Imported tree merges without duplicates.                           | Could    |
  \| 4   | Any                  | Search and filter nodes by name/type/attribute                                            | - Live filter collapses non‑matches but shows parent path.
* “No results” message.                              | Must     |
  \| 5   | Alex                 | Offline snapshots + sync                                                                   | - Local commits queue.
* Offline badge & pending count.
* Sync button triggers push/pull.                                 | Could    |
  \| 6   | Alex                 | Conflict alert & merge UI                                                                  | - Visual conflict dialog.
* Options: Accept ours/theirs or merge manually.                         | Won’t    |
  \| 7   | Sophia/Alex          | Custom commit messages on snapshot                                                        | - Modal to edit message.
* Preview auto‑generated then save.                                       | Should   |
  \| 8   | Pat                  | Initialize Git repo at known path with `.gitignore`                                       | - Configurable path.
* `.gitignore` covers temp files by default.                            | Should   |
  \| 9   | Jordan               | Filter diff results by change type                                                        | - Checkboxes for Adds/Deletes/Modifies.
* Filter applies instantly.                                           | Could    |
  \| 10  | Jordan               | Side‑by‑side, color‑coded tree diff                                                        | - Trees scroll sync’d.
* Legend for colors.                                                     | Should   |
  \| 11  | Jordan               | Export diffs (Markdown, PDF, JSON)                                                        | - Export button.
* File auto-download with correct format.                                | Should   |
  \| 12  | Jordan               | Export full snapshot as human‑readable report                                              | - Export full config.
* Format option (PDF/Text).                                            | Could    |
  \| 13  | Alex                 | Inline edit properties with dirty indicators                                               | - Dirty dot gutter.
* Red outline + tooltip for invalid values.                                   | Should   |
  \| 14  | Sophia               | Bulk‑update selected nodes                                                                 | - Bulk toolbar appears on multi‑select.
* Dry‑run preview modal.
* Apply button.                 | Could    |
  \| 15  | Sophia               | Undo/discard pending changes                                                               | - “Discard All” in toolbar.
* Confirmation prompt.                                               | Should   |
  \| 16  | Any                  | Validate saved JSON against schema                                                         | - `$schema` field present.
* Validation run on import/save with UI feedback.                          | Must     |
  \| 17  | Any                  | Native executables for Windows/macOS/Linux                                                | - Bundled via Wails build.
* Zero external dependencies.                                          | Could    |
  \| 18  | Pat                  | Role‑based access control                                                                  | - Settings UI to assign roles.
* UI/CLI hide disabled actions.                                            | Should   |
  \| 19  | Sophia               | Legacy CSV import                                                                          | - CSV→JSON mapping UI.
* Field mapping dialog & validation.                                     | Could    |
  \| 20  | Developer            | CLI/REST API for scripting                                                                 | - `archon snapshot`, `archon diff`.
* OpenAPI spec generated.                                             | Should   |

---

## 3. Architecture & Code Layout

### 3.1 Directory Structure

```bash
configvault/
├── backend/
│   ├── main.go                 # Wails bootstrap
│   ├── plugins/                # Plugin manager & hooks
│   │   └── hook_manager.go
│   ├── models/                 # Go structs mirror JSON schema
│   │   ├── node.go
│   │   └── attribute.go
│   ├── repository/             # Interfaces + impl (file, memory)
│   ├── service/                # Business logic
│   ├── util/                   # JSON schema, Git wrappers
│   └── schema/archon-v1.json   # JSON schema file
└── frontend/
    ├── src/
    │   ├── api/                # Generated RPC clients
    │   ├── components/         # Svelte components
    │   ├── plugin-sdk.js       # JS plugin host API
    │   ├── stores/             # Svelte stores
    │   └── App.svelte          # Main layout
    ├── package.json
    └── tailwind.config.js
```

### 3.2 Key Backend Interfaces

```go
// repository/node_repo.go
// Defines CRUD and hierarchical operations for nodes.
type NodeRepository interface {
    Save(ctx context.Context, node *models.Node) (*models.Node, error)
    Move(ctx context.Context, nodeID, newParentID string, pos int) error
    Search(ctx context.Context, query string) ([]*models.Node, error)
    Filter(ctx context.Context, filters []models.Filter) ([]*models.Node, error)
}

// repository/attribute_type_repo.go
// Manages custom attribute types and patterns.
type AttributeTypeRepository interface {
    Save(ctx context.Context, t *models.AttributeType) error
    List(ctx context.Context) ([]*models.AttributeType, error)
    GetByName(ctx context.Context, name string) (*models.AttributeType, error)
}
```

### 3.3 JSON Schema & Git Utilities

* **JSON Schema Validation** (`util/jsonschema.go`)

  * Validates configuration JSON against `schema/archon-v1.json` using gojsonschema.
  * Error details include node path and expected pattern.
* **Git Wrapper** (`util/git.go`)

  * Initializes embedded repository.
  * Exposes functions: `InitRepo(path)`, `CommitAll(message)`, `ListCommits()`, `DiffCommits(base, head)`.
  * Queues commits offline and syncs on demand.

---

## 4. UI Specification

### 4.1 Three‑Column Editor View

**Hierarchy Panel**

* Search input with live filtering and parent-path retention.
* Filter dropdown for attribute criteria.
* Tree with ARIA roles, ⋮ context menus (Add, Rename, Move, Copy, Paste, Delete).
* Drag‑and‑drop with ghost preview and undo toast.

**Details Panel**

* Header with “Details” and an **Edit** toggle.
* Property sheet table: Name, Type, and each Attribute on its own row.
* **View Mode**: static text.
* **Edit Mode**: inline inputs, dirty dot gutter, red outline + error icon tooltips.
* Bulk‑Update toolbar on multi‑select: Set Attribute, Add/Remove Attribute, Bump Firmware; dry-run diff preview; **Apply** button.

**Snapshots Panel**

* Header: “Snapshots” + offline badge with pending count.
* Filters: Author, Date Range, Message Keyword.
* Actions: **Save Snapshot** (editable message modal), **Diff Selected** (enabled when two selected).
* Commit list with checkboxes: single click for Base (blue), ⇧‑click for Compare (teal).

### 4.2 Diff & Reporting View

* **Navigation**: back button to Editor.
* **Filter Toggles**: Adds, Deletes, Modifies; live update.
* **Side‑by‑Side Trees**: sync scroll, color‑coded highlights (+ green, - red, \~ yellow).
* **Export Menu**: options for Markdown, PDF, JSON, and Copy to Clipboard.
* **Pagination/Lazy Load**: for large diffs.

### 4.3 Settings & Management

* **Tabs**: Attribute Types, User Management.
* **Attribute Types**: table with Name & Regex Pattern; Add/Edit/Delete actions; confirmation on delete if in use.
* **User Management**: table of Users with Role dropdown; Add/Edit/Delete with form modals; password reset.
* **JSON Schema Errors**: modal listing validation issues with links to offending nodes.

### 4.4 Responsiveness & Accessibility

* **<1400px**: stack Hierarchy + Details; Snapshots becomes collapsible sidebar.
* **<800px**: tabbed full‑screen view.
* Keyboard nav: focus management, ARIA roles for tree/table, shortcuts (Ctrl+S, Ctrl+Z).
* High‑contrast mode support; color‑blind palettes.

---

## 5. Svelte Component Contracts

Refer to `frontend/src/components/Svelte_Component_Contracts.md` for detailed props/events for:

* `HierarchyTree.svelte`
* `NodeEditor.svelte`
* `SnapshotTimeline.svelte`
* `DiffView.svelte`
* `AttributeTypeSettings.svelte`
* `UserManagement.svelte`

---

## 6. Plugin Architecture

1. **Packaging & Discovery**: `~/.configvault/plugins/*/plugin.json` manifests.
2. **Hook Points**: Pre/Post NodeCreate, Pre/Post Snapshot, OnSearch, OnDiffCompute, Pre/Post BulkUpdate.
3. **Engines**: WASM & JS; host APIs for node access, logging, cancelation.
4. **UI Extensions**: inject DetailPanel widgets, new Settings pages, SnapshotPanel slots.
5. **Security**: SemVer plugin versioning, sandboxed WASM, signed manifests optional.

---

## 7. Naming & Positioning

* **Product Name**: ConfigVault
* **Tagline**: “Your secure vault for equipment configurations.”
* **Market Fit**: unique cross‑platform desktop CMDB with Git snapshots & inline diffs.

---

## 8. Next Steps

1. **Finalize Backend**: implement services & plugin manager.
2. **Frontend Scaffold**: create Svelte components per contracts & wire API.
3. **CLI & API**: build CLI commands and OpenAPI spec.
4. **Plugin SDK**: draft JS/WASM examples and CLI scaffolding (`configvault plugin init`).
5. **Open‑Source Launch**: repo setup, docs, contribution guidelines, community outreach.
