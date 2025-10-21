# 🚀 Parallax SDK: Datadome & Perimeterx (Go)

**High-performance Go SDK for bypassing DataDome and PerimeterX anti-bot protection.**

Built for speed and efficiency, this SDK leverages Go's concurrency features to deliver lightning-fast cookie generation and seamless integration into your Go applications. Perfect for high-throughput scraping, automation, and testing workflows.

## ✨ Why Choose the Go SDK?

- **⚡ Lightning Fast**: Native Go performance with goroutine support for concurrent operations
- **🔧 Simple API**: Clean, idiomatic Go interfaces with comprehensive error handling
- **🎯 Type-Safe**: Full type safety with struct-based configuration
- **🔄 Flexible**: Multiple initialization options and custom timeout support
- **🌐 Production-Ready**: Battle-tested in high-volume production environments
- **📦 Zero Dependencies**: Minimal external dependencies for easy maintenance

---

## 🚀 Quick Start

Get started with ParallaxAPIs SDK's in under 5 minutes:

1. **Join our [Discord](https://www.parallaxsystems.io/join?s=gh)** - Connect with our community
2. **Create a ticket** - Request your API key
3. **Get your free trial** - Start testing immediately
4. **[Install the SDK](#-installation)** - Choose your preferred language
5. **Solve all anti-bots in seconds** - Start bypassing DataDome, PerimeterX & more

---

## 📦 Installation

```bash
go get github.com/parallaxapis/parallaxapis-sdk-go
```

---

## 🧑‍💻 Datadome Usage

### ⚡ SDK Initialization

```go
import (
    "time"
    "github.com/parallaxapis/parallaxapis-sdk-go"
)

// Basic initialization with API key
sdk := parallaxsdk.NewDatadomeSDK("key", "")

// Custom host
sdk := parallaxsdk.NewDatadomeSDK("key", "https://example.host.com")

// With custom timeout (default is 30 seconds)
sdk := parallaxsdk.NewDatadomeSDK("key", "", parallaxsdk.WithCustomTimeout(60*time.Second))

// With HTTP proxy for client requests
sdk := parallaxsdk.NewDatadomeSDK("key", "", parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"))

// Multiple options combined
sdk := parallaxsdk.NewDatadomeSDK("key", "https://example.host.com",
    parallaxsdk.WithCustomTimeout(45*time.Second),
    parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"))
```

### 🕵️‍♂️ Generate New User Agent

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
userAgent, err := sdk.GenerateUserAgent(parallaxsdk.TaskGenUserAgent{
    Region: "pl",
    Site: "vinted",
})
if err != nil {
    panic(err)
}
fmt.Println(userAgent)
// Output:
// &parallaxsdk.UserAgentResponse{
//     Message: "New device successfully created.",
//     UserAgent: "Mozilla/5.0 ...",
//     SecHeader: "...",
//     SecFullVersionList: "...",
//     SecPlatform: "...",
//     SecArch: "...",
// }
```

### 🔍 Get Task Data

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
challengeURL := "https://geo.captcha-delivery.com/captcha/?initialCid=initialCid&cid=cid&referer=referer&hash=hash&t=t&s=s&e=e"
cookie := "cookie"
taskData, productType, err := parallaxsdk.ParseChallengeURL(challengeURL, cookie)
if err != nil {
    panic(err)
}
fmt.Println(taskData, productType)
// Output:
// &parallaxsdk.TaskDatadomeCookieData{
//     Cid: "cookie",
//     B: "",
//     E: "e",
//     S: "s",
//     InitialCid: "initialCid",
// }, "captcha"
```

### 🍪 Generate Cookie

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
challengeURL := "https://geo.captcha-delivery.com/captcha/?initialCid=initialCid&cid=cid&referer=referer&hash=hash&t=t&s=s&e=e"
cookie := "cookie"
taskData, productType, err := parallaxsdk.ParseChallengeURL(challengeURL, cookie)
if err != nil {
    panic(err)
}
cookieResp, err := sdk.GenerateDatadomeCookie(parallaxsdk.TaskDatadomeCookie{
    Site: "vinted",
    Region: "pl",
    Data: *taskData,
    Pd: productType,
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
})
if err != nil {
    panic(err)
}
fmt.Println(cookieResp)
// Output:
// &parallaxsdk.DatadomeCookieResponse{
//     Message: "datadome=cookie_value",
//     UserAgent: "Mozilla/5.0 ...",
// }
```

### 🏷️ Generate Tags Cookie

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
cookieResp, err := sdk.GenerateDatadomeTagsCookie(parallaxsdk.TaskDatadomeTagsCookie{
    Site: "vinted",
    Region: "pl",
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
})
if err != nil {
    panic(err)
}
fmt.Println(cookieResp)
// Output:
// &parallaxsdk.DatadomeCookieResponse{
//     Message: "datadome=cookie_value",
//     UserAgent: "Mozilla/5.0 ...",
// }
```

### 🔍 Detect and Parse Challenge

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
responseBody := "<html>...</html>" // Response body from website
prevCookie := "previous_datadome_cookie"

isBlocked, taskData, productType, err := parallaxsdk.DetectChallengeAndParse(responseBody, prevCookie)
if err != nil {
    panic(err)
}
if isBlocked {
    fmt.Printf("Datadome challenge detected: %s\n", productType)
    // Use taskData with GenerateDatadomeCookie to solve the challenge
    cookieResp, err := sdk.GenerateDatadomeCookie(parallaxsdk.TaskDatadomeCookie{
        Site: "vinted",
        Region: "pl",
        Data: *taskData,
        Pd: productType,
        Proxy: "http://user:pas@addr:port",
        Proxyregion: "eu",
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(cookieResp)
}
```

### 📄 Parse Challenge HTML

```go
htmlBody := "<html><script>dd={cid:'abc',t:'fe',s:123,e:'xyz',b:'1'}</script></html>"
prevCookie := "previous_datadome_cookie"

taskData, productType, err := parallaxsdk.ParseChallengeHTML(htmlBody, prevCookie)
if err != nil {
    panic(err)
}
fmt.Println(taskData, productType)
// Output:
// &parallaxsdk.TaskDatadomeCookieData{
//     Cid: "previous_datadome_cookie",
//     B: "1",
//     E: "xyz",
//     S: "123",
//     InitialCid: "abc",
// }, "captcha"
```

---

## 🛡️ Perimeterx Usage

### ⚡ SDK Initialization

```go
import (
    "time"
    "github.com/parallaxapis/parallaxapis-sdk-go"
)

// Basic initialization with API key
sdk := parallaxsdk.NewPerimeterxSDK("key", "")

// Custom host
sdk := parallaxsdk.NewPerimeterxSDK("key", "example.host.com")

// With custom timeout (default is 30 seconds)
sdk := parallaxsdk.NewPerimeterxSDK("key", "", parallaxsdk.WithCustomTimeout(60*time.Second))

// With HTTP proxy for client requests
sdk := parallaxsdk.NewPerimeterxSDK("key", "", parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"))

// Multiple options combined
sdk := parallaxsdk.NewPerimeterxSDK("key", "example.host.com",
    parallaxsdk.WithCustomTimeout(45*time.Second),
    parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"))
```

### 🍪 Generate PX Cookie

```go
sdk := parallaxsdk.NewPerimeterxSDK("key", "")
result, err := sdk.GenerateCookies(parallaxsdk.TaskGeneratePXCookies{
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
    Region: "com",
    Site: "stockx",
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", result)
// Output:
// &parallaxsdk.PxCookieResponse{
//     Cookie: "_px3=d3sswjaltwxgAd...",
//     Vid: "514d7e11-6962-11f0-810f-88cc16043287",
//     Cts: "514d8e28-6962-11f0-810f-51b6xf2786b0",
//     IsFlagged: false,
//     IsMaybeFlagged: true,
//     UserAgent: "Mozilla/5.0 ...",
//     Data: "==WlrBti6vpO6rshP1CFtBsiocoO8...",
// }

holdCaptchaResult, err := sdk.GenerateHoldCaptcha(parallaxsdk.TaskGenerateHoldCaptcha{
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
    Region: "com",
    Site: "stockx",
    Data: result.Data,
    PowPro: "",
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", holdCaptchaResult)
// Output:
// &parallaxsdk.PxCookieResponse{
//     Cookie: "_px3=d3sswjaltwxgAd...",
//     Vid: "514d7e11-6962-11f0-810f-88cc16043287",
//     FlaggedPow: false,
// }
```

---

## 📚 Documentation & Help

- Full API docs & support: [Discord](https://www.parallaxsystems.io/join?s=gh)



## 🌟 Contributing

Got feedback or found a bug? Feel free to open an issue or send us a pull request!



## 🏢 Enterprise

Unlock enterprise-grade performance with custom solutions, expanded limits, and expert support. Contact us to learn more.



## 📝 License

MIT

---

## 🔑 Keywords

**DataDome bypass** • **PerimeterX bypass** • **Anti-bot bypass** • **Bot detection bypass** • **CAPTCHA solver** • **Cookie generator** • **Go web scraping** • **Go bot automation** • **Golang anti-bot** • **DataDome Go SDK** • **PerimeterX Go SDK** • **Headless browser alternative** • **Request-based bypass** • **Go automation** • **Web scraping Go** • **Bot mitigation bypass** • **Sensor data generation** • **Challenge solver**
