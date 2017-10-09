# apixu-go
Golang library for Apixu Weather API http://www.apixu.com

## Installation
```bash
$ go get github.com/mohan3d/apixu-go
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/mohan3d/apixu-go"
)

func main() {
	client := apixu.NewClient("<APIXU_KEY>")

	current, err := client.Current("Paris")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v °C, %v °F\n", current.Current.TempC, current.Current.TempF)

	forecast, err := client.Forecast("London", 3)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", forecast.Location.Country)
}
```

## Testing
**APIXU_KEY** must be exported to environment variables before running tests.

```bash
$ export APIXU_KEY=<YOUR_APIXU_KEY>
```