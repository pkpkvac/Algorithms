# Go Algorithms

A collection of algorithm solutions implemented in Go with comprehensive test suites, following Go best practices for backend engineering.

## Project Structure

```
go/
├── go.mod                    # Go module definition
├── README.md                 # This file
├── array_hashing/           # Category: Array & Hashing problems
│   ├── remove_element/      # Individual problem folder
│   │   ├── solution.go      # Main solution implementation
│   │   ├── solution_test.go # Test suite (Go convention)
│   │   └── README.md        # Problem description
│   ├── two_sum/
│   └── contains_duplicate/
├── dynamic_programming/      # Category: DP problems
├── graphs/                  # Category: Graph algorithms
└── ...
```

## Naming Conventions

### Folders

- **Categories**: `snake_case` (e.g., `array_hashing`, `dynamic_programming`)
- **Problems**: `snake_case` (e.g., `remove_element`, `two_sum`)

### Files

- **Solution**: `solution.go` (contains the main algorithm)
- **Tests**: `solution_test.go` (Go convention for test files)
- **Documentation**: `README.md` (problem description and examples)

### Go Code

- **Functions**: `PascalCase` for exported functions (e.g., `RemoveElement`)
- **Variables**: `camelCase` for local variables
- **Constants**: `PascalCase` for exported constants

## Running Tests

### Run all tests in a specific problem:

```bash
cd array_hashing/remove_element
go test -v
```

### Run tests with coverage:

```bash
go test -v -cover
```

### Run benchmarks:

```bash
go test -bench=.
```

### Run all tests in the project:

```bash
go test ./...
```

## Go Best Practices for Backend Engineering

1. **Test Coverage**: Every solution includes comprehensive test cases
   - **Why**: Production backend code requires tests for reliability
   - **Practice**: Write tests for every function you implement
   - **Interview Prep**: Backend interviews often test your testing skills
2. **Benchmarking**: Performance tests for optimization
3. **Documentation**: Clear problem descriptions and complexity analysis
4. **Module Management**: Proper `go.mod` for dependency management
5. **Error Handling**: Robust error handling patterns
6. **Code Organization**: Clean separation of concerns

## Why Tests Are Essential (Not Overkill!)

### **For Backend Engineering:**

- **Production Systems**: Every backend service needs comprehensive tests
- **Go Culture**: The Go community expects high test coverage
- **Debugging**: Tests catch edge cases before they reach production
- **Refactoring**: Tests let you optimize code with confidence
- **Team Collaboration**: Tests document expected behavior

### **For Interviews:**

- **System Design**: You'll need to discuss testing strategies
- **Code Reviews**: Interviewers expect you to write tests
- **Best Practices**: Shows you understand production requirements

## What Makes a Go Algorithm "Complete"?

A Go algorithm problem is considered complete when it includes:

### **Required Files:**

1. **`solution.go`** - Clean, well-documented implementation
2. **`solution_test.go`** - Comprehensive test suite with:
   - Edge cases (empty arrays, single elements, etc.)
   - LeetCode examples
   - Performance benchmarks
3. **`README.md`** - Problem description and complexity analysis

### **Test Requirements:**

- **At least 5-8 test cases** covering different scenarios
- **Edge cases** (empty inputs, single elements, all same values)
- **LeetCode examples** (if applicable)
- **Benchmark tests** for performance analysis
- **Table-driven tests** (Go best practice)

### **Code Quality:**

- **Time/Space complexity** documented
- **Clean variable names** and comments
- **Error handling** where appropriate
- **Go idioms** and best practices

## Adding New Problems

1. Create a new folder under the appropriate category
2. Add `solution.go` with your implementation
3. Add `solution_test.go` with comprehensive tests
4. Add `README.md` with problem description
5. Run tests to ensure everything works

## Example Usage

```bash
# Navigate to a problem
cd array_hashing/remove_element

# Run tests
go test -v

# Run with coverage
go test -v -cover

# Run benchmarks
go test -bench=.
```

This structure will serve you well as you transition to backend engineering with Go! 🚀
