package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	// Test with valid input and banner
	os.Args = []string{"test", "test string", "thinkertoy"}
	main()
	if os.Args[0] != "test" {
		t.Errorf("Expected input to be 'test', but got '%s'", os.Args[0])
	}

	// Test with invalid banner
	os.Args = []string{"test", "invalid", "standardo"}
	main()
	if os.Args[1] != "invalid" {
		t.Errorf("Expected banner to be 'invalid', but got '%s'", os.Args[1])
	}

	// Test with no input
	os.Args = []string{""}
	main()
	if os.Args[0] != "" {
		t.Errorf("Expected no input, but got '%s'", os.Args[0])
	}
}
