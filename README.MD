# ESEP Grade Calculator

This library is a simple utility to calculate final grades for educational courses. It supports three types of grades: assignments, exams, and essays, with configurable weights.

## Features
- Calculate weighted average grades across different grade types
- Convert numerical grades to letter grades (A, B, C, D, F)
- Support for assignments (50% weight), exams (35% weight), and essays (15% weight)

## Grade Scale
- A: >=90
- B: >=80
- C: >=70
- D: >=60
- F: <60

## Tech Stack
- **Language**: Go 1.24+
- **Testing**: Go's built-in testing package

### Running Tests
Run unit tests:
```bash
cd grade-calculator/
go test -v
```

Run unit tests with code coverage:
```bash
cd grade-calculator/
go test -coverprofile=coverage.txt -v
```

View coverage report:
```bash
go tool cover -html=coverage.txt
```

### Usage Example
```go
calculator := NewGradeCalculator()
calculator.AddGrade("Homework 1", 95, Assignment)
calculator.AddGrade("Midterm", 88, Exam)
calculator.AddGrade("Final Essay", 92, Essay)
finalGrade := calculator.GetFinalGrade() // Returns "A"


## File Structure
- `grade_calculator.go`: Main implementation of the grade calculator containing functions and structs
- `grade_calculator_test.go`: Unit tests for the calculator
- `go.mod`: Go module configuration specifying version number of Go
- `README.MD`: This documentation file