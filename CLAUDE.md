# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go SDK for TAPD (腾讯敏捷产品研发平台) Open API. Standalone Go module, **standard library only — no external dependencies**. Go 1.24+. Apache 2.0.

Package layout: top-level `tapd` package (HTTP client + per-resource methods) and `model/` package (request/response structs).

## Common commands

```bash
# Run all tests with race detector + coverage (matches CI)
go test -race -coverprofile=coverage.out ./...

# Run all tests in one package
go test ./model/...

# Run single test by name (regex)
go test -run TestNewClient_BearerAuth ./...
go test -run '^TestListStories$' .

# Static analysis (CI gate)
go vet ./...

# Formatting check (CI fails on any unformatted file)
gofmt -l .
gofmt -w .   # apply
```

CI (`.github/workflows/ci.yml`) runs gofmt check, `go vet ./...`, then `go test -race -coverprofile=coverage.out ./...` on Go 1.24.12. Keep these green locally before pushing.

## Architecture

### Request flow

Every resource method follows the same pipeline — understand it once, applies everywhere:

1. Caller builds a `*model.XxxRequest` struct and calls `client.XxxAction(ctx, req)`.
2. Method calls `req.ToParams()` to get `map[string]string`.
3. `Client.doGet` / `Client.doPost` (in `client.go`) builds the HTTP request, attaches `Authorization` header (`Bearer <token>` or `Basic <base64>`), sends it.
4. `doRequest` reads body, maps non-2xx HTTP to `*TAPDError` with an `ExitCode` (see `mapHTTPError`), then unmarshals into `model.TAPDResponse{Status, Data, Info}`. If `Status != 1`, returns `*TAPDError` with `ExitCode: 4`.
5. Resource method calls `parseList[T]` / `parseOne[T]` / `parseCount` (generic helpers in `parse.go`) to unwrap the TAPD-specific envelope `[{"Story": {...}}, ...]` or `{"Story": {...}}`.
6. For single-entity responses, methods often set `entity.URL` using `c.webURL` so the model carries a clickable link.

This means: when adding a new resource, you almost never write JSON parsing — `parseList[model.Foo](data, "Foo")` is the pattern.

### Adding a new API method

1. Define request struct in `model/<resource>.go` with a `ToParams() map[string]string` method. Use `setOptional` / `setOptionalInt` (defined in `model/wiki.go`) to skip empty values.
2. If the response wraps a new entity, add the model struct in the same file.
3. Add the method on `*Client` in `<resource>.go` (top-level package). Choose `doGet` for read, `doPost` (form-urlencoded) for write. JSON POSTs (e.g. WeCom webhook) use `doPostJSON`.
4. Use `parseList[T]` / `parseOne[T]` / `parseCount` to decode `data`.
5. Add a test in `<resource>_test.go` using `newMockServer` (see `client_test.go`) — handler returns `{"status":1,"data":...,"info":"success"}`.

### File splitting (extras pattern)

Large resources are split across multiple files to keep each under control:

- `<resource>.go` — base CRUD (Create/Get/List/Update/Count).
- `<resource>_extras.go` — secondary endpoints (categories, links, relations, batch ops, copy, templates, view-conf queries, etc.).
- `tcase_test_plan.go` / `tcase_instance.go` — domain-specific sub-modules of tcase.

Same convention in `model/` (e.g. `model/story_extras.go`, `model/bug_extras.go`, `model/wiki_extras.go`). When adding many endpoints to one resource, prefer a new `_extras.go` file over bloating the base file. **Never** dump new types into `model/request.go` or `model/model.go` — those are shared and crowded.

### Custom fields (important quirk)

`Story`, `Task`, `Bug` (and `model/story_extras.go`, `model/bug_extras.go`, etc.) implement custom `UnmarshalJSON` / `MarshalJSON` that route any key starting with `custom_field_` or `custom_plan_field_` into the model's `CustomFields map[string]string` instead of dropping it. Helpers live in `model/custom_fields.go`:

- `IsCustomField(key)` — prefix check.
- `ExtractCustomFields(raw)` — pull custom keys from a raw JSON map during unmarshal.
- `MergeCustomFields(params, custom)` — inject back into request params.

When adding new entities that may carry custom fields, follow the same UnmarshalJSON/MarshalJSON pattern (see `model/story.go` for the canonical example).

### Auth

`NewClient(accessToken, apiUser, apiPassword)` — accessToken non-empty wins (Bearer); otherwise Basic auth from user/pass. `NewClientWithBaseURL` allows pointing at a private TAPD deployment by overriding `apiURL` and `webURL` (passing `""` falls back to `https://api.tapd.cn` / `https://www.tapd.cn`).

`Client` is goroutine-safe. The `nick` field (current user, fetched via `FetchNick`) is guarded by `sync.RWMutex` — use `GetNick`/`SetNick`, don't access `c.nick` directly.

### Errors

Errors returned to callers are either `*TAPDError` (HTTP non-2xx, or `status != 1` from TAPD body) or wrapped stdlib errors. `TAPDError.ExitCode` mapping: 401→1, 404→2, 422→3, others/TAPD logical errors→4. Callers can type-assert with `errors.As` or `err.(*tapd.TAPDError)`.

## Conventions

- All public APIs take `context.Context` as first arg.
- Doc comments on exported identifiers are in Chinese — match existing style. Each resource method links to its TAPD API doc URL in the comment.
- Request structs use `string` fields for IDs (TAPD returns IDs as strings) and explicit zero-omission via `setOptional`.
- Tests are colocated (`foo.go` ↔ `foo_test.go`) and use `httptest.NewServer` mock servers — never hit the real TAPD API in tests.
- Zero-dep rule is load-bearing: do not add `require` entries to `go.mod`. Anything beyond stdlib breaks the project's design contract.
