// pathutils/stack_test.go

package pathutils

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Basic stack operations", func(t *testing.T) {
		stack := NewStack(5)

		// Test empty stack
		if !stack.IsEmpty() {
			t.Error("New stack should be empty")
		}

		// Test push and pop
		stack.Push("first")
		stack.Push("second")

		if stack.IsEmpty() {
			t.Error("Stack should not be empty after push")
		}

		if item := stack.Pop(); item != "second" {
			t.Errorf("Expected 'second', got %q", item)
		}

		if item := stack.Pop(); item != "first" {
			t.Errorf("Expected 'first', got %q", item)
		}

		if !stack.IsEmpty() {
			t.Error("Stack should be empty after popping all items")
		}

		// Test pop on empty stack
		if item := stack.Pop(); item != "" {
			t.Errorf("Expected empty string from empty stack, got %q", item)
		}
	})

	t.Run("ToSlice conversion", func(t *testing.T) {
		stack := NewStack(3)
		expected := []string{"first", "second", "third"}

		for _, item := range expected {
			stack.Push(item)
		}

		result := stack.ToSlice()

		if len(result) != len(expected) {
			t.Errorf("Expected slice length %d, got %d", len(expected), len(result))
		}

		for i, item := range result {
			if item != expected[i] {
				t.Errorf("At index %d: expected %q, got %q", i, expected[i], item)
			}
		}
	})
}
