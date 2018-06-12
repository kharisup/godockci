package main

import "testing"

// test function names begin with uppercase Test
func TestMessage(t *testing.T) {
	project := new(Project)
	project.PushedAt = "now"
	project.Name = "TestProject"
	msg := Message(project)
	if msg != "<p>TestProject: Latest commit: now</p>" {
		t.Errorf("Incorrect message: %s", msg)
	}
}

// test function names begin with uppercase Test
func Test2Message(t *testing.T) {
	project := new(Project)
	project.PushedAt = "yesterday"
	project.Name = "Test2Project"
	msg := Message(project)
	if msg != "<p>Test2Project: Latest commit: yesterday</p>" {
		t.Errorf("Incorrect message: %s", msg)
	}
}
