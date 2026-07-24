# DevSecOps learning path

## Goal

Build a small Go app that can be developed, tested, packaged, scanned, and deployed like a real CI/CD project.

## Phase 1: Development

- Keep the Go code simple
- Add one small function at a time
- Write a test for each new helper
- Use `go fmt` and `go test` often

## Phase 2: Continuous Integration

- Run tests on every push
- Run tests on every pull request
- Add linting
- Fail fast when code breaks

## Phase 3: Security

- Scan the code for common mistakes
- Scan dependencies for vulnerabilities
- Scan Docker images before release
- Look for secrets before pushing

## Phase 4: Delivery

- Build a Docker image
- Push the image to a registry
- Deploy to staging
- Deploy to production after review

## What to practice in this repo

- adding small Go features
- moving code into packages
- writing tests
- wiring those tests into CI
- adding Docker and security tooling
