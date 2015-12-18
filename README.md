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


Split
-----

```go
segments, err := versiongo.Split("0.0.1")

if err != nil {
    panic(err)
}

fmt.Println(segments) // [0, 0, 1]

```
