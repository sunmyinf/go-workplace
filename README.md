# go-workplace
go-workplace provides packages for workplace app.

## Usage
```go
package main

import (
  "log"

  "github.com/sunmyinf/go-workplace/webhook"
  "github.com/sunmyinf/go-workplace/decode"
)

func main() {
  server := webhook.NewServer("App Secret", "Access Token", "Verify Token")
  server.HandleObjectFunc(decode.ObjectPage, func(req decode.Request) error {
    log.Println("hello page!")
    return nil
  })
  if err := server.ListenAndServe(); err != nil {
    log.Printf("failed to launch server: %v\n", err)
  }
}

```
