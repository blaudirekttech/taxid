# Tax ID

Syntactically validate TaxID based on country specific validations.

## Supported Countries

Following countries are supported for validations currently:

- Germany
- Austria

## Example Usage

```go
package main

import (
	"fmt"

	"github.com/blaudirekttech/taxid"
)

func main() {
	if err := taxid.Validate(taxid.Germany, "valid tax id"); err != nil {
		fmt.Println("Tax ID is valid")
	}
}
```