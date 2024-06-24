# GoFM

A [Go](https://go.dev/) interface to [Last.fm](https://last.fm/) API.

# Installation

This assumes you already have a working Go environment.

`go get` *will always pull the latest tagged release from the master branch.*

```bash
go get github.com/notiku/gofm
```

# Features

- Simple public interface.
- ~~Access to all the data exposed by the Last.fm web services~~. (Working on it)

## Usage

Import the package in your project.

```go
import "github.com/notiku/gofm"
```

Construct a new Last.fm `Network` which can be used to access the [Last.fm API](https://last.fm/api).

```go
n := gofm.New("YOUR_API_KEY", "YOUR_API_SECRET")
```

Examples can be found [here](https://github.com/notiku/gofm/examples/).

# Documentation

**NOTICE:** This library is still unfinished and in Development. Because of that there may be major changes in the future.
