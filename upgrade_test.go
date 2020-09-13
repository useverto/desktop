package main

import "testing"

func TestUpgrade(t *testing.T) {
	if NeedsUpgrade() != false {
		t.Errorf("No release exists but NeedsUpgrade() returned true")
	}
}
