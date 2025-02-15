// pathutils/converter_test.go

package pathutils

import (
	"testing"
)

func TestConvertToAbsolutePath(t *testing.T) {
	tests := []struct {
		name         string
		currentDir   string
		relativePath string
		expected     string
	}{
		{
			name:         "Simple relative path",
			currentDir:   "/home/myhome/myfolder1",
			relativePath: "mydocument.txt",
			expected:     "/home/myhome/myfolder1/mydocument.txt",
		},
		{
			name:         "Path with current directory notation",
			currentDir:   "/home/myhome/myfolder1",
			relativePath: "./mydocument.txt",
			expected:     "/home/myhome/myfolder1/mydocument.txt",
		},
		{
			name:         "Path with parent directory notation",
			currentDir:   "/home/myhome/myfolder1",
			relativePath: "../myfolder2/mydocument.txt",
			expected:     "/home/myhome/myfolder2/mydocument.txt",
		},
		{
			name:         "Multiple parent directories",
			currentDir:   "/home/myhome/myfolder1/subfolder",
			relativePath: "../../myfolder2/mydocument.txt",
			expected:     "/home/myhome/myfolder2/mydocument.txt",
		},
		{
			name:         "Handle multiple consecutive slashes",
			currentDir:   "/home///myhome////myfolder1",
			relativePath: "mydocument.txt",
			expected:     "/home/myhome/myfolder1/mydocument.txt",
		},
		{
			name:         "Navigate to root",
			currentDir:   "/home/myhome",
			relativePath: "../..",
			expected:     "/",
		},
		{
			name:         "Empty relative path",
			currentDir:   "/home/myhome",
			relativePath: "",
			expected:     "/home/myhome",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertToAbsolutePath(tt.currentDir, tt.relativePath)
			if result != tt.expected {
				t.Errorf("ConvertToAbsolutePath(%q, %q) = %q, want %q",
					tt.currentDir, tt.relativePath, result, tt.expected)
			}
		})
	}
}
