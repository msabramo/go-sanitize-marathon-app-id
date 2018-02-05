[![Build Status](https://travis-ci.org/msabramo/go-sanitize-marathon-app-id.svg?branch=master)](https://travis-ci.org/msabramo/go-sanitize-marathon-app-id)
[![Coverage Status](https://coveralls.io/repos/github/msabramo/go-sanitize-marathon-app-id/badge.svg?branch=travis-ci)](https://coveralls.io/github/msabramo/go-sanitize-marathon-app-id?branch=travis-ci)

# go-sanitize-marathon-app-id

> :put_litter_in_its_place: Sanitize marathon app ids. In Golang.

Inspired by https://github.com/holidaycheck/sanitize-marathon-app-id, which does the same for Javascript.

[Marathon](https://mesosphere.github.io/marathon) only allows a [limited character set](https://mesosphere.github.io/marathon/docs/rest-api.html#id) to be used in an app id.
If you generate these ids automatically based on e.g. current system information it is useful to sanitize the app id and filter all characters which are not allowed.

## Install

```
$ go get github.com/msabramo/go-sanitize-marathon-app-id
```

## Usage

### Programmatic usage

```go
import (
	sanitizeMarathon "github.com/msabramo/go-sanitize-marathon-app-id"
)

appId := sanitizeMarathon.Sanitize("FOO_λ_BAR_1.0.0") // → foo-bar-1-0-0
```

### CLI

```
$ go get -u -v github.com/msabramo/go-sanitize-marathon-app-id/cmd/...

$ sanitizemarathonappid a.b.c.d.e a_b_c_d aBcDeF foo.λ___BAR
a-b-c-d-e
a-b-c-d
abcdef
foo-bar
```
