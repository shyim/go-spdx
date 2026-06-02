package spdx

import (
	"testing"
)

func TestNewSpdxLicenses(t *testing.T) {
	t.Parallel()
	_, err := NewSpdxLicenses()
	if err != nil {
		t.Fatalf("NewSpdxLicenses() returned error: %v", err)
	}
}

func TestSpdxLicenses_Validate(t *testing.T) {
	t.Parallel()
	s, err := NewSpdxLicenses()
	if err != nil {
		t.Fatalf("NewSpdxLicenses() returned error: %v", err)
	}

	tests := []struct {
		identifier string
	}{
		{"MIT"},
		{"mit"},
		{"LGPL-2.1-only"},
		{"GPL-3.0-or-later"},
		{"(LGPL-2.1-only or GPL-3.0-or-later)"},
	}

	for _, tt := range tests {
		t.Run(tt.identifier, func(t *testing.T) {
			t.Parallel()
			valid, err := s.Validate(tt.identifier)
			if err != nil {
				t.Fatalf("Validate(%q) returned error: %v", tt.identifier, err)
			}
			if !valid {
				t.Fatalf("Validate(%q) = false, want true", tt.identifier)
			}
		})
	}
}
