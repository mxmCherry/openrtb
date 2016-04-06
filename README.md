# openrtb

[![GoDoc](https://godoc.org/github.com/mxmCherry/openrtb?status.svg)](https://godoc.org/github.com/mxmCherry/openrtb)

[OpenRTB](//github.com/openrtb/OpenRTB) [v2.3](//github.com/openrtb/OpenRTB/blob/master/OpenRTB-API-Specification-Version-2-3-FINAL.pdf) types for Go programming language (golang)

**Warning!** This package is considered unstable, consider using [glide](https://github.com/Masterminds/glide) to vendor specific version (commit hash).

# Goals

Provide base for OpenRTB-related projects, focusing on:
- Extensive documentation
- Strict specification (using unsigned numeric types for values, that are not meant to be signed; don't overuse pointers to avoid nil dereferencing etc.)
- Efficient memory usage (using numeric types large enough just to hold intended values etc.)

# Guidelines

## Naming convention
- [UpperCamelCase](http://en.wikipedia.org/wiki/CamelCase)
- Capitalized abbreviations (e.g., AT, COPPA, PMP etc.)
- Capitalized ID keys

## Types
- Key types should be chosen according to OpenRTB v2.3 specification (attribute types)
- Numeric types:
	- architecture-independent, e.g., ```int32``` instead of ```int```
	- signed integral types should be used only when absolutely needed (value may contain negative numbers), unsigned integral types are preferred
	- enumerations should be represented with minimal integral types, e.g., ```uint8``` or ```int8``` for enumerations with <= 256 variants
	- for floating-point attributes only ```float64``` type should be used

## Documentation
- [Godoc: documenting Go code](http://blog.golang.org/godoc-documenting-go-code)
- Each entity (type or struct key) should be documented
- Comments for entities should be copy-pasted "as-is" from OpenRTB specification

## Code organization
- Each RTB type should be kept in its own file, named after type
- File names are in underscore_case, e.g., ```type BidRequest``` should be declared in ```bid_request.go```
- [go fmt your code](https://blog.golang.org/go-fmt-your-code)
