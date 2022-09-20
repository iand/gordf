# gordf 

An ultra minimal RDF library for Go

[![Check Status](https://github.com/iand/gordf/actions/workflows/check.yml/badge.svg)](https://github.com/iand/gordf/actions/workflows/check.yml)
[![Test Status](https://github.com/iand/gordf/actions/workflows/test.yml/badge.svg)](https://github.com/iand/gordf/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/iand/gordf)](https://goreportcard.com/report/github.com/iand/gordf)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/iand/gordf)


## Overview

This is a minimalistic approach to a core RDF library. It is intended to contain only the basic types required for 
interoperability between RDF packages and some useful pre-defined, well-known IRIs and other constants. No attempt
is made to use Go's type systemt to enforce constraints on terms or their composition into triples or graphs. Validation
of these constraints is the responsibility of the importing package.

## Author

* [Ian Davis](http://github.com/iand) - <http://iandavis.com/>

# License

This is free and unencumbered software released into the public domain. For more
information, see <http://unlicense.org/> or the accompanying [`UNLICENSE`](UNLICENSE) file.

