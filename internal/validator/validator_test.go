package validator

import (
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	v := New()

	if v == nil {
		t.Fatal("New() returned nil")
	}

	if v.Errors == nil {
		t.Error("Errors map was not initialized")
	}

	if len(v.Errors) != 0 {
		t.Error("Errors map should be empty initially")
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		name     string
		errors   map[string]string
		expected bool
	}{
		{
			name:     "empty errors map",
			errors:   map[string]string{},
			expected: true,
		},
		{
			name:     "nil errors map",
			errors:   nil,
			expected: true,
		},
		{
			name: "non-empty errors map",
			errors: map[string]string{
				"field": "error message",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{Errors: tt.errors}

			if v.Valid() != tt.expected {
				t.Errorf("Valid() = %v, want %v", v.Valid(), tt.expected)
			}
		})
	}
}

func TestAddError(t *testing.T) {
	v := New()

	v.AddError("field1", "error message 1")

	if len(v.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(v.Errors))
	}

	if v.Errors["field1"] != "error message 1" {
		t.Errorf("Expected 'error message 1', got %q", v.Errors["field1"])
	}

	v.AddError("field2", "error message 2")

	if len(v.Errors) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(v.Errors))
	}

	v.AddError("field1", "new error message")

	if len(v.Errors) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(v.Errors))
	}

	if v.Errors["field1"] != "error message 1" {
		t.Errorf("Expected original error message, got %q", v.Errors["field1"])
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name           string
		ok             bool
		key            string
		message        string
		shouldAddError bool
	}{
		{
			name:           "condition is true",
			ok:             true,
			key:            "field",
			message:        "error message",
			shouldAddError: false,
		},
		{
			name:           "condition is false",
			ok:             false,
			key:            "field",
			message:        "error message",
			shouldAddError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			v.Check(tt.ok, tt.key, tt.message)

			if tt.shouldAddError {
				if len(v.Errors) != 1 {
					t.Errorf("Expected 1 error, got %d", len(v.Errors))
				}
				if v.Errors[tt.key] != tt.message {
					t.Errorf("Expected %q, got %q", tt.message, v.Errors[tt.key])
				}
			} else {
				if len(v.Errors) != 0 {
					t.Errorf("Expected no errors, got %d", len(v.Errors))
				}
			}
		})
	}
}

func TestIn(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		list     []string
		expected bool
	}{
		{
			name:     "value in list",
			value:    "apple",
			list:     []string{"apple", "banana", "cherry"},
			expected: true,
		},
		{
			name:     "value not in list",
			value:    "grape",
			list:     []string{"apple", "banana", "cherry"},
			expected: false,
		},
		{
			name:     "empty list",
			value:    "apple",
			list:     []string{},
			expected: false,
		},
		{
			name:     "empty value in list",
			value:    "",
			list:     []string{"", "apple"},
			expected: true,
		},
		{
			name:     "empty value not in list",
			value:    "",
			list:     []string{"apple", "banana"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := In(tt.value, tt.list...)
			if result != tt.expected {
				t.Errorf("In(%q, %v) = %v, want %v", tt.value, tt.list, result, tt.expected)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	emailRegex := EmailRX

	tests := []struct {
		name     string
		value    string
		rx       *regexp.Regexp
		expected bool
	}{
		{
			name:     "valid email",
			value:    "test@example.com",
			rx:       emailRegex,
			expected: true,
		},
		{
			name:     "invalid email",
			value:    "invalid-email",
			rx:       emailRegex,
			expected: false,
		},
		{
			name:     "empty string",
			value:    "",
			rx:       emailRegex,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Matches(tt.value, tt.rx)
			if result != tt.expected {
				t.Errorf("Matches(%q, regex) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected bool
	}{
		{
			name:     "all unique values",
			values:   []string{"apple", "banana", "cherry"},
			expected: true,
		},
		{
			name:     "duplicate values",
			values:   []string{"apple", "banana", "apple"},
			expected: false,
		},
		{
			name:     "empty slice",
			values:   []string{},
			expected: true,
		},
		{
			name:     "single value",
			values:   []string{"apple"},
			expected: true,
		},
		{
			name:     "all same values",
			values:   []string{"apple", "apple", "apple"},
			expected: false,
		},
		{
			name:     "empty strings",
			values:   []string{"", ""},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unique(tt.values)
			if result != tt.expected {
				t.Errorf("Unique(%v) = %v, want %v", tt.values, result, tt.expected)
			}
		})
	}
}

func TestEmailRX(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "valid email",
			email:    "test@example.com",
			expected: true,
		},
		{
			name:     "valid email with numbers",
			email:    "user123@example.com",
			expected: true,
		},
		{
			name:     "valid email with special chars",
			email:    "user.name+tag@example-domain.com",
			expected: true,
		},
		{
			name:     "invalid email without @",
			email:    "invalid.email.com",
			expected: false,
		},
		{
			name:     "invalid email without domain",
			email:    "user@",
			expected: false,
		},
		{
			name:     "empty string",
			email:    "",
			expected: false,
		},
		{
			name:     "just @",
			email:    "@",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EmailRX.MatchString(tt.email)
			if result != tt.expected {
				t.Errorf("EmailRX.MatchString(%q) = %v, want %v", tt.email, result, tt.expected)
			}
		})
	}
}
