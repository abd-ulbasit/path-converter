// pathutils/converter.go

package pathutils

import (
	"strings"
)

// ConvertToAbsolutePath converts a relative path to an absolute path given the current directory
func ConvertToAbsolutePath(currentDir, relativePath string) string {
	// Normalize paths by removing redundant slashes
	currentDir = normalizeSlashes(currentDir)
	relativePath = normalizeSlashes(relativePath)

	// Handle empty paths
	if relativePath == "" {
		return currentDir
	}

	// Split paths into components
	dirParts := strings.Split(currentDir, "/")
	pathParts := strings.Split(relativePath, "/")

	// Initialize stack with approximate capacity
	stack := NewStack(len(dirParts) + len(pathParts))

	// Add directory parts to stack
	for _, part := range dirParts {
		if part != "" {
			stack.Push(part)
		}
	}

	// Process each part of the relative path
	for _, part := range pathParts {
		switch part {
		case "", ".":
			continue
		case "..":
			if !stack.IsEmpty() {
				stack.Pop()
			}
		default:
			stack.Push(part)
		}
	}

	// Construct final path
	if stack.IsEmpty() {
		return "/"
	}

	return "/" + strings.Join(stack.ToSlice(), "/")
}

// normalizeSlashes removes redundant slashes and handles edge cases
func normalizeSlashes(path string) string {
	// Replace multiple consecutive slashes with a single slash
	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}
	return path
}
