package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"quicksearch/internal"
	"os"
	"os/exec"
)

type DesktopEntries struct {
	ctx     context.Context
	Entries []*internal.DesktopEntry `json:"entries"`
}

func NewEntries() *DesktopEntries {
	return &DesktopEntries{
		Entries: internal.GetDesktopEntries(),
	}
}

func (m *DesktopEntries) RunCommand(s string) {
	cmd := exec.Command("gtk-launch", s)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

  os.Exit(0)
}

func (m *DesktopEntries) Search(pattern string) string {
	if len(pattern) <= 1 {
		return "[]"
	}

	var filtered []internal.DesktopEntry
	for _, e := range m.Entries {
		if internal.Search(e.Name, pattern) > -1 {
			filtered = append(filtered, (*e))
		}
	}

	if len(filtered) == 0 {
		return "[]"
	}

	data, err := json.Marshal(filtered)
	if err != nil {
		log.Println(err)
	}

	return string(data)
}



// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) Exit() {
  os.Exit(0)
}

func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Ctx() context.Context {
	return a.ctx
}
