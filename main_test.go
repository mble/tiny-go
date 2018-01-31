package main

import "testing"

func TestCreateExtensions(t *testing.T) {
	err := CreateExtensions()
	if err != nil {
		t.Fatalf("Error creating extensions: %v", err)
	}
}
