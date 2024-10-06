package main

import (
	"os"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()

	if app.Name != "dknathalage" {
		t.Errorf("expected app name 'dknathalage', got '%s'", app.Name)
	}

	if app.Usage != "Utility cli for dknathalage" {
		t.Errorf("expected app usage 'Utility cli for dknathalage', got '%s'", app.Usage)
	}

	if app.Commands == nil {
		t.Errorf("expected commands to be initialized, got nil")
	}

	found := false
	for _, command := range app.Commands {
		if command.Name == "cloudrun" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'cloudrun' command to be present in app's commands")
	}
}

func TestAppRun(t *testing.T) {
	app := NewApp()

	// Override os.Args to simulate command-line arguments
	// Simulate running "./app cloudrun"
	os.Args = []string{"./app", "cloudrun"}

	err := app.Run(os.Args)

	// Check that no errors occurred during app execution
	if err != nil {
		t.Errorf("expected no error, got '%v'", err)
	}
}
