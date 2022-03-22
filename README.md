# Sendchamp with Go
Go library for interacting with the [Sendchamp API](https://sendchamp.com)

## Usage
```go
package main

import (
  "log"
  "fmt"

  "github.com/fuadop/sendchamp"
)

var publicKey string = "your-public-key"
var mode string = sendchamp.ModeLive // can be set to test mode (sendchamp.ModeTest) too

func main() {
  client := sendchamp.NewClient(publicKey, mode)
  res, err := client.NewSms().Send("sendchamp", []string{"2348023456087"}, "My sms message", sendchamp.RouteInternational)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(res)
  // res field contains various values from res
  // like res.Status, res.Message, res.Code, etc.
}
```
