package utils

import (
	"testing"
)

func TestRemoveSlashesAndConvertToString(t *testing.T) {
    // Define test cases
    tests := []struct {
        input    any
        expected string
    }{
        // Test cases with various input types and expected results
        {"/path/to/file", "pathtofile"},
        {"C:\\path\\to\\file", "C:pathtofile"},
        {"simple/string", "simplestring"},
        {"", ""},
        {12345, "12345"},
        {3.14, "3.14"},
        {true, "true"},
        {false, "false"},
    }

    for _, tt := range tests {
        t.Run(tt.expected, func(t *testing.T) {
            result := RemoveSlashesAndConvertToString(tt.input)
            if result != tt.expected {
                t.Errorf("removeSlashesAndConvertToString(%v) = %v; want %v", tt.input, result, tt.expected)
            }
        })
    }
}
