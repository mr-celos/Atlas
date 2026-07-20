# Atlas

Atlas is a backend platform for multiplayer games, built as a long-term
learning project focused on backend engineering fundamentals in Go.

The idea is to build a simplified Backend-as-a-Service (BaaS) for games —
something in the spirit of Firebase, PlayFab, or Supabase, but scoped
specifically to game backend needs: accounts, inventories, matchmaking,
leaderboards, economy, and the infrastructure around them.

The game itself is not the point. Atlas is the product — games are just
clients that talk to it over HTTP (and, eventually, WebSockets).

## Why this project exists

This is primarily a vehicle for learning backend engineering properly:
Go idioms, HTTP, database design, caching, concurrency, testing, and
eventually deployment and distributed systems — built incrementally,
the way a real product would be, rather than as a tutorial.

## Tech stack

- **Language:** Go
- **Database:** PostgreSQL
- **Planned:** Redis, Docker, REST APIs, WebSockets

## Running locally

```bash
cd cmd/atlas
go run .
```

Configuration is read from environment variables (with sensible defaults
if unset):

| Variable | Default | Description |
|---|---|---|
| `PORT` | `:8080` | Address the server listens on |
| `PRODUCT_VERSION` | `Development` | Version string reported by `/health` |
| `DATABASE_HOST` | `localhost` | PostgreSQL database host address |
| `DATABASE_PORT` | `5432` | PostgreSQL database port number |
| `DATABASE_USER` | `postgres` | PostgreSQL database username |
| `DATABASE_PASSWORD` | *(empty)* | PostgreSQL database password |
| `DATABASE_NAME` | `atlas` | PostgreSQL database name |

## Running tests

```bash
go test ./...
```

## License

See [LICENSE](./LICENSE).