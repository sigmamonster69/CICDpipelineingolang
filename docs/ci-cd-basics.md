# CI/CD basics

This repository is a small learning scaffold for understanding how a CI/CD pipeline fits together.

## What the pieces do

- `go.mod` defines the Go module and toolchain version.
- `cmd/cicd-demo/main.go` is a tiny application the pipeline can build.
- `.github/workflows/ci.yml` runs on push and pull requests.
- `docs/ci-cd-basics.md` explains the workflow at a high level.

## Pipeline flow

1. A developer pushes code or opens a pull request.
2. GitHub Actions checks out the repository.
3. Go is installed.
4. Dependencies are prepared.
5. The app is built.
6. Tests run.

## Why this is a good starter layout

- It keeps the app small so you can focus on the pipeline.
- It uses GitHub Actions, which is the easiest place to learn CI on GitHub.
- It gives you a place to add linting, security scans, and deployments later.

## Next things to learn

- Add unit tests around the Go app.
- Add `golangci-lint` or `go vet`.
- Add Docker build and image publishing.
- Add separate deploy jobs for staging and production.
