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

## License - MIT
MIT License

Copyright (c) 2020 Phillip Shipley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
