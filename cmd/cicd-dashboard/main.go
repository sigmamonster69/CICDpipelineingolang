package main

import (
\t"log"
\t"net/http"
\t"os"

\t"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
\t"github.com/sigmamonster69/CICDpipelineingolang/internal/ui"
)

func main() {
\tstateDir, err := os.Getwd()
\tif err != nil {
\t\tlog.Fatal(err)
\t}

\thttp.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
\t\ttasks, err := ui.LoadTasks(stateDir)
\t\tif err != nil {
\t\t\thttp.Error(w, err.Error(), http.StatusInternalServerError)
\t\t\treturn
\t\t}

\t\ttaskName := r.URL.Query().Get("task")
\t\tif taskName != "" {
\t\t\tif r.URL.Query().Get("done") == "1" {
\t\t\t\ttasks = ui.UpdateTaskStatus(tasks, taskName, true)
\t\t\t}
\t\t\tif r.URL.Query().Get("undo") == "1" {
\t\t\t\ttasks = ui.UpdateTaskStatus(tasks, taskName, false)
\t\t\t}
\t\t\tif err := ui.SaveTasks(stateDir, tasks); err != nil {
\t\t\t\thttp.Error(w, err.Error(), http.StatusInternalServerError)
\t\t\t\treturn
\t\t\t}
\t\t\thttp.Redirect(w, r, "/", http.StatusSeeOther)
\t\t\treturn
\t\t}

\t\thtml, err := ui.RenderDashboard(ui.PageData{
\t\t\tTitle:       "CI/CD learning dashboard",
\t\t\tDescription: "A tiny browser view for the Go CI/CD learning project.",
\t\t\tMessage:     app.Message(),
\t\t\tReport:      app.BuildReport("sam", 2, 3),
\t\t\tTasks:       tasks,
\t\t})
\t\tif err != nil {
\t\t\thttp.Error(w, err.Error(), http.StatusInternalServerError)
\t\t\treturn
\t\t}

\t\tw.Header().Set("Content-Type", "text/html; charset=utf-8")
\t\t_, _ = w.Write([]byte(html))
\t})

\tlog.Println("dashboard running on http://localhost:8080")
\tlog.Fatal(http.ListenAndServe(":8080", nil))
}
