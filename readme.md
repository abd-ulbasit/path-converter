# Path Converter

This Go package provides functionality to convert relative paths to absolute paths given a current directory. The implementation handles various path scenarios including `.`, `..`, and multiple consecutive slashes.

## Features

- Converts relative paths to absolute paths
- Handles current directory (`.`) and parent directory (`..`) notations
- Normalizes paths by removing redundant slashes
- No dependencies on built-in path conversion functions

## Installation

Clone the repository:

```bash
git clone https://github.com/abd-ulbasit/path-converter
cd path-converter
```

## Usage

```go
import "github.com/abd-ulbasit/path-converter/pathutils"

currentDir := "/home/myhome/myfolder1"
relativePath := "../myfolder2/mydocument.txt"

absolutePath := pathutils.ConvertToAbsolutePath(currentDir, relativePath)
// Result: "/home/myhome/myfolder2/mydocument.txt"
```

## Running Tests

To run the tests, navigate to the project directory and execute:

```bash
go test ./pathutils -run TestConvertToAbsolutePath -v
```

This will run all test cases and show detailed output for each test.

## Project Structure

```
path-converter/
├── pathutils/
│   ├── converter.go     # Main implementation
│   └── converter_test.go # Test cases
└── README.md
```

## Test Cases

The package includes comprehensive tests covering various scenarios:
- Simple relative paths
- Paths with current directory notation (./)
- Paths with parent directory notation (../)
- Multiple parent directories
- Multiple consecutive slashes
- Navigation to root directory
- Empty paths

