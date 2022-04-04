# Sendchamp with Go
Go library for interacting with the [Sendchamp API](https://sendchamp.com)

## Table of contents
- [Get started](#get-started)
- [Basic Usage](#basic-usage)
- [SMS](#sms)
- [Voice](#voice)
- [Wallet](#wallet)
- [Examples](#examples)
- [Contributing](#contributing)
- [Todo List](#todo-list)

## Get Started

### Install 
```bash
go get github.com/fuadop/sendchamp
```

### Import 
```go
import "github.com/fuadop/sendchamp"
```

### Basic Usage
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
  
  sender := "sendchamp"
  to := []string{"2348023456087"}
  message := "my sms message"
  route := sendchamp.RouteInternational

  res, err := client.NewSms().Send(sender, to, message, route)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(res)
  // res field contains various values from res
  // like res.Status, res.Message, res.Code, etc.
}
```

## Examples
Checkout more examples in the test files and [examples](/examples) folder.

## SMS
### Initialize
```go
 publicKey := "public_key"
 mode := sendchamp.ModeLive

 client := sendchamp.NewClient(publicKey, mode)
 sms := client.NewSms()
```

### Constants
```go
// routes
sendchamp.RouteNonDND = "non_dnd"
sendchamp.RouteDND = "dnd"
sendchamp.RouteInternational = "international"

// use cases (for creation of sender ID)
sendchamp.UseCaseTransactional = "transactional"
sendchamp.UseCaseMarketing = "marketing"
sendchamp.UseCaseTransactionalAndMarketing = "transaction_marketing"

```

### Methods
- `Send`
  > Send a message to one of more phone number
  ```go
  sender := "sendchamp" 
  // slice of phone numbers
  to := []string{"2348143775374"}
  message := "Holla amigo!"
  route := sendchamp.RouteInternational

	res, err := sms.Send(sender, to, message, route)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.
  ```
- `CreateSenderID`
  > Create a sender ID
  ```go
  name := "mySenderId"
  sample := "Your otp is 200"
  useCase := sendchamp.UseCaseTransactional

	res, err := sms.CreateSenderID(name, sample, useCase)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.
  ```
- `GetDeliveryReport`
  > Get an SMS delivery report with its ID, refer to line 40 in [sms_test.go](sms_test.go)
  ```go
  // res.Data.ID is the ID gotten from the send sms (sms.Send) method.
	res, err := sms.GetDeliveryReport(res.Data.ID.(string))
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.DeliveredAt, etc.
  ```

## Voice
### Initialize
```go
 publicKey := "public_key"
 mode := sendchamp.ModeLive

 client := sendchamp.NewClient(publicKey, mode)
 voice := client.NewVoice()
```

### Constants
```go
sendchamp.VoiceTypeOutgoing = "outgoing"
```

### Methods

- `Send`
  > Send a voice message to a phone number. Refer to [voice_test.go](voice_test.go).
  ```go
  customerMobileNumber := "2348153207998"
	message := "Test from golang test suite."
	voiceType := sendchamp.VoiceTypeOutgoing // only supported type currently
	var repeat uint = 3 // repeat the voice 3 times

  res, err := voice.Send(customerMobileNumber, message, voiceType, repeat)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  ```

## Wallet
### Initialize
```go
 publicKey := "public_key"
 mode := sendchamp.ModeLive

 client := sendchamp.NewClient(publicKey, mode)
```

### Methods

- `WalletBalance`
  > Get your sendchamp wallet balance. Refer to [wallet_test.go](wallet_test.go).
  ```go
  res, err := client.WalletBalance()
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  

  fmt.Println(res.Data.AvailableBalance) // Balance in usd (type string)

  fmt.Println(res.Data.Details.BusinessAmount) // Balance in ngn (type float64)
  ```

## Contributing
PRs are greatly appreciated, help us build this hugely needed tool so anyone else can easily integrate sendchamp into their Go based projects and applications.
<br/>
1. Create a fork
2. Create your feature branch: git checkout -b my-feature
3. Commit your changes: git commit -am 'Add some feature'
4. Push to the branch: git push origin my-new-feature
5. Submit a pull request 🚀

### Todo List
* [ ] Report Service
  * [ ] Get Wallet Balance method
* [ ] Verification Service
  * [ ] Send OTP Method
  * [ ] Confirm OTP Method
* [ ] Customer Service
* [ ] Customer Group Service
* [ ] Whatsapp Service
* [ ] Email Service
