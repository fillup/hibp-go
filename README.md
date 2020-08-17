# hibp-go: A simple HaveIBeenPwned library for Go
_(seo help: golang)_

I needed a simple way to check if a given password is known according to 
HaveIBeenPwned. This library does just that. I don't plan to add other HIBP
API support, but I'm happy to accept PR's if others would like to. 

## Usage

```go
package main

import (
    "log"

    "github.com/fillup/hibp-go"
)

func main() {
    isPwned, err := hibp.IsPwned("pass123")
    if err != nil {
        log.Fatal(err) 
    }
    
    if isPwned {
        // do something
    }
}
```
