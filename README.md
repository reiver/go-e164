# go-e164

Package **e164** implements tools for working with E.164 formatted phone-numbers, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-e164

[![GoDoc](https://godoc.org/github.com/reiver/go-e164?status.svg)](https://godoc.org/github.com/reiver/go-e164)

## Example

Here is an example of parsing a phone-number if E.164 format:

```golang
countryCode, nationalDestinationCode, subscriberNumber, err := e164.ParseTolerantly("+16045551234")
```

## Import

To import package **e164** use `import` code like the following:
```
import "github.com/reiver/go-e164"
```

## Installation

To install package **e164** do the following:
```
GOPROXY=direct go get github.com/reiver/go-e164
```

## Author

Package **e164** was written by [Charles Iliya Krempeaux](http://reiver.link)
