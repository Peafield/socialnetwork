package helpers_test

import (
	"socialnetwork/pkg/helpers"
	"strings"
	"testing"
)

func TestCreateUUID(t *testing.T) {
	uuid, err := helpers.CreateUUID()
	if err != nil {
		t.Errorf("Error = %v, wantErr = nil", err)
		return
	}

	if uuid == "" {
		t.Errorf("An empty string was returned, want a non-empty string")
		return
	}

	if len(uuid) != 36 {
		t.Errorf("A string length of %v was returned, want a string length of 36", len(uuid))
		return
	}

	hyphenCount := strings.Count(uuid, "-")
	if hyphenCount != 4 {
		t.Errorf("The string contains only %v hyphens, want 4", hyphenCount)
		return
	}
}
