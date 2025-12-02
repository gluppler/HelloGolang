# Go Fundamentals

This directory contains comprehensive examples of all fundamental Go concepts, following clean code principles and secure coding practices.

## Files Overview

1. **01_variables_and_types.go** - Variables, types, constants, pointers, type conversions, zero values
2. **02_functions.go** - Function declarations, parameters, return values, closures, recursion
3. **03_control_structures.go** - If/else, switch, for loops, range, break/continue, defer, goto
4. **04_arrays_and_slices.go** - Arrays, slices, operations, appending, copying, filtering
5. **05_maps.go** - Map operations, iteration, safety patterns, common patterns
6. **06_structs.go** - Struct types, fields, methods, embedding, tags, comparison
7. **07_interfaces.go** - Interface types, implementation, polymorphism, type assertions
8. **08_error_handling.go** - Error creation, custom errors, error wrapping, patterns
9. **09_goroutines.go** - Goroutines, synchronization, WaitGroup, Mutex, patterns
10. **10_channels.go** - Channel operations, select, patterns, safety
11. **11_packages_and_modules.go** - Package organization, visibility, modules
12. **12_generics.go** - Generic types and functions (Go 1.18+)
13. **13_reflection.go** - Reflection capabilities (use sparingly)
14. **14_standard_library.go** - Common standard library packages
15. **15_testing.go** - Testing framework, benchmarks, examples
16. **16_context.go** - Context for cancellation and timeouts

## Security Features

All code follows secure coding principles:

- ✅ **Input Validation**: All user inputs are validated before use
- ✅ **Bounds Checking**: Array/slice access is always bounds-checked
- ✅ **Division by Zero**: Protected against division by zero errors
- ✅ **Nil Pointer Checks**: Proper nil checks before dereferencing
- ✅ **Error Handling**: Comprehensive error handling throughout
- ✅ **Resource Cleanup**: Proper use of defer for cleanup
- ✅ **Concurrency Safety**: Proper synchronization for concurrent access
- ✅ **Type Safety**: Strong typing and type assertions where needed

## Clean Code Principles

- **Single Responsibility**: Each function has a clear, single purpose
- **Clear Naming**: Descriptive function and variable names
- **DRY (Don't Repeat Yourself)**: Reusable functions and patterns
- **Comments**: Clear documentation for complex logic
- **Error Handling**: Explicit error handling, no silent failures
- **Consistent Style**: Follows Go conventions and formatting

## Running the Examples

Each file is a standalone Go program. To run:

```bash
go run 01_variables_and_types.go
go run 02_functions.go
# ... etc
```

Or compile all:

```bash
go build *.go
```

## Key Concepts Covered

### Basic Concepts
- Variables and types
- Functions and methods
- Control flow
- Data structures (arrays, slices, maps, structs)

### Advanced Concepts
- Interfaces and polymorphism
- Error handling patterns
- Concurrency (goroutines, channels)
- Generics (Go 1.18+)
- Reflection
- Context management

### Best Practices
- Package organization
- Testing strategies
- Standard library usage
- Security considerations
- Performance optimization

## Notes

- All code is production-ready and follows Go best practices
- Security vulnerabilities have been identified and mitigated
- Code is well-documented and maintainable
- Examples demonstrate both basic and advanced patterns
- All error cases are properly handled

## Go Version

These examples require Go 1.18+ (for generics support in file 12). Most examples work with Go 1.16+.

## Contributing

When adding new examples:
1. Follow the existing code style
2. Include security checks
3. Add proper error handling
4. Include comments for complex logic
5. Test the code before committing
