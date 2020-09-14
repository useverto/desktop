package main

import "testing"

func TestFind(t *testing.T) {
	total := FindIcon("verto_desktop.png")
	if total != false {
		t.Errorf("Icon does not exist. Yet, FindIcon() finds it")
	}
}
