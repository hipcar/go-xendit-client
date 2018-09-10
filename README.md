# go-xendit-client
Xendit Client written in Go. Un-official Xendit API Wrapper. 
- [Have trouble ?](https://github.com/hipcar/go-xendit-client/issues)
- [Submit changes/features ?](https://github.com/hipcar/go-xendit-client/pulls)

Documentation
=============

## Getting Started
```
go get github.com/hipcar/go-xendit-client
```

## Init
```go
package main

import (
	"github.com/hipcar/go-xendit-client"
)

func main() {
  xenditClient := xendit.NewClient()
  xenditClient.EnableLog = true // logging is false by default
  xenditClient.SecretKey = "YOUR_XENDIT_SECRET_KEY"
}
```

## Balance

### Account Type
|Supported Account Type| Value |
|--|--|
|xendit.AccountTypeCash| CASH |
|xendit.AccountTypeHolding| HOLDING |
|xendit.AccountTypeTax| TAX |

### Get Balance
```go
res, err := client.Balance.GetBalance(xendit.AccountTypeCash)
```

## Credit Card

## Virtual Account

## Retail Outlet

## Invoice

## Disbursement

## Forex

## Name Validator

## NPWP Validator
