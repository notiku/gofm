# GoFM

GoFM is a [Go](https://go.dev/) package that provides bindings to the [Last.fm](https://last.fm/) API.

# Getting Started

## Installing

This assumes you already have a working Go environment.

`go get` *will always pull the latest tagged release from the master branch.*

```go get github.com/notiku/gofm```

## Usage

Import the package in your project.

```go
import "github.com/notiku/gofm"
```

Construct a new Last.fm client which can be used to access the Last.fm API functions.

```go
client := gofm.New("YOUR_API_KEY")
```

Examples can be found [here](https://github.com/notiku/gofm/examples).

# Documentation

**NOTICE:** This library is still unfinished and in Development. Because of that there may be major changes in the future.
