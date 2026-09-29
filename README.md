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

1. Add APIs in backend.
2. Updaet APIs in postman and export the json file
3. Update the openapi.yaml with your postman json
4. Use `make openapi` to generate the api typescript file in frontend
5. Go to make!
