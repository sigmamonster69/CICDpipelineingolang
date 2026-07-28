# CICDpipelineingolang

A small Go-based CI/CD learning scaffold.

## What is in this repo

- `cmd/cicd-demo/main.go` - a tiny Go app to build in CI
- `cmd/cicd-dashboard/main.go` - a tiny browser dashboard for the same project
- `practice.go` - extra small Go practice helpers
- `practice_test.go` - tests for the practice helpers
- `internal/app/app.go` - the basic app logic and helper functions
- `internal/app/app_test.go` - unit tests for the helper functions
- `internal/ui/dashboard.go` - HTML rendering for the dashboard
- `internal/ui/dashboard_test.go` - unit test for the dashboard renderer
- `go.mod` - module definition for the Go project
- `.github/workflows/ci.yml` - starter GitHub Actions CI workflow

## What this teaches

- How CI runs on every push and pull request
- How a Go build fits into automation
- How tests make the pipeline more useful than a compile-only check
- How simple Go functions and tests fit into a pipeline
- How to organize Go code into small packages
- How a tiny browser UI can sit on top of the same Go logic
- Where to add tests, linting, Docker, and deployment later

## Suggested next step

Use the dashboard task board as your reading and checklist surface, then add a lint step and a Docker build so you can see the pipeline evolve one piece at a time.

## How to make more contributions

- Pick one small file.
- Change one tiny thing.
- Commit it.
- Push it to GitHub.
- Check the update on GitHub.
- Repeat with another small change.

## Dashboard

Run `cmd/cicd-dashboard` to open a tiny task board in the browser.

- Click a task to mark it done.
- The dashboard saves completion state in `.task-state.json`.
- Done tasks show the date they were completed.
