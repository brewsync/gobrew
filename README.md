# gobrew - Homebrew client package for Go

Based on https://github.com/micnncim/homebrew-go, but unmarshaling the structured JSON output of `brew info`.

[![GoDoc][godoc-badge]][godoc]

## Usage

Import the same way as any other Go package:

```
import "github.com/brewsync/gobrew"
```

Be sure to run `go get github.com/brewsync/gobrew`

### Methods

Currently only listing installed packages and querying an individual package as a go struct are supported. See [examples](#examples) below.

### Examples

Check out [examples](examples/) for more, or simply:

```
import "github.com/brewsync/gobrew"

...

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

Brewsync Copyright @balintb
