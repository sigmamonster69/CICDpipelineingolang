# CICDpipelineingolang

A small Go-based CI/CD learning scaffold.

## What is in this repo

- `cmd/cicd-demo/main.go` - a tiny Go app to build in CI
- `internal/demo/demo.go` - the testable logic behind the demo app
- `internal/demo/demo_test.go` - a basic unit test for the demo message
- `go.mod` - module definition for the Go project
- `.github/workflows/ci.yml` - starter GitHub Actions CI workflow
- `docs/ci-cd-basics.md` - explanation of the pipeline pieces

## What this teaches

- How CI runs on every push and pull request
- How a Go build fits into automation
- How tests make the pipeline more useful than a compile-only check
- Where to add tests, linting, Docker, and deployment later

## Suggested next step

Read `docs/ci-cd-basics.md`, then add a lint step and a Docker build so you can see the pipeline evolve one piece at a time.
