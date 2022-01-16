<p align="center"><a href="https://www.logharvestor.com" target="_blank" rel="noopener" referrerpolicy='origin'><img width="70%" src="https://i.ibb.co/gwFL3jk/logo-drk.png" alt="LogHarvestor Logo"></a></p>


<p align="center">
  <a href="https://www.linkedin.com/company/log-harvestor" rel="nofollow">
    <img src="https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white" alt="Log Harvestor - LinkedIn"> 
  </a> &nbsp; 
  <a href="https://twitter.com/LogHarvestor" rel="nofollow">
    <img src="https://img.shields.io/badge/Twitter-%231DA1F2.svg?style=for-the-badge&logo=Twitter&logoColor=white" alt="LogHarvestor - Twitter">
  </a> &nbsp; 
  <a href="https://www.youtube.com/channel/UCS9BdZPla9UbUQ3AZJEzVvw" rel="nofollow">
    <img src="https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white" alt="Log Harvestor - You Tube">
  </a>
</p>

# log-harvestor-go

## Documentation
See [API Docs](https://www.logharvestor.com/docs/api) for Log-Harvestor.

This package is specific to `golang`. Please see our docs for other supported languages, or use standard HTTP requests.

## Installation
______________
```
go get github.com/logharvestor/log-harvestor-go
```

## Usage
_____________
This package requires that you have a **Log Harvestor** account, and *Forwarder's* created.
If you have not done this yet:
1. Go to [LogHarvestor.com](https://www.logharvestor.com)
2. Register for a new Account (This is free) [Register](https://app.logharvestor.com/register)  
3. Create a new Forwarder - [Link](https://app.logharvestor.com/forwarder)
4. Generate a Forwarder Token

Now you can use this forwarder token to send logs, by adding it to the config:
```Go
  pvt_token := "your_forwarder_token"
  fwdr := *NewForwarder(Config{Token: pvt_token})
	success, msg := suite.forwarder.log(Log{Type: "test", Msg: bson.M{title: "Hello World"}}})
```
## Configuration
___________

| Option              | Default                                 | Description                                     |
| :---                | :----                                   | :---                                            |
| **Token**           | ""                                      | The JWT token assigned to your forwarder        |
| **ApiUrl**          | https://app.logharvestor.com/log        | This should never change unless using proxies   |
| **Verbose**         | false                                   | Verbose mode prints info to the console         |

### *Note*: The `LogHarvestorGo v.1.x.x` mod versions do not support batching or compression

___

## Examples
- ### **Configuring**
```Go
  pvt_token := "your_forwarder_token"
  fwdr := *NewForwarder(Config{Token: pvt_token, Verbose: true})
```
- ### **Sending Logs**
```Go
  //  After config/init
  fwdr.log(Log{Type: "whatever", Msg: bson.M{title: "Hello World"}}})
  fwdr.log(Log{Type: "you", Msg: "GoodbyWorld"})

  type CustomLogMessageStruct struct{
    TreeName      string
    Family        string
    Age           int
    HeightFeet    int
  }
  customMsg := CustomLogMessageStruct{
    TreeType: "white pine",
    Family:   "inaceae",
    Age: 16,
    HeightFeet: 56
  }

  fwdr.log(Log{Type: "want", Msg: bson.M{customMsg}})
```

## Recomendations
________________
1. Keep your Logging specific, and consise. This makes searching faster and more accurate
2. No need to add timestamps or info about the forwarder. This information is automatically included with the log.


<p align="center"><a href="https://www.logharvestor.com" target="_blank" rel="noopener" referrerpolicy='origin'><img width="100" src="https://i.ibb.co/80sThNP/icon-drk.png" alt="LogHarvestor Logo"></a></p>
