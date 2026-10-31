# CLAUDE.md

Full-stack web app: Go (gin) backend + React (Vite, TypeScript) frontend. The backend serves both the API under `/api` and the built frontend, so they share one origin.

## Layout

| Path | What |
| - | - |
| `web/openapi.yaml` | API spec, the single source of truth for every endpoint |
| `web/frontend/src/api/` | Generated typescript-axios client. Never edit by hand |
| `web/frontend/src/apiClient.ts` | The `api` instance every page uses, JWT + 401 handling |
| `web/frontend/src/page/<name>/` | One folder per page (`XxxPage.tsx` + `xxx-page.module.css`) |
| `web/frontend/src/components/` | Shared components (button, modal, pager, sidebar, ...) |
| `web/backend/internal/api_<area>.go` | Routes and gin handlers of one area |
| `web/backend/internal/processor/` | Business logic, returns `*model.ErrorDetail` on failure |
| `web/backend/internal/context/` | State and storage (`dbContext`, `DbIf`) |
| `web/backend/model/` | Request / response structs |
| `config.yaml` / `docker/config.yaml` | Local / docker config, keep both in sync with `config/config.go` |

## Adding an API

1. Add the path and schemas to `web/openapi.yaml`. Paths behind login get `security: [ { bearerAuth: [] } ]`.
2. Run `make openapi` to regenerate `web/frontend/src/api`.
3. Backend:
   - request / response structs in `model/`
   - logic in `processor/<area>.go`
   - a `getXxxRoutes()` + handlers in `internal/api_<area>.go`, each `HandlerFunc` wrapped with `withLogging(name, logger, handler)`
   - register it in `addServices` (`internal/backend.go`): `authGroup` if it needs login, `apiGroup` otherwise
4. Frontend: call it through `api.xxx()` from `apiClient.ts`, show errors with `extractErrorMessage`.

Postman is not used; do not add a collection.

## Rules

- **401 means "our own JWT is invalid", nothing else.** The frontend logs the user out on any 401 (`apiClient.ts`). When a call the backend makes to another service is rejected, answer 502 (or 403), never pass its 401 through.
- **Never hard-code the API host or port in the frontend.** Use `apiBaseUrl` from `apiClient.ts` (the page's own origin). For a WebSocket, derive it: `apiBaseUrl.replace(/^http/, 'ws')`.
- A new logger tag needs both `logger/tag.go` and a field in `logger/logger.go`.
- New config fields go in `config/config.go`, `config.yaml` and `docker/config.yaml`.
- Files embedded with `//go:embed` under `web/backend` trigger a rebuild in `make` like `.go` files do.
- Commit messages follow Conventional Commits (`feat:`, `fix:`, `chore:`, ...), checked in CI.

## Commands

| Task | Command |
| - | - |
| Build backend + frontend | `make` |
| Run | `make run` (serves on the `backend.port` in `config.yaml`, default 8888) |
| Frontend with hot reload | `make run`, then `yarn dev` in `web/frontend` (proxies `/api` to :8888, or `VITE_DEV_API_TARGET`) |
| Regenerate API client | `make openapi` |
| Backend tests | `cd web/backend && go test ./...` |
| Lint | `make lint` (backend), `yarn lint` in `web/frontend` |

Before calling a change done, the backend must build, test and lint clean, and `yarn build` must pass in `web/frontend`.
