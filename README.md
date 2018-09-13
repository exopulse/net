# exopulse net package
Golang network related utilities.

[![CircleCI](https://circleci.com/gh/exopulse/net.svg?style=svg)](https://circleci.com/gh/exopulse/net)
[![GitHub license](https://img.shields.io/github/license/exopulse/unit.svg)](https://github.com/exopulse/unit/blob/master/LICENSE)

# Overview

This package contains types, helpers and utilities to make network programming easier.

## Features

- Address - simple container containing network host address and port with related operations

# Using net package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/net
 
Include net in your application.

```go
import "github.com/exopulse/net"
```

## Use Address type to parse and handle network addresses

```go
a, err := net.ParseAddress("127.0.0.1:3000")
```

## Access address parts

```go
host := a.Host()
port := a.Port()
```

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

Net package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/net/blob/master/LICENSE)
