package ui

import (
\t"bytes"
\t"encoding/json"
\t"fmt"
\t"html/template"
\t"os"
\t"path/filepath"
\t"time"
)

const stateFileName = ".task-state.json"

// Task holds a single checklist item shown on the dashboard.
type Task struct {
\tName      string `json:"name"`
\tDone      bool   `json:"done"`
\tDoneAt    string `json:"done_at,omitempty"`
\tSuggested bool   `json:"suggested,omitempty"`
}

// PageData holds the basic values shown on the dashboard.
type PageData struct {
\tTitle       string
\tDescription string
\tMessage     string
\tReport      string
\tTasks       []Task
}

var dashboardTmpl = template.Must(template.New("dashboard").Parse(`
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <style>
    :root {
      --bg: #f6f8fb;
      --card: #ffffff;
      --text: #1f2937;
      --muted: #6b7280;
      --line: #e5e7eb;
      --accent: #0f766e;
      --done: #16a34a;
    }
    body {
      font-family: Arial, sans-serif;
      margin: 0;
      padding: 40px;
      background: var(--bg);
      color: var(--text);
    }
    .card {
      max-width: 860px;
      margin: 0 auto;
      background: var(--card);
      border-radius: 16px;
      padding: 32px;
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
    }
    h1 { margin-top: 0; }
    .muted { color: var(--muted); }
    .report {
      background: #111827;
      color: #f9fafb;
      padding: 16px;
      border-radius: 12px;
      overflow-x: auto;
      margin: 16px 0 24px;
    }
    .task-list {
      list-style: none;
      padding: 0;
      margin: 0;
    }
    .task {
      border: 1px solid var(--line);
      border-radius: 12px;
      padding: 14px 16px;
      margin-bottom: 12px;
      display: flex;
      justify-content: space-between;
      gap: 16px;
      align-items: start;
    }
    .task.done {
      border-color: rgba(22, 163, 74, 0.35);
      background: rgba(22, 163, 74, 0.05);
    }
    .task strong {
      display: block;
      margin-bottom: 4px;
    }
    .task-meta {
      font-size: 14px;
      color: var(--muted);
      margin-top: 6px;
    }
    .status {
      font-size: 13px;
      border-radius: 999px;
      padding: 6px 10px;
      white-space: nowrap;
      border: 1px solid var(--line);
      color: var(--muted);
      text-decoration: none;
    }
    .status.done {
      color: var(--done);
      border-color: rgba(22, 163, 74, 0.35);
    }
    .status:hover { border-color: var(--accent); color: var(--accent); }
    .section-title {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 16px;
    }
  </style>
</head>
<body>
  <main class="card">
    <h1>{{.Title}}</h1>
    <p class="muted">{{.Description}}</p>
    <p><strong>Status:</strong> {{.Message}}</p>
    <div class="report">{{.Report}}</div>
    <div class="section-title">
      <h2>Learning tasks</h2>
      <p class="muted">Click a task to mark it done or reopen it.</p>
    </div>
    <ul class="task-list">
      {{range .Tasks}}
      <li class="task {{if .Done}}done{{end}}">
        <div>
          <strong>{{if .Done}}[x]{{else}}[ ]{{end}} {{.Name}}</strong>
          {{if .Done}}<div class="task-meta">Done on {{.DoneAt}}</div>{{else}}<div class="task-meta">Not done yet</div>{{end}}
        </div>
        {{if .Done}}
          <a class="status done" href="/?task={{urlquery .Name}}&undo=1">Mark not done</a>
        {{else}}
          <a class="status" href="/?task={{urlquery .Name}}&done=1">Mark done</a>
        {{end}}
      </li>
      {{end}}
    </ul>
  </main>
</body>
</html>
`))

// DefaultTasks returns the checklist to show in the dashboard.
func DefaultTasks() []Task {
\treturn []Task{
\t\t{Name: "Read ROADMAP.txt"},
\t\t{Name: "Read TOOLS.txt"},
\t\t{Name: "Read NEXT_STEPS.txt"},
\t\t{Name: "Open .github/workflows/ci.yml"},
\t\t{Name: "Open internal/app/app.go"},
\t\t{Name: "Open internal/app/app_test.go"},
\t\t{Name: "Try to explain each file in your own words"},
\t\t{Name: "Add one tiny Go function yourself"},
\t\t{Name: "Add a test for that function"},
\t\t{Name: "Push the change and watch the CI pipeline"},
\t\t{Name: "Add Dockerfile"},
\t\t{Name: "Add docker-compose.yml"},
\t\t{Name: "Add Makefile"},
\t\t{Name: "Add a lint step to CI"},
\t\t{Name: "Add a security scan step after that"},
\t}
}

// LoadTasks loads task state from disk if available.
func LoadTasks(dir string) ([]Task, error) {
\ttasks := DefaultTasks()
\tpath := filepath.Join(dir, stateFileName)
\tdata, err := os.ReadFile(path)
\tif err != nil {
\t\tif os.IsNotExist(err) {
\t\t\treturn tasks, nil
\t\t}
\t\treturn nil, err
\t}
\tif err := json.Unmarshal(data, &tasks); err != nil {
\t\treturn nil, err
\t}
\treturn tasks, nil
}

// SaveTasks writes task state to disk.
func SaveTasks(dir string, tasks []Task) error {
\tpath := filepath.Join(dir, stateFileName)
\tdata, err := json.MarshalIndent(tasks, "", "  ")
\tif err != nil {
\t\treturn err
\t}
\treturn os.WriteFile(path, data, 0o644)
}

// UpdateTaskStatus updates a task and stamps the current date when done.
func UpdateTaskStatus(tasks []Task, taskName string, done bool) []Task {
\tfor i := range tasks {
\t\tif tasks[i].Name == taskName {
\t\t\ttasks[i].Done = done
\t\t\tif done {
\t\t\t\ttasks[i].DoneAt = time.Now().Format("2006-01-02")
\t\t\t} else {
\t\t\t\ttasks[i].DoneAt = ""
\t\t\t}
\t\t}
\t}
\treturn tasks
}

// RenderDashboard returns a simple HTML dashboard.
func RenderDashboard(data PageData) (string, error) {
\tif data.Title == "" {
\t\tdata.Title = "CI/CD dashboard"
\t}
\tif data.Description == "" {
\t\tdata.Description = "A small browser view for the CI/CD learning project."
\t}
\tif data.Message == "" {
\t\tdata.Message = "Ready"
\t}
\tif data.Report == "" {
\t\tdata.Report = fmt.Sprintf("No report data yet for %s.", data.Title)
\t}
\tif len(data.Tasks) == 0 {
\t\tfor _, task := range DefaultTasks() {
\t\t\tdata.Tasks = append(data.Tasks, task)
\t\t}
\t}

\tvar buf bytes.Buffer
\tif err := dashboardTmpl.Execute(&buf, data); err != nil {
\t\treturn "", err
\t}
\treturn buf.String(), nil
}
