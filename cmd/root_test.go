package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	// Capture the output
	output := captureOutput(Execute)

	expected := "=== Village Quest ===\n"
	if output != expected {
		t.Errorf("Expected %q but got %q", expected, output)
	}
}

// Helper function to capture output
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
