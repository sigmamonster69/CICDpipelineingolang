# CI/CD Demo Application

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-educational-green.svg)](./LICENSE)
[![CI Pipeline](https://github.com/yourusername/cicd-demo/actions/workflows/ci.yml/badge.svg)](.github/workflows/ci.yml)

> 🎓 **A hands-on learning environment for mastering Continuous Integration and Continuous Deployment (CI/CD) concepts using GitHub Actions and Docker.**

---

## 📑 Table of Contents

- [Overview](#-overview)
- [Features](#-features)
- [Quick Start](#-quick-start)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [CI/CD Pipeline](#-ci-cd-pipeline)
- [Learning Resources](#-learning-resources)
- [Common Commands](#-common-commands)
- [Troubleshooting](#-troubleshooting)
- [Contributing](#-contributing)

---

## 📖 Overview

This project provides a **hands-on learning environment** for understanding CI/CD pipelines. It includes a basic Go application with automated testing, containerization, and GitHub Actions workflow configuration.

Perfect for developers who want to:
- ✅ Learn CI/CD fundamentals through practice
- ✅ Understand GitHub Actions workflows
- ✅ Master Docker containerization
- ✅ Build confidence in automated testing

---

## 🚀 Features

| Feature | Description | Benefit |
|---------|-------------|---------|
| **Automated Testing** | Go tests run automatically on every push | Catch bugs before they reach production |
| **Docker Support** | Containerized application for consistent deployments | "Works everywhere" guarantee |
| **GitHub Actions** | CI pipeline configured for automated builds and tests | Free CI/CD for public repositories |
| **Educational** | Well-commented code explaining CI/CD concepts | Learn by doing, not just reading |
| **Modular Structure** | Clean separation of app logic and UI | Best practices for real-world projects |

---

## 🏃 Quick Start

```bash
# 1. Clone the repository
git clone <repository-url>
cd cicd-demo

# 2. Download dependencies
go mod download

# 3. Run tests (verify everything works)
go test -v ./...

# 4. Build and run the application
go build -v ./... && ./cicd-demo
```

**Expected Output:**
```
Welcome to our CI/CD Demo App!
5 + 3 = 8
10 - 4 = 6
```

---

## 📁 Project Structure

```
cicd-demo/
├── .github/
│   └── workflows/
│       └── ci.yml              # GitHub Actions CI pipeline configuration
├── cmd/
│   ├── cicd-dashboard/         # Dashboard application entry point
│   └── cicd-demo/              # Demo application entry point
├── internal/
│   ├── app/                    # Core application logic
│   │   ├── app.go              # Main application functions
│   │   └── app_test.go         # Unit tests for app logic
│   └── ui/                     # User interface components
│       ├── dashboard.go        # Dashboard UI code
│       └── dashboard_test.go   # UI component tests
├── main.go                     # Application entry point
├── practice.go                 # Practice exercises and examples
├── go.mod                      # Go module dependencies
├── Dockerfile                  # Docker container configuration
├── README.md                   # This file - project overview
├── PROJECT_SUMMARY.md          # Complete project documentation
├── TASKS.md                    # Step-by-step learning exercises
└── README_CI_CD_GUIDE.md       # Beginner's guide to CI/CD
```

**Key Directories:**
- `.github/workflows/` - Automation pipeline definitions
- `cmd/` - Application binaries (one per command)
- `internal/` - Private application code (not importable by other projects)

---

## 🛠️ Prerequisites

Before you begin, ensure you have the following installed:

| Tool | Version | Purpose | Installation Link |
|------|---------|---------|-------------------|
| **Go** | 1.19+ | Build and run the application | [golang.org/dl](https://golang.org/dl/) |
| **Git** | Latest | Version control | [git-scm.com](https://git-scm.com/) |
| **Docker** | Latest (optional) | Containerization | [docker.com/get-started](https://www.docker.com/get-started/) |

### Verify Installation

```bash
# Check Go version
go version

# Check Git version
git --version

# Check Docker version (optional)
docker --version
```

---

## 💻 Getting Started

### Clone the Repository

```bash
git clone <repository-url>
cd cicd-demo
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
# Run all tests with verbose output
go test -v ./...

# Run tests with coverage report
go test -v -cover ./...
```

### Build Docker Image

```bash
# Build the Docker image
docker build -t cicd-demo .

# Run the container
docker run cicd-demo
```

---

## 🔄 CI/CD Pipeline

### Continuous Integration (CI)

The GitHub Actions workflow (`.github/workflows/ci.yml`) automatically triggers on every push or pull request to the `main` branch:

| Step | Action | Purpose |
|------|--------|---------|
| 1️⃣ | Checkout Code | Downloads your code to a fresh Ubuntu runner |
| 2️⃣ | Setup Go | Installs Go 1.19 on the runner |
| 3️⃣ | Download Dependencies | Fetches all required Go modules |
| 4️⃣ | Build | Compiles the application (catches syntax errors) |
| 5️⃣ | Test | Runs all unit tests (catches logic bugs) |

**If any step fails**, the pipeline stops and notifies you immediately!

### Continuous Deployment (CD)

After successful CI:
1. Docker image is built from the Dockerfile
2. Image can be pushed to a registry (Docker Hub, GitHub Container Registry)
3. Deploy to production servers (manual or automated)

### Workflow Diagram

```
┌─────────────────────────────────────────────────────────────┐
│ Developer pushes code → GitHub Actions triggers             │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│ Fresh Ubuntu environment → Install Go → Download deps       │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│ Build succeeds? → Run tests → All tests pass?               │
└─────────────────────────────────────────────────────────────┘
                            ↓
                    ┌──────────────┐
                    │ All checks   │
                    │ passed?      │
                    └──────────────┘
                      YES ↓    ↓ NO
                         │     └──→ ❌ Stop & Notify
                         │         Show error message
                         │         Block deployment
                         ↓
                    ✅ Build Docker image
                         ↓
                    ✅ Push to registry
                         ↓
                    ✅ Deploy to production
```

---

## 📚 Learning Resources

This project includes comprehensive learning materials:

| Resource | Description | Best For |
|----------|-------------|----------|
| [TASKS.md](./TASKS.md) | Step-by-step hands-on exercises | Learning by doing |
| [PROJECT_SUMMARY.md](./PROJECT_SUMMARY.md) | Complete project documentation | Understanding the big picture |
| [README_CI_CD_GUIDE.md](./README_CI_CD_GUIDE.md) | Beginner-friendly CI/CD guide | First-time CI/CD learners |

### Recommended Learning Path

1. **Start Here**: Read this README for an overview
2. **Hands-On**: Follow the exercises in `TASKS.md`
3. **Deep Dive**: Review `PROJECT_SUMMARY.md` for detailed explanations
4. **Reference**: Use `README_CI_CD_GUIDE.md` when you need clarification

---

## 🔧 Common Commands

### Development

```bash
# Run all tests with coverage
go test -v -cover ./...

# Run specific test function
go test -v -run TestAdd ./...

# Build with verbose output
go build -v ./...

# Run the application
go run main.go
```

### Docker

```bash
# Build Docker image with tag
docker build -t cicd-demo:latest .

# Build with no cache (fresh build)
docker build --no-cache -t cicd-demo .

# Run container interactively
docker run -it cicd-demo

# Run container and remove after exit
docker run --rm cicd-demo

# View running containers
docker ps

# View container logs
docker logs <container-id>
```

### CI/CD Simulation

```bash
# Simulate the CI pipeline locally
go mod download && go build -v ./... && go test -v ./...

# Check if code compiles without running
go build -o /dev/null ./...
```

---

## 🐛 Troubleshooting

### Common Issues and Solutions

| Issue | Possible Cause | Solution |
|-------|---------------|----------|
| **Tests failing** | Logic error or incorrect test expectations | Check `*_test.go` files for expected behavior |
| **Build errors** | Syntax error or missing dependencies | Run `go mod tidy` and check error messages |
| **Go version mismatch** | Installed Go version < 1.19 | Update Go from [golang.org/dl](https://golang.org/dl/) |
| **Docker build fails** | Docker daemon not running or Dockerfile error | Verify Docker is running: `docker info` |
| **Pipeline not triggering** | Wrong branch or workflow disabled | Check `.github/workflows/ci.yml` and Actions tab |

### Debug Tips

```bash
# Verbose test output
go test -v ./...

# Show which tests are running
go test -v -run . ./...

# Clean build (remove cached artifacts)
go clean -cache && go build -v ./...

# Check Go environment
go env
```

### Getting Help

1. Read error messages carefully - they often tell you exactly what's wrong
2. Check the [TASKS.md](./TASKS.md) troubleshooting section
3. Review comments in source files for explanations
4. Run commands one at a time to isolate issues

---

## 🤝 Contributing

We welcome contributions! Here's how to get involved:

### Contribution Process

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Make** your changes
4. **Run** tests: `go test -v ./...`
5. **Commit** your changes (`git commit -m 'Add amazing feature'`)
6. **Push** to the branch (`git push origin feature/amazing-feature`)
7. **Open** a Pull Request

### Code Style Guidelines

- Write clear, commented code
- Add tests for new functionality
- Follow Go best practices ([Effective Go](https://golang.org/doc/effective_go))
- Keep commits small and focused

### Questions?

- Open an issue for bugs or feature requests
- Check existing documentation before asking
- Be respectful and helpful to other contributors

---

## 📄 License

This project is for **educational purposes**. Feel free to use, modify, and share for learning!

---

## 🎯 What You'll Learn

By completing this project, you'll gain practical experience with:

- ✅ **CI/CD Fundamentals** - Understand the why and how of continuous integration
- ✅ **GitHub Actions** - Create and customize automation workflows
- ✅ **Go Testing** - Write effective unit tests
- ✅ **Docker Basics** - Containerize applications for consistent deployment
- ✅ **Debugging** - Interpret CI/CD pipeline failures
- ✅ **Best Practices** - Industry-standard development workflows

---

## 🚀 Next Steps

Ready to level up? Try these extensions:

1. **Add a Linter** - Integrate `golangci-lint` for code quality checks
2. **Code Coverage** - Set up coverage thresholds (e.g., minimum 80%)
3. **Multi-Platform Builds** - Build for Linux, macOS, and Windows
4. **Deploy to Cloud** - Connect to AWS, GCP, or Azure
5. **Notifications** - Add Slack/Discord alerts for pipeline status
6. **Environment Variables** - Handle secrets securely with GitHub Secrets
7. **Rollback Strategy** - Implement automatic rollback on deployment failure

---

## 📞 Support

**Learning Path:** Start with [TASKS.md](./TASKS.md) to begin your CI/CD journey!

**Remember:** Every developer started where you are now. Breaking things is how you learn! 🔧

---

**Happy Learning!** 🚀
