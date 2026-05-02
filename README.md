# LeetCode Solutions in Go

A structured collection of LeetCode problem solutions implemented in Go, organized by topic with comprehensive tests and progress tracking.

## Repository Structure

Problems are organized by topic/category for easy navigation:

```
leetcode/
├── arrays/              # Array problems
├── strings/             # String manipulation
├── linked_lists/        # Linked list problems
├── trees/               # Binary trees, BST, etc.
├── dynamic_programming/ # DP problems
├── graphs/              # Graph algorithms
├── sorting/             # Sorting algorithms
├── binary_search/       # Binary search variants
├── backtracking/        # Backtracking problems
├── math/                # Mathematical problems
├── hash_table/          # Hash table problems
├── stack/               # Stack problems
└── queue/               # Queue problems
```

## Progress Tracking

| # | Problem Name | Difficulty | Topic(s) | Solution | Status |
|---|--------------|------------|----------|----------|--------|
| 1 | Two Sum | Easy | Arrays, Hash Table | [Solution](./arrays/two_sum.go) | ✅ |

**Stats**: 1 solved | 0 in progress

## Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific topic:
```bash
go test ./arrays -v
```

Run tests with coverage:
```bash
go test ./... -cover
```

Generate coverage report:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Problem Solution Pattern

Each problem follows a consistent structure:

**Solution file** (`topic/problem_name.go`):
- Package declaration
- Problem description with LeetCode link
- Examples and constraints
- Function implementation(s)

**Test file** (`topic/problem_name_test.go`):
- Table-driven tests
- Multiple test cases including edge cases
- Descriptive test names

## Adding New Problems

1. Identify the primary topic category
2. Create `topic/problem_name.go` with the solution
3. Create `topic/problem_name_test.go` with tests
4. Update the progress tracking table in this README
5. Run tests to verify: `go test ./topic -v`

## Topics Coverage

- **Arrays**: 1 problem
- **Strings**: 0 problems
- **Linked Lists**: 0 problems
- **Trees**: 0 problems
- **Dynamic Programming**: 0 problems
- **Graphs**: 0 problems
- **Other**: 0 problems

## Resources

- [LeetCode](https://leetcode.com/)
- [Go Documentation](https://golang.org/doc/)
- [Go Testing Package](https://golang.org/pkg/testing/)
