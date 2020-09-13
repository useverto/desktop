package main

import "testing"

func TestFind(t *testing.T) {
	total := FindIcon()
	if total != false {
		t.Errorf("Icon does not exist. Yet, FindIcon() finds it")
	}
}
