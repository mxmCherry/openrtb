# openrtb [![GoDoc](https://godoc.org/github.com/mxmCherry/openrtb?status.svg)](https://godoc.org/github.com/mxmCherry/openrtb) [![Build Status](https://travis-ci.org/mxmCherry/openrtb.svg?branch=master)](https://travis-ci.org/mxmCherry/openrtb)

[OpenRTB](https://www.iab.com/guidelines/real-time-bidding-rtb-project/) [v2.5](https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf) types for Go programming language (Golang)

Also includes [OpenRTB](https://www.iab.com/guidelines/real-time-bidding-rtb-project/) [Dynamic Native Ads API
Specification Version 1.1](https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-Native-Ads-Specification-1-1_2016.pdf) types:
- [4 Native Ad Request Markup Details](native/request/)
- [5 Native Ad Response Markup Details](native/response/)

# Using

```bash
go get -u "github.com/mxmCherry/openrtb"
```

```go
import "github.com/mxmCherry/openrtb"
```

This repo follows [semver](http://semver.org/) - see [releases](https://github.com/mxmCherry/openrtb/releases).
Master always contains latest code, so better use some package manager to vendor specific version.

# Guidelines

## Naming convention
- [UpperCamelCase](http://en.wikipedia.org/wiki/CamelCase)
- Capitalized abbreviations (e.g., `AT`, `COPPA`, `PMP` etc.)
- Capitalized `ID` keys

## Types
- Key types should be chosen according to OpenRTB specification (attribute types)
- Numeric types:
	- `int64` - time, duration, length, unbound enums (like `BidRequest.at` - exchange-specific auctions types are > 500)
	- `int8` - short enums (with values <= 127), boolean-like attributes (like `BidRequest.test`)
	- `uint64` - width, height, bitrate etc. (unbound positive numbers)
	- `float64` - coordinates, prices etc.
- Enums:
	- all enums, described in section 5, must be typed with section name singularized (e.g., "5.2 Banner Ad Types" -> `type BannerAdType int8`)
	- all typed enums must have constants for each element, prefixed with type name (e.g., "5.2 Banner Ad Types - XHTML Text Ad (usually mobile)" -> `const BannerAdTypeXHTMLTextAd BannerAdType = 1`)
	- never use `iota` for enum constants
	- section "5.1 Content Categories" should remain untyped and have no constants

## Documentation ([godoc](https://godoc.org/github.com/mxmCherry/openrtb))
- [Godoc: documenting Go code](http://blog.golang.org/godoc-documenting-go-code)
- Each entity (type, struct key or constant) should be documented
- Comments for entities should be copy-pasted "as-is" from OpenRTB specification (except section 5 - replace "table" with "list" there; ideally, each sentence must be on a new line)

## Code organization
- Each RTB type should be kept in its own file, named after type
- File names are in underscore_case, e.g., `type BidRequest` should be declared in `bid_request.go`
- [go fmt your code](https://blog.golang.org/go-fmt-your-code)

# TODO
- [ ] Review all integral types, probably, switch everything to signed ones or just to `int`?
- [ ] Consider switching back to `encoding/json.RawMessage`, as Go 1.8 fixed serialisation for non-ptr (probably, when Go 1.9 or even 1.10 is out)
- [x] Review enum types (typed enum attributes + constants)
- [ ] Review types, that are enums (or "open enums" like `BidRequest.at`) themselves, but not described in section 5 - make them typed
