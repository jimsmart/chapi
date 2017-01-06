# chapi

[![MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE.md) [![Build Status](https://img.shields.io/travis/jimsmart/chapi/master.svg?style=flat)](https://travis-ci.org/jimsmart/chapi) [![codecov](https://codecov.io/gh/jimsmart/chapi/branch/master/graph/badge.svg)](https://codecov.io/gh/jimsmart/chapi) [![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jimsmart/chapi)

Package chapi is a [Go](https://golang.org) package providing clients and data structures for working with the [Companies House API](https://developer.companieshouse.gov.uk/api/docs/).

## Installation
```bash
$ go get github.com/jimsmart/chapi
```

```go
import "github.com/jimsmart/chapi"

chapi.APIKey = "[YOUR_COMPANIES_HOUSE_API_KEY]"
```

You must provide a valid Companies House API key.
1. (Register a user account with Companies House)[https://developer.companieshouse.gov.uk/developer/signin].
2. (Follow these instructions)[https://developer.companieshouse.gov.uk/api/docs/index/gettingStarted/apikey_authorisation.html] to get a key.
3. Either use the key directly in your code, as shown above - or stash it externally in your `.zshrc` (or equivalent)

```bash
export COMPANIES_HOUSE_API_KEY=paste_your_key_here
```

Use some setup code like to this to reference the key from the external environment variable:

```go
func init() {
	apiKey = os.Getenv("COMPANIES_HOUSE_API_KEY")
	if len(apiKey) == 0 {
		panic("COMPANIES_HOUSE_API_KEY environment variable not set")
	}
	// This will use the same key globally throughout the whole chapi package.
	// See docs if you wish to set a unique key per client.
	chapi.APIKey = apiKey
}
```

### Dependencies

- Standard library.

## Example

```go
import "github.com/jimsmart/chapi"

// TODO
```

## Documentation

GoDocs [https://godoc.org/github.com/jimsmart/chapi](https://godoc.org/github.com/jimsmart/chapi)

## Testing

Package chapi currently doesn't include any useful tests or example code - pull requests welcome.

## License

Package chapi is copyright 2016-2017 by Jim Smart and released under the [MIT License](LICENSE.md)
