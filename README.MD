# ESEP Grade Calculator

This package provides a Grade Calcualtor for calculating weighted final scores based on assignments, exams, and essays and outputting as a letter grade.

Runtime requirements:
Golang 1.24+

Ensure Golang is installed and in PATH

Usage:
Input student scores through the AddGrade function 
Calculate using calculateNumericalGrade()
Get letter grade using GetFinalGrade()

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
