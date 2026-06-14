package parser

import (
	"testing"
)

func TestJSONParser_ParseValidJSON(t *testing.T) {
	parser := NewJSONParser()

	tests := []struct {
		name              string
		input             string
		expectedLevel     string
		expectedComponent string
	}{
		{
			name:              "basic info log",
			input:             `{"level":"INFO","component":"auth","message":"User logged in"}`,
			expectedLevel:     "INFO",
			expectedComponent: "auth",
		},
		{
			name:              "error with timestamp",
			input:             `{"level":"ERROR","component":"db","message":"Connection refused","timestamp":"2026-06-08T10:00:00Z"}`,
			expectedLevel:     "ERROR",
			expectedComponent: "db",
		},
		{
			name:              "warning",
			input:             `{"level":"WARNING","component":"cache","message":"Cache miss"}`,
			expectedLevel:     "WARN",
			expectedComponent: "cache",
		},
		{
			name:              "fatal",
			input:             `{"level":"FATAL","component":"app","message":"Critical error"}`,
			expectedLevel:     "FATAL",
			expectedComponent: "app",
		},
		{
			name:              "critical",
			input:             `{"level":"CRITICAL","component":"app","message":"Critical error"}`,
			expectedLevel:     "CRITICAL",
			expectedComponent: "app",
		},
		{
			name:              "emergency",
			input:             `{"level":"EMERGENCY","component":"app","message":"Critical error"}`,
			expectedLevel:     "EMERGENCY",
			expectedComponent: "app",
		},
		{
			name:              "lowercase level normalized",
			input:             `{"level":"info","component":"auth","message":"User logged in"}`,
			expectedLevel:     "INFO",
			expectedComponent: "auth",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := parser.Parse(tt.input)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if entry == nil {
				t.Fatal("expected non-nil entry")
			}
			if entry.Level != tt.expectedLevel {
				t.Errorf("expected level %q, got %q", tt.expectedLevel, entry.Level)
			}
			if entry.Component != tt.expectedComponent {
				t.Errorf("expected component %q, got %q", tt.expectedComponent, entry.Component)
			}
		})
	}
}

func TestJSONParser_ParseInvalidJSON(t *testing.T) {
	parser := NewJSONParser()

	tests := []struct {
		name  string
		input string
	}{
		{"not json", "not valid json"},
		{"empty object", "{}"},
		{"missing level", `{"component":"auth","message":"User logged in"}`},
		{"missing component", `{"level":"INFO","message":"User logged in"}`},
		{"empty string", ""},
		{"partial json", `{"level":"INFO"}`},
		{"partial json 2", `{"component":"auth"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := parser.Parse(tt.input)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if entry != nil {
				t.Errorf("expected nil entry on error, got %v", entry)
			}
		})
	}
}

func TestJSONParser_ParseEmptyLevel(t *testing.T) {
	parser := NewJSONParser()
	input := `{"level":"","component":"auth","message":"User logged in"}`

	entry, err := parser.Parse(input)
	if err == nil {
		t.Fatalf("expected error for empty level, got nil")
	}
	if entry != nil {
		t.Errorf("expected nil entry on error, got %v", entry)
	}
}

func TestJSONParser_ParseEmptyComponent(t *testing.T) {
	parser := NewJSONParser()
	input := `{"level":"INFO","component":"","message":"User logged in"}`

	entry, err := parser.Parse(input)
	if err == nil {
		t.Fatalf("expected error for empty component, got nil")
	}
	if entry != nil {
		t.Errorf("expected nil entry on error, got %v", entry)
	}
}
