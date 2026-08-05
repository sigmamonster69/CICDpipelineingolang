package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

const stateFileName = ".task-state.json"

type Task struct {
	Name      string `json:"name"`
	Done      bool   `json:"done"`
	DoneAt    string `json:"done_at,omitempty"`
	Suggested bool   `json:"suggested,omitempty"`
}

type PageData struct {
	Title       string
	Description string
	Message     string
	Report      string
	Notes       []string
	Tasks       []Task
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
      --bg: #f3efe8;
      --panel: #fffaf2;
      --panel-2: #f7f0e3;
      --text: #1c1a17;
      --muted: #6d6459;
      --line: #dccfbf;
      --accent: #b45309;
      --accent-2: #0f766e;
      --done: #15803d;
    }
    * { box-sizing: border-box; }
    body {
      font-family: Georgia, "Times New Roman", serif;
      margin: 0;
      min-height: 100vh;
      background: var(--bg);
      color: var(--text);
    }
    .shell {
      max-width: 1120px;
      margin: 0 auto;
      padding: 28px 20px 40px;
    }
    .hero {
      background: linear-gradient(135deg, #fffaf2 0%, #f7efe1 100%);
      border: 1px solid var(--line);
      border-radius: 24px;
      padding: 28px;
      box-shadow: 0 18px 35px rgba(62, 42, 10, 0.08);
    }
    h1, h2, h3, p { margin-top: 0; }
    h1 {
      font-size: clamp(2rem, 5vw, 3.25rem);
      line-height: 1;
      margin-bottom: 12px;
    }
    .muted { color: var(--muted); }
    .report {
      background: #1f2937;
      color: #f8fafc;
      padding: 16px 18px;
      border-radius: 16px;
      overflow-x: auto;
      margin: 18px 0 18px;
      font-family: "SFMono-Regular", Menlo, Consolas, monospace;
      font-size: 14px;
    }
    .grid {
      display: grid;
      grid-template-columns: 1.25fr 0.75fr;
      gap: 18px;
      margin-top: 18px;
    }
    .panel {
      background: var(--panel);
      border: 1px solid var(--line);
      border-radius: 20px;
      padding: 20px;
      box-shadow: 0 10px 24px rgba(62, 42, 10, 0.05);
    }
    .panel.alt {
      background: var(--panel-2);
    }
    .section-head {
      display: flex;
      justify-content: space-between;
      gap: 12px;
      align-items: baseline;
      margin-bottom: 14px;
    }
    .pill {
      display: inline-flex;
      align-items: center;
      gap: 8px;
      border-radius: 999px;
      border: 1px solid var(--line);
      background: rgba(255, 255, 255, 0.7);
      padding: 8px 12px;
      font-size: 13px;
      color: var(--muted);
    }
    .task-list, .note-list {
      list-style: none;
      padding: 0;
      margin: 0;
    }
    .task {
      border: 1px solid var(--line);
      border-radius: 16px;
      padding: 14px 16px;
      margin-bottom: 12px;
      display: flex;
      justify-content: space-between;
      gap: 16px;
      align-items: start;
      background: rgba(255,255,255,0.65);
    }
    .task.done {
      border-color: rgba(21, 128, 61, 0.32);
      background: rgba(21, 128, 61, 0.05);
    }
    .task strong {
      display: block;
      margin-bottom: 4px;
      font-size: 16px;
    }
    .task-meta, .note-list li {
      font-size: 14px;
      color: var(--muted);
      line-height: 1.5;
    }
    .status {
      font-size: 13px;
      border-radius: 999px;
      padding: 7px 11px;
      white-space: nowrap;
      border: 1px solid var(--line);
      color: var(--muted);
      text-decoration: none;
      background: #fff;
      align-self: center;
    }
    .status.done {
      color: var(--done);
      border-color: rgba(21, 128, 61, 0.32);
    }
    .status:hover {
      border-color: var(--accent);
      color: var(--accent);
    }
    .steps {
      display: grid;
      gap: 10px;
    }
    .step {
      border-left: 4px solid var(--accent-2);
      padding: 10px 12px;
      background: rgba(255,255,255,0.72);
      border-radius: 12px;
    }
    @media (max-width: 860px) {
      .grid { grid-template-columns: 1fr; }
      .shell { padding-inline: 14px; }
    }
  </style>
</head>
<body>
  <main class="shell">
    <section class="hero">
      <div class="pill">CI/CD learning dashboard</div>
      <h1>{{.Title}}</h1>
      <p class="muted">{{.Description}}</p>
      <p><strong>Status:</strong> {{.Message}}</p>
      <div class="report">{{.Report}}</div>
    </section>

    <div class="grid">
      <section class="panel">
        <div class="section-head">
          <h2>Task board</h2>
          <p class="muted">Tick items on and off as you learn.</p>
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
      </section>

      <aside class="panel alt">
        <div class="section-head">
          <h2>Docs</h2>
          <p class="muted">Short notes for the project.</p>
        </div>
        <div class="steps">
          {{range .Notes}}
          <div class="step">{{.}}</div>
          {{end}}
        </div>
      </aside>
    </div>
  </main>
</body>
</html>
`))

// DefaultTasks returns the default checklist of learning activities for the dashboard.
func DefaultTasks() []Task {
	return []Task{
		{Name: "Read the task board in the website"},
		{Name: "Pick one small file and improve it"},
		{Name: "Make a commit with that small change"},
		{Name: "Push your commit to GitHub"},
		{Name: "Open the repo on GitHub and check the update"},
		{Name: "Repeat with one more tiny change"},
		{Name: "Open practice.go"},
		{Name: "Open internal/app/app.go"},
		{Name: "Open internal/app/app_test.go"},
		{Name: "Open .github/workflows/ci.yml"},
		{Name: "Explain each file in your own words"},
		{Name: "Add one tiny Go function yourself"},
		{Name: "Add a test for that function"},
		{Name: "Push the change and watch the CI pipeline"},
		{Name: "Add Dockerfile"},
		{Name: "Add docker-compose.yml"},
		{Name: "Add Makefile"},
		{Name: "Add a lint step to CI"},
		{Name: "Add a security scan step after that"},
	}
}

// DefaultNotes returns the short documentation panel content shown on the dashboard.
func DefaultNotes() []string {
	return []string{
		"Keep the app small and readable.",
		"Use Go for the server and helpers.",
		"Use the dashboard as the main learning surface.",
		"Mark tasks done as you finish them.",
		"To make another GitHub contribution: change one small file, commit it, push it, then mark the task done.",
		"Add one small change at a time.",
	}
}

// LoadTasks loads task completion state from disk if the state file exists.
func LoadTasks(dir string) ([]Task, error) {
	tasks := DefaultTasks()
	path := filepath.Join(dir, stateFileName)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// SaveTasks writes the current task state to disk as JSON.
func SaveTasks(dir string, tasks []Task) error {
	path := filepath.Join(dir, stateFileName)

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

// UpdateTaskStatus updates a specific task's completion status and timestamp.
func UpdateTaskStatus(tasks []Task, taskName string, done bool) []Task {
	for i := range tasks {
		if tasks[i].Name == taskName {
			tasks[i].Done = done
			if done {
				tasks[i].DoneAt = time.Now().Format("2006-01-02")
			} else {
				tasks[i].DoneAt = ""
			}
		}
	}
	return tasks
}

// RenderDashboard generates an HTML dashboard page from the provided data.
func RenderDashboard(data PageData) (string, error) {
	if data.Title == "" {
		data.Title = "CI/CD dashboard"
	}
	if data.Description == "" {
		data.Description = "A small browser view for the CI/CD learning project."
	}
	if data.Message == "" {
		data.Message = "Ready"
	}
	if data.Report == "" {
		data.Report = fmt.Sprintf("No report data yet for %s.", data.Title)
	}

	if len(data.Tasks) == 0 {
		for _, task := range DefaultTasks() {
			data.Tasks = append(data.Tasks, task)
		}
	}

	if len(data.Notes) == 0 {
		data.Notes = DefaultNotes()
	}

	var buf bytes.Buffer
	if err := dashboardTmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
