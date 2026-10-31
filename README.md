# full-stack-framework

This is framework project for quickly build a website with full-stack (React + go).

## Develop Environment

| DevOpts | Version |
| - | - |
| OS | Ubuntu 25.04 |
| go | 1.26.2 |
| nodejs | v22.23.3 |
| yarn | 1.22.22 |

## Make

| Type | Command |
| - | - |
| Make all | `make` |
| Backend | `make backend` |
| Frontend | `make frontend` |
| Run | `make run` |
| Tidy | `make tidy` |
| Lint | `make lint` |
| Generate frontend openapi | `make openapi` |
| Docker image | `make docker` |
| Clean (build + db) | `make clean` |

`make` builds the backend binary and the frontend resources under the `build` directory.

## Execute

Setup the configuration file: [config.yaml](./config.yaml)

And run:

```bash
make run
```

## Install - Docker Compose

1. Check the config

    Modify system settings at `./docker/config.yaml`, e.g. the default login credential:

    ```yaml
    username: "admin"
    password: "0000"
    ```

    For other settings, make sure the change is reflected in `docker-compose.yaml` too.

2. Up the compose

    ```bash
    cd docker
    docker compose up -d
    ```

    The default db is stored at `/var/lib/system/db`, mounted in the compose file.

3. Down the compose

    ```bash
    docker compose down
    ```

## Develop Steps

`web/openapi.yaml` is the single source of truth for the API.

1. Add or change the paths and schemas in `web/openapi.yaml`.
2. Run `make openapi` to regenerate the typescript client in `web/frontend/src/api` (never edit it by hand).
3. Implement the routes in the backend (`web/backend/internal/api_xxx.go` + `processor`).
4. Call them from the frontend through `api` in `web/frontend/src/apiClient.ts`.
5. Go to make!

For frontend development with hot reload, run the backend (`make backend && make run`) and then `yarn dev` in `web/frontend`: Vite proxies `/api` to `http://localhost:8888`, or to `VITE_DEV_API_TARGET` if set.

See [CLAUDE.md](./CLAUDE.md) for the project conventions.
