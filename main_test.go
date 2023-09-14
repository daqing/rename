package main

import "testing"

func TestNewName(t *testing.T) {
	name := newName("xyz.tar.gz")

	if name != "xyz.tar" {
		t.Errorf("New name is wrong, got %s\n", name)
	}
}
