# chapi

[![MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE.md)
[![Build Status](https://github.com/jimsmart/chapi/actions/workflows/main.yml/badge.svg)](https://github.com/jimsmart/chapi/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/jimsmart/chapi/branch/master/graph/badge.svg)](https://codecov.io/gh/jimsmart/chapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimsmart/chapi)](https://goreportcard.com/report/github.com/jimsmart/chapi)
[![Used By](https://img.shields.io/sourcegraph/rrc/github.com/jimsmart/chapi.svg)](https://sourcegraph.com/github.com/jimsmart/chapi)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jimsmart/chapi)

chapi is a [Go](https://golang.org) package providing clients and data structures for working with the [Companies House API](https://developer.companieshouse.gov.uk/api/docs/).

This API consists of:

- `Client` — higher-level API, with all methods returning structs.
- `RestClient` — lower-level API, all methods return raw JSON bytes.

Because the Companies House API is rate-limited, it may be preferable to use the `RestClient` and persist the returned data for later use. Resource structs to unmarshal the JSON into can be found in the subpackage ch.

## Installation

```bash
go get github.com/jimsmart/chapi
```

```go
import "github.com/jimsmart/chapi"

chapi.APIKey = "your_Companies_House_API_key"
```

You must provide a valid Companies House API key.

### Get an API key

1. [Register a user account with Companies House](https://developer.companieshouse.gov.uk/developer/signin).
2. Follow [these instructions](https://developer.companieshouse.gov.uk/api/docs/index/gettingStarted/apikey_authorisation.html) to get a key.

#### Keeping your key secret

Either use the key directly in your code, as shown above — or, preferably, keep it outside your code by stashing it externally in your `.zshrc` (or equivalent). For example:

```bash
export COMPANIES_HOUSE_API_KEY=your_Companies_House_API_key
```

Then get the API key from the environment variable, using code similar to this:

```go
func init() {
    // Get the key from the environment variable.
    apiKey := os.Getenv("COMPANIES_HOUSE_API_KEY")
    if len(apiKey) == 0 {
        panic("COMPANIES_HOUSE_API_KEY environment variable not set")
    }
    // Setting chapi.APIKey provides a default key for all clients.
    // If instead you wish to use a unique key per client, see chapi.NewClientWithKey.
    chapi.APIKey = apiKey
}
```

### Dependencies

- Standard library.
- [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) if you wish to run the tests.

## Example

```go
import "github.com/jimsmart/chapi"

chapi.APIKey = "your_Companies_House_API_key"

func main() {
    ch := chapi.NewClient()
    res, err := ch.Search("Richard Branson", 1, -1)
    if err != nil {
        panic(err)
    }
    // TODO do something with results
}
```

## Documentation

GoDocs [https://godoc.org/github.com/jimsmart/chapi](https://godoc.org/github.com/jimsmart/chapi)

## Testing

Package chapi includes a partial test suite but no example code at present - pull requests welcome.

To run the tests execute `go test` inside the project folder.

## License

Package chapi is copyright 2016-2017 by Jim Smart and released under the [MIT License](LICENSE.md)

## History

- v0.0.1 (2021-04-19) Use Go modules. Enable CI using GitHub Actions. Remove Travis.
