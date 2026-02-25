---
title: Examples
sidebar_position: 1
---

Full real-world examples of supported input formats.

Each example uses the same project structure but expressed using different formats:

- `structure.md`
- `structure.yml`
- `structure.json`

> **Note:** All formats produce the same internal AST.
>
> - JSON uses object entries to support metadata like `ignore`.
> - YAML supports inline comments; JSON does not.
> - Markdown is the fastest format to write manually.

## `structure.md`

```markdown
# backend

- cmd/
  - api/
    - main.go
  - worker/
    - main.go

- internal/
  - domain/
    - user/
      - entity.go
      - repository.go
      - errors.go
    - auth/
      - token.go
      - claims.go

  - application/
    - user_service.go
    - auth_service.go

  - infra/
    - db/
      - migrate.go
      - sqlite.go
      - sqlcgenerated/ #ignore
    - http/
      - router.go
      - middleware.go

- pkg/
  - logger/
    - logger.go
  - config/
    - loader.go

- scripts/
  - dev.sh
  - release.sh
```

## `structure.yml`

```yaml
backend/:
  cmd/:
    api/:
      - main.go
    worker/:
      - main.go

  internal/:
    domain/:
      user/:
        - entity.go
        - repository.go
        - errors.go
      auth/:
        - token.go
        - claims.go

    application/:
      - user_service.go
      - auth_service.go

    infra/:
      db/:
        - migrate.go
        - sqlite.go
        - sqlcgenerated/ #ignore
      http/:
        - router.go
        - middleware.go

  pkg/:
    logger/:
      - logger.go
    config/:
      - loader.go

  scripts/:
    - dev.sh
    - release.sh
```

## `structure.json`

```json
{
  "backend/": {
    "cmd/": {
      "api/": ["main.go"],
      "worker/": ["main.go"]
    },
    "internal/": {
      "domain/": {
        "user/": ["entity.go", "repository.go", "errors.go"],
        "auth/": ["token.go", "claims.go"]
      },
      "application/": ["user_service.go", "auth_service.go"],
      "infra/": {
        "db/": [
          "migrate.go",
          "sqlite.go",
          { "name": "sqlcgenerated/", "ignore": true }
        ],
        "http/": ["router.go", "middleware.go"]
      }
    },
    "pkg/": {
      "logger/": ["logger.go"],
      "config/": ["loader.go"]
    },
    "scripts/": ["dev.sh", "release.sh"]
  }
}
```

## Advanced: Template Directives & Language Features

### Multi-Language Project (JSON)

```json
{
  "webapp/": {
    "backend/": {
      "src/": [
        { "name": "main.go", "template": "go" },
        { "name": "config.go", "template": "go" },
        { "name": "database.go", "template": "go" }
      ],
      "tests/": [{ "name": "main_test.go", "template": "go" }]
    },
    "frontend/": {
      "src/": [
        { "name": "App.tsx", "template": "typescript" },
        { "name": "main.ts", "template": "typescript" },
        { "name": "utils.ts", "template": "typescript" }
      ],
      "styles/": [{ "name": "main.css", "template": "none" }]
    },
    "scripts/": [
      { "name": "build.sh", "template": "none" },
      { "name": "deploy.sh", "template": "none" }
    ]
  }
}
```

### With Ignore Rules (YAML)

```yaml
backend/:
  src/:
    - main.go
    - database.go

  migrations/:
    - 001_init.sql
    - 002_users.sql

  generated/: #ignore
    - sqlc.gen.go
    - models.gen.go

  vendor/: #ignore

  tests/:
    - main_test.go
```

### Markdown with Directives

```markdown
# myapp

- backend/ #template:go
  - main.go
  - config.go
  - setup.py #template:python
  - setup.sh #template:none

- frontend/ #template:typescript
  - app.tsx
  - utils.ts

- .generated/ #ignore
- node_modules/ #ignore
```

## Reverse Mode: Scan Existing Project

```bash
# Scan backend directory and export as markdown
scaffoldgen reverse --in ./backend --format md --out structure.md

# Scan and output as JSON
scaffoldgen reverse --in ./src --format json > structure.json

# Output to stdout (YAML)
scaffoldgen reverse --in ./project --format yml
```

Useful for:

- Documenting existing projects
- Migrating from one format to another
- Creating scaffold templates from real layouts
