# Tic-Tac-Toe API

All application code is inside `src/`. Each layer is a separate Go module,
connected locally by `go.work`. Run all commands from `src/`.

## Setup

Use Go 1.23 or newer.

```sh
cd src
go mod download -modfile=di/go.mod
```

To initialize dependency checksums for each module, use:

```sh
cd di && go mod tidy && cd ..
cd cmd/server && go mod tidy && cd ../..
```

Start the server with:

```sh
go run ./cmd/server
```

The service listens on http://localhost:8080.

## API

- `POST /game` creates an empty game and returns its UUID.
- `GET /game/{UUID}` retrieves an existing game.
- `POST /game/{UUID}` submits the current board with exactly one new human move.
  The server returns the updated board after the computer's Minimax move.
  A new UUID may also be supplied directly for a game's first move.

Board values: `0` empty, `1` human, `2` computer.

Example request (replace UUID with the ID from POST /game):

```json
{
  "board": [
    [1, 0, 0],
    [0, 0, 0],
    [0, 0, 0]
  ]
}
```

Example command:

```sh
curl -X POST http://localhost:8080/game
curl -X POST http://localhost:8080/game/YOUR_UUID \
  -H 'Content-Type: application/json' \
  -d '{"board":[[1,0,0],[0,0,0],[0,0,0]]}'
```

The response includes `id`, `board`, `status` (`playing`, `finished`,
`draw`) and `winner` (0 for no winner, 1 for human, 2 for computer).

Incorrect moves return HTTP 400 with `{"error":"..."}`.

## Checks

```sh
cd src
go fmt ./domain/... ./datasource/... ./web/... ./di/... ./cmd/server/...
go test ./domain/... ./datasource/... ./web/... ./di/... ./cmd/server/...
go test -race ./domain/... ./datasource/... ./web/... ./di/... ./cmd/server/...
go vet ./domain/... ./datasource/... ./web/... ./di/... ./cmd/server/...
```

The game state is kept in memory and will be lost when the process stops.
A mutex serializes the complete game update, while sync.Map makes individual
storage operations safe under concurrent access.
