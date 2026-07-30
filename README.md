# CI/CD Demo Application

A simple Go application demonstrating Continuous Integration and Continuous Deployment (CI/CD) concepts using GitHub Actions and Docker.

## 📖 Overview

This project provides a hands-on learning environment for understanding CI/CD pipelines. It includes a basic Go application with automated testing, containerization, and GitHub Actions workflow configuration.

## 🚀 Features

- **Automated Testing**: Go tests run automatically on every push
- **Docker Support**: Containerized application for consistent deployments
- **GitHub Actions**: CI pipeline configured for automated builds and tests
- **Educational**: Well-commented code explaining CI/CD concepts

## 📁 Project Structure

```
├── main.go                 # Main application code
├── main_test.go            # Automated tests
├── .github/workflows/ci.yml # GitHub Actions CI pipeline
├── Dockerfile              # Docker container configuration
├── go.mod                  # Go module dependencies
├── TASKS.md                # Step-by-step learning exercises
└── PROJECT_SUMMARY.md      # Complete project documentation
```

## 🛠️ Prerequisites

- Go 1.19 or later
- Docker (optional, for containerization)
- Git

## 💻 Getting Started

### Clone the Repository

```bash
git clone <repository-url>
cd <project-directory>
```

### Run Locally

```bash
# Download dependencies
go mod download

# Build the application
go build -v ./...

# Run the application
./cicd-demo
```

### Run Tests

```bash
go test -v ./...
```

### Build Docker Image

```bash
docker build -t cicd-demo .
docker run cicd-demo
```

## 🔄 CI/CD Pipeline

### Continuous Integration (CI)

The GitHub Actions workflow (`.github/workflows/ci.yml`) automatically:

1. Triggers on push to `main` branch
2. Sets up a fresh Ubuntu environment
3. Installs Go 1.19
4. Downloads dependencies
5. Builds the application
6. Runs all tests

### Continuous Deployment (CD)

After successful CI:
1. Docker image is built from the Dockerfile
2. Image can be pushed to a registry (Docker Hub, etc.)
3. Deploy to production servers

## 📊 Workflow Diagram

```
Developer pushes code → GitHub Actions triggers → Tests run → Build succeeds? 
    ↓ No  ──────────────────────────────────────→ Stop & Notify
    ↓ Yes
Build Docker image → Push to registry → Deploy to production
```

## 🎯 Learning Objectives

By working with this project, you'll learn:

- How CI/CD pipelines work
- Writing testable Go code
- Creating GitHub Actions workflows
- Building Docker containers
- Automating software delivery

## 📝 Key Files Explained

| File | Purpose |
|------|---------|
| `main.go` | Application code with testable functions |
| `main_test.go` | Unit tests for automatic validation |
| `.github/workflows/ci.yml` | CI pipeline configuration |
| `Dockerfile` | Container build instructions |
| `TASKS.md` | Hands-on learning exercises |

## 🔧 Common Commands

```bash
# Run all tests with coverage
go test -v -cover ./...

# Build Docker image with tag
docker build -t cicd-demo:latest .

# Run container interactively
docker run -it cicd-demo

# Simulate CI pipeline locally
go mod download && go build -v ./... && go test -v ./...
```

## 🐛 Troubleshooting

- **Tests failing?** Check `main_test.go` for expected behavior
- **Build errors?** Ensure Go 1.19+ is installed
- **Docker issues?** Verify Docker daemon is running

## 📚 Additional Resources

- [PROJECT_SUMMARY.md](./PROJECT_SUMMARY.md) - Detailed project documentation
- [TASKS.md](./TASKS.md) - Step-by-step learning exercises
- [GitHub Actions Docs](https://docs.github.com/en/actions)
- [Docker Documentation](https://docs.docker.com/)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `go test -v ./...`
5. Submit a pull request

## 📄 License

This project is for educational purposes.

---

**Happy Learning!** 🚀 Start with `TASKS.md` to begin your CI/CD journey.
