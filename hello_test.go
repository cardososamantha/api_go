package main

import "testing"

func TestHello(t *testing.T) {
    result := Hello("Samantha")
    expected := "Hello, Samantha"

    if result != expected {
        t.Errorf("resultado '%s', esperado '%s'", result, expected)
    }
}