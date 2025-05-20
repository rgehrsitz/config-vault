**Svelte Component Contracts for Archon**

Below are the core Svelte components for Archon’s MVP slice, their file locations, props, and dispatched events. These contracts ensure a clear contract between UI and backend.

---

## 1. `HierarchyTree.svelte`

**Location:** `frontend/src/components/HierarchyTree.svelte`

**Props:**

```ts
export let nodes: Node[];
export let selectedNodeIds: string[];
export let searchQuery: string;
export let filters: Filter[];
```

**Events:**

* `on:select` → `{ nodeId: string }`
* `on:multiSelect` → `{ nodeIds: string[] }`
* `on:createNode` → `{ parentId: string; newNodeData?: Partial<Node> }`
  *Emits intent to create; parent component may collect new node fields.*
* `on:copy` → `{ nodeIds: string[] }`
* `on:paste` → `{ targetParentId: string }`
* `on:move` → `{ nodeId: string; newParentId: string; position: number }`
* `on:rename` → `{ nodeId: string; newName: string }`
* `on:delete` → `{ nodeId: string }`

---

## 2. `NodeEditor.svelte`

**Location:** `frontend/src/components/NodeEditor.svelte`

**Props:**

```ts
export let node: Node | Node[];
export let mode: 'view' | 'edit';
export let attributeTypes: AttributeType[];
export let bulkMode: boolean;
```

**Events:**

* `on:editToggle` → `{}`  *Toggle between view/edit modes.*
* `on:change` → `{ field: string; value: any }`  *Field-level change.*
* `on:save` → `{}`  *Commit all pending changes.*
* `on:discard` → `{}`  *Discard unsaved changes.*
* `on:bulkUpdate` → `{ updates: BulkUpdatePayload }`  *Perform batch updates.*

---

## 3. `SnapshotTimeline.svelte`

**Location:** `frontend/src/components/SnapshotTimeline.svelte`

**Props:**

```ts
export let commits: Commit[];
export let selectedBase: string | null;
export let selectedCompare: string | null;
export let offlineCount: number;
```

**Events:**

* `on:selectBase` → `{ sha: string }`
* `on:selectCompare` → `{ sha: string }`
* `on:saveSnapshot` → `{ message: string }`
* `on:diff` → `{ baseSha: string; compareSha: string }`
* `on:sync` → `{}`  *Trigger offline sync attempt.*

---

## 4. `DiffView.svelte`

**Location:** `frontend/src/components/DiffView.svelte`

**Props:**

```ts
export let beforeTree: Node[];
export let afterTree: Node[];
export let filters: { adds: boolean; deletes: boolean; modifies: boolean };
```

**Events:**

* `on:export` → `{ format: 'markdown' | 'pdf' | 'json' }`
* `on:back` → `{}`

---

## 5. `AttributeTypeSettings.svelte`

**Location:** `frontend/src/components/AttributeTypeSettings.svelte`

**Props:**

```ts
export let attributeTypes: AttributeType[];
```

**Events:**

* `on:addType` → `{ newAttributeType?: AttributeType }`  *Intent to add; parent may open modal.*
* `on:editType` → `{ originalTypeName: string; updatedAttributeType?: AttributeType }`  *Signal edit or completion.*
* `on:deleteType` → `{ typeName: string }`
* `on:saveSettings` → `{}`

---

## 6. `UserManagement.svelte`

**Location:** `frontend/src/components/UserManagement.svelte`

**Props:**

```ts
export let users: User[];
export let roles: Role[];
```

**Events:**

* `on:addUser` → `{ newUserData?: User }`  *Intent to add; parent may open form.*
* `on:editUser` → `{ userId: string; updatedUserData?: Partial<User> }`
* `on:deleteUser` → `{ userId: string }`
* `on:assignRole` → `{ userId: string; role: string }`
