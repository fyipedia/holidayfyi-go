# holidayfyi

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/holidayfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/holidayfyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Holiday dates, Easter calculation, and 200+ country calendars — API client for [holidayfyi.com](https://holidayfyi.com).

> **Try the interactive tools at [holidayfyi.com](https://holidayfyi.com)**

## Install

`go get github.com/fyipedia/holidayfyi-go`

## Quick Start

```go
package main

import (
    "fmt"
    "fyipedia/holidayfyi-go"
)

func main() {
    client := holidayfyi.NewClient()
    result, err := client.Search("christmas")
    if err != nil {
        panic(err)
    }
    fmt.Println(result.Total, "results")
}
```

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install holidayfyi` | [PyPI](https://pypi.org/project/holidayfyi/) |
| **npm** | `npm install holidayfyi` | [npm](https://www.npmjs.com/package/holidayfyi) |
| **Go** | `go get github.com/fyipedia/holidayfyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/holidayfyi-go) |
| **Rust** | `cargo add holidayfyi` | [crates.io](https://crates.io/crates/holidayfyi) |
| **Ruby** | `gem install holidayfyi` | [rubygems](https://rubygems.org/gems/holidayfyi) |

## Embed Widget

Embed [HolidayFYI](https://holidayfyi.com) widgets on any website with [holidayfyi-embed](https://widget.holidayfyi.com):

```html
<script src="https://cdn.jsdelivr.net/npm/holidayfyi-embed@1/dist/embed.min.js"></script>
<div data-holidayfyi="entity" data-slug="christmas"></div>
```

Zero dependencies · Shadow DOM · 4 themes (light/dark/sepia/auto) · [Widget docs](https://widget.holidayfyi.com)

## Links

- [HolidayFYI](https://holidayfyi.com) — Main site
- [API Documentation](https://holidayfyi.com/developers/)
- [OpenAPI Spec](https://holidayfyi.com/api/openapi.json)
- [Glossary](https://holidayfyi.com/glossary/)

## License

MIT
