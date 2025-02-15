// pathutils/stack.go

package pathutils

// Stack represents a simple stack data structure for path components
type Stack struct {
	items []string
}

// NewStack creates a new stack with optional initial capacity
func NewStack(capacity int) *Stack {
	return &Stack{
		items: make([]string, 0, capacity),
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// Returns empty string if stack is empty
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// IsEmpty returns true if the stack has no items
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// ToSlice returns the stack items as a slice
func (s *Stack) ToSlice() []string {
	result := make([]string, len(s.items))
	copy(result, s.items)
	return result
}
