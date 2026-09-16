# CLAUDE.md

Guidance for Claude Code / AI assistants working in this repository. Keep it concise and prefer pointing to the canonical docs over duplicating them.

## What this is

The Solace PubSub+ OpenTelemetry Integration API for Go — handles injection and extraction of tracing context on carrier objects (Solace messages). Go module path: `solace.dev/go/messaging-trace/opentelemetry`.

It has a hard dependency on the [Solace PubSub+ Messaging API for Go](https://docs.solace.com/API/Messaging-APIs/Go-API/go-home.htm) (a Cgo wrapper over the Solace C API), which constrains the supported operating systems — see the [README](README.md) "OS Support".

- **Go version:** 1.20+ (per `go.mod` and the README).
- **User-facing usage, install, OS support:** [README.md](README.md).
- **Contributor workflow (build / test / docs / prerequisites):** [CONTRIBUTING.md](CONTRIBUTING.md) → "Developer Guide".

## Repository layout

Per [CONTRIBUTING.md](CONTRIBUTING.md) "Repository Contents":

- `./` — the API's main entrypoints (`tracing.go`, `tracing_property_name.go`, `version.go`, `doc.go`).
- `carrier/` — the API's carrier interfaces.
- `internal/` — implementation of the integration API.
- `logging/` — logging interfaces, constants, and basic data structures.
- `test/` — integration tests. **Separate Go module** (see note under Building & testing).

## Building & testing

Follow [CONTRIBUTING.md](CONTRIBUTING.md) "Prerequisites" / "Testing" and [test/README.md](test/README.md) rather than restating commands here. Key points:

- Build / vet the library: `go build ./...` from the repo root.
- Formatting is enforced in CI — run `go fmt ./...` before pushing.
- Unit tests: `go test ./...` from the repo root.
- Integration tests live in the `test/` module (Ginkgo + Gomega) and need Docker / docker-compose. They require a one-time `go generate` in `test/sempclient` to produce the (uncommitted) SEMPv2 client. See [test/README.md](test/README.md) for the full flow, focused runs, coverage, and running against an external broker.
- `test/` is its own Go module — open it in its own editor workspace (`gopls` supports a single module per workspace).

## CI

- Public CI runs on **GitHub Actions**: [`.github/workflows/test.yml`](.github/workflows/test.yml) — `go build`, a `go fmt` check, unit tests, then the `test/` integration suite on the latest Go on Linux. Dependabot ([`.github/dependabot.yml`](.github/dependabot.yml)) keeps the Go modules current.
- A `Jenkinsfile` is also present for Solace's internal maintainer build/release pipeline; external contributors can ignore it — the GitHub Actions workflow above is the source of truth for public CI.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the issue/PR process and developer guide, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md). Any new feature must add corresponding tests in `test/` and must not decrease overall test coverage.
