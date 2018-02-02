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
