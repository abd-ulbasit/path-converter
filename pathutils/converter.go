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

	// Remove empty parts
	dirParts = removeEmptyParts(dirParts)

	// Initialize result stack
	result := make([]string, 0, len(dirParts))
	result = append(result, dirParts...)

	// Process each part of the relative path
	for _, part := range pathParts {
		switch part {
		case "":
			continue
		case ".":
			continue
		case "..":
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		default:
			result = append(result, part)
		}
	}

	// Construct final path
	if len(result) == 0 {
		return "/"
	}

	return "/" + strings.Join(result, "/")
}

// normalizeSlashes removes redundant slashes and handles edge cases
func normalizeSlashes(path string) string {
	// Replace multiple consecutive slashes with a single slash
	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}
	return path
}

// removeEmptyParts removes empty strings from the path parts
func removeEmptyParts(parts []string) []string {
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}
