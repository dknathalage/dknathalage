package main

import "testing"

func TestNewApp(t *testing.T) {
	app := NewApp()
	if app.Name != "dknathalage" {
		t.Errorf("Expected app name to be dknathalage, but got %s", app.Name)
	}
	if app.Commands[0].Name != "cloudrun" {
		t.Errorf("Expected app command name to be cloudrun, but got %s", app.Commands[0].Name)
	}
}
