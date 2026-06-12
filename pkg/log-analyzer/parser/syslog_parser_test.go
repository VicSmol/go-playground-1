package parser

import (
	"testing"
)

func TestSyslogParser_ParseValidSyslog(t *testing.T) {
	parser := NewSyslogParser()

	tests := []struct {
		name           string
		input          string
		expectedLevel  string
		expectedComponent string
	}{
		{
			name:           "basic info log with pid",
			input:          "Jun  8 10:00:00 host nginx[123]: [INFO] User logged in",
			expectedLevel:  "INFO",
			expectedComponent: "nginx",
		},
		{
			name:           "error without pid",
			input:          "Jun  8 10:00:01 host nginx: [ERROR] Connection refused",
			expectedLevel:  "ERROR",
			expectedComponent: "nginx",
		},
		{
			name:           "warning normalized",
			input:          "Jun  8 10:00:02 host auth[456]: [WARNING] Session expiring",
			expectedLevel:  "WARN",
			expectedComponent: "auth",
		},
		{
			name:           "fatal normalized",
			input:          "Jun  8 10:00:03 host app: [FATAL] Critical error",
			expectedLevel:  "FATAL",
			expectedComponent: "app",
		},
		{
			name:           "critical normalized",
			input:          "Jun  8 10:00:04 host app[789]: [CRITICAL] Critical error",
			expectedLevel:  "FATAL",
			expectedComponent: "app",
		},
		{
			name:           "level without brackets",
			input:          "Jun  8 10:00:05 host db: INFO Query executed",
			expectedLevel:  "INFO",
			expectedComponent: "db",
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

func TestSyslogParser_ParseInvalidSyslog(t *testing.T) {
	parser := NewSyslogParser()

	tests := []struct {
		name string
		input string
	}{
		{"not syslog format", "not a syslog line"},
		{"empty string", ""},
		{"partial syslog", "Jun  8 10:00:00 host"},
		{"invalid month", "XYZ  8 10:00:00 host nginx: [INFO] User logged in"},
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

func TestSyslogParser_ParseDifferentMonths(t *testing.T) {
	parser := NewSyslogParser()

	tests := []struct {
		name string
		input string
	}{
		{"January", "Jan  1 00:00:00 host nginx: [INFO] User logged in"},
		{"February", "Feb  2 00:00:00 host nginx: [INFO] User logged in"},
		{"December", "Dec 31 23:59:59 host nginx: [INFO] User logged in"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, err := parser.Parse(tt.input)
			if err != nil {
				t.Fatalf("expected no error for %s, got %v", tt.name, err)
			}
			if entry == nil {
				t.Fatal("expected non-nil entry")
			}
		})
	}
}
