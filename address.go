// Package net contains general network related helpers like Address which serves as a container for host and port.
package net

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Address holds network host and port.
type Address struct {
	host string
	port int
}

// Host returns host.
func (a Address) Host() string {
	return a.host
}

// Port returns port.
func (a Address) Port() int {
	return a.port
}

// String returns string presentation of this address. If 0 port is omitted from the output.
func (a Address) String() string {
	if a.port == 0 {
		return a.host
	}

	return fmt.Sprintf("%s:%d", a.host, a.port)
}

// WithDefaultPort returns address with specified port only if original port is not assigned (it has value of 0).
func (a Address) WithDefaultPort(port int) Address {
	if a.port == 0 {
		return Address{a.host, port}
	}

	return a
}

// MarshalJSON marshals this address into JSON value.
func (a *Address) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(a.String())), nil
}

// UnmarshalJSON un-marshals bytes into Address value.
func (a *Address) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))

	if err != nil {
		return err
	}

	host, port, err := parseAddress(s)

	if err != nil {
		return err
	}

	a.host = host
	a.port = port

	return nil
}

// ParseAddress parses string and returns an Address.
func ParseAddress(value string) (Address, error) {
	host, port, err := parseAddress(value)

	if err != nil {
		return Address{}, err
	}

	return Address{host, port}, nil
}

func parseAddress(value string) (string, int, error) {
	var host string
	var portStr string

	colonAt := strings.LastIndex(value, ":")

	if colonAt == -1 {
		host, portStr = strings.TrimSpace(value), "0"
	} else {
		host, portStr = strings.TrimSpace(value[:colonAt]), strings.TrimSpace(value[colonAt+1:])
	}

	if len(host) == 0 {
		return "", 0, errors.New("empty address")
	}

	port, err := strconv.ParseInt(portStr, 10, 16)

	if err != nil {
		return "", 0, err
	}

	return host, int(port), nil
}

// MustParseAddress is a helper method that wraps a call to ParseAddress() and panics if the error is non-nil.
func MustParseAddress(value string) Address {
	addr, e := ParseAddress(value)

	if e != nil {
		log.Panicf("failed to parse address [%s]: %s", value, e)
	}

	return addr
}
