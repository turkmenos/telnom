# telnom

Go library for parsing, validating, and formatting Turkmen phone numbers.

## Install

```bash
go get github.com/turkmenos/telnom
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/turkmenos/telnom"
)

func main() {
	fmt.Println(telnom.Validate("+99371123456"))

	number, err := telnom.Parse("+99371123456")
	if err != nil {
		return
	}

	fmt.Println(number.CountryCode)
	fmt.Println(number.National)
	fmt.Println(number.E164())

	formatted, _ := telnom.Format("+99371123456")
	fmt.Println(formatted)

	normalized, _ := telnom.Normalize("+993-71-12-34-56")
	fmt.Println(normalized)

	fmt.Println(telnom.OperatorOf("+99371123456"))
	fmt.Println(telnom.IsMobile("+99371123456"))
	fmt.Println(telnom.IsFixedLine("+99312123456"))
	fmt.Println(telnom.TypeOf("+99371123456"))
}
```

## Features

- Parse Turkmen phone numbers
- Validate `+993` numbers
- Format phone numbers
- Normalize phone numbers to E.164 format
- Detect mobile numbers
- Detect fixed-line numbers
- Classify numbers as mobile, fixed-line, or unknown
- Detect supported operators

## License

MIT
