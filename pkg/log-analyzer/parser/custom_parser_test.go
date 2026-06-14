package parser

import (
	"testing"
)

func TestCustomParser_ParseValidCustom(t *testing.T) {
	parser := NewCustomParser()

	tests := []struct {
		name              string
		input             string
		expectedLevel     string
		expectedComponent string
	}{
		{
			name:              "basic info log",
			input:             "INFO auth: User logged in",
			expectedLevel:     "INFO",
			expectedComponent: "auth",
		},
		{
			name:              "error with colon in message",
			input:             "ERROR db: Connection: refused",
			expectedLevel:     "ERROR",
			expectedComponent: "db",
		},
		{
			name:              "warning",
			input:             "WARNING auth: Session expiring",
			expectedLevel:     "WARN",
			expectedComponent: "auth",
		},
		{
			name:              "fatal",
			input:             "FATAL app: Critical error",
			expectedLevel:     "FATAL",
			expectedComponent: "app",
		},
		{
			name:              "critical",
			input:             "CRITICAL app: Critical error",
			expectedLevel:     "CRITICAL",
			expectedComponent: "app",
		},
		{
			name:              "emergency",
			input:             "EMERGENCY app: Critical error",
			expectedLevel:     "EMERGENCY",
			expectedComponent: "app",
		},
		{
			name:              "lowercase level normalized",
			input:             "info auth: User logged in",
			expectedLevel:     "INFO",
			expectedComponent: "auth",
		},
		{
			name:              "warn vs warning",
			input:             "WARN auth: Warning message",
			expectedLevel:     "WARN",
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

func TestCustomParser_ParseInvalidCustom(t *testing.T) {
	parser := NewCustomParser()

	tests := []struct {
		name  string
		input string
	}{
		{"not custom format", "not a custom line"},
		{"missing colon", "INFO auth"},
		{"empty string", ""},
		{"only level", "INFO"},
		{"level without message", "INFO auth:"},
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

func TestCustomParser_ParseComponentWithSpecialChars(t *testing.T) {
	parser := NewCustomParser()

	tests := []struct {
		name              string
		input             string
		expectedComponent string
	}{
		{"component with underscore", "INFO auth_service: User logged in", "auth_service"},
		{"component with dash", "ERROR api-gateway: Connection failed", "api-gateway"},
		{"component with number", "WARN cache2: Cache miss", "cache2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := parser.Parse(tt.input)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if entry.Component != tt.expectedComponent {
				t.Errorf("expected component %q, got %q", tt.expectedComponent, entry.Component)
			}
		})
	}
}
