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
