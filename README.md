# gopan

## GoPAN

Validate and Generate PAN (PrimaryAccountNumber a.k.a CreditCardNumber).

## Installation

```bash
go get github.com/ken-ton/gopan
```

## Quick Start

```go
import "github.com/ken-ton/gopan"

// Generate
pan := gopan.Generate("AMERICAN EXPRESS") // 378282246310005

// Validate
isValid := gopan.IsValid("4012-8888-8888-1881") // true

// Get Brand or IssuingNetwork
brand := gopan.GetBrand("3566 0020 2036 0505") // JCB
```
