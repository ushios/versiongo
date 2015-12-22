VersionGo
=========

Compare versions

[![Build Status](https://travis-ci.org/ushios/versiongo.svg?branch=master)](https://travis-ci.org/ushios/versiongo)
[![Coverage Status](https://coveralls.io/repos/ushios/versiongo/badge.svg?branch=master&service=github)](https://coveralls.io/github/ushios/versiongo?branch=master)

GettingStarted
===============

Install
--------

```
go get github.com/ushios/versiongo
```

Documentation
-------------

[![GoDoc](https://godoc.org/github.com/ushios/versiongo?status.svg)](https://godoc.org/github.com/ushios/versiongo)

Table of Contents
==================

- [Split version string to int list](#split)
- [Compare version strings](#compare)

Split
-----

```go
segments, err := versiongo.Split("0.0.1")

if err != nil {
    panic(err)
}

fmt.Println(segments) // [0, 0, 1]

```

Compare
--------

```go
result1, err1 := Compare("1.0", "1.1")
fmt.Println(result1) // LessThan

result2, err2 := Compare("0.9", "1.0")
fmt.Println(result2) // GreaterThan

result3, err3 := Compare("1.0", "1.0")
fmt.Println(result3) // Equals
```

### Errors

```go
// Compare 3 segments and 2 segments.
result, err := Compare("1.0.0", "1.0")
fmt.Println(result) // UnKnown

// Version number not found.
reuslt, err := Compare("one", "1.0.0")
fmt.Println(result) // UnKnown
```
