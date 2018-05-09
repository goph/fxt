# Extensions for Uber's [fx](https://github.com/uber-go/fx)

[![Build Status](https://img.shields.io/travis/goph/fxt.svg?style=flat-square)](https://travis-ci.org/goph/fxt)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/fxt?style=flat-square)](https://goreportcard.com/report/github.com/goph/fxt)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/fxt)


**fxt is an application builder kit mainly based on [go.uber.org/fx](https://go.uber.org/fx).**


## Installation

Minimum Go version: 1.9

Since this library uses [Glide](http://glide.sh/) I recommend using it in your
project as well.

```bash
$ glide get github.com/goph/fxt
```

## Usage

As you will see, the top level (and some deeper level) packages have the `fx` prefix in their package name,
but not in their directory name. The reason for the prefix is simple: avoiding name collisions.
This is however not really idiomatic Go: although the package name can be different from the directory name,
it is recommended that they are the same for clarity. In this case the prefix is consistent though and I found it
would probably be harder to read directory names, so I left the original versions there. Ideally an IDE
should be able to handle this.

I'm going to continue using and testing it and might change the directory names as well if it turns out
to be too confusing. 


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
