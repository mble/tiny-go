package main

import "testing"

func TestCreateExtensions(t *testing.T) {
	err := CreateExtensions()
	if err != nil {
		t.Fatalf("Error creating extensions: %v", err)
	}
}

func TestPokeRedis(t *testing.T) {
	err := PokeRedis()
	if err != nil {
		t.Fatalf("Error poking redis: %v", err)
	}
}
