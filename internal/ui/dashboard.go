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

// stateFileName is the name of the JSON file used to persist task completion state.
// This file stores which tasks have been marked as done by the user.
const stateFileName = ".task-state.json"

// Task represents a single checklist item displayed on the dashboard.
// Each task tracks its name, completion status, and optional completion date.
// The Suggested field indicates whether this is a recommended learning activity.
type Task struct {
	Name      string `json:"name"`                // Display name of the task
	Done      bool   `json:"done"`                // Whether the task has been completed
	DoneAt    string `json:"done_at,omitempty"`   // Date when task was completed (YYYY-MM-DD format)
	Suggested bool   `json:"suggested,omitempty"` // Whether this is a suggested task
}

// PageData holds all the values needed to render the dashboard HTML page.
// This struct is passed to the template engine for generating the UI.
type PageData struct {
	Title       string   // Page title shown in browser tab and header
	Description string   // Subtitle/description of the dashboard
	Message     string   // Current status message from the application
	Report      string   // Formatted report string from app.BuildReport()
	Notes       []string // List of documentation notes/tips to display
	Tasks       []Task   // Slice of tasks to display in the checklist
}

// dashboardTmpl is the HTML template for the CI/CD learning dashboard.
// It includes embedded CSS for styling and uses Go's html/template package.
// The template displays: hero section with status, task board, and documentation notes.
// Features responsive design that adapts to mobile and desktop screens.
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
// These tasks guide users through the CI/CD learning journey step by step.
// Each task represents a small, achievable learning goal.
// Returns:
//   - []Task: slice of predefined tasks covering basic to advanced CI/CD concepts
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
// These notes provide guidance and tips for navigating the learning project.
// Returns:
//   - []string: slice of helpful notes and instructions
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
// If no state file is found, it returns the default task list.
// The state is stored as JSON in the .task-state.json file.
// Parameters:
//   - dir: directory path where the state file should be located
//
// Returns:
//   - []Task: loaded tasks with their completion state, or defaults if file doesn't exist
//   - error: any error encountered during file reading or JSON parsing
func LoadTasks(dir string) ([]Task, error) {
	// Start with default tasks
	tasks := DefaultTasks()
	path := filepath.Join(dir, stateFileName)

	// Try to read the state file
	data, err := os.ReadFile(path)
	if err != nil {
		// Return defaults if file doesn't exist (not an error condition)
		if os.IsNotExist(err) {
			return tasks, nil
		}
		// Return error for other read failures
		return nil, err
	}

	// Parse JSON data into tasks slice
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// SaveTasks writes the current task state to disk as JSON.
// This persists user progress so it survives page refreshes and server restarts.
// Parameters:
//   - dir: directory path where the state file should be written
//   - tasks: slice of tasks with their current completion state
//
// Returns:
//   - error: any error encountered during JSON marshaling or file writing
func SaveTasks(dir string, tasks []Task) error {
	path := filepath.Join(dir, stateFileName)

	// Marshal tasks to formatted JSON
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	// Write to file with standard permissions (readable/writable by owner, readable by others)
	return os.WriteFile(path, data, 0o644)
}

// UpdateTaskStatus updates a specific task's completion status and timestamp.
// When marking a task as done, it records the current date. When undoing, it clears the date.
// Parameters:
//   - tasks: slice of tasks to update (modified in place)
//   - taskName: name of the task to update (exact match required)
//   - done: true to mark complete, false to mark incomplete
//
// Returns:
//   - []Task: the updated tasks slice (same reference as input)
func UpdateTaskStatus(tasks []Task, taskName string, done bool) []Task {
	for i := range tasks {
		if tasks[i].Name == taskName {
			tasks[i].Done = done
			if done {
				// Record completion date in YYYY-MM-DD format
				tasks[i].DoneAt = time.Now().Format("2006-01-02")
			} else {
				// Clear completion date when marking as not done
				tasks[i].DoneAt = ""
			}
		}
	}
	return tasks
}

// RenderDashboard generates an HTML dashboard page from the provided data.
// It applies default values for any empty fields and executes the HTML template.
// Parameters:
//   - data: PageData struct containing all information to display
//
// Returns:
//   - string: rendered HTML as a string
//   - error: any template execution error
func RenderDashboard(data PageData) (string, error) {
	// Apply sensible defaults for empty fields
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

	// Use default tasks if none provided
	if len(data.Tasks) == 0 {
		for _, task := range DefaultTasks() {
			data.Tasks = append(data.Tasks, task)
		}
	}

	// Use default notes if none provided
	if len(data.Notes) == 0 {
		data.Notes = DefaultNotes()
	}

	// Execute template and capture output
	var buf bytes.Buffer
	if err := dashboardTmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
