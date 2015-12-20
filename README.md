VersionGo
=========

Compare versions

[![Build Status](https://travis-ci.org/ushios/versiongo.svg?branch=master)](https://travis-ci.org/ushios/versiongo)


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

- [Split version string to int list](#Split)
- [Compare version strings](#Compare)

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
