# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

## Repository Files and Folders
* <b>README.MD</b> | <i>A markdown file to help summarize the repository and its contents, and providing information on "using" the repository</i>

* <b>grade-calculator</b> | <i>A folder that houses the setup, code implementation, and tests for the Grade Calculator</i>
  * <b>go.mod</b> | <i>A file that provides relevant information on the grade-calculator module, such as the module path and minimum required version of Go; it seems that the grade-calculator directory is meant to be a Go module</i>
  * <b>grade_calculator.go</b> | <i>This file contains functionality for the Grade and GradeCalculator structs, and functions to aid in storing and organizing grades and grade types and calculating the final grade of a GradeCalculator object</i>
  * <b>grade_calculator_test.go</b> | <i>This file contains tests to test the functionality and correctness of the grade_calculator.go file</i>

## Using the Repository
Runtime requirements:
Golang 1.24+
* [Go Download and Install Overview](https://go.dev/doc/install)
* [All Go Releases](https://go.dev/dl/)

Clone the repository:
```
git clone https://github.com/Nicholas-Lucky/esep-grade-calculator.git
cd esep-grade-calculator
```

Check that Go is installed, and its version is at least 1.24.4
```
go version
```

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

## Resources
* <b>[Go Download and Install Overview](https://go.dev/doc/install)</b> | <i>Provides a good overview on downloading Go</i>
* <b>[All Go Releases](https://go.dev/dl/)</b> | <i>Provides the featured Go releases and downloads for Windows, Apple, Linux, and Source</i>
* <b>[Official Go Documentation](https://go.dev/doc/)</b> | <i>Provides a good introduction to creating Go files and modules</i>
* <b>[W3Schools Go Tutorial](https://www.w3schools.com/go/)</b> | <i>Provides some useful information (for me) on the Go syntax</i>