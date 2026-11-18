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

func TestSpdxLicenses_GetLicenseByIdentifier(t *testing.T) {
	t.Parallel()
	s, err := NewSpdxLicenses()
	if err != nil {
		t.Fatalf("NewSpdxLicenses() returned error: %v", err)
	}

	tests := []struct {
		identifier string
		want       License
		found      bool
	}{
		{"MIT", License{Identifier: "MIT", Name: "MIT License", OSIApproved: true}, true},
		{"mit", License{Identifier: "MIT", Name: "MIT License", OSIApproved: true}, true},
		{"apache-2.0", License{Identifier: "Apache-2.0", Name: "Apache License 2.0", OSIApproved: true}, true},
		{"GPL-2.0+", License{Identifier: "GPL-2.0+", Name: "GNU General Public License v2.0 or later", OSIApproved: true, Deprecated: true}, true},
		{"NONE", License{}, false},
		{"NOASSERTION", License{}, false},
		{"LicenseRef-foo", License{}, false},
		{"MIT OR Apache-2.0", License{}, false},
		{"proprietary", License{}, false},
		{"", License{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.identifier, func(t *testing.T) {
			t.Parallel()
			got, found := s.GetLicenseByIdentifier(tt.identifier)
			if found != tt.found {
				t.Fatalf("GetLicenseByIdentifier(%q) found = %v, want %v", tt.identifier, found, tt.found)
			}
			if got != tt.want {
				t.Fatalf("GetLicenseByIdentifier(%q) = %+v, want %+v", tt.identifier, got, tt.want)
			}
		})
	}
}
