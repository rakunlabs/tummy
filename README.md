# Tummy Time ⏳

[![License](https://img.shields.io/github/license/rakunlabs/tummy?color=red&style=flat-square)](https://raw.githubusercontent.com/rakunlabs/tummy/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/rakunlabs_tummy?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=rakunlabs_tummy)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/rakunlabs/tummy/test.yml?branch=main&logo=github&style=flat-square&label=ci)](https://github.com/rakunlabs/tummy/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/rakunlabs/tummy?style=flat-square)](https://goreportcard.com/report/github.com/rakunlabs/tummy)
[![Go PKG](https://raw.githubusercontent.com/rakunlabs/.github/main/assets/badges/gopkg.svg)](https://pkg.go.dev/github.com/rakunlabs/tummy)

Tummy is a Go library that provides a simple way to manipulate time in your tests.  
It allows you to control the flow of time, making it easier to test time-dependent code.

```sh
go get github.com/rakunlabs/tummy
```

## Usage

```go
// enable tummy time manipulation, usually in the test setup
tummy.Enable()

tummy.SetTime(time.Date(1919, 5, 19, 12, 0, 0, 0, time.UTC))

fmt.Println("Current time:", tummy.Now())
// Current time: 1919-05-19 12:00:00.000000119 +0000 UTC

tummy.AddDuration(2 * time.Hour)
fmt.Println("After 2 hours:", tummy.Now())
// After 2 hours: 1919-05-19 14:00:00.000019847 +0000 UTC

tummy.AddDate(4, 4, 4)
fmt.Println("After 1 day:", tummy.Now())
// After 1 day: 1923-10-23 14:00:00.000015641 +0000 UTC
```
