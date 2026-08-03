APP_NAME := cicd-dashboard

.PHONY: build test vet run dashboard lint security docker-build

build:
	go build ./...

test:
	go test ./...

vet:
	go vet ./...

run:
	go run ./cmd/cicd-demo

dashboard:
	go run ./cmd/cicd-dashboard

lint: vet

security:
	@echo "Run Trivy or gosec in CI for a real security scan."

docker-build:
	docker build -t $(APP_NAME):latest .
