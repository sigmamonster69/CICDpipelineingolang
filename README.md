# CICDpipelineingolang

A small Go-based CI/CD learning scaffold.

## What is in this repo

- `cmd/cicd-demo/main.go` - a tiny Go app to build in CI
- `go.mod` - module definition for the Go project
- `.github/workflows/ci.yml` - starter GitHub Actions CI workflow
- `docs/ci-cd-basics.md` - explanation of the pipeline pieces

## What this teaches

- How CI runs on every push and pull request
- How a Go build fits into automation
- Where to add tests, linting, Docker, and deployment later

## Suggested next step

Read `docs/ci-cd-basics.md`, then add a test file and a lint step so you can see the pipeline evolve one piece at a time.
