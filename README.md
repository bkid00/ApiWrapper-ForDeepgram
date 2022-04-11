# ApiWrapper-ForDeepgram


## Installation
Install this package with:
```
go get github.com/agouil/deepgram-go
```


## Usage
```go
package main

import (
    "fmt"
    "github.com/agouil/deepgram-go"
)

func main() {
    deepgramClient := deepgram.Deepgram{
        ApiKey: <DEEPGRAM_API_KEY>,
    }
    result, err := deepgramClient.CheckBalance()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(
        "Balance:", result.Balance,
        "UserID:", result.UserId,
    )
}
```


### Testing
Run tests with:
```
go test -v
```

## License
[MIT](LICENSE)
