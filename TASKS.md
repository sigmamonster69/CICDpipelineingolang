# 🎓 CI/CD Learning Tasks - Step by Step

## Your Mission: Build and Understand a CI/CD Pipeline

Complete these tasks in order. Each one teaches you a key concept!

---

## ✅ Phase 1: Local Testing (Do This First)

### Task 1.1: Run Tests Locally
**Goal:** Simulate what GitHub Actions does on your computer

```bash
cd /workspace
go mod download
go build -v ./...
go test -v ./...
```

**What to observe:**
- All tests should pass ✅
- You should see "Welcome to our CI/CD Demo App!" output
- This is exactly what the CI pipeline runs automatically!

**Expected Output:**
```
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract
--- PASS: TestSubtract (0.00s)
=== RUN   TestAddNegativeNumbers
--- PASS: TestAddNegativeNumbers (0.00s)
PASS
ok      cicd-demo    0.002s
```

---

### Task 1.2: Break the Code Intentionally
**Goal:** See how CI catches bugs

**Steps:**
1. Open `main.go`
2. Change line 9 from `return a + b` to `return a - b` (introduce a bug!)
3. Run `go test -v ./...` again
4. Watch the test fail ❌

**What to observe:**
- TestAdd should fail because 2 - 3 ≠ 5
- Error message shows exactly what went wrong
- **This is why CI is valuable!** It catches bugs before deployment

**Fix it back:** Change it back to `return a + b` after you understand

---

### Task 1.3: Add Your Own Test
**Goal:** Learn how to extend the test suite

**Steps:**
1. Open `main_test.go`
2. Add a new test function for Subtract with negative numbers:

```go
func TestSubtractNegativeNumbers(t *testing.T) {
	result := Subtract(-5, -3)
	expected := -2
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
```

3. Run `go test -v ./...` again
4. Verify your new test passes

**Why this matters:** CI pipelines grow as your codebase grows!

---

## 🐳 Phase 2: Docker (Optional but Recommended)

### Task 2.1: Create a Dockerfile
**Goal:** Package your app in a container

**Create a file called `Dockerfile`** (no extension) with:

```dockerfile
# Use official Go image as base
FROM golang:1.19-alpine

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (for better caching)
COPY go.mod ./
RUN go mod download

# Copy all source code
COPY . .

# Build the application
RUN go build -o myapp .

# Command to run when container starts
CMD ["./myapp"]
```

**What this does:** Creates a portable package with your app + Go runtime

---

### Task 2.2: Build and Run Docker Container
**Goal:** Test your containerized app

```bash
# Build the Docker image
docker build -t cicd-demo .

# Run the container
docker run cicd-demo
```

**Expected Output:**
```
Welcome to our CI/CD Demo App!
5 + 3 = 8
10 - 4 = 6
```

**Why this matters:** Same container runs everywhere - your laptop, CI server, production!

---

## 🚀 Phase 3: Understanding the Full Pipeline

### Task 3.1: Trace the Complete Flow
**Draw this diagram on paper or in a note:**

```
[You write code] 
       ↓
[git push to GitHub]
       ↓
[GitHub Actions triggers]
       ↓
[Ubuntu server spins up]
       ↓
[Code downloaded]
       ↓
[Go installed]
       ↓
[Dependencies downloaded]
       ↓
[App built] → If fails, STOP and notify
       ↓
[Tests run] → If fails, STOP and notify
       ↓
[Docker image built] (Phase 2)
       ↓
[Deployed to server] (Phase 3)
       ↓
[Team notified of success]
```

**Understanding check:** Can you explain each step out loud?

---

### Task 3.2: Modify the CI Workflow
**Goal:** Customize your pipeline

**Edit `.github/workflows/ci.yml`:**
1. Add a step that prints "Starting tests..." before running tests
2. Add a step that prints "All tests passed!" after tests succeed

Add these lines to the workflow:

```yaml
- name: Notify Start
  run: echo "🚀 Starting tests..."

# ... existing test step ...

- name: Notify Success
  run: echo "✅ All tests passed!"
```

---

## 📝 Knowledge Check Questions

Answer these before moving to Phase 4:

1. **What triggers the CI pipeline?**
2. **What happens if a test fails?**
3. **Why run tests locally before pushing?**
4. **What problem does Docker solve?**
5. **Where do you see CI results in GitHub?**

*(Answers at bottom of this file)*

---

## 🎯 Phase 4: Next Steps (Choose Your Adventure)

After completing above, choose one:

**A)** Push to GitHub and watch it work live
**B)** Add automatic deployment to a cloud service
**C)** Add code quality checks (linting)
**D)** Set up notifications (Slack/Email on failure)

---

## 🔑 Answers to Knowledge Check

1. **What triggers the CI pipeline?** 
   → Pushing code to main/master branch or creating a pull request

2. **What happens if a test fails?**
   → Pipeline stops, shows error, blocks deployment, notifies team

3. **Why run tests locally before pushing?**
   → Catch bugs early, save CI minutes, faster feedback loop

4. **What problem does Docker solve?**
   → "Works on my machine" problem - ensures same environment everywhere

5. **Where do you see CI results in GitHub?**
   → In the "Actions" tab, or as checkmarks on commits/PRs

---

## 💡 Pro Tips

- **Small commits** = easier to debug when CI fails
- **Read error messages** - they tell you exactly what's wrong
- **Green CI = confidence to deploy** anytime
- **Red CI = stop everything** and fix immediately

---

**Ready?** Start with Task 1.1 and report back your results! 🚀
