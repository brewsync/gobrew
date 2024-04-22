# gobrew - Homebrew client package for Go

Based loosely on https://github.com/micnncim/homebrew-go, but works by unmarshaling the structured JSON output of `brew info` to Go structs.

[![Go Reference](https://pkg.go.dev/badge/github.com/brewsync/gobrew.svg)](https://pkg.go.dev/github.com/brewsync/gobrew)
[![Go Report Card](https://goreportcard.com/badge/github.com/brewsync/gobrew)](https://goreportcard.com/report/github.com/brewsync/gobrew)

`gobrew` requires [Homebrew](https://brew.sh/) be installed and linked, as it uses `brew` under the hood. See [installation instructions](https://docs.brew.sh/Installation) for how to setup Homebrew.

## Usage

> [!CAUTION]
> This package is in alpha state and not considered production-ready yet. It has not been extensively tested with Homebrew for Linux, only MacOS. There is a high possibility the interface will change. Please use caution and [report issues](https://github.com/brewsync/gobrew/labels/bug).

Import the same way as any other Go package:

```go
import "github.com/brewsync/gobrew"
```

Be sure to run

```bash
go get github.com/brewsync/gobrew
```

### Methods

Currently only listing installed packages and querying an individual package as a go struct are supported. See [examples](#examples) below.

### Examples

Check out [examples](examples/) for more, or simply:

```go
import "github.com/brewsync/gobrew"

//...

brew := gobrew.New()
ctx, cancel := context.WithTimeout(
    context.Background(),
    time.Minute,
)
defer cancel()

packages, err := brew.List(ctx)
```

## License

[MIT](LICENSE)

Brewsync Copyright [@balintb](https://balint.click/github)
