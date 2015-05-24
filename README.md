# go-rtb

[![GoDoc](https://godoc.org/github.com/mxmCherry/go-rtb/2.3/rtb?status.svg)](https://godoc.org/github.com/mxmCherry/go-rtb/2.3/rtb)

Common [OpenRTB](//github.com/openrtb/OpenRTB) [v2.3](//github.com/openrtb/OpenRTB/blob/master/OpenRTB-API-Specification-Version-2-3-FINAL.pdf) Object types for Go programming language

**Warning!** Currently, this package is quite unstable (types and names may change). Also, even after v1.0 this repo will contain only recent version (v1+). So, if you need to use this code in production, take a look on [godep](//github.com/tools/godep).


# Progress
- [x] 3.2 Object Specifications
	- [x] 3.2.1 Object: BidRequest
	- [x] 3.2.2 Object: Imp
	- [x] 3.2.3 Object: Banner
	- [x] 3.2.4 Object: Video
	- [x] 3.2.5 Object: Native
	- [x] 3.2.6 Object: Site
	- [x] 3.2.7 Object: App
	- [x] 3.2.8 Object: Publisher
	- [x] 3.2.9 Object: Content
	- [x] 3.2.10 Object: Producer
	- [x] 3.2.11 Object: Device
	- [x] 3.2.12 Object: Geo
	- [x] 3.2.13 Object: User
	- [x] 3.2.14 Object: Data
	- [x] 3.2.15 Object: Segment
	- [x] 3.2.16 Object: Regs
	- [x] 3.2.17 Object: Pmp
	- [x] 3.2.18 Object: Deal
- [x] 4.2 Object Specifications
	- [x] 4.2.1 Object: BidResponse
	- [x] 4.2.2 Object: SeatBid
	- [x] 4.2.3 Object: Bid
- [ ] Code quality/review
	- [x] Review struct key types, which refer to section "5. Enumerated Lists Specification"
	- [ ] Check types for struct keys (see Guidelines - Types)
	- [x] Use struct pointers for optional keys (e.g., App.Content)
	- [ ] Add json directive "omitempty" for optional keys


# Guidelines

## Naming convention
- [UpperCamelCase](http://en.wikipedia.org/wiki/CamelCase)
- Capitalized abbreviations (e.g., AT, COPPA, PMP etc.)
- Capitalized ID keys

## Types
- Key types should be chosen according to OpenRTB specification (attribute types)
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
