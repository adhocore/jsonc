# adhocore/jsonc

[![Latest Version](https://img.shields.io/github/release/adhocore/jsonc.svg?style=flat-square)](https://github.com/adhocore/jsonc/releases)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Go Report](https://goreportcard.com/badge/github.com/adhocore/jsonc)](https://goreportcard.com/report/github.com/adhocore/jsonc)
[![Test](https://github.com/adhocore/jsonc/actions/workflows/test-action.yml/badge.svg)](https://github.com/adhocore/jsonc/actions/workflows/test-action.yml)
[![Codecov](https://img.shields.io/codecov/c/github/adhocore/jsonc/main.svg?style=flat-square)](https://codecov.io/gh/adhocore/jsonc)
[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Lightweight+fast+and+deps+free+commented+json+parser+for+Golang&url=https://github.com/adhocore/jsonc&hashtags=go,golang,parser,json,json-comment)


- Lightweight JSON comment stripper library for Go.
- Makes possible to have comment in any form of JSON data.
- Supported comments: single line `// comment` or multi line `/* comment */`.
- Also strips trailing comma at the end of array or object, eg:
    - `[1,2,,]` => `[1,2]`
    - `{"x":1,,}` => `{"x":1}`
- Handles literal LF (newline/linefeed) within string notation so that we can have multiline string
- Supports JSON string inside JSON string
- Zero dependency (no vendor bloat).

Find jsonc in [pkg.go.dev](https://pkg.go.dev/github.com/adhocore/jsonc).

## Installation

```sh
go get -u github.com/adhocore/jsonc
```

## Usage

Import and init library:
```go
import (
	"fmt"

	"github.com/adhocore/jsonc"
)

j := jsonc.New()
```

Strip and parse:
```go
json := []byte(`{
	// single line comment
	"a'b": "apple'ball",
	/* multi line
	   comment */
	"cat": [
		"dog",
		"// not a comment",
		"/* also not a comment */",
	],
	"longtext": "long text in
	multple lines",
}`)

var out map[string]interface{}

j.Unmarshall(json, &out)
fmt.Printf("%+v\n", out)
```

Strip comments/commas only:
```go
json := []byte(`{"some":"json",}`)
json = j.Strip(json)
```

Using strings instead of byte array:
```go
json := `{"json": "some
	text",// comment
	"array": ["a",]
}`
json = j.StripS(json)
```

Parsing from JSON file directly:
```go
var out map[string]interface{}

j.UnmarshalFile("./examples/test.json5", &out)
fmt.Printf("%+v\n", out)
```

> Run working [examples](./examples/main.go) with `go run examples/main.go`.

---
## License

> &copy; [MIT](./LICENSE) | 2021-2099, Jitendra Adhikari

---
### Other projects
My other golang projects you might find interesting and useful:

- [**gronx**](https://github.com/adhocore/gronx) - Lightweight, fast and dependency-free Cron expression parser (due checker, next run finder), task scheduler and/or daemon for Golang (tested on v1.13 and above) and standalone usage.
- [**urlsh**](https://github.com/adhocore/urlsh) - URL shortener and bookmarker service with UI, API, Cache, Hits Counter and forwarder using postgres and redis in backend, bulma in frontend; has [web](https://urlssh.xyz) and cli client.
- [**fast**](https://github.com/adhocore/fast) - Check your internet speed with ease and comfort right from the terminal.
- [**goic**](https://github.com/adhocore/goic) - Go Open ID Connect, is OpenID connect client library for Golang, supports the Authorization Code Flow of OpenID Connect specification.
- [**chin**](https://github.com/adhocore/chin) - A Golang command line tool to show a spinner as user waits for some long running jobs to finish.
