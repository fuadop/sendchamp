# Sendchamp with Go
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-4-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
Go library for interacting with the [Sendchamp API](https://sendchamp.com)

## Table of contents
- [Get started](#get-started)
- [Basic Usage](#basic-usage)
- [SMS](#sms)
- [Voice](#voice)
- [Wallet](#wallet)
- [Verification](#verification)
- [Whatsapp](#whatsapp)
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

var publicKey = &sendchamp.Keys{
	PublicKey: "your-public-key",
} //pass as nil to use the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
var mode string = sendchamp.ModeLive // can be set to test mode (sendchamp.ModeTest) too

func main() {
  //if the first parameter is nil, the package would look for the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
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
Checkout more examples in the test files and this example http server repo https://github.com/fuadop/my_server.

## SMS
### Initialize
```go
 publicKey :=  &sendchamp.Keys{
	PublicKey: "your-public-key",
 } //pass as nil to use the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
 mode := sendchamp.ModeLive
 //if the first parameter is nil, the package would look for the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
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
 publicKey :=  &sendchamp.Keys{
	PublicKey: "your-public-key",
 } //pass as nil to use the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
 mode := sendchamp.ModeLive
 //if the first parameter is nil, the package would look for the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
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
	customerMobileNumbers := []string{"2348153207998"}
	message := "Test from golang test suite."
	voiceType := sendchamp.VoiceTypeOutgoing // only supported type currently
	var repeat uint = 3 // repeat the voice 3 times

  res, err := voice.Send(customerMobileNumbers, message, voiceType, repeat)
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
publicKey :=  &sendchamp.Keys{
	PublicKey: "your-public-key",
 } //pass as nil to use the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
 mode := sendchamp.ModeLive
 //if the first parameter is nil, the package would look for the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
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

  // use vscode intellisense/auto-complete to see more fields.
  ```

## Verification
### Initialize
```go
 publicKey :=  &sendchamp.Keys{
	PublicKey: "your-public-key",
 } //pass as nil to use the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
 mode := sendchamp.ModeLive
 //if the first parameter is nil, the package would look for the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
 client := sendchamp.NewClient(publicKey, mode)
 verification := client.NewVerification()
```

### Constants
```go
// otp channels
sendchamp.OTPChannelSMS = "sms"
sendchamp.OTPChannelEmail = "email"

// otp token types
sendchamp.OTPTokenTypeNumeric = "numeric"
sendchamp.OTPTypeAlphaNumeric = "alphanumeric"
```

### Methods

- `SendOTP`
  > Send OTP to customer with sms or email channel. Refer to [verification_test.go](verification_test.go).
  ```go
  // create this outside the main function
  // to contain any fields of your type
  type metadata struct {
    FirstName string // important - export fields and add json annotations
    LastName  string
  }

  payload := sendchamp.SendOTPPayload{
    // specify channel , "sms" or "email"
		Channel:              sendchamp.OTPChannelSMS,
		Sender:               "Sendchamp",
    // specify "numeric" or "alphanumeric"
		TokenType:            sendchamp.OTPTokenTypeNumeric,
		TokenLength:          "4",
    // expiration time in minutes
		ExpirationTime:       6,
    // mobile number in the 13 digit format
		CustomerMobileNumber: "2348143222998",
		CustomerEmailAddress: "abc@gmail.com",
    // your metadata struct
		MetaData:             metadata{"Shina", "Ebuka"},
	}

  res, err := verification.SendOTP(payload)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  // use vscode autocomplete to see more values
  ```

- `ConfirmOTP`
  > Confirm OTP sent to user. Refer to [verification_test.go](verification_test.go).
  ```go
	code, reference := "01799", "de858be1-6240-48fb-916c-4d07d8c9f79d"
  res, err := verification.ConfirmOTP(code, reference)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  // use vscode autocomplete/intellisense to see more properties
  ```

## Whatsapp
### Initialize
```go
publicKey :=  &sendchamp.Keys{
	PublicKey: "your-public-key",
 } //pass as nil to use the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
 mode := sendchamp.ModeLive
 //if the first parameter is nil, the package would look for the enviroment variable 'SENDCHAMP_PUBLIC_KEY'
 client := sendchamp.NewClient(publicKey, mode)
 whatsapp := client.NewWhatsapp()
```

### Methods
- `SendTemplate`
  > Send a whatsapp message using template created on dashboard. Refer to [whatsapp_test.go](whatsapp_test.go).
  ```go
	sender := "2348120678278"
	recipient := "2348153207998"
	templateCode := "912671fe-5f20-4b59-92ee-a33a62ea6a19"
	data := map[string]string{
		"1": "Test",
		"2": "1234",
		"3": "10",
	}

	res, err := whatsapp.SendTemplate(sender, recipient, templateCode, data)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  ```

- `SendText`
  > Send a whatsapp text. Refer to [whatsapp_test.go](whatsapp_test.go).
  ```go
	sender := "2348120678278"
	recipient := "2348153207998"
	message := "Hello World"

	res, err := whatsapp.SendText(sender, recipient, message)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  ```

- `SendAudio`
  > Send a whatsapp audio message. Refer to [whatsapp_test.go](whatsapp_test.go).
  ```go
	sender := "2348120678278"
	recipient := "2348153207998"
	message := "I am the best"
	link := "https://sample-videos.com/audio/mp3/crowd-cheering.mp3"

	res, err := whatsapp.SendAudio(sender, recipient, message, link)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  ```

- `SendVideo`
  > Send a whatsapp video message. Refer to [whatsapp_test.go](whatsapp_test.go).
  ```go
	sender := "2348120678278"
	recipient := "2348153207998"
	link := "https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4"

	res, err := whatsapp.SendVideo(sender, recipient, link)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  ```

- `SendSticker`
  > Send a whatsapp sticker message. Refer to [whatsapp_test.go](whatsapp_test.go).
  ```go
	sender := "2348120678278"
	recipient := "2348153207998"
	link := "https://studio.posit.us/api/samples/sticker.webp"

	res, err := whatsapp.SendSticker(sender, recipient, link)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
  ```

- `SendLocation`
  > Send a location via whatsapp. Refer to [whatsapp_test.go](whatsapp_test.go).
  ```go
	sender := "2348120678278"
	recipient := "2348153207998"
	longitude := -46.662787
	latitude := -23.553610
	name := "Robbu Brazil"
	address := "Av. Angélica, 2530 - Bela Vista, São Paulo - SP, 01228-200"

	res, err := whatsapp.SendLocation(sender, recipient, longitude, latitude, name, address)
  // use err variables to check for errors like network errors, etc.
  if err != nil {
    // handle
  }
  // use res for api response
  // res.Status, res.Code, res.Message, res.Data.ID, res.Data.CreatedAt, etc.  
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
* [ ] Customer Service
* [ ] Customer Group Service

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center"><a href="https://github.com/sadiqful"><img src="https://avatars.githubusercontent.com/u/39303081?v=4?s=100" width="100px;" alt="Aliyu Abubakar"/><br /><sub><b>Aliyu Abubakar</b></sub></a><br /><a href="#business-sadiqful" title="Business development">💼</a> <a href="#content-sadiqful" title="Content">🖋</a> <a href="#financial-sadiqful" title="Financial">💵</a> <a href="#ideas-sadiqful" title="Ideas, Planning, & Feedback">🤔</a> <a href="#mentoring-sadiqful" title="Mentoring">🧑‍🏫</a> <a href="#platform-sadiqful" title="Packaging/porting to new platform">📦</a> <a href="#question-sadiqful" title="Answering Questions">💬</a> <a href="#tutorial-sadiqful" title="Tutorials">✅</a> <a href="#talk-sadiqful" title="Talks">📢</a></td>
      <td align="center"><a href="http://fuadolatunji.me"><img src="https://avatars.githubusercontent.com/u/65264054?v=4?s=100" width="100px;" alt="Fuad Olatunji"/><br /><sub><b>Fuad Olatunji</b></sub></a><br /><a href="https://github.com/fuadop/sendchamp/issues?q=author%3Afuadop" title="Bug reports">🐛</a> <a href="https://github.com/fuadop/sendchamp/commits?author=fuadop" title="Code">💻</a> <a href="#infra-fuadop" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="#maintenance-fuadop" title="Maintenance">🚧</a> <a href="#design-fuadop" title="Design">🎨</a> <a href="#example-fuadop" title="Examples">💡</a> <a href="#security-fuadop" title="Security">🛡️</a> <a href="#tool-fuadop" title="Tools">🔧</a> <a href="#userTesting-fuadop" title="User Testing">📓</a></td>
      <td align="center"><a href="https://sayopaul.tech"><img src="https://avatars.githubusercontent.com/u/21235901?v=4?s=100" width="100px;" alt="Sayo Paul"/><br /><sub><b>Sayo Paul</b></sub></a><br /><a href="https://github.com/fuadop/sendchamp/commits?author=sayopaul" title="Code">💻</a> <a href="https://github.com/fuadop/sendchamp/issues?q=author%3Asayopaul" title="Bug reports">🐛</a> <a href="https://github.com/fuadop/sendchamp/commits?author=sayopaul" title="Documentation">📖</a></td>
      <td align="center"><a href="https://showbaba.github.io/samuelsx2/"><img src="https://avatars.githubusercontent.com/u/28683674?v=4?s=100" width="100px;" alt="Sam"/><br /><sub><b>Sam</b></sub></a><br /><a href="https://github.com/fuadop/sendchamp/commits?author=ShowBaba" title="Code">💻</a> <a href="https://github.com/fuadop/sendchamp/issues?q=author%3AShowBaba" title="Bug reports">🐛</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!