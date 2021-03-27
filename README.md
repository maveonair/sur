# SUR (Solus User Repository)

Self-hosted Solus repository.

## Configuration

### Environment variables

| Variable         | Description             | Example               |
| ---------------- | ----------------------- | --------------------- |
| SUR_BASE_URL     | Base URL of running app | http://localhost:8080 |
| SUR_PACKAGES_DIR | Packages directory      | /packages             |

## Developer Setup

### Run tests

```
$ make test
```

### Run application

1. Run the make dev command

```
$ make dev
```

2. Open [http://localhost:8080](http://localhost)
