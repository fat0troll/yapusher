# UUID

[![GoDoc](http://godoc.org/gitlab.com/pztrn/uuid?status.svg)](http://godoc.org/gitlab.com/pztrn/uuid)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/pztrn/uuid)](https://goreportcard.com/report/gitlab.com/pztrn/uuid)

Package uuid provides a pure Go implementation of Universally Unique Identifiers (UUID) variant as defined in RFC-4122. This package supports both the creation and parsing of UUIDs in different formats.

This package supports the following UUID versions:

* Version 1, based on timestamp and MAC address (RFC-4122)
* Version 2, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1)
* Version 3, based on MD5 hashing of a named value (RFC-4122)
* Version 4, based on random numbers (RFC-4122)
* Version 5, based on SHA-1 hashing of a named value (RFC-4122)

## Project History

This project was originally forked from the [github.com/satori/go.uuid](https://github.com/satori/go.uuid) repository after it appeared to be no longer maintained, while exhibiting [critical flaws](https://github.com/satori/go.uuid/issues/73). We have decided to take over this project to ensure it receives regular maintenance for the benefit of the larger Go community.

We'd like to thank Maxim Bublis for his hard work on the original iteration of the package.

After that it was forked again when developers [don't want to make everyone happy](https://github.com/gofrs/uuid/issues/48) because of things that out of their control. Also, package was "kind of" refactored to allow some great features be implemented.

There is no guarantee that ``github.com/gofrs/uuid`` changes will be ported.

## License

This source code of this package is released under the MIT License. Please see the [LICENSE](https://gitlab.com/pztrn/uuid/blob/master/LICENSE) for the full content of the license.

## Installation

It is recommended to use a package manager like `dep` that understands tagged releases of a package, as well as semantic versioning.

If you are unable to make use of a dependency manager with your project, you can use the `go get` command to download it directly:

```Shell
$ go get gitlab.com/pztrn/uuid
```

## Requirements

Due to subtests not being supported in older versions of Go, this package is only regularly tested against Go 1.7+. This package may work perfectly fine with Go 1.2+, but support for these older versions is not actively maintained.

## Go 1.11 Modules

This package should be compatible with Go's modules thing. I'm not using it extensively ATM due to "really beta" stage IMO, so PRs that fixes any problems are encouraged.

## Usage

Here is a quick overview of how to use this package. For more detailed
documentation, please see the [GoDoc Page](http://godoc.org/gitlab.com/pztrn/uuid).

```go
package main

import (
	"log"

	"gitlab.com/pztrn/uuid"
)

// Create a Version 4 UUID, panicking on error.
// Use this form to initialize package-level variables.
var u1 = uuid.Must(uuid.NewV4())

func main() {
	// Create a Version 4 UUID.
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u2)

	// Parse a UUID from a string.
	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	u3, err := uuid.FromString(s)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", s, err)
	}
	log.Printf("successfully parsed UUID %v", u3)
}
```

## References

* [RFC-4122](https://tools.ietf.org/html/rfc4122)
* [DCE 1.1: Authentication and Security Services](http://pubs.opengroup.org/onlinepubs/9696989899/chap5.htm#tagcjh_08_02_01_01)
