package main

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		wantErr  bool
		errMatch string
	}{
		{
			name: "Valid User",
			input: User{
				Name:  "Tahsin",
				Email: "tahsin@example.com",
			},
			wantErr: false,
		},
		{
			name: "Empty Name (too short)",
			input: User{
				Name:  "",
				Email: "tahsin@example.com",
			},
			wantErr:  true,
			errMatch: "must be at least 2 characters",
		},
		{
			name: "Name Too short (1 char)",
			input: User{
				Name:  "A",
				Email: "tahsin@example.com",
			},
			wantErr:  true,
			errMatch: "must be at least 2 characters",
		},
		{
			name: "Name Too Long",
			input: User{
				Name:  "A very long name exceeding ten chars",
				Email: "tahsin@example.com",
			},
			wantErr:  true,
			errMatch: "must be at most 10 characters",
		},
		{
			name: "Empty Email (required)",
			input: User{
				Name:  "Tahsin",
				Email: "",
			},
			wantErr:  true,
			errMatch: "is required",
		},
		{
			name: "Invalid Email (missing @)",
			input: User{
				Name:  "Tahsin",
				Email: "tahsin-at-example.com",
			},
			wantErr:  true,
			errMatch: "must be a valid email",
		},
		{
			name:     "Non-struct input",
			input:    "not a struct",
			wantErr:  true,
			errMatch: "expected a struct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil && !strings.Contains(err.Error(), tt.errMatch) {
				t.Errorf("validate() error = %v, want errMatch %v", err, tt.errMatch)
			}
		})
	}
}
